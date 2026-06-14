// crawl.mjs
//
// Recursively drives the UniFi Network Application SPA and records every
// request/response to a HAR file for OpenAPI generation.
//
// Strategy (see README.md for the full design):
//   1. Log in (simulation-mode admin/admin) and dismiss any first-run wizard.
//   2. BFS over every in-app route reachable from the side nav + a known-route
//      seed list  -> exhaustive GET coverage of the read API surface.
//   3. On every page, "exercise" the forms (when MUTATE=1):
//        a. submit the empty form               -> 400/validation XHR
//        b. fill the form with valid dummy data  -> POST (create)
//        c. re-open the created item, tweak it   -> PUT/PATCH (update)
//        d. delete the created item              -> DELETE
//      All requests are captured whether they succeed or fail; even a
//      validation 400 documents the endpoint + payload shape.
//   4. Close the context -> Playwright flushes the HAR; we also emit a
//      hand-rolled api-endpoints.json summary of the distinct XHR seen.
//
// The controller is disposable (simulation mode) so destructive mutation is safe.

import { chromium } from 'playwright';
import { mkdirSync, writeFileSync } from 'node:fs';
import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));

// ---------------------------------------------------------------- config ----
const cfg = {
  baseURL: (process.env.UNIFI_API || 'https://localhost:8443').replace(/\/$/, ''),
  user: process.env.UNIFI_USERNAME || 'admin',
  pass: process.env.UNIFI_PASSWORD || 'admin',
  site: process.env.UNIFI_SITE || 'default',
  headless: process.env.HEADLESS !== '0',
  mutate: process.env.MUTATE !== '0', // default on
  outDir: resolve(process.env.OUT_DIR || resolve(__dirname, 'out')),
  // hard wall-clock budget so the crawl always terminates and flushes the HAR
  budgetMs: Number(process.env.BUDGET_MS || 20 * 60 * 1000),
  navTimeout: Number(process.env.NAV_TIMEOUT_MS || 30000),
  idleTimeout: Number(process.env.IDLE_TIMEOUT_MS || 8000),
};

mkdirSync(cfg.outDir, { recursive: true });
const harPath = resolve(cfg.outDir, 'unifi-ui.har');
const summaryPath = resolve(cfg.outDir, 'api-endpoints.json');
const routesPath = resolve(cfg.outDir, 'routes-visited.json');

const t0 = Date.now();
const overBudget = () => Date.now() - t0 > cfg.budgetMs;
const log = (...a) => console.log(`[crawl +${((Date.now() - t0) / 1000).toFixed(0)}s]`, ...a);
const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

// Known UniFi Network settings/page routes, merged with dynamic discovery so
// coverage does not depend solely on the nav DOM (which varies by version).
const seedRoutes = (site) =>
  [
    `/manage/site/${site}/dashboard`,
    `/manage/site/${site}/devices`,
    `/manage/site/${site}/clients`,
    `/manage/site/${site}/insights`,
    `/manage/site/${site}/topology`,
    `/manage/site/${site}/traffic`,
    // settings tree (covers v8/v9 UI)
    'wifi', 'networks', 'internet', 'vpn', 'teleport', 'profiles',
    'routing', 'firewall', 'traffic-management', 'traffic-rules',
    'port-forwarding', 'security', 'multicast-dns', 'controls',
    'system', 'admins', 'auto-advanced', 'guest_hotspot', 'radius',
    'network-isolation', 'global-network',
  ].map((r) => (r.startsWith('/') ? r : `/manage/site/${site}/settings/${r}`));

// --------------------------------------------------------- xhr summary ------
/** url(no-query) -> { methods:Set, statuses:Set, sample } */
const endpoints = new Map();
function recordXhr(method, url, status, resourceType) {
  try {
    const u = new URL(url);
    // only the controller's own API/XHR surface
    if (!/\/(api|proxy|v2|status|wss)\b/.test(u.pathname) && resourceType !== 'xhr' && resourceType !== 'fetch')
      return;
    if (!/^(xhr|fetch)$/.test(resourceType) && !u.pathname.startsWith('/api') && !u.pathname.startsWith('/proxy') && !u.pathname.startsWith('/v2'))
      return;
    const key = `${u.pathname}`;
    if (!endpoints.has(key)) endpoints.set(key, { path: key, methods: new Set(), statuses: new Set(), examples: [] });
    const e = endpoints.get(key);
    e.methods.add(method);
    e.statuses.add(status);
    if (e.examples.length < 3) e.examples.push(`${method} ${u.pathname}${u.search} -> ${status}`);
  } catch {
    /* ignore */
  }
}

// ------------------------------------------------------- DOM interaction ----
// Fill a single input/select/textarea with a plausible value derived from its
// name/label/placeholder/type. Returns true if it set something.
const fillScript = `
(root) => {
  const rand = String(Math.floor(Math.random()*9000)+1000);
  const guess = (el) => {
    const hint = [el.name, el.id, el.getAttribute('placeholder'), el.getAttribute('aria-label'),
      (el.labels && el.labels[0] && el.labels[0].textContent) || ''].join(' ').toLowerCase();
    const type = (el.type || '').toLowerCase();
    if (type === 'email') return 'crawler' + rand + '@example.com';
    if (type === 'number') return '10';
    if (type === 'password') return 'Passw0rd!' + rand;
    if (/\\bip\\b|address|gateway|subnet/.test(hint)) return '10.99.' + (rand.slice(0,1)) + '.1';
    if (/mac/.test(hint)) return '00:11:22:33:44:5' + rand.slice(-1);
    if (/vlan/.test(hint)) return '99';
    if (/port/.test(hint)) return '8081';
    if (/url|http/.test(hint)) return 'https://example.com';
    if (/email/.test(hint)) return 'crawler' + rand + '@example.com';
    if (/pass|secret|psk|key/.test(hint)) return 'Passw0rd!' + rand;
    if (/name|ssid|desc|label/.test(hint)) return 'crawler-' + rand;
    return 'crawler-' + rand;
  };
  let touched = 0;
  const scope = root || document;
  scope.querySelectorAll('input, textarea, select').forEach((el) => {
    if (el.disabled || el.readOnly || el.offsetParent === null) return;
    const type = (el.type || '').toLowerCase();
    try {
      if (el.tagName === 'SELECT') {
        const opts = [...el.options].filter(o => o.value && !o.disabled);
        if (opts.length) { el.value = opts[opts.length-1].value; el.dispatchEvent(new Event('change',{bubbles:true})); touched++; }
      } else if (type === 'checkbox' || type === 'radio') {
        if (!el.checked) { el.click(); touched++; }
      } else if (['hidden','submit','button','file'].includes(type)) {
        /* skip */
      } else {
        const v = guess(el);
        const setter = Object.getOwnPropertyDescriptor(window.HTMLInputElement.prototype,'value') ||
                       Object.getOwnPropertyDescriptor(window.HTMLTextAreaElement.prototype,'value');
        setter.set.call(el, v);
        el.dispatchEvent(new Event('input',{bubbles:true}));
        el.dispatchEvent(new Event('change',{bubbles:true}));
        touched++;
      }
    } catch (e) {}
  });
  return touched;
}`;

async function clickByText(scope, regex, { timeout = 1500 } = {}) {
  // Click the first visible button/link whose text matches.
  const els = scope.locator('button, a, [role=button], [role=tab]');
  const n = Math.min(await els.count().catch(() => 0), 60);
  for (let i = 0; i < n; i++) {
    const el = els.nth(i);
    try {
      if (!(await el.isVisible())) continue;
      const txt = ((await el.innerText({ timeout: 500 }).catch(() => '')) || '').trim();
      if (regex.test(txt)) {
        await el.click({ timeout });
        return txt || true;
      }
    } catch {
      /* keep scanning */
    }
  }
  return false;
}

async function waitIdle(page) {
  await page.waitForLoadState('networkidle', { timeout: cfg.idleTimeout }).catch(() => {});
}

// Exercise the forms on the current page: validation-probe, create, update, delete.
async function exerciseForms(page) {
  if (!cfg.mutate || overBudget()) return;
  try {
    // 1. Open a "create / add new" affordance if present.
    const opened = await clickByText(page, /^(\+|add|create|new|create new|add new)\b/i, { timeout: 2000 });
    await sleep(500);

    // The form may be a modal/drawer; scope to a dialog if one is present.
    const dialog = page.locator('[role=dialog], .modal, [class*=drawer], [class*=panel]').first();
    const formScope = (await dialog.count().catch(() => 0)) && (await dialog.isVisible().catch(() => false))
      ? dialog
      : page;

    // 2a. Validation probe: submit empty -> capture 400/validation XHR.
    await clickByText(formScope, /^(save|apply|add|create|update|confirm|done)\b/i, { timeout: 1500 });
    await waitIdle(page);

    // 2b. Fill with dummy data and submit -> POST (create).
    const handle = (await formScope.elementHandle?.().catch(() => null)) || null;
    const touched = await page.evaluate(fillScript, handle).catch(() => 0);
    if (touched) {
      await clickByText(formScope, /^(save|apply|add|create|next|update|confirm|done)\b/i, { timeout: 2000 });
      await waitIdle(page);
      log(`  exercised form (filled ${touched} field(s))`);
    }

    // 2c. Update: re-open the first row/item, tweak, save -> PUT/PATCH.
    const row = page.locator('table tbody tr, [class*=list] [class*=row], li[class*=item]').first();
    if (await row.count().catch(() => 0)) {
      await row.click({ timeout: 1500 }).catch(() => {});
      await sleep(400);
      await page.evaluate(fillScript, null).catch(() => {});
      await clickByText(page, /^(save|apply|update|confirm|done)\b/i, { timeout: 1500 });
      await waitIdle(page);
    }

    // 2d. Delete: open the item and remove it -> DELETE.
    await clickByText(page, /^(delete|remove|forget|trash)\b/i, { timeout: 1500 });
    await sleep(300);
    await clickByText(page, /^(delete|remove|confirm|yes|ok)\b/i, { timeout: 1500 }); // confirm dialog
    await waitIdle(page);

    // Close any lingering modal so it does not block the next route.
    await page.keyboard.press('Escape').catch(() => {});
  } catch (e) {
    log(`  exerciseForms note: ${e.message?.split('\n')[0]}`);
  }
}

// Click every "safe" button/tab on the page to reveal sub-panels and their XHR,
// without leaving the route. Returns nothing; XHR is captured passively.
async function clickTabs(page) {
  if (overBudget()) return;
  const tabs = page.locator('[role=tab], [class*=tab]:not([class*=table]), [class*=segment] button');
  const n = Math.min(await tabs.count().catch(() => 0), 25);
  for (let i = 0; i < n && !overBudget(); i++) {
    try {
      const el = tabs.nth(i);
      if (await el.isVisible()) {
        await el.click({ timeout: 1000 });
        await page.waitForLoadState('networkidle', { timeout: 3000 }).catch(() => {});
      }
    } catch {
      /* ignore */
    }
  }
}

// Discover in-app navigation targets from the current DOM.
async function discoverLinks(page) {
  return page
    .evaluate(() => {
      const out = new Set();
      document.querySelectorAll('a[href]').forEach((a) => {
        const href = a.getAttribute('href') || '';
        if (href.startsWith('/manage') || /\/site\/[^/]+\//.test(href)) out.add(href);
      });
      return [...out];
    })
    .catch(() => []);
}

// ----------------------------------------------------------------- login ----
async function login(page) {
  log('navigating to login...');
  await page.goto(`${cfg.baseURL}/manage/account/login`, { waitUntil: 'domcontentloaded', timeout: cfg.navTimeout });
  await sleep(2000);
  // Already authenticated? (simulation mode may auto-session)
  if (!/login/.test(page.url())) {
    log('no login required (already authenticated)');
    return;
  }
  const userSel = 'input[name=username], input[type=text], input[autocomplete=username]';
  const passSel = 'input[name=password], input[type=password]';
  await page.fill(userSel, cfg.user, { timeout: cfg.navTimeout }).catch(() => {});
  await page.fill(passSel, cfg.pass, { timeout: cfg.navTimeout }).catch(() => {});
  await Promise.all([
    page.waitForNavigation({ timeout: cfg.navTimeout }).catch(() => {}),
    (async () => {
      if (!(await clickByText(page, /^(sign in|log ?in|login|continue|submit)\b/i, { timeout: 3000 }))) {
        await page.press(passSel, 'Enter').catch(() => {});
      }
    })(),
  ]);
  await waitIdle(page);
  log(`post-login url: ${page.url()}`);
}

// Dismiss first-run setup wizard / cookie / "what's new" dialogs.
async function dismissWizards(page) {
  for (let i = 0; i < 4; i++) {
    const hit =
      (await clickByText(page, /^(skip|maybe later|later|not now|got it|accept|dismiss|close|next|finish|done|continue)\b/i, { timeout: 1200 }));
    if (!hit) break;
    await sleep(600);
  }
}

// ------------------------------------------------------------------ main ----
async function main() {
  log(`launching chromium (headless=${cfg.headless}, mutate=${cfg.mutate})`);
  const browser = await chromium.launch({ headless: cfg.headless, args: ['--ignore-certificate-errors'] });
  const context = await browser.newContext({
    ignoreHTTPSErrors: true,
    recordHar: {
      path: harPath,
      mode: 'full',
      content: 'embed',
      // Record only the XHR/API surface (not static JS/CSS/img bundles) so the
      // HAR stays small and is directly consumable for OpenAPI generation.
      urlFilter: /\/(api|proxy|v2|wss|status|guest)\b/,
    },
    viewport: { width: 1600, height: 1000 },
  });
  const page = await context.newPage();
  page.setDefaultTimeout(cfg.navTimeout);

  context.on('response', (resp) => {
    const req = resp.request();
    recordXhr(req.method(), req.url(), resp.status(), req.resourceType());
  });
  page.on('dialog', (d) => d.accept().catch(() => {})); // auto-accept native confirm()

  const visited = new Set();
  const queue = [];
  const enqueue = (href) => {
    if (!href) return;
    let path;
    try {
      path = new URL(href, cfg.baseURL).pathname;
    } catch {
      return;
    }
    if (!path.startsWith('/manage')) return;
    if (visited.has(path) || queue.includes(path)) return;
    queue.push(path);
  };

  try {
    await login(page);
    await dismissWizards(page);

    // Seed the queue: known routes + whatever the post-login DOM exposes.
    seedRoutes(cfg.site).forEach(enqueue);
    (await discoverLinks(page)).forEach(enqueue);
    log(`seeded ${queue.length} routes`);

    while (queue.length && !overBudget()) {
      const path = queue.shift();
      if (visited.has(path)) continue;
      visited.add(path);
      const url = `${cfg.baseURL}${path}`;
      log(`(${visited.size}) visiting ${path}  [queue ${queue.length}]`);
      try {
        await page.goto(url, { waitUntil: 'domcontentloaded', timeout: cfg.navTimeout });
        await waitIdle(page);
        await dismissWizards(page);
        // discover more routes from this page
        (await discoverLinks(page)).forEach(enqueue);
        // reveal sub-panels
        await clickTabs(page);
        // mutate
        await exerciseForms(page);
      } catch (e) {
        log(`  visit error: ${e.message?.split('\n')[0]}`);
      }
    }

    log(`crawl finished: ${visited.size} routes, ${endpoints.size} distinct endpoints` + (overBudget() ? ' (budget hit)' : ''));
  } finally {
    // Flush HAR + summaries no matter what.
    await context.close().catch(() => {});
    await browser.close().catch(() => {});

    const summary = [...endpoints.values()]
      .map((e) => ({ path: e.path, methods: [...e.methods].sort(), statuses: [...e.statuses].sort(), examples: e.examples }))
      .sort((a, b) => a.path.localeCompare(b.path));
    writeFileSync(summaryPath, JSON.stringify(summary, null, 2));
    writeFileSync(routesPath, JSON.stringify([...visited].sort(), null, 2));
    log(`HAR   -> ${harPath}`);
    log(`API   -> ${summaryPath} (${summary.length} endpoints)`);
    log(`ROUTES-> ${routesPath} (${visited.size} routes)`);
  }
}

main().catch((e) => {
  console.error('[crawl] fatal', e);
  process.exit(1);
});
