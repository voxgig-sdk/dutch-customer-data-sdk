<?php
declare(strict_types=1);

// DutchCustomerData SDK configuration

class DutchCustomerDataConfig
{
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "DutchCustomerData",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://free.bedrijfsdata.nl/v1.1",
                "auth" => [
                    "prefix" => "Bearer",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "eu_ap_i" => [],
                    "global_ap_i" => [],
                    "netherlands_ap_i" => [],
                ],
            ],
            "entity" => [
        'eu_ap_i' => [
          'fields' => [
            [
              'name' => 'buyer',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'buyer_country',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'contract_nature',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'html',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
            [
              'name' => 'id',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 4,
            ],
            [
              'name' => 'link',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 5,
            ],
            [
              'name' => 'notice_type',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 6,
            ],
            [
              'name' => 'official_language',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 7,
            ],
            [
              'name' => 'pdf',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 8,
            ],
            [
              'name' => 'place_of_performance',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 9,
            ],
            [
              'name' => 'procedure_type',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 10,
            ],
            [
              'name' => 'publication_date',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 11,
            ],
            [
              'name' => 'status',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 12,
            ],
            [
              'name' => 'title',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 13,
            ],
            [
              'name' => 'vat',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 14,
            ],
          ],
          'name' => 'eu_ap_i',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => '37080091',
                        'kind' => 'query',
                        'name' => 'q',
                        'orig' => 'q',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/tender',
                  'parts' => [
                    'tender',
                  ],
                  'select' => [
                    'exist' => [
                      'q',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'NL001672022B01',
                        'kind' => 'query',
                        'name' => 'vat',
                        'orig' => 'vat',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/vat',
                  'parts' => [
                    'vat',
                  ],
                  'select' => [
                    'exist' => [
                      'vat',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'global_ap_i' => [
          'fields' => [
            [
              'name' => 'addition',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'admin1',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'admin2',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'admin3',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
            [
              'name' => 'bic',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 4,
            ],
            [
              'name' => 'city',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 5,
            ],
            [
              'name' => 'currency',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 6,
            ],
            [
              'name' => 'date',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 7,
            ],
            [
              'name' => 'dns',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 8,
            ],
            [
              'name' => 'email',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 9,
            ],
            [
              'name' => 'found',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 10,
            ],
            [
              'name' => 'freeformaddress',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 11,
            ],
            [
              'name' => 'from_currency',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 12,
            ],
            [
              'name' => 'iban',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 13,
            ],
            [
              'name' => 'ip',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 14,
            ],
            [
              'name' => 'lat',
              'req' => false,
              'type' => '`$NUMBER`',
              'active' => true,
              'index$' => 15,
            ],
            [
              'name' => 'lei',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 16,
            ],
            [
              'name' => 'letter',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 17,
            ],
            [
              'name' => 'lon',
              'req' => false,
              'type' => '`$NUMBER`',
              'active' => true,
              'index$' => 18,
            ],
            [
              'name' => 'municipality',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 19,
            ],
            [
              'name' => 'number',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 20,
            ],
            [
              'name' => 'password',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 21,
            ],
            [
              'name' => 'phone',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 22,
            ],
            [
              'name' => 'population',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 23,
            ],
            [
              'name' => 'postcode',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 24,
            ],
            [
              'name' => 'province',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 25,
            ],
            [
              'name' => 'province_code',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 26,
            ],
            [
              'name' => 'score',
              'req' => false,
              'type' => '`$NUMBER`',
              'active' => true,
              'index$' => 27,
            ],
            [
              'name' => 'status',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 28,
            ],
            [
              'name' => 'street',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 29,
            ],
            [
              'name' => 'type',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 30,
            ],
            [
              'name' => 'url',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 31,
            ],
            [
              'name' => 'webrank',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 32,
            ],
          ],
          'name' => 'global_ap_i',
          'op' => [
            'create' => [
              'name' => 'create',
              'points' => [
                [
                  'method' => 'POST',
                  'orig' => '/password',
                  'parts' => [
                    'password',
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'args' => [],
                  'select' => [],
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'create',
            ],
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'Haarlem',
                        'kind' => 'query',
                        'name' => 'city',
                        'orig' => 'city',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'NL',
                        'kind' => 'query',
                        'name' => 'country_code',
                        'orig' => 'country_code',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'full',
                        'orig' => 'full',
                        'reqd' => false,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/city',
                  'parts' => [
                    'city',
                  ],
                  'select' => [
                    'exist' => [
                      'city',
                      'country_code',
                      'full',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'NL',
                        'kind' => 'query',
                        'name' => 'country_code',
                        'orig' => 'country_code',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '1013PN',
                        'kind' => 'query',
                        'name' => 'postcode',
                        'orig' => 'postcode',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/postcode',
                  'parts' => [
                    'postcode',
                  ],
                  'select' => [
                    'exist' => [
                      'country_code',
                      'postcode',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'nl',
                        'kind' => 'query',
                        'name' => 'country_code',
                        'orig' => 'country_code',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'Kalverstraat 1, 1012NX',
                        'kind' => 'query',
                        'name' => 'q',
                        'orig' => 'q',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/geocoding',
                  'parts' => [
                    'geocoding',
                  ],
                  'select' => [
                    'exist' => [
                      'country_code',
                      'q',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 2,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'AKZONL2AXXX',
                        'kind' => 'query',
                        'name' => 'bic',
                        'orig' => 'bic',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '724500XYIJUGXAA5QD70',
                        'kind' => 'query',
                        'name' => 'lei',
                        'orig' => 'lei',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '09007809',
                        'kind' => 'query',
                        'name' => 'local_id',
                        'orig' => 'local_id',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/lei',
                  'parts' => [
                    'lei',
                  ],
                  'select' => [
                    'exist' => [
                      'bic',
                      'lei',
                      'local_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'password',
                        'orig' => 'password',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '2c4c3891e2ac6958e9810a1e49c6705784fbfa1a',
                        'kind' => 'query',
                        'name' => 'password_sha1',
                        'orig' => 'password_sha1',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 25,
                        'kind' => 'query',
                        'name' => 'threshold',
                        'orig' => 'threshold',
                        'reqd' => false,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/password',
                  'parts' => [
                    'password',
                  ],
                  'select' => [
                    'exist' => [
                      'password',
                      'password_sha1',
                      'threshold',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'nl',
                        'kind' => 'query',
                        'name' => 'country_code',
                        'orig' => 'country_code',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '+31207895050',
                        'kind' => 'query',
                        'name' => 'phone',
                        'orig' => 'phone',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/phone',
                  'parts' => [
                    'phone',
                  ],
                  'select' => [
                    'exist' => [
                      'country_code',
                      'phone',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 2,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'AKZONL2A',
                        'kind' => 'query',
                        'name' => 'bic',
                        'orig' => 'bic',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/bic',
                  'parts' => [
                    'bic',
                  ],
                  'select' => [
                    'exist' => [
                      'bic',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 3,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'eur',
                        'kind' => 'query',
                        'name' => 'currency',
                        'orig' => 'currency',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/currency',
                  'parts' => [
                    'currency',
                  ],
                  'select' => [
                    'exist' => [
                      'currency',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 4,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'bedrijfsdata.nl',
                        'kind' => 'query',
                        'name' => 'domain',
                        'orig' => 'domain',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/dns',
                  'parts' => [
                    'dns',
                  ],
                  'select' => [
                    'exist' => [
                      'domain',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 5,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'google.com',
                        'kind' => 'query',
                        'name' => 'domain',
                        'orig' => 'domain',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/webrank',
                  'parts' => [
                    'webrank',
                  ],
                  'select' => [
                    'exist' => [
                      'domain',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 6,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'piet@bedrijfsdata.nl',
                        'kind' => 'query',
                        'name' => 'email',
                        'orig' => 'email',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/email',
                  'parts' => [
                    'email',
                  ],
                  'select' => [
                    'exist' => [
                      'email',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 7,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'NL17ADYB2017400505',
                        'kind' => 'query',
                        'name' => 'iban',
                        'orig' => 'iban',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/iban',
                  'parts' => [
                    'iban',
                  ],
                  'select' => [
                    'exist' => [
                      'iban',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 8,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'ua',
                        'orig' => 'ua',
                        'reqd' => false,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/ip',
                  'parts' => [
                    'ip',
                  ],
                  'select' => [
                    'exist' => [
                      'ua',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 9,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'http://www.bedrijfsdata.nl',
                        'kind' => 'query',
                        'name' => 'url',
                        'orig' => 'url',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/url',
                  'parts' => [
                    'url',
                  ],
                  'select' => [
                    'exist' => [
                      'url',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 10,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'netherlands_ap_i' => [
          'fields' => [
            [
              'name' => 'active',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'addition',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'city',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'coc',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
            [
              'name' => 'construction_year',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 4,
            ],
            [
              'name' => 'floor_area',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 5,
            ],
            [
              'name' => 'freeformaddress',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 6,
            ],
            [
              'name' => 'id',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 7,
            ],
            [
              'name' => 'lat',
              'req' => false,
              'type' => '`$NUMBER`',
              'active' => true,
              'index$' => 8,
            ],
            [
              'name' => 'letter',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 9,
            ],
            [
              'name' => 'lon',
              'req' => false,
              'type' => '`$NUMBER`',
              'active' => true,
              'index$' => 10,
            ],
            [
              'name' => 'municipality',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 11,
            ],
            [
              'name' => 'name',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 12,
            ],
            [
              'name' => 'number',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 13,
            ],
            [
              'name' => 'postcode',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 14,
            ],
            [
              'name' => 'province',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 15,
            ],
            [
              'name' => 'province_code',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 16,
            ],
            [
              'name' => 'purpose',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 17,
            ],
            [
              'name' => 'street',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 18,
            ],
            [
              'name' => 'type',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 19,
            ],
            [
              'name' => 'vestiging',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 20,
            ],
          ],
          'name' => 'netherlands_ap_i',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => '3',
                        'kind' => 'query',
                        'name' => 'number',
                        'orig' => 'number',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => '1011PN',
                        'kind' => 'query',
                        'name' => 'postcode',
                        'orig' => 'postcode',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'a',
                        'kind' => 'query',
                        'name' => 'suffix',
                        'orig' => 'suffix',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/bag',
                  'parts' => [
                    'bag',
                  ],
                  'select' => [
                    'exist' => [
                      'number',
                      'postcode',
                      'suffix',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => '89395808',
                        'kind' => 'query',
                        'name' => 'kvk',
                        'orig' => 'kvk',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/kvk',
                  'parts' => [
                    'kvk',
                  ],
                  'select' => [
                    'exist' => [
                      'kvk',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return DutchCustomerDataFeatures::make_feature($name);
    }
}
