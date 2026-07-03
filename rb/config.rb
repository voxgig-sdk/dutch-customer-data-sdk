# DutchCustomerData SDK configuration

module DutchCustomerDataConfig
  def self.make_config
    {
      "main" => {
        "name" => "DutchCustomerData",
      },
      "feature" => {
        "test" => {
          "options" => {
            "active" => false,
          },
        },
      },
      "options" => {
        "base" => "https://free.bedrijfsdata.nl/v1.1",
        "auth" => {
          "prefix" => "Bearer",
        },
        "headers" => {
          "content-type" => "application/json",
        },
        "entity" => {
          "eu_ap_i" => {},
          "global_ap_i" => {},
          "netherlands_ap_i" => {},
        },
      },
      "entity" => {
        "eu_ap_i" => {
          "fields" => [
            {
              "active" => true,
              "name" => "buyer",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "buyer_country",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "contract_nature",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "html",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "link",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "notice_type",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "official_language",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "pdf",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "place_of_performance",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "procedure_type",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "publication_date",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "status",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "title",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "vat",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 14,
            },
          ],
          "name" => "eu_ap_i",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "37080091",
                        "kind" => "query",
                        "name" => "q",
                        "orig" => "q",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/tender",
                  "parts" => [
                    "tender",
                  ],
                  "select" => {
                    "exist" => [
                      "q",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "NL001672022B01",
                        "kind" => "query",
                        "name" => "vat",
                        "orig" => "vat",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/vat",
                  "parts" => [
                    "vat",
                  ],
                  "select" => {
                    "exist" => [
                      "vat",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "global_ap_i" => {
          "fields" => [
            {
              "active" => true,
              "name" => "addition",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "admin1",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "admin2",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "admin3",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "bic",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "city",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "currency",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "date",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "dns",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "email",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "found",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "freeformaddress",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "from_currency",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "iban",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "ip",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 14,
            },
            {
              "active" => true,
              "name" => "lat",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 15,
            },
            {
              "active" => true,
              "name" => "lei",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 16,
            },
            {
              "active" => true,
              "name" => "letter",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 17,
            },
            {
              "active" => true,
              "name" => "lon",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 18,
            },
            {
              "active" => true,
              "name" => "municipality",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 19,
            },
            {
              "active" => true,
              "name" => "number",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 20,
            },
            {
              "active" => true,
              "name" => "password",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 21,
            },
            {
              "active" => true,
              "name" => "phone",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 22,
            },
            {
              "active" => true,
              "name" => "population",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 23,
            },
            {
              "active" => true,
              "name" => "postcode",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 24,
            },
            {
              "active" => true,
              "name" => "province",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 25,
            },
            {
              "active" => true,
              "name" => "province_code",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 26,
            },
            {
              "active" => true,
              "name" => "score",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 27,
            },
            {
              "active" => true,
              "name" => "status",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 28,
            },
            {
              "active" => true,
              "name" => "street",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 29,
            },
            {
              "active" => true,
              "name" => "type",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 30,
            },
            {
              "active" => true,
              "name" => "url",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 31,
            },
            {
              "active" => true,
              "name" => "webrank",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 32,
            },
          ],
          "name" => "global_ap_i",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {},
                  "method" => "POST",
                  "orig" => "/password",
                  "parts" => [
                    "password",
                  ],
                  "select" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "Haarlem",
                        "kind" => "query",
                        "name" => "city",
                        "orig" => "city",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "NL",
                        "kind" => "query",
                        "name" => "country_code",
                        "orig" => "country_code",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "full",
                        "orig" => "full",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/city",
                  "parts" => [
                    "city",
                  ],
                  "select" => {
                    "exist" => [
                      "city",
                      "country_code",
                      "full",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "NL",
                        "kind" => "query",
                        "name" => "country_code",
                        "orig" => "country_code",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "1013PN",
                        "kind" => "query",
                        "name" => "postcode",
                        "orig" => "postcode",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/postcode",
                  "parts" => [
                    "postcode",
                  ],
                  "select" => {
                    "exist" => [
                      "country_code",
                      "postcode",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "nl",
                        "kind" => "query",
                        "name" => "country_code",
                        "orig" => "country_code",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "Kalverstraat 1, 1012NX",
                        "kind" => "query",
                        "name" => "q",
                        "orig" => "q",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/geocoding",
                  "parts" => [
                    "geocoding",
                  ],
                  "select" => {
                    "exist" => [
                      "country_code",
                      "q",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "AKZONL2AXXX",
                        "kind" => "query",
                        "name" => "bic",
                        "orig" => "bic",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "724500XYIJUGXAA5QD70",
                        "kind" => "query",
                        "name" => "lei",
                        "orig" => "lei",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "09007809",
                        "kind" => "query",
                        "name" => "local_id",
                        "orig" => "local_id",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/lei",
                  "parts" => [
                    "lei",
                  ],
                  "select" => {
                    "exist" => [
                      "bic",
                      "lei",
                      "local_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "password",
                        "orig" => "password",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "2c4c3891e2ac6958e9810a1e49c6705784fbfa1a",
                        "kind" => "query",
                        "name" => "password_sha1",
                        "orig" => "password_sha1",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => 25,
                        "kind" => "query",
                        "name" => "threshold",
                        "orig" => "threshold",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/password",
                  "parts" => [
                    "password",
                  ],
                  "select" => {
                    "exist" => [
                      "password",
                      "password_sha1",
                      "threshold",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "nl",
                        "kind" => "query",
                        "name" => "country_code",
                        "orig" => "country_code",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "+31207895050",
                        "kind" => "query",
                        "name" => "phone",
                        "orig" => "phone",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/phone",
                  "parts" => [
                    "phone",
                  ],
                  "select" => {
                    "exist" => [
                      "country_code",
                      "phone",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "AKZONL2A",
                        "kind" => "query",
                        "name" => "bic",
                        "orig" => "bic",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/bic",
                  "parts" => [
                    "bic",
                  ],
                  "select" => {
                    "exist" => [
                      "bic",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 3,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "eur",
                        "kind" => "query",
                        "name" => "currency",
                        "orig" => "currency",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/currency",
                  "parts" => [
                    "currency",
                  ],
                  "select" => {
                    "exist" => [
                      "currency",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 4,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "bedrijfsdata.nl",
                        "kind" => "query",
                        "name" => "domain",
                        "orig" => "domain",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/dns",
                  "parts" => [
                    "dns",
                  ],
                  "select" => {
                    "exist" => [
                      "domain",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 5,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "google.com",
                        "kind" => "query",
                        "name" => "domain",
                        "orig" => "domain",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/webrank",
                  "parts" => [
                    "webrank",
                  ],
                  "select" => {
                    "exist" => [
                      "domain",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 6,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "piet@bedrijfsdata.nl",
                        "kind" => "query",
                        "name" => "email",
                        "orig" => "email",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/email",
                  "parts" => [
                    "email",
                  ],
                  "select" => {
                    "exist" => [
                      "email",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 7,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "NL17ADYB2017400505",
                        "kind" => "query",
                        "name" => "iban",
                        "orig" => "iban",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/iban",
                  "parts" => [
                    "iban",
                  ],
                  "select" => {
                    "exist" => [
                      "iban",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 8,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "ua",
                        "orig" => "ua",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/ip",
                  "parts" => [
                    "ip",
                  ],
                  "select" => {
                    "exist" => [
                      "ua",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 9,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "http://www.bedrijfsdata.nl",
                        "kind" => "query",
                        "name" => "url",
                        "orig" => "url",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/url",
                  "parts" => [
                    "url",
                  ],
                  "select" => {
                    "exist" => [
                      "url",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 10,
                },
              ],
              "key$" => "load",
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "netherlands_ap_i" => {
          "fields" => [
            {
              "active" => true,
              "name" => "active",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "addition",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "city",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "coc",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "construction_year",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "floor_area",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "freeformaddress",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "lat",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "letter",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "lon",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "municipality",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "number",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "postcode",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 14,
            },
            {
              "active" => true,
              "name" => "province",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 15,
            },
            {
              "active" => true,
              "name" => "province_code",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 16,
            },
            {
              "active" => true,
              "name" => "purpose",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 17,
            },
            {
              "active" => true,
              "name" => "street",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 18,
            },
            {
              "active" => true,
              "name" => "type",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 19,
            },
            {
              "active" => true,
              "name" => "vestiging",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 20,
            },
          ],
          "name" => "netherlands_ap_i",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "3",
                        "kind" => "query",
                        "name" => "number",
                        "orig" => "number",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "1011PN",
                        "kind" => "query",
                        "name" => "postcode",
                        "orig" => "postcode",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "a",
                        "kind" => "query",
                        "name" => "suffix",
                        "orig" => "suffix",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/bag",
                  "parts" => [
                    "bag",
                  ],
                  "select" => {
                    "exist" => [
                      "number",
                      "postcode",
                      "suffix",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "query" => [
                      {
                        "active" => true,
                        "example" => "89395808",
                        "kind" => "query",
                        "name" => "kvk",
                        "orig" => "kvk",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/kvk",
                  "parts" => [
                    "kvk",
                  ],
                  "select" => {
                    "exist" => [
                      "kvk",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
              ],
              "key$" => "list",
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
      },
    }
  end


  def self.make_feature(name)
    require_relative 'features'
    DutchCustomerDataFeatures.make_feature(name)
  end
end
