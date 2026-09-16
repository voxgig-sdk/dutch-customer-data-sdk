"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('EuApIEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when DUTCH_CUSTOMER_DATA_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('DUTCH_CUSTOMER_DATA_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.DutchCustomerDataSDK.test();
        const ent = testsdk.EuApI();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.DUTCH_CUSTOMER_DATA_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'eu_ap_i.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "active", "req": false, "type": "`$INTEGER`", "index$": 0 }, { "active": true, "name": "address", "req": false, "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "buyer", "req": false, "type": "`$STRING`", "index$": 2 }, { "active": true, "name": "buyer_country", "req": false, "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "city", "req": false, "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "contract_nature", "req": false, "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "country", "req": false, "type": "`$STRING`", "index$": 6 }, { "active": true, "format": "uri", "name": "html", "req": false, "type": "`$STRING`", "index$": 7 }, { "active": true, "name": "id", "req": false, "type": "`$STRING`", "index$": 8 }, { "active": true, "format": "uri", "name": "link", "req": false, "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "name", "req": false, "type": "`$STRING`", "index$": 10 }, { "active": true, "name": "notice_type", "req": false, "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "official_language", "req": false, "type": "`$STRING`", "index$": 12 }, { "active": true, "format": "uri", "name": "pdf", "req": false, "type": "`$STRING`", "index$": 13 }, { "active": true, "name": "place_of_performance", "req": false, "type": "`$STRING`", "index$": 14 }, { "active": true, "name": "postcode", "req": false, "type": "`$STRING`", "index$": 15 }, { "active": true, "name": "procedure_type", "req": false, "type": "`$STRING`", "index$": 16 }, { "active": true, "format": "date", "name": "publication_date", "req": false, "type": "`$STRING`", "index$": 17 }, { "active": true, "format": "date-time", "name": "response_date", "req": false, "type": "`$STRING`", "index$": 18 }, { "active": true, "name": "title", "req": false, "type": "`$STRING`", "index$": 19 }, { "active": true, "name": "vat", "req": false, "type": "`$STRING`", "index$": 20 }], "id": { "field": "id", "name": "id" }, "name": "eu_ap_i", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": { "query": [{ "active": true, "example": "37080091", "kind": "query", "name": "q", "orig": "q", "reqd": true, "type": "`$STRING`", "index$": 0 }] }, "contract": { "id": "GET /tender", "json": "{\"operationId\":\"searchTenders\",\"parameters\":[{\"description\":\"Search query (name, KvK number, etc.)\",\"in\":\"query\",\"name\":\"q\",\"required\":true,\"schema\":{\"example\":\"37080091\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"found\":{\"type\":\"integer\"},\"status\":{\"type\":\"string\"},\"tender\":{\"items\":{\"properties\":{\"buyer\":{\"type\":\"string\"},\"buyer_country\":{\"type\":\"string\"},\"contract_nature\":{\"type\":\"string\"},\"html\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"type\":\"string\"},\"link\":{\"format\":\"uri\",\"type\":\"string\"},\"notice_type\":{\"type\":\"string\"},\"official_language\":{\"type\":\"string\"},\"pdf\":{\"format\":\"uri\",\"type\":\"string\"},\"place_of_performance\":{\"type\":\"string\"},\"procedure_type\":{\"type\":\"string\"},\"publication_date\":{\"format\":\"date\",\"type\":\"string\"},\"title\":{\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"Successful response\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/tender", "segments": [{ "lit": "tender" }], "select": { "exist": ["q"] }, "transform": { "req": "`reqdata`", "res": "`body.tender`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "query": [{ "active": true, "example": "NL001672022B01", "kind": "query", "name": "vat", "orig": "vat", "reqd": true, "type": "`$STRING`", "index$": 0 }] }, "contract": { "id": "GET /vat", "json": "{\"operationId\":\"verifyVat\",\"parameters\":[{\"description\":\"EU VAT number to verify\",\"in\":\"query\",\"name\":\"vat\",\"required\":true,\"schema\":{\"example\":\"NL001672022B01\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"status\":{\"type\":\"string\"},\"vat\":{\"properties\":{\"active\":{\"type\":\"integer\"},\"address\":{\"type\":\"string\"},\"city\":{\"type\":\"string\"},\"country\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"postcode\":{\"type\":\"string\"},\"response_date\":{\"format\":\"date-time\",\"type\":\"string\"},\"vat\":{\"type\":\"string\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"description\":\"Successful response\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/vat", "segments": [{ "lit": "vat" }], "select": { "exist": ["vat"] }, "transform": { "req": "`reqdata`", "res": "`body.vat`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "eu_ap_i", "name__orig": "eu_ap_i", "Name": "EuApI", "name_": "eu_ap_i", "name-": "eu-ap-i", "NAME": "EU_AP_I", "index$": 0 }, { "active": true, "entity": "eu_ap_i", "key$": "BasicEuApIFlow", "kind": "basic", "name": "BasicEuApIFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "eu_ap_i_ref01" } }], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "eu_ap_i_ref01", "srcdatavar": "eu_ap_i_ref01_data", "suffix": "_dt0" }, "match": {}, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-eu_ap_i_ref01" } }], "index$": 1 }] }, 'EuApI');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let eu_ap_i_ref01_data = Object.values(setup.data.existing.eu_ap_i)[0];
        // LIST
        const eu_ap_i_ref01_ent = client.EuApI();
        const eu_ap_i_ref01_match = {};
        const eu_ap_i_ref01_list = (await eu_ap_i_ref01_ent.list(eu_ap_i_ref01_match)).map((e) => e.data());
        // LOAD
        const eu_ap_i_ref01_match_dt0 = {};
        eu_ap_i_ref01_match_dt0.id = eu_ap_i_ref01_data.id;
        const eu_ap_i_ref01_data_dt0 = (await eu_ap_i_ref01_ent.load(eu_ap_i_ref01_match_dt0)).data();
        (0, node_assert_1.default)(eu_ap_i_ref01_data_dt0.id === eu_ap_i_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/eu_ap_i/EuApITestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.DutchCustomerDataSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['eu_ap_i01', 'eu_ap_i02', 'eu_ap_i03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'DUTCH_CUSTOMER_DATA_TEST_EU_AP_I_ENTID': idmap,
        'DUTCH_CUSTOMER_DATA_TEST_LIVE': 'FALSE',
        'DUTCH_CUSTOMER_DATA_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['DUTCH_CUSTOMER_DATA_TEST_EU_AP_I_ENTID'];
    const live = 'TRUE' === env.DUTCH_CUSTOMER_DATA_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['DUTCH_CUSTOMER_DATA_TEST_EU_AP_I_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.DutchCustomerDataSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {},
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.DUTCH_CUSTOMER_DATA_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=EuApIEntity.test.js.map