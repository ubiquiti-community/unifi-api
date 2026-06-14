// setup.mjs
//
// Mirrors the readiness/bootstrap logic from
//   terraform-provider-unifi/unifi/provider_test.go :: waitForUniFiAPI
// but as a standalone, dependency-free Node poller.
//
// The controller is started by run.sh in *simulation mode* (see
// scripts/init.d/demo-mode), which seeds an admin/admin account plus demo
// sites, devices and networks. This script just waits until the API is fully
// answering, ensures a default WAN network exists, then exits 0 so the crawler
// can drive the already-initialised UI.
//
// Usage: node setup.mjs   (honours UNIFI_API / UNIFI_USERNAME / UNIFI_PASSWORD)

import { Agent } from 'node:https';

const BASE = (process.env.UNIFI_API || 'https://localhost:8443').replace(/\/$/, '');
const USER = process.env.UNIFI_USERNAME || 'admin';
const PASS = process.env.UNIFI_PASSWORD || 'admin';
const SITE = process.env.UNIFI_SITE || 'default';
const MAX_RETRIES = Number(process.env.UNIFI_MAX_RETRIES || 60);
const RETRY_DELAY_MS = Number(process.env.UNIFI_RETRY_DELAY_MS || 3000);

// Accept the controller's self-signed cert.
const agent = new Agent({ rejectUnauthorized: false, keepAlive: true });

const sleep = (ms) => new Promise((r) => setTimeout(r, ms));
const log = (...a) => console.log('[setup]', ...a);

// A tiny cookie jar so the session survives across calls.
let cookie = '';

async function req(path, { method = 'GET', body } = {}) {
  const res = await fetch(`${BASE}${path}`, {
    method,
    agent, // node-fetch style; node's global fetch ignores this, so we also set NODE_TLS_REJECT_UNAUTHORIZED in run.sh
    headers: {
      'content-type': 'application/json',
      ...(cookie ? { cookie } : {}),
    },
    body: body ? JSON.stringify(body) : undefined,
  });
  const setCookie = res.headers.getSetCookie?.() || [];
  if (setCookie.length) cookie = setCookie.map((c) => c.split(';')[0]).join('; ');
  let json = null;
  const text = await res.text();
  try {
    json = text ? JSON.parse(text) : null;
  } catch {
    /* non-JSON (HTML login page, etc.) */
  }
  return { status: res.status, json, text };
}

async function login() {
  // Classic controller login endpoint (matches go-unifi loginPath).
  const { status } = await req('/api/login', {
    method: 'POST',
    body: { username: USER, password: PASS, remember: true },
  });
  return status >= 200 && status < 300;
}

async function listSites() {
  const { json } = await req('/api/self/sites');
  return json?.data ?? [];
}

async function listNetworks() {
  const { json } = await req(`/api/s/${SITE}/rest/networkconf`);
  return json?.data ?? [];
}

async function ensureWan(networks) {
  const hasWan = networks.some(
    (n) => n.purpose === 'wan' && n.wan_networkgroup === 'WAN',
  );
  if (hasWan) return;
  log('No default WAN network found, creating "Internet 1"...');
  const { status } = await req(`/api/s/${SITE}/rest/networkconf`, {
    method: 'POST',
    body: {
      name: 'Internet 1',
      purpose: 'wan',
      wan_networkgroup: 'WAN',
      wan_type: 'dhcp',
    },
  });
  log(status < 300 ? '✓ Created default WAN network' : `WAN create returned ${status}`);
}

async function main() {
  log(`Waiting for UniFi API at ${BASE} (max ${MAX_RETRIES} attempts)...`);
  let loggedIn = false;
  for (let i = 1; i <= MAX_RETRIES; i++) {
    try {
      if (!(await login())) throw new Error('login not successful');
      if (!loggedIn) {
        log(`✓ Login successful after ${i} attempts`);
        loggedIn = true;
      }
      const sites = await listSites();
      if (!sites.length) throw new Error('no sites yet');
      const networks = await listNetworks();
      if (!networks.length) throw new Error('no networks yet');
      await ensureWan(networks);
      log(`✓ UniFi API fully ready (${sites.length} site(s), ${networks.length} network(s)) after ${i} attempts`);
      return;
    } catch (err) {
      if (i % 10 === 0 || i === 1) {
        log(`still waiting... (attempt ${i}/${MAX_RETRIES}): ${err.message}`);
      }
      if (i === MAX_RETRIES) {
        console.error(`[setup] UniFi API did not become ready after ${MAX_RETRIES} attempts`);
        process.exit(1);
      }
      await sleep(RETRY_DELAY_MS);
    }
  }
}

main().catch((e) => {
  console.error('[setup] fatal', e);
  process.exit(1);
});
