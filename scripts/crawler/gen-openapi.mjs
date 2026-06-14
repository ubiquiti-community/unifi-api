// gen-openapi.mjs
//
// Turn a captured HAR (every XHR the UI made) into an OpenAPI 3.0.3 spec.
//
//   node gen-openapi.mjs <har-file> [out-prefix]
//   node gen-openapi.mjs out-deep/unifi-ui.har out-deep/openapi
//
// What it does:
//   * keeps only the controller's JSON API calls (/api, /v2, /proxy)
//   * templatizes volatile path segments -> path params:
//       /api/s/default/rest/networkconf/65abc...   ->  /api/s/{site}/rest/networkconf/{id}
//   * groups by (templated path, method) -> one operation
//   * infers request/response JSON Schemas by merging every observed sample
//   * records observed query params and response status codes
//   * emits <prefix>.json and <prefix>.yaml
//
// No external dependencies; the YAML is emitted by a tiny local serializer.

import { readFileSync, writeFileSync } from 'node:fs';

// Usage: node gen-openapi.mjs <har...> [--out <prefix>]
// Multiple HARs are merged (union of all observed traffic). Default output
// prefix is derived from the first HAR's directory.
const argv = process.argv.slice(2);
const outIdx = argv.indexOf('--out');
const outPrefix = outIdx >= 0 ? argv[outIdx + 1] : null;
const harFiles = argv.filter((a, i) => a.endsWith('.har') && i !== outIdx + 1);
if (!harFiles.length) harFiles.push('out-deep/unifi-ui.har');
const finalPrefix = outPrefix || harFiles[0].replace(/\.har$/, '').replace(/[^/]+$/, 'openapi');

const entries = [];
for (const f of harFiles) {
  try {
    const har = JSON.parse(readFileSync(f, 'utf8'));
    const es = har.log?.entries ?? [];
    entries.push(...es);
    console.log(`[openapi] loaded ${es.length} entries from ${f}`);
  } catch (e) {
    console.warn(`[openapi] skip ${f}: ${e.message}`);
  }
}

// ---- path templatization -------------------------------------------------
const isObjectId = (s) => /^[0-9a-f]{24}$/i.test(s);
const isUuid = (s) => /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i.test(s);
const isMac = (s) => /^[0-9a-f]{2}([:-][0-9a-f]{2}){5}$/i.test(s);
const isNum = (s) => /^\d+$/.test(s);

function templatize(pathname) {
  const params = [];
  const segs = pathname.split('/').map((seg, i, arr) => {
    // /api/s/<site>/...  and  /v2/api/site/<site>/...  -> {site}
    if ((arr[i - 1] === 's' || arr[i - 1] === 'site') && seg && !isObjectId(seg) && !isNum(seg) && seg !== 'self') {
      if (seg === 'default' || /^[a-z0-9]{6,8}$/i.test(seg)) {
        params.push({ name: 'site', sample: seg });
        return '{site}';
      }
    }
    if (isObjectId(seg)) { params.push({ name: 'id', sample: seg }); return '{id}'; }
    if (isUuid(seg)) { params.push({ name: 'id', sample: seg }); return '{id}'; }
    if (isMac(seg)) { params.push({ name: 'mac', sample: seg }); return '{mac}'; }
    if (isNum(seg) && i > 2) { params.push({ name: 'id', sample: seg }); return '{id}'; }
    return seg;
  });
  // de-dup repeated param names: id, id2, ...
  const seen = {};
  const finalParams = params.map((p) => {
    seen[p.name] = (seen[p.name] || 0) + 1;
    const name = seen[p.name] === 1 ? p.name : `${p.name}${seen[p.name]}`;
    return { ...p, name };
  });
  let pi = 0;
  const path = segs.map((s) => (/^\{/.test(s) ? `{${finalParams[pi++].name}}` : s)).join('/');
  return { path, params: finalParams };
}

// ---- JSON Schema inference + merge ---------------------------------------
function infer(v) {
  if (v === null) return { nullable: true };
  if (Array.isArray(v)) return { type: 'array', items: v.length ? infer(v[0]) : {} };
  switch (typeof v) {
    case 'string': return { type: 'string', example: v.length > 60 ? undefined : v };
    case 'number': return { type: Number.isInteger(v) ? 'integer' : 'number' };
    case 'boolean': return { type: 'boolean' };
    case 'object': {
      const props = {};
      for (const k of Object.keys(v)) props[k] = infer(v[k]);
      return { type: 'object', properties: props };
    }
    default: return {};
  }
}
function merge(a, b) {
  if (!a) return b;
  if (!b) return a;
  if (a.nullable && !b.nullable) b.nullable = true;
  if (b.nullable && a.type) { a.nullable = true; return mergeTyped(a, b); }
  if (!a.type) return b;
  if (!b.type) return a;
  return mergeTyped(a, b);
}
function mergeTyped(a, b) {
  if (a.type !== b.type) {
    // widen integer<->number; otherwise fall back to untyped
    if ((a.type === 'integer' && b.type === 'number') || (a.type === 'number' && b.type === 'integer')) {
      return { type: 'number', nullable: a.nullable || b.nullable };
    }
    return { nullable: a.nullable || b.nullable };
  }
  const out = { type: a.type, nullable: a.nullable || b.nullable };
  if (a.type === 'object') {
    out.properties = { ...(a.properties || {}) };
    for (const k of Object.keys(b.properties || {})) out.properties[k] = merge(out.properties[k], b.properties[k]);
  } else if (a.type === 'array') {
    out.items = merge(a.items || {}, b.items || {});
  } else if (a.example !== undefined) {
    out.example = a.example;
  } else if (b.example !== undefined) {
    out.example = b.example;
  }
  return out;
}
const clean = (s) => {
  if (!s || typeof s !== 'object') return s;
  if (s.nullable === false) delete s.nullable;
  if (s.properties) for (const k of Object.keys(s.properties)) clean(s.properties[k]);
  if (s.items) clean(s.items);
  return s;
};

// ---- walk the HAR --------------------------------------------------------
const ops = new Map(); // `${method} ${path}` -> op aggregate
let kept = 0;

for (const e of entries) {
  const req = e.request, res = e.response;
  if (!req) continue;
  let u;
  try { u = new URL(req.url); } catch { continue; }
  if (!/^\/(api|v2|proxy)\b/.test(u.pathname)) continue;
  const method = req.method.toUpperCase();
  const { path, params } = templatize(u.pathname);
  const opKey = `${method} ${path}`;
  if (!ops.has(opKey)) {
    ops.set(opKey, { method, path, params: {}, query: {}, reqSchema: null, responses: {}, count: 0 });
  }
  const op = ops.get(opKey);
  op.count++;
  kept++;
  for (const p of params) op.params[p.name] = p.sample;
  for (const q of req.queryString || []) op.query[q.name] = q.value;

  // request body
  const reqText = req.postData?.text;
  if (reqText) {
    try { op.reqSchema = merge(op.reqSchema, infer(JSON.parse(reqText))); } catch {}
  }
  // response body by status
  const status = String(res?.status || 0);
  const resText = res?.content?.text;
  let resSchema = null;
  if (resText) {
    let body = resText;
    if (res.content.encoding === 'base64') { try { body = Buffer.from(resText, 'base64').toString('utf8'); } catch {} }
    try { resSchema = infer(JSON.parse(body)); } catch {}
  }
  if (!op.responses[status]) op.responses[status] = null;
  op.responses[status] = merge(op.responses[status], resSchema);
}

// ---- assemble OpenAPI ----------------------------------------------------
const paths = {};
for (const op of [...ops.values()].sort((a, b) => (a.path + a.method).localeCompare(b.path + b.method))) {
  paths[op.path] = paths[op.path] || {};
  const parameters = [];
  for (const name of Object.keys(op.params)) {
    parameters.push({
      name, in: 'path', required: true,
      schema: { type: name === 'id' || name === 'site' || name === 'mac' ? 'string' : 'string' },
      example: op.params[name],
    });
  }
  for (const name of Object.keys(op.query)) {
    parameters.push({ name, in: 'query', required: false, schema: { type: 'string' }, example: op.query[name] });
  }

  const operation = {
    summary: `${op.method} ${op.path}`,
    operationId: `${op.method.toLowerCase()}${op.path.replace(/[^a-zA-Z0-9]+/g, '_')}`,
    description: `Observed ${op.count}x during UI crawl.`,
    ...(parameters.length ? { parameters } : {}),
    responses: {},
  };
  if (['POST', 'PUT', 'PATCH'].includes(op.method) && op.reqSchema) {
    operation.requestBody = { content: { 'application/json': { schema: clean(op.reqSchema) } } };
  }
  const statuses = Object.keys(op.responses).filter((s) => s !== '0');
  if (!statuses.length) statuses.push('200');
  for (const s of statuses) {
    const sch = op.responses[s];
    operation.responses[s] = {
      description: s.startsWith('2') ? 'Success' : s.startsWith('4') ? 'Client error / validation' : 'Response',
      ...(sch ? { content: { 'application/json': { schema: clean(sch) } } } : {}),
    };
  }
  paths[op.path][op.method.toLowerCase()] = operation;
}

const spec = {
  openapi: '3.0.3',
  info: {
    title: 'UniFi Network Application API (reverse-engineered)',
    version: '0.1.0',
    description:
      'Auto-generated from a HAR capture of the UniFi Network Application UI. ' +
      `Built from ${kept} XHR calls across ${ops.size} operations. Schemas are inferred from observed traffic.`,
  },
  servers: [{ url: 'https://localhost:8443', description: 'UniFi controller (simulation mode)' }],
  paths,
};

// ---- emit JSON + YAML ----------------------------------------------------
writeFileSync(`${finalPrefix}.json`, JSON.stringify(spec, null, 2));
try {
  writeFileSync(`${finalPrefix}.yaml`, toYaml(spec));
} catch (e) {
  console.warn(`[openapi] YAML emit skipped (${e.message}); JSON spec is authoritative`);
}

const byMethod = {};
for (const op of ops.values()) byMethod[op.method] = (byMethod[op.method] || 0) + 1;
console.log(`[openapi] ${kept} XHR -> ${ops.size} operations`, byMethod);
console.log(`[openapi] wrote ${finalPrefix}.json and ${finalPrefix}.yaml`);

// ---- minimal YAML serializer --------------------------------------------
function toYaml(obj, indent = 0) {
  const pad = '  '.repeat(indent);
  if (indent > 80) return JSON.stringify(obj); // depth guard: inline very deep nodes
  if (obj === null || obj === undefined) return 'null';
  if (typeof obj === 'string') return needsQuote(obj) ? JSON.stringify(obj) : obj;
  if (typeof obj === 'number' || typeof obj === 'boolean') return String(obj);
  if (Array.isArray(obj)) {
    if (!obj.length) return '[]';
    return obj.map((v) => {
      if (v && typeof v === 'object' && Object.keys(v).length) {
        // render the object's lines, then hang the first under "- "
        const lines = toYaml(v, indent + 1).split('\n');
        const itemPad = '  '.repeat(indent);
        lines[0] = `${itemPad}- ${lines[0].slice((indent + 1) * 2)}`;
        return lines.join('\n');
      }
      return `${pad}- ${toYaml(v, indent + 1)}`;
    }).join('\n');
  }
  const keys = Object.keys(obj);
  if (!keys.length) return '{}';
  return keys.map((k) => {
    const v = obj[k];
    const key = needsQuote(k) ? JSON.stringify(k) : k;
    if (v && typeof v === 'object' && (Array.isArray(v) ? v.length : Object.keys(v).length)) {
      return `${pad}${key}:\n${toYaml(v, indent + 1)}`;
    }
    return `${pad}${key}: ${toYaml(v, indent + 1)}`;
  }).join('\n');
}
function needsQuote(s) {
  return typeof s === 'string' && (/[:#{}\[\],&*?|<>=!%@`"']/.test(s) || /^\s|\s$/.test(s) || /^(true|false|null|yes|no|\d)/i.test(s) || s === '');
}
