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
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
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
						"name": "html",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
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
						"name": "publication_date",
						"type": "`$STRING`",
					},
					map[string]any{
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
								"parts": []any{
									"tender",
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
								"parts": []any{
									"vat",
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
						"name": "renewal_date",
						"type": "`$STRING`",
					},
					map[string]any{
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
								"parts": []any{
									"password",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.password`",
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
								"parts": []any{
									"city",
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
								"parts": []any{
									"postcode",
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
								"parts": []any{
									"geocoding",
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
								"parts": []any{
									"lei",
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
								"parts": []any{
									"password",
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
								"parts": []any{
									"phone",
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
								"parts": []any{
									"bic",
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
								"parts": []any{
									"currency",
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
								"parts": []any{
									"dns",
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
								"parts": []any{
									"webrank",
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
								"parts": []any{
									"email",
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
								"parts": []any{
									"iban",
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
								"parts": []any{
									"ip",
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
								"parts": []any{
									"url",
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
						"name": "lat",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "letter",
						"type": "`$STRING`",
					},
					map[string]any{
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
								"parts": []any{
									"bag",
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
								"parts": []any{
									"kvk",
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
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
