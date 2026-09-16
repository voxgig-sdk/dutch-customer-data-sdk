package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "DutchCustomerData",
			"slug": "dutch-customer-data",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://free.bedrijfsdata.nl/v1.1",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"eu_ap_i": map[string]any{},
				"global_ap_i": map[string]any{},
				"netherlands_ap_i": map[string]any{},
			},
		},
		"entity": map[string]any{
			"eu_ap_i": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "active",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "buyer",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "buyer_country",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "city",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "contract_nature",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "uri",
						"name": "html",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "uri",
						"name": "link",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notice_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "official_language",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "uri",
						"name": "pdf",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_of_performance",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "procedure_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date",
						"name": "publication_date",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "response_date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vat",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "eu_ap_i",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "37080091",
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tender",
								"segments": []any{
									map[string]any{
										"lit": "tender",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"q",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.tender`",
								},
								"parts": []any{
									"tender",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "NL001672022B01",
											"kind": "query",
											"name": "vat",
											"orig": "vat",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/vat",
								"segments": []any{
									map[string]any{
										"lit": "vat",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"vat",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.vat`",
								},
								"parts": []any{
									"vat",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"global_ap_i": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "addition",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "admin1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "admin2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "admin3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "bank",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "bic",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "browser",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "builtwith",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "carrier",
						"short": "Carrier name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "city",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cloudflare",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "commoncrawl",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "content_length",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "content_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country",
						"short": "ISO country code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country_code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "crux",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "device_family",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "device_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "device_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "disposable",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dns_a",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "dns_mx",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "dns_ns",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "dns_soa",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "dns_txt",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "dns_www_a",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "dnsserver",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "domain",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "domcop",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "found",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "free",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "freeformaddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "host",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "host_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hostio",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "http_code",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "iban",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iban_human",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "int",
						"short": "International format without plus sign",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "international",
						"short": "International formatted phone number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ip",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ipint",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ismobile",
						"short": "1 if mobile, 0 if not",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"format": "double",
						"name": "lat",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "lei",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "letter",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "local_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "double",
						"name": "lon",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "mailserver",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "majestic",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "message",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "municipality",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mx_host",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mx_ip",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "national",
						"short": "National formatted phone number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "number",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ocid",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pagerank",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "platform",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "population",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "postcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "province",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "province_code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "redirect_count",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "region",
						"short": "Geographic region",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "register_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "renewal_date",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "double",
						"name": "score",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sepa",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "spf",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "street",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "success",
						"short": "1 if successful, 0 if not",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "swift",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "tranco",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "umbrella",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "user",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "user_agent",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "valid",
						"short": "1 if valid, 0 if not",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "verified",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "verified_checksum",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "webrank",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "wrong_email",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "wrong_format",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "wrong_password",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "wrong_phone",
						"short": "1 if wrong, 0 if correct",
						"type": "`$INTEGER`",
					},
				},
				"name": "global_ap_i",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/password",
								"segments": []any{
									map[string]any{
										"lit": "password",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.password`",
								},
								"parts": []any{
									"password",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "Haarlem",
											"kind": "query",
											"name": "city",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "NL",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "full",
											"orig": "full",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/city",
								"segments": []any{
									map[string]any{
										"lit": "city",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"city",
										"country_code",
										"full",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.city`",
								},
								"parts": []any{
									"city",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "NL",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "1013PN",
											"kind": "query",
											"name": "postcode",
											"orig": "postcode",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/postcode",
								"segments": []any{
									map[string]any{
										"lit": "postcode",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"country_code",
										"postcode",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.postcode`",
								},
								"parts": []any{
									"postcode",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "nl",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Kalverstraat 1, 1012NX",
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/geocoding",
								"segments": []any{
									map[string]any{
										"lit": "geocoding",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"country_code",
										"q",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.geocoding`",
								},
								"parts": []any{
									"geocoding",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "AKZONL2AXXX",
											"kind": "query",
											"name": "bic",
											"orig": "bic",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "724500XYIJUGXAA5QD70",
											"kind": "query",
											"name": "lei",
											"orig": "lei",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "09007809",
											"kind": "query",
											"name": "local_id",
											"orig": "local_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/lei",
								"segments": []any{
									map[string]any{
										"lit": "lei",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"bic",
										"lei",
										"local_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.lei`",
								},
								"parts": []any{
									"lei",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "password",
											"orig": "password",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "2c4c3891e2ac6958e9810a1e49c6705784fbfa1a",
											"kind": "query",
											"name": "password_sha1",
											"orig": "password_sha1",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 25,
											"kind": "query",
											"name": "threshold",
											"orig": "threshold",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/password",
								"segments": []any{
									map[string]any{
										"lit": "password",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"password",
										"password_sha1",
										"threshold",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.password`",
								},
								"parts": []any{
									"password",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "nl",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "+31207895050",
											"kind": "query",
											"name": "phone",
											"orig": "phone",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/phone",
								"segments": []any{
									map[string]any{
										"lit": "phone",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"country_code",
										"phone",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.phone`",
								},
								"parts": []any{
									"phone",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "AKZONL2A",
											"kind": "query",
											"name": "bic",
											"orig": "bic",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/bic",
								"segments": []any{
									map[string]any{
										"lit": "bic",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"bic",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.bic`",
								},
								"parts": []any{
									"bic",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "eur",
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/currency",
								"segments": []any{
									map[string]any{
										"lit": "currency",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"currency",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.currency`",
								},
								"parts": []any{
									"currency",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "bedrijfsdata.nl",
											"kind": "query",
											"name": "domain",
											"orig": "domain",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/dns",
								"segments": []any{
									map[string]any{
										"lit": "dns",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"domain",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.dns`",
								},
								"parts": []any{
									"dns",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "google.com",
											"kind": "query",
											"name": "domain",
											"orig": "domain",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/webrank",
								"segments": []any{
									map[string]any{
										"lit": "webrank",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"domain",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.webrank`",
								},
								"parts": []any{
									"webrank",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "piet@bedrijfsdata.nl",
											"kind": "query",
											"name": "email",
											"orig": "email",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/email",
								"segments": []any{
									map[string]any{
										"lit": "email",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.email`",
								},
								"parts": []any{
									"email",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "NL17ADYB2017400505",
											"kind": "query",
											"name": "iban",
											"orig": "iban",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/iban",
								"segments": []any{
									map[string]any{
										"lit": "iban",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"iban",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.iban`",
								},
								"parts": []any{
									"iban",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "ua",
											"orig": "ua",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/ip",
								"segments": []any{
									map[string]any{
										"lit": "ip",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"ua",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.ip`",
								},
								"parts": []any{
									"ip",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "http://www.bedrijfsdata.nl",
											"kind": "query",
											"name": "url",
											"orig": "url",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/url",
								"segments": []any{
									map[string]any{
										"lit": "url",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"url",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.url`",
								},
								"parts": []any{
									"url",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"netherlands_ap_i": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "active",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "addition",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "city",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "coc",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "construction_year",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "floor_area",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "freeformaddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "double",
						"name": "lat",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "letter",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "double",
						"name": "lon",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "municipality",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "province",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "province_code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "purpose",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "street",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vestiging",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "netherlands_ap_i",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "3",
											"kind": "query",
											"name": "number",
											"orig": "number",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "1011PN",
											"kind": "query",
											"name": "postcode",
											"orig": "postcode",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "a",
											"kind": "query",
											"name": "suffix",
											"orig": "suffix",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/bag",
								"segments": []any{
									map[string]any{
										"lit": "bag",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"number",
										"postcode",
										"suffix",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.bag`",
								},
								"parts": []any{
									"bag",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "89395808",
											"kind": "query",
											"name": "kvk",
											"orig": "kvk",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/kvk",
								"segments": []any{
									map[string]any{
										"lit": "kvk",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"kvk",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.kvk`",
								},
								"parts": []any{
									"kvk",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
