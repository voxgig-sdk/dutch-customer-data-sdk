# DutchCustomerData SDK configuration


def make_config():
    return {
        "main": {
            "name": "DutchCustomerData",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://free.bedrijfsdata.nl/v1.1",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "eu_ap_i": {},
                "global_ap_i": {},
                "netherlands_ap_i": {},
            },
        },
        "entity": {
      "eu_ap_i": {
        "fields": [
          {
            "active": True,
            "name": "buyer",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "buyer_country",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "contract_nature",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "html",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "link",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "notice_type",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "official_language",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "pdf",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "place_of_performance",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "procedure_type",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "publication_date",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "title",
            "req": False,
            "type": "`$STRING`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "vat",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 14,
          },
        ],
        "name": "eu_ap_i",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "37080091",
                      "kind": "query",
                      "name": "q",
                      "orig": "q",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/tender",
                "parts": [
                  "tender",
                ],
                "select": {
                  "exist": [
                    "q",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "NL001672022B01",
                      "kind": "query",
                      "name": "vat",
                      "orig": "vat",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/vat",
                "parts": [
                  "vat",
                ],
                "select": {
                  "exist": [
                    "vat",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "global_ap_i": {
        "fields": [
          {
            "active": True,
            "name": "addition",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "admin1",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "admin2",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "admin3",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "bic",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "city",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "currency",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "date",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "dns",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "email",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "found",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "freeformaddress",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "from_currency",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "iban",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "ip",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "lat",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "lei",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "letter",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "lon",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "municipality",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "number",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "password",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "phone",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "population",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "postcode",
            "req": False,
            "type": "`$STRING`",
            "index$": 24,
          },
          {
            "active": True,
            "name": "province",
            "req": False,
            "type": "`$STRING`",
            "index$": 25,
          },
          {
            "active": True,
            "name": "province_code",
            "req": False,
            "type": "`$STRING`",
            "index$": 26,
          },
          {
            "active": True,
            "name": "score",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 27,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 28,
          },
          {
            "active": True,
            "name": "street",
            "req": False,
            "type": "`$STRING`",
            "index$": 29,
          },
          {
            "active": True,
            "name": "type",
            "req": False,
            "type": "`$STRING`",
            "index$": 30,
          },
          {
            "active": True,
            "name": "url",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 31,
          },
          {
            "active": True,
            "name": "webrank",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 32,
          },
        ],
        "name": "global_ap_i",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {},
                "method": "POST",
                "orig": "/password",
                "parts": [
                  "password",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "Haarlem",
                      "kind": "query",
                      "name": "city",
                      "orig": "city",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "NL",
                      "kind": "query",
                      "name": "country_code",
                      "orig": "country_code",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "full",
                      "orig": "full",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/city",
                "parts": [
                  "city",
                ],
                "select": {
                  "exist": [
                    "city",
                    "country_code",
                    "full",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "NL",
                      "kind": "query",
                      "name": "country_code",
                      "orig": "country_code",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "1013PN",
                      "kind": "query",
                      "name": "postcode",
                      "orig": "postcode",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/postcode",
                "parts": [
                  "postcode",
                ],
                "select": {
                  "exist": [
                    "country_code",
                    "postcode",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "nl",
                      "kind": "query",
                      "name": "country_code",
                      "orig": "country_code",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "Kalverstraat 1, 1012NX",
                      "kind": "query",
                      "name": "q",
                      "orig": "q",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/geocoding",
                "parts": [
                  "geocoding",
                ],
                "select": {
                  "exist": [
                    "country_code",
                    "q",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "AKZONL2AXXX",
                      "kind": "query",
                      "name": "bic",
                      "orig": "bic",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "724500XYIJUGXAA5QD70",
                      "kind": "query",
                      "name": "lei",
                      "orig": "lei",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "09007809",
                      "kind": "query",
                      "name": "local_id",
                      "orig": "local_id",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/lei",
                "parts": [
                  "lei",
                ],
                "select": {
                  "exist": [
                    "bic",
                    "lei",
                    "local_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "password",
                      "orig": "password",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "2c4c3891e2ac6958e9810a1e49c6705784fbfa1a",
                      "kind": "query",
                      "name": "password_sha1",
                      "orig": "password_sha1",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": 25,
                      "kind": "query",
                      "name": "threshold",
                      "orig": "threshold",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/password",
                "parts": [
                  "password",
                ],
                "select": {
                  "exist": [
                    "password",
                    "password_sha1",
                    "threshold",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "nl",
                      "kind": "query",
                      "name": "country_code",
                      "orig": "country_code",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "+31207895050",
                      "kind": "query",
                      "name": "phone",
                      "orig": "phone",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/phone",
                "parts": [
                  "phone",
                ],
                "select": {
                  "exist": [
                    "country_code",
                    "phone",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "AKZONL2A",
                      "kind": "query",
                      "name": "bic",
                      "orig": "bic",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/bic",
                "parts": [
                  "bic",
                ],
                "select": {
                  "exist": [
                    "bic",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "eur",
                      "kind": "query",
                      "name": "currency",
                      "orig": "currency",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/currency",
                "parts": [
                  "currency",
                ],
                "select": {
                  "exist": [
                    "currency",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 4,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "bedrijfsdata.nl",
                      "kind": "query",
                      "name": "domain",
                      "orig": "domain",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/dns",
                "parts": [
                  "dns",
                ],
                "select": {
                  "exist": [
                    "domain",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 5,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "google.com",
                      "kind": "query",
                      "name": "domain",
                      "orig": "domain",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/webrank",
                "parts": [
                  "webrank",
                ],
                "select": {
                  "exist": [
                    "domain",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 6,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "piet@bedrijfsdata.nl",
                      "kind": "query",
                      "name": "email",
                      "orig": "email",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/email",
                "parts": [
                  "email",
                ],
                "select": {
                  "exist": [
                    "email",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 7,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "NL17ADYB2017400505",
                      "kind": "query",
                      "name": "iban",
                      "orig": "iban",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/iban",
                "parts": [
                  "iban",
                ],
                "select": {
                  "exist": [
                    "iban",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 8,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "ua",
                      "orig": "ua",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/ip",
                "parts": [
                  "ip",
                ],
                "select": {
                  "exist": [
                    "ua",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 9,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "http://www.bedrijfsdata.nl",
                      "kind": "query",
                      "name": "url",
                      "orig": "url",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/url",
                "parts": [
                  "url",
                ],
                "select": {
                  "exist": [
                    "url",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 10,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "netherlands_ap_i": {
        "fields": [
          {
            "active": True,
            "name": "active",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "addition",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "city",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "coc",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "construction_year",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "floor_area",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "freeformaddress",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "lat",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "letter",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "lon",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "municipality",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "number",
            "req": False,
            "type": "`$STRING`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "postcode",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "province",
            "req": False,
            "type": "`$STRING`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "province_code",
            "req": False,
            "type": "`$STRING`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "purpose",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "street",
            "req": False,
            "type": "`$STRING`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "type",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "vestiging",
            "req": False,
            "type": "`$STRING`",
            "index$": 20,
          },
        ],
        "name": "netherlands_ap_i",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "3",
                      "kind": "query",
                      "name": "number",
                      "orig": "number",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "1011PN",
                      "kind": "query",
                      "name": "postcode",
                      "orig": "postcode",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "a",
                      "kind": "query",
                      "name": "suffix",
                      "orig": "suffix",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/bag",
                "parts": [
                  "bag",
                ],
                "select": {
                  "exist": [
                    "number",
                    "postcode",
                    "suffix",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "example": "89395808",
                      "kind": "query",
                      "name": "kvk",
                      "orig": "kvk",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/kvk",
                "parts": [
                  "kvk",
                ],
                "select": {
                  "exist": [
                    "kvk",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
