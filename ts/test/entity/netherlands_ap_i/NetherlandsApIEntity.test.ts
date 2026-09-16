

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { DutchCustomerDataSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('NetherlandsApIEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DUTCH_CUSTOMER_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('DUTCH_CUSTOMER_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DutchCustomerDataSDK.test()
    const ent = testsdk.NetherlandsApI()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DUTCH_CUSTOMER_DATA_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'netherlands_ap_i.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"active","req":false,"type":"`$INTEGER`","index$":0},{"active":true,"name":"addition","req":false,"type":"`$STRING`","index$":1},{"active":true,"name":"city","req":false,"type":"`$STRING`","index$":2},{"active":true,"name":"coc","req":false,"type":"`$STRING`","index$":3},{"active":true,"name":"construction_year","req":false,"type":"`$INTEGER`","index$":4},{"active":true,"name":"floor_area","req":false,"type":"`$INTEGER`","index$":5},{"active":true,"name":"freeformaddress","req":false,"type":"`$STRING`","index$":6},{"active":true,"name":"id","req":false,"type":"`$STRING`","index$":7},{"active":true,"format":"double","name":"lat","req":false,"type":"`$NUMBER`","index$":8},{"active":true,"name":"letter","req":false,"type":"`$STRING`","index$":9},{"active":true,"format":"double","name":"lon","req":false,"type":"`$NUMBER`","index$":10},{"active":true,"name":"municipality","req":false,"type":"`$STRING`","index$":11},{"active":true,"name":"name","req":false,"type":"`$STRING`","index$":12},{"active":true,"name":"number","req":false,"type":"`$STRING`","index$":13},{"active":true,"name":"postcode","req":false,"type":"`$STRING`","index$":14},{"active":true,"name":"province","req":false,"type":"`$STRING`","index$":15},{"active":true,"name":"province_code","req":false,"type":"`$STRING`","index$":16},{"active":true,"name":"purpose","req":false,"type":"`$STRING`","index$":17},{"active":true,"name":"street","req":false,"type":"`$STRING`","index$":18},{"active":true,"name":"type","req":false,"type":"`$STRING`","index$":19},{"active":true,"name":"vestiging","req":false,"type":"`$STRING`","index$":20}],"id":{"field":"id","name":"id"},"name":"netherlands_ap_i","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"query":[{"active":true,"example":"3","kind":"query","name":"number","orig":"number","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"example":"1011PN","kind":"query","name":"postcode","orig":"postcode","reqd":true,"type":"`$STRING`","index$":1},{"active":true,"example":"a","kind":"query","name":"suffix","orig":"suffix","reqd":false,"type":"`$STRING`","index$":2}]},"contract":{"id":"GET /bag","json":"{\"operationId\":\"verifyBag\",\"parameters\":[{\"description\":\"Dutch postcode\",\"in\":\"query\",\"name\":\"postcode\",\"required\":true,\"schema\":{\"example\":\"1011PN\",\"type\":\"string\"}},{\"description\":\"House number\",\"in\":\"query\",\"name\":\"number\",\"required\":true,\"schema\":{\"example\":\"3\",\"type\":\"string\"}},{\"description\":\"House number suffix (letter or addition)\",\"in\":\"query\",\"name\":\"suffix\",\"required\":false,\"schema\":{\"example\":\"a\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"bag\":{\"items\":{\"properties\":{\"addition\":{\"type\":\"string\"},\"city\":{\"type\":\"string\"},\"construction_year\":{\"type\":\"integer\"},\"floor_area\":{\"type\":\"integer\"},\"freeformaddress\":{\"type\":\"string\"},\"lat\":{\"format\":\"double\",\"type\":\"number\"},\"letter\":{\"type\":\"string\"},\"lon\":{\"format\":\"double\",\"type\":\"number\"},\"municipality\":{\"type\":\"string\"},\"number\":{\"type\":\"string\"},\"postcode\":{\"type\":\"string\"},\"province\":{\"type\":\"string\"},\"province_code\":{\"type\":\"string\"},\"purpose\":{\"type\":\"string\"},\"street\":{\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"},\"found\":{\"type\":\"integer\"},\"status\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Successful response\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/bag","segments":[{"lit":"bag"}],"select":{"exist":["number","postcode","suffix"]},"transform":{"req":"`reqdata`","res":"`body.bag`"},"index$":0},{"active":true,"args":{"query":[{"active":true,"example":"89395808","kind":"query","name":"kvk","orig":"kvk","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /kvk","json":"{\"operationId\":\"verifyKvk\",\"parameters\":[{\"description\":\"Dutch Chamber of Commerce (KvK) number\",\"in\":\"query\",\"name\":\"kvk\",\"required\":true,\"schema\":{\"example\":\"89395808\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"found\":{\"type\":\"integer\"},\"kvk\":{\"items\":{\"properties\":{\"active\":{\"type\":\"integer\"},\"city\":{\"type\":\"string\"},\"coc\":{\"type\":\"string\"},\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"type\":{\"type\":\"string\"},\"vestiging\":{\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"},\"status\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Successful response\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/kvk","segments":[{"lit":"kvk"}],"select":{"exist":["kvk"]},"transform":{"req":"`reqdata`","res":"`body.kvk`"},"index$":1}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"netherlands_ap_i","name__orig":"netherlands_ap_i","Name":"NetherlandsApI","name_":"netherlands_ap_i","name-":"netherlands-ap-i","NAME":"NETHERLANDS_AP_I","index$":2}, {"active":true,"entity":"netherlands_ap_i","key$":"BasicNetherlandsApIFlow","kind":"basic","name":"BasicNetherlandsApIFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"netherlands_ap_i_ref01"}}],"index$":0}]}, 'NetherlandsApI')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let netherlands_ap_i_ref01_data = Object.values(setup.data.existing.netherlands_ap_i)[0] as any

    // LIST
    const netherlands_ap_i_ref01_ent = client.NetherlandsApI()
    const netherlands_ap_i_ref01_match: any = {}

    const netherlands_ap_i_ref01_list = (await netherlands_ap_i_ref01_ent.list(netherlands_ap_i_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/netherlands_ap_i/NetherlandsApITestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = DutchCustomerDataSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['netherlands_ap_i01','netherlands_ap_i02','netherlands_ap_i03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DUTCH_CUSTOMER_DATA_TEST_NETHERLANDS_AP_I_ENTID': idmap,
    'DUTCH_CUSTOMER_DATA_TEST_LIVE': 'FALSE',
    'DUTCH_CUSTOMER_DATA_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['DUTCH_CUSTOMER_DATA_TEST_NETHERLANDS_AP_I_ENTID']

  const live = 'TRUE' === env.DUTCH_CUSTOMER_DATA_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DUTCH_CUSTOMER_DATA_TEST_NETHERLANDS_AP_I_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new DutchCustomerDataSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
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
  }

  return setup
}
  
