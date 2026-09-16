"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FEATURE_PLUGINS = exports.config = void 0;
const RatelimitFeature_1 = require("./feature/ratelimit/RatelimitFeature");
const RetryFeature_1 = require("./feature/retry/RetryFeature");
const TestFeature_1 = require("./feature/test/TestFeature");
const TimeoutFeature_1 = require("./feature/timeout/TimeoutFeature");
const FEATURE_CLASS = {
    ratelimit: RatelimitFeature_1.RatelimitFeature,
    retry: RetryFeature_1.RetryFeature,
    test: TestFeature_1.TestFeature,
    timeout: TimeoutFeature_1.TimeoutFeature,
};
// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS = {};
exports.FEATURE_PLUGINS = FEATURE_PLUGINS;
class Config {
    makeFeature(fn) {
        const fc = FEATURE_CLASS[fn];
        const fi = new fc();
        // TODO: errors etc
        return fi;
    }
    // False for a feature added at runtime via options.extend (station's
    // adopt path) - the constructor uses this to skip makeFeature for names
    // no generated class backs.
    hasFeature(fn) {
        return null != FEATURE_CLASS[fn];
    }
    main = {
        name: 'DutchCustomerData',
        slug: "dutch-customer-data",
        version: "0.0.1",
        target: "ts",
    };
    feature = {
        ratelimit: {
            "options": {
                "active": false,
                "burst": 5,
                "rate": 5
            },
            "optspec": {
                "now": "`$FUNCTION`",
                "sleep": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "wrap"
        },
        retry: {
            "options": {
                "active": false,
                "factor": 2,
                "maxDelay": 2000,
                "minDelay": 50,
                "retries": 2,
                "statuses": [
                    408,
                    425,
                    429,
                    500,
                    502,
                    503,
                    504
                ]
            },
            "optspec": {
                "jitter": "`$BOOLEAN`",
                "sleep": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "wrap"
        },
        test: {
            "options": {
                "active": false
            },
            "optspec": {
                "entity": "`$MAP`",
                "net": "`$MAP`"
            },
            "strict": false,
            "transport": "base"
        },
        timeout: {
            "options": {
                "active": false,
                "ms": 30000
            },
            "optspec": {
                "clearTimer": "`$FUNCTION`",
                "setTimer": "`$FUNCTION`"
            },
            "strict": false,
            "transport": "wrap"
        },
    };
    options = {
        base: "https://free.bedrijfsdata.nl/v1.1",
        headers: {
            "content-type": "application/json"
        },
        entity: {
            eu_ap_i: {},
            global_ap_i: {},
            netherlands_ap_i: {},
        }
    };
    entity = {
        "eu_ap_i": {
            "fields": [
                {
                    "name": "active",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "address",
                    "type": "`$STRING`"
                },
                {
                    "name": "buyer",
                    "type": "`$STRING`"
                },
                {
                    "name": "buyer_country",
                    "type": "`$STRING`"
                },
                {
                    "name": "city",
                    "type": "`$STRING`"
                },
                {
                    "name": "contract_nature",
                    "type": "`$STRING`"
                },
                {
                    "name": "country",
                    "type": "`$STRING`"
                },
                {
                    "format": "uri",
                    "name": "html",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "format": "uri",
                    "name": "link",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "type": "`$STRING`"
                },
                {
                    "name": "notice_type",
                    "type": "`$STRING`"
                },
                {
                    "name": "official_language",
                    "type": "`$STRING`"
                },
                {
                    "format": "uri",
                    "name": "pdf",
                    "type": "`$STRING`"
                },
                {
                    "name": "place_of_performance",
                    "type": "`$STRING`"
                },
                {
                    "name": "postcode",
                    "type": "`$STRING`"
                },
                {
                    "name": "procedure_type",
                    "type": "`$STRING`"
                },
                {
                    "format": "date",
                    "name": "publication_date",
                    "type": "`$STRING`"
                },
                {
                    "format": "date-time",
                    "name": "response_date",
                    "type": "`$STRING`"
                },
                {
                    "name": "title",
                    "type": "`$STRING`"
                },
                {
                    "name": "vat",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "eu_ap_i",
            "op": {
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "37080091",
                                        "kind": "query",
                                        "name": "q",
                                        "orig": "q",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/tender",
                            "segments": [
                                {
                                    "lit": "tender"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "q"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.tender`"
                            },
                            "parts": [
                                "tender"
                            ]
                        }
                    ]
                },
                "load": {
                    "input": "data",
                    "name": "load",
                    "points": [
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "NL001672022B01",
                                        "kind": "query",
                                        "name": "vat",
                                        "orig": "vat",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/vat",
                            "segments": [
                                {
                                    "lit": "vat"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "vat"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.vat`"
                            },
                            "parts": [
                                "vat"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "global_ap_i": {
            "fields": [
                {
                    "name": "addition",
                    "type": "`$STRING`"
                },
                {
                    "name": "address",
                    "type": "`$STRING`"
                },
                {
                    "name": "admin1",
                    "type": "`$STRING`"
                },
                {
                    "name": "admin2",
                    "type": "`$STRING`"
                },
                {
                    "name": "admin3",
                    "type": "`$STRING`"
                },
                {
                    "name": "bank",
                    "type": "`$STRING`"
                },
                {
                    "name": "bic",
                    "type": "`$STRING`"
                },
                {
                    "name": "browser",
                    "type": "`$STRING`"
                },
                {
                    "name": "builtwith",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "carrier",
                    "short": "Carrier name",
                    "type": "`$STRING`"
                },
                {
                    "name": "city",
                    "type": "`$STRING`"
                },
                {
                    "name": "cloudflare",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "commoncrawl",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "content_length",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "content_type",
                    "type": "`$STRING`"
                },
                {
                    "name": "country",
                    "short": "ISO country code",
                    "type": "`$STRING`"
                },
                {
                    "name": "country_code",
                    "type": "`$STRING`"
                },
                {
                    "name": "crux",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "device_family",
                    "type": "`$STRING`"
                },
                {
                    "name": "device_name",
                    "type": "`$STRING`"
                },
                {
                    "name": "device_type",
                    "type": "`$STRING`"
                },
                {
                    "name": "disposable",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "dns_a",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "dns_mx",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "dns_ns",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "dns_soa",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "dns_txt",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "dns_www_a",
                    "type": "`$ARRAY`"
                },
                {
                    "name": "dnsserver",
                    "type": "`$STRING`"
                },
                {
                    "name": "domain",
                    "type": "`$STRING`"
                },
                {
                    "name": "domcop",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "email",
                    "type": "`$STRING`"
                },
                {
                    "name": "found",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "free",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "freeformaddress",
                    "type": "`$STRING`"
                },
                {
                    "name": "host",
                    "type": "`$STRING`"
                },
                {
                    "name": "host_type",
                    "type": "`$STRING`"
                },
                {
                    "name": "hostio",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "http_code",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "iban",
                    "type": "`$STRING`"
                },
                {
                    "name": "iban_human",
                    "type": "`$STRING`"
                },
                {
                    "name": "int",
                    "short": "International format without plus sign",
                    "type": "`$STRING`"
                },
                {
                    "name": "international",
                    "short": "International formatted phone number",
                    "type": "`$STRING`"
                },
                {
                    "name": "ip",
                    "type": "`$STRING`"
                },
                {
                    "name": "ipint",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "ismobile",
                    "short": "1 if mobile, 0 if not",
                    "type": "`$INTEGER`"
                },
                {
                    "format": "double",
                    "name": "lat",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "lei",
                    "type": "`$STRING`"
                },
                {
                    "name": "letter",
                    "type": "`$STRING`"
                },
                {
                    "name": "local_id",
                    "type": "`$STRING`"
                },
                {
                    "format": "double",
                    "name": "lon",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "mailserver",
                    "type": "`$STRING`"
                },
                {
                    "name": "majestic",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "message",
                    "type": "`$STRING`"
                },
                {
                    "name": "municipality",
                    "type": "`$STRING`"
                },
                {
                    "name": "mx_host",
                    "type": "`$STRING`"
                },
                {
                    "name": "mx_ip",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "type": "`$STRING`"
                },
                {
                    "name": "national",
                    "short": "National formatted phone number",
                    "type": "`$STRING`"
                },
                {
                    "name": "number",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "ocid",
                    "type": "`$STRING`"
                },
                {
                    "name": "pagerank",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "platform",
                    "type": "`$STRING`"
                },
                {
                    "name": "population",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "postcode",
                    "type": "`$STRING`"
                },
                {
                    "name": "province",
                    "type": "`$STRING`"
                },
                {
                    "name": "province_code",
                    "type": "`$STRING`"
                },
                {
                    "name": "redirect_count",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "region",
                    "short": "Geographic region",
                    "type": "`$STRING`"
                },
                {
                    "name": "register_id",
                    "type": "`$STRING`"
                },
                {
                    "format": "date-time",
                    "name": "renewal_date",
                    "type": "`$STRING`"
                },
                {
                    "format": "double",
                    "name": "score",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "sepa",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "spf",
                    "type": "`$STRING`"
                },
                {
                    "name": "status",
                    "type": "`$STRING`"
                },
                {
                    "name": "street",
                    "type": "`$STRING`"
                },
                {
                    "name": "success",
                    "short": "1 if successful, 0 if not",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "swift",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "tranco",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "type",
                    "type": "`$STRING`"
                },
                {
                    "name": "umbrella",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "url",
                    "type": "`$STRING`"
                },
                {
                    "name": "user",
                    "type": "`$STRING`"
                },
                {
                    "name": "user_agent",
                    "type": "`$STRING`"
                },
                {
                    "name": "valid",
                    "short": "1 if valid, 0 if not",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "verified",
                    "type": "`$BOOLEAN`"
                },
                {
                    "name": "verified_checksum",
                    "type": "`$BOOLEAN`"
                },
                {
                    "name": "webrank",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "wrong_email",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "wrong_format",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "wrong_password",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "wrong_phone",
                    "short": "1 if wrong, 0 if correct",
                    "type": "`$INTEGER`"
                }
            ],
            "name": "global_ap_i",
            "op": {
                "create": {
                    "input": "data",
                    "name": "create",
                    "points": [
                        {
                            "args": {},
                            "kind": "http",
                            "method": "POST",
                            "orig": "/password",
                            "segments": [
                                {
                                    "lit": "password"
                                }
                            ],
                            "select": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.password`"
                            },
                            "parts": [
                                "password"
                            ]
                        }
                    ]
                },
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "Haarlem",
                                        "kind": "query",
                                        "name": "city",
                                        "orig": "city",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "NL",
                                        "kind": "query",
                                        "name": "country_code",
                                        "orig": "country_code",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "kind": "query",
                                        "name": "full",
                                        "orig": "full",
                                        "type": "`$INTEGER`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/city",
                            "segments": [
                                {
                                    "lit": "city"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "city",
                                    "country_code",
                                    "full"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.city`"
                            },
                            "parts": [
                                "city"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "NL",
                                        "kind": "query",
                                        "name": "country_code",
                                        "orig": "country_code",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "1013PN",
                                        "kind": "query",
                                        "name": "postcode",
                                        "orig": "postcode",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/postcode",
                            "segments": [
                                {
                                    "lit": "postcode"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "country_code",
                                    "postcode"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.postcode`"
                            },
                            "parts": [
                                "postcode"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "nl",
                                        "kind": "query",
                                        "name": "country_code",
                                        "orig": "country_code",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "Kalverstraat 1, 1012NX",
                                        "kind": "query",
                                        "name": "q",
                                        "orig": "q",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/geocoding",
                            "segments": [
                                {
                                    "lit": "geocoding"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "country_code",
                                    "q"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.geocoding`"
                            },
                            "parts": [
                                "geocoding"
                            ]
                        }
                    ]
                },
                "load": {
                    "input": "data",
                    "name": "load",
                    "points": [
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "AKZONL2AXXX",
                                        "kind": "query",
                                        "name": "bic",
                                        "orig": "bic",
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "724500XYIJUGXAA5QD70",
                                        "kind": "query",
                                        "name": "lei",
                                        "orig": "lei",
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "09007809",
                                        "kind": "query",
                                        "name": "local_id",
                                        "orig": "local_id",
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/lei",
                            "segments": [
                                {
                                    "lit": "lei"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "bic",
                                    "lei",
                                    "local_id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.lei`"
                            },
                            "parts": [
                                "lei"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "kind": "query",
                                        "name": "password",
                                        "orig": "password",
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "2c4c3891e2ac6958e9810a1e49c6705784fbfa1a",
                                        "kind": "query",
                                        "name": "password_sha1",
                                        "orig": "password_sha1",
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": 25,
                                        "kind": "query",
                                        "name": "threshold",
                                        "orig": "threshold",
                                        "type": "`$INTEGER`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/password",
                            "segments": [
                                {
                                    "lit": "password"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "password",
                                    "password_sha1",
                                    "threshold"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.password`"
                            },
                            "parts": [
                                "password"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "nl",
                                        "kind": "query",
                                        "name": "country_code",
                                        "orig": "country_code",
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "+31207895050",
                                        "kind": "query",
                                        "name": "phone",
                                        "orig": "phone",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/phone",
                            "segments": [
                                {
                                    "lit": "phone"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "country_code",
                                    "phone"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.phone`"
                            },
                            "parts": [
                                "phone"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "AKZONL2A",
                                        "kind": "query",
                                        "name": "bic",
                                        "orig": "bic",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/bic",
                            "segments": [
                                {
                                    "lit": "bic"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "bic"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.bic`"
                            },
                            "parts": [
                                "bic"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "eur",
                                        "kind": "query",
                                        "name": "currency",
                                        "orig": "currency",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/currency",
                            "segments": [
                                {
                                    "lit": "currency"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "currency"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.currency`"
                            },
                            "parts": [
                                "currency"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "bedrijfsdata.nl",
                                        "kind": "query",
                                        "name": "domain",
                                        "orig": "domain",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/dns",
                            "segments": [
                                {
                                    "lit": "dns"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "domain"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.dns`"
                            },
                            "parts": [
                                "dns"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "google.com",
                                        "kind": "query",
                                        "name": "domain",
                                        "orig": "domain",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/webrank",
                            "segments": [
                                {
                                    "lit": "webrank"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "domain"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.webrank`"
                            },
                            "parts": [
                                "webrank"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "piet@bedrijfsdata.nl",
                                        "kind": "query",
                                        "name": "email",
                                        "orig": "email",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/email",
                            "segments": [
                                {
                                    "lit": "email"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "email"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.email`"
                            },
                            "parts": [
                                "email"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "NL17ADYB2017400505",
                                        "kind": "query",
                                        "name": "iban",
                                        "orig": "iban",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/iban",
                            "segments": [
                                {
                                    "lit": "iban"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "iban"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.iban`"
                            },
                            "parts": [
                                "iban"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "kind": "query",
                                        "name": "ua",
                                        "orig": "ua",
                                        "type": "`$INTEGER`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/ip",
                            "segments": [
                                {
                                    "lit": "ip"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "ua"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.ip`"
                            },
                            "parts": [
                                "ip"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "http://www.bedrijfsdata.nl",
                                        "kind": "query",
                                        "name": "url",
                                        "orig": "url",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/url",
                            "segments": [
                                {
                                    "lit": "url"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "url"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.url`"
                            },
                            "parts": [
                                "url"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "netherlands_ap_i": {
            "fields": [
                {
                    "name": "active",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "addition",
                    "type": "`$STRING`"
                },
                {
                    "name": "city",
                    "type": "`$STRING`"
                },
                {
                    "name": "coc",
                    "type": "`$STRING`"
                },
                {
                    "name": "construction_year",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "floor_area",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "freeformaddress",
                    "type": "`$STRING`"
                },
                {
                    "name": "id",
                    "type": "`$STRING`"
                },
                {
                    "format": "double",
                    "name": "lat",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "letter",
                    "type": "`$STRING`"
                },
                {
                    "format": "double",
                    "name": "lon",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "municipality",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "type": "`$STRING`"
                },
                {
                    "name": "number",
                    "type": "`$STRING`"
                },
                {
                    "name": "postcode",
                    "type": "`$STRING`"
                },
                {
                    "name": "province",
                    "type": "`$STRING`"
                },
                {
                    "name": "province_code",
                    "type": "`$STRING`"
                },
                {
                    "name": "purpose",
                    "type": "`$STRING`"
                },
                {
                    "name": "street",
                    "type": "`$STRING`"
                },
                {
                    "name": "type",
                    "type": "`$STRING`"
                },
                {
                    "name": "vestiging",
                    "type": "`$STRING`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "netherlands_ap_i",
            "op": {
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "3",
                                        "kind": "query",
                                        "name": "number",
                                        "orig": "number",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "1011PN",
                                        "kind": "query",
                                        "name": "postcode",
                                        "orig": "postcode",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": "a",
                                        "kind": "query",
                                        "name": "suffix",
                                        "orig": "suffix",
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/bag",
                            "segments": [
                                {
                                    "lit": "bag"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "number",
                                    "postcode",
                                    "suffix"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.bag`"
                            },
                            "parts": [
                                "bag"
                            ]
                        },
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": "89395808",
                                        "kind": "query",
                                        "name": "kvk",
                                        "orig": "kvk",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/kvk",
                            "segments": [
                                {
                                    "lit": "kvk"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "kvk"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body.kvk`"
                            },
                            "parts": [
                                "kvk"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        }
    };
}
const config = new Config();
exports.config = config;
//# sourceMappingURL=Config.js.map