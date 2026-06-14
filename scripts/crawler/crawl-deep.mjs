// crawl-deep.mjs
//
// Exhaustive UniFi Network Application UI recursion.
//
// Unlike crawl.mjs (fast BFS for GET coverage), this pass genuinely *clicks
// every interactive element* and *submits every form* on every route, so the
// HAR captures the full write surface (POST / PUT / PATCH / DELETE).
//
// How it stays correct on a custom-component Angular SPA:
//   * Every clickable element is stamped with a persistent data-crawl-id, so it
//     is clicked exactly once even as the DOM mutates (stamps survive
//     SPA-internal re-renders; only a full reload resets them).
//   * Route-changing <a href>s are left to the BFS — here we click the buttons,
//     tabs, menu items, switches and icons that mutate in-page state or open
//     dialogs, which is where the write API lives.
//   * When a click opens a dialog/drawer, we run the form sequence inside it:
//     submit empty (validation 4xx) -> fill -> submit (create) -> close.
//   * A destructive denylist (logout / restart / factory-reset / delete-site)
//     keeps the session and the controller alive; if we still get bounced to
//     the login page we re-authenticate and continue.
//
// The controller runs in disposable simulation mode, so create/update/delete
// against it is safe.

import { chromium } from 'playwright';
import { mkdirSync, writeFileSync } from 'node:fs';
import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));

const cfg = {
  baseURL: (process.env.UNIFI_API || 'https://localhost:8443').replace(/\/$/, ''),
  user: process.env.UNIFI_USERNAME || 'admin',
  pass: process.env.UNIFI_PASSWORD || 'admin',
  site: process.env.UNIFI_SITE || 'default',
  headless: process.env.HEADLESS !== '0',
  outDir: resolve(process.env.OUT_DIR || resolve(__dirname, 'out-deep')),
  budgetMs: Number(process.env.BUDGET_MS || 30 * 60 * 1000),
  pageMs: Number(process.env.PAGE_MS || 75 * 1000), // per-route interaction cap
  maxClicksPerPage: Number(process.env.MAX_CLICKS || 140),
  navTimeout: Number(process.env.NAV_TIMEOUT_MS || 30000),
  idleTimeout: Number(process.env.IDLE_TIMEOUT_MS || 6000),
};

mkdirSync(cfg.outDir, { recursive: true });
const harPath = resolve(cfg.outDir, 'unifi-ui.har');
const summaryPath = resolve(cfg.outDir, 'api-endpoints.json');
const routesPath = resolve(cfg.outDir, 'routes-visited.json');

const t0 = Date.now();
const overBudget = () => Date.now() - t0 > cfg.budgetMs;
const log = (...a) => console.log(`[deep +${((Date.now() - t0) / 1000).toFixed(0)}s]`, ...a);
const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

const seedRoutes = (s) =>
  [
    // Settings forms first — that's where create/update/delete (POST/PUT/DELETE)
    // lives, and the per-page budget is finite. Monitoring pages last.
    'wifi', 'networks', 'internet', 'vpn', 'teleport', 'profiles', 'routing',
    'firewall', 'traffic-management', 'traffic-and-firewall-rules', 'port-forwarding',
    'security', 'multicast-dns', 'controls', 'system', 'admins', 'radius',
    'network-isolation', 'global-network', 'settings_overview',
    `/manage/${s}/devices`, `/manage/${s}/clients/main`, `/manage/${s}/insights/flows`,
    `/manage/${s}/ports`, `/manage/${s}/topology`, `/manage/${s}/dashboard`,
  ].map((r) => (r.startsWith('/') ? r : `/manage/${s}/settings/${r}`));

// ---- xhr summary ----
const endpoints = new Map();
function recordXhr(method, url, status, rtype) {
  try {
    const u = new URL(url);
    const isApi = /^\/(api|proxy|v2|status|guest)\b/.test(u.pathname);
    if (!isApi && rtype !== 'xhr' && rtype !== 'fetch') return;
    const key = u.pathname;
    if (!endpoints.has(key)) endpoints.set(key, { path: key, methods: new Set(), statuses: new Set(), examples: [] });
    const e = endpoints.get(key);
    e.methods.add(method);
    e.statuses.add(status);
    if (e.examples.length < 3) e.examples.push(`${method} ${u.pathname}${u.search} -> ${status}`);
  } catch {}
}

// ---- generic field filler (runs in page) ----
const FILL = `
(rootSel) => {
  const root = rootSel ? document.querySelector(rootSel) : document;
  if (!root) return 0;
  const rand = String(Math.floor(Math.random()*9000)+1000);
  const set = (el, v) => {
    const proto = el.tagName === 'TEXTAREA' ? window.HTMLTextAreaElement.prototype : window.HTMLInputElement.prototype;
    const d = Object.getOwnPropertyDescriptor(proto, 'value');
    d.set.call(el, v);
    el.dispatchEvent(new Event('input',{bubbles:true}));
    el.dispatchEvent(new Event('change',{bubbles:true}));
  };
  const guess = (el) => {
    const h = [el.name, el.id, el.getAttribute('placeholder'), el.getAttribute('aria-label'),
      (el.labels&&el.labels[0]&&el.labels[0].textContent)||''].join(' ').toLowerCase();
    const t = (el.type||'').toLowerCase();
    if (t==='email'||/email/.test(h)) return 'crawler'+rand+'@example.com';
    if (t==='number') return '10';
    if (/\\bip\\b|address|gateway|subnet|dns/.test(h)) return '10.99.'+rand.slice(0,1)+'.2';
    if (/mac/.test(h)) return '00:11:22:33:44:'+('5'+rand.slice(-1)).slice(-2);
    if (/vlan/.test(h)) return '99';
    if (/port/.test(h)) return '8081';
    if (/url|http/.test(h)) return 'https://example.com';
    if (/pass|secret|psk|key|wpa/.test(h)||t==='password') return 'Passw0rd!'+rand;
    return 'crawler-'+rand;
  };
  let n=0;
  root.querySelectorAll('input,textarea,select').forEach((el)=>{
    if (el.disabled||el.readOnly) return;
    const r = el.getBoundingClientRect(); if (r.width===0&&r.height===0) return;
    const t=(el.type||'').toLowerCase();
    try {
      if (el.tagName==='SELECT'){const o=[...el.options].filter(x=>x.value&&!x.disabled);if(o.length){el.value=o[o.length-1].value;el.dispatchEvent(new Event('change',{bubbles:true}));n++;}}
      else if (t==='checkbox'||t==='radio'){if(!el.checked){el.click();n++;}}
      else if (['hidden','submit','button','file','image'].includes(t)){}
      else {set(el, guess(el)); n++;}
    } catch(e){}
  });
  return n;
}`;

// Stamp every not-yet-seen clickable element; returns the freshly stamped ones.
const STAMP = `
() => {
  if (!window.__cn) window.__cn = 0;
  const deny = /log\\s?out|sign\\s?out|restart|reboot|shut\\s?down|power\\s?off|factory|reset to default|\\bformat\\b|restore|erase|delete site|remove site|forget site|leave|wipe/i;
  // route-changing anchors are the BFS's job; skip them here
  const sel = 'button,[role=button],[role=tab],[role=menuitem],[role=menuitemcheckbox],[role=switch],[data-testid],[class*=button],[class*=btn],[class*=Toggle],[class*=toggle],svg[class*=icon]';
  const out=[];
  document.querySelectorAll(sel).forEach((el)=>{
    if (el.hasAttribute('data-crawl-id')) return;
    if (el.closest('a[href]') && /^\\/?manage/.test(el.closest('a[href]').getAttribute('href')||'')) return;
    const r = el.getBoundingClientRect();
    if (r.width===0||r.height===0||r.bottom<0||r.top>window.innerHeight*4) return;
    const text=(el.innerText||el.getAttribute('aria-label')||el.title||el.getAttribute('data-testid')||'').trim().slice(0,50);
    if (deny.test(text)) return;
    el.setAttribute('data-crawl-id', String(window.__cn));
    out.push({id: window.__cn, text});
    window.__cn++;
  });
  return out;
}`;

// Invoke a function-source string in the page WITH args. (page.evaluate(string)
// treats the string as an expression and would NOT call the function, so we
// build an invoked IIFE expression with JSON-encoded args.)
const evalFn = (page, src, arg) =>
  page.evaluate(`(${src})(${arg === undefined ? '' : JSON.stringify(arg)})`);

async function settle(page) {
  await page.waitForLoadState('networkidle', { timeout: cfg.idleTimeout }).catch(() => {});
}

async function dialogOpen(page) {
  const d = page.locator('[role=dialog],[class*=modal],[class*=drawer],[class*=Drawer],[class*=Modal]').first();
  return (await d.count().catch(() => 0)) > 0 && (await d.isVisible().catch(() => false));
}

async function clickByText(scope, regex, timeout = 1500) {
  const els = scope.locator('button,a,[role=button]');
  const n = Math.min(await els.count().catch(() => 0), 40);
  for (let i = 0; i < n; i++) {
    const el = els.nth(i);
    try {
      if (!(await el.isVisible())) continue;
      const txt = ((await el.innerText({ timeout: 400 }).catch(() => '')) || '').trim();
      if (regex.test(txt)) { await el.click({ timeout }); return txt || true; }
    } catch {}
  }
  return false;
}

// Submit the form inside the current dialog: empty -> filled -> (multi-step) -> close.
async function submitForm(page) {
  const dlgSel = '[role=dialog],[class*=modal],[class*=drawer],[class*=Drawer],[class*=Modal]';
  const dlg = page.locator(dlgSel).first();
  const SAVE = /^(save|apply|add|create|next|continue|update|confirm|finish|done|provision|enable)\b/i;
  try {
    // 1. validation probe (submit as-is)
    await clickByText(dlg, SAVE, 1200);
    await settle(page);
    // 2. fill + submit, walking up to 3 wizard steps
    for (let step = 0; step < 3; step++) {
      if (!(await dialogOpen(page))) break;
      await evalFn(page, FILL, dlgSel).catch(() => {});
      const clicked = await clickByText(dlg, SAVE, 1500);
      await settle(page);
      if (!clicked) break;
    }
  } catch {}
  // 3. close whatever remains
  await clickByText(page, /^(cancel|close|dismiss)\b/i, 800);
  await page.keyboard.press('Escape').catch(() => {});
  await sleep(150);
}

async function ensureLoggedIn(page, routeUrl) {
  if (!/\/account\/login|\/login/.test(page.url())) return;
  log('  bounced to login - re-authenticating');
  await page.fill('input[name=username],input[type=text]', cfg.user).catch(() => {});
  await page.fill('input[name=password],input[type=password]', cfg.pass).catch(() => {});
  await clickByText(page, /^(sign in|log ?in|login|continue)\b/i, 2000);
  await settle(page);
  await page.goto(routeUrl, { waitUntil: 'domcontentloaded' }).catch(() => {});
  await settle(page);
}

// Exhaustively click everything on the current route + submit every dialog form.
async function exhaustPage(page, routeUrl) {
  const pageStart = Date.now();
  const pageUp = () => Date.now() - pageStart > cfg.pageMs;
  let clicks = 0;
  for (let pass = 0; pass < 8 && !pageUp() && !overBudget(); pass++) {
    const fresh = (await evalFn(page, STAMP).catch(() => [])) || [];
    if (!fresh.length) break;
    for (const c of fresh) {
      if (pageUp() || overBudget() || clicks >= cfg.maxClicksPerPage) break;
      clicks++;
      const before = page.url();
      try {
        await page.locator(`[data-crawl-id="${c.id}"]`).click({ timeout: 1200 });
      } catch {
        continue;
      }
      await settle(page);
      await ensureLoggedIn(page, routeUrl);
      if (await dialogOpen(page)) {
        await submitForm(page);
      }
      // A click that navigated to a different route: BFS owns those routes, so
      // return here and keep clicking the rest of this one.
      if (page.url() !== before && !page.url().startsWith(routeUrl.split('#')[0])) {
        await page.goto(routeUrl, { waitUntil: 'domcontentloaded' }).catch(() => {});
        await settle(page);
        break; // DOM/tags reset on reload; restart the pass loop
      }
    }
  }
  return clicks;
}

async function discoverLinks(page) {
  return page
    .evaluate(() => {
      const out = new Set();
      document.querySelectorAll('a[href]').forEach((a) => {
        const h = a.getAttribute('href') || '';
        if (h.startsWith('/manage')) out.add(h);
      });
      return [...out];
    })
    .catch(() => []);
}

async function login(page) {
  await page.goto(`${cfg.baseURL}/manage/account/login`, { waitUntil: 'domcontentloaded', timeout: cfg.navTimeout });
  await sleep(2000);
  if (/login/.test(page.url())) {
    await page.fill('input[name=username],input[type=text]', cfg.user).catch(() => {});
    await page.fill('input[name=password],input[type=password]', cfg.pass).catch(() => {});
    await Promise.all([
      page.waitForNavigation({ timeout: cfg.navTimeout }).catch(() => {}),
      clickByText(page, /^(sign in|log ?in|login|continue|submit)\b/i, 3000),
    ]);
    await settle(page);
  }
  for (let i = 0; i < 4; i++) {
    if (!(await clickByText(page, /^(skip|later|not now|got it|accept|dismiss|close|next|finish|done|continue)\b/i, 1000))) break;
    await sleep(500);
  }
  log(`logged in: ${page.url()}`);
}

async function main() {
  log(`launching chromium (headless=${cfg.headless}) budget=${cfg.budgetMs / 1000}s/page=${cfg.pageMs / 1000}s`);
  const browser = await chromium.launch({ headless: cfg.headless, args: ['--ignore-certificate-errors'] });
  const context = await browser.newContext({
    ignoreHTTPSErrors: true,
    recordHar: { path: harPath, mode: 'full', content: 'embed', urlFilter: /\/(api|proxy|v2|wss|status|guest)\b/ },
    viewport: { width: 1600, height: 1000 },
  });
  const page = await context.newPage();
  page.setDefaultTimeout(cfg.navTimeout);
  context.on('response', (r) => recordXhr(r.request().method(), r.request().url(), r.status(), r.request().resourceType()));
  page.on('dialog', (d) => d.accept().catch(() => {}));

  const visited = new Set();
  const queue = [];
  const enqueue = (href) => {
    let p;
    try { p = new URL(href, cfg.baseURL).pathname; } catch { return; }
    if (!p.startsWith('/manage') || visited.has(p) || queue.includes(p)) return;
    queue.push(p);
  };

  try {
    await login(page);
    seedRoutes(cfg.site).forEach(enqueue);
    (await discoverLinks(page)).forEach(enqueue);
    log(`seeded ${queue.length} routes`);

    while (queue.length && !overBudget()) {
      const path = queue.shift();
      if (visited.has(path)) continue;
      visited.add(path);
      const routeUrl = `${cfg.baseURL}${path}`;
      try {
        await page.goto(routeUrl, { waitUntil: 'domcontentloaded', timeout: cfg.navTimeout });
        await settle(page);
        await ensureLoggedIn(page, routeUrl);
        (await discoverLinks(page)).forEach(enqueue);
        const clicks = await exhaustPage(page, routeUrl);
        log(`(${visited.size}) ${path}  clicked ${clicks}  [endpoints ${endpoints.size}, queue ${queue.length}]`);
      } catch (e) {
        log(`  ${path} error: ${e.message?.split('\n')[0]}`);
      }
    }
    log(`deep crawl done: ${visited.size} routes, ${endpoints.size} endpoints${overBudget() ? ' (budget hit)' : ''}`);
  } finally {
    await context.close().catch(() => {});
    await browser.close().catch(() => {});
    const summary = [...endpoints.values()]
      .map((e) => ({ path: e.path, methods: [...e.methods].sort(), statuses: [...e.statuses].sort(), examples: e.examples }))
      .sort((a, b) => a.path.localeCompare(b.path));
    writeFileSync(summaryPath, JSON.stringify(summary, null, 2));
    writeFileSync(routesPath, JSON.stringify([...visited].sort(), null, 2));
    const m = {};
    summary.forEach((e) => e.methods.forEach((x) => (m[x] = (m[x] || 0) + 1)));
    log(`HAR -> ${harPath}`);
    log(`API -> ${summaryPath} (${summary.length} endpoints, methods ${JSON.stringify(m)})`);
  }
}

main().catch((e) => { console.error('[deep] fatal', e); process.exit(1); });
