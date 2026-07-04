package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "DutchCustomerData",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
						"active": true,
						"name": "buyer",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "buyer_country",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "contract_nature",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "html",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "link",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "notice_type",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "official_language",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "pdf",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "place_of_performance",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "procedure_type",
						"req": false,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "publication_date",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "vat",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 14,
					},
				},
				"name": "eu_ap_i",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "37080091",
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "NL001672022B01",
											"kind": "query",
											"name": "vat",
											"orig": "vat",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"global_ap_i": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "addition",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "admin1",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "admin2",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "admin3",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "bic",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "city",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "currency",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "date",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "dns",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "found",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "freeformaddress",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "from_currency",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "iban",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "ip",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "lat",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "lei",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "letter",
						"req": false,
						"type": "`$STRING`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "lon",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "municipality",
						"req": false,
						"type": "`$STRING`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "number",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "password",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 21,
					},
					map[string]any{
						"active": true,
						"name": "phone",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 22,
					},
					map[string]any{
						"active": true,
						"name": "population",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 23,
					},
					map[string]any{
						"active": true,
						"name": "postcode",
						"req": false,
						"type": "`$STRING`",
						"index$": 24,
					},
					map[string]any{
						"active": true,
						"name": "province",
						"req": false,
						"type": "`$STRING`",
						"index$": 25,
					},
					map[string]any{
						"active": true,
						"name": "province_code",
						"req": false,
						"type": "`$STRING`",
						"index$": 26,
					},
					map[string]any{
						"active": true,
						"name": "score",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 27,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": false,
						"type": "`$STRING`",
						"index$": 28,
					},
					map[string]any{
						"active": true,
						"name": "street",
						"req": false,
						"type": "`$STRING`",
						"index$": 29,
					},
					map[string]any{
						"active": true,
						"name": "type",
						"req": false,
						"type": "`$STRING`",
						"index$": 30,
					},
					map[string]any{
						"active": true,
						"name": "url",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 31,
					},
					map[string]any{
						"active": true,
						"name": "webrank",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 32,
					},
				},
				"name": "global_ap_i",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/password",
								"parts": []any{
									"password",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "Haarlem",
											"kind": "query",
											"name": "city",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "NL",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "full",
											"orig": "full",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "NL",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "1013PN",
											"kind": "query",
											"name": "postcode",
											"orig": "postcode",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "nl",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "Kalverstraat 1, 1012NX",
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 2,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "AKZONL2AXXX",
											"kind": "query",
											"name": "bic",
											"orig": "bic",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "724500XYIJUGXAA5QD70",
											"kind": "query",
											"name": "lei",
											"orig": "lei",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "09007809",
											"kind": "query",
											"name": "local_id",
											"orig": "local_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "password",
											"orig": "password",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "2c4c3891e2ac6958e9810a1e49c6705784fbfa1a",
											"kind": "query",
											"name": "password_sha1",
											"orig": "password_sha1",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 25,
											"kind": "query",
											"name": "threshold",
											"orig": "threshold",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "nl",
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "+31207895050",
											"kind": "query",
											"name": "phone",
											"orig": "phone",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 2,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "AKZONL2A",
											"kind": "query",
											"name": "bic",
											"orig": "bic",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 3,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "eur",
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 4,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "bedrijfsdata.nl",
											"kind": "query",
											"name": "domain",
											"orig": "domain",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 5,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "google.com",
											"kind": "query",
											"name": "domain",
											"orig": "domain",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 6,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "piet@bedrijfsdata.nl",
											"kind": "query",
											"name": "email",
											"orig": "email",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 7,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "NL17ADYB2017400505",
											"kind": "query",
											"name": "iban",
											"orig": "iban",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 8,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "ua",
											"orig": "ua",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 9,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "http://www.bedrijfsdata.nl",
											"kind": "query",
											"name": "url",
											"orig": "url",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 10,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"netherlands_ap_i": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "active",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "addition",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "city",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "coc",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "construction_year",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "floor_area",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "freeformaddress",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "lat",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "letter",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "lon",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "municipality",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "number",
						"req": false,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "postcode",
						"req": false,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "province",
						"req": false,
						"type": "`$STRING`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "province_code",
						"req": false,
						"type": "`$STRING`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "purpose",
						"req": false,
						"type": "`$STRING`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "street",
						"req": false,
						"type": "`$STRING`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "type",
						"req": false,
						"type": "`$STRING`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "vestiging",
						"req": false,
						"type": "`$STRING`",
						"index$": 20,
					},
				},
				"name": "netherlands_ap_i",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "3",
											"kind": "query",
											"name": "number",
											"orig": "number",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "1011PN",
											"kind": "query",
											"name": "postcode",
											"orig": "postcode",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "a",
											"kind": "query",
											"name": "suffix",
											"orig": "suffix",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "89395808",
											"kind": "query",
											"name": "kvk",
											"orig": "kvk",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 1,
							},
						},
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
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
