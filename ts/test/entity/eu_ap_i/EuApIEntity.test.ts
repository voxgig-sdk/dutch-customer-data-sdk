
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { DutchCustomerDataSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('EuApIEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DUTCHCUSTOMERDATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('DUTCHCUSTOMERDATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DutchCustomerDataSDK.test()
    const ent = testsdk.EuApI()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DUTCH_CUSTOMER_DATA_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (maybeSkipControl(t, 'entityOp', 'eu_ap_i.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set DUTCH_CUSTOMER_DATA_TEST_EU_AP_I_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let eu_ap_i_ref01_data = Object.values(setup.data.existing.eu_ap_i)[0] as any

    // LIST
    const eu_ap_i_ref01_ent = client.EuApI()
    const eu_ap_i_ref01_match: any = {}

    const eu_ap_i_ref01_list = await eu_ap_i_ref01_ent.list(eu_ap_i_ref01_match)


    // LOAD
    const eu_ap_i_ref01_match_dt0: any = {}
    eu_ap_i_ref01_match_dt0.id = eu_ap_i_ref01_data.id
    const eu_ap_i_ref01_data_dt0 = await eu_ap_i_ref01_ent.load(eu_ap_i_ref01_match_dt0)
    assert(eu_ap_i_ref01_data_dt0.id === eu_ap_i_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/eu_ap_i/EuApITestData.json')

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
    ['eu_ap_i01','eu_ap_i02','eu_ap_i03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['DUTCH_CUSTOMER_DATA_TEST_EU_AP_I_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'DUTCH_CUSTOMER_DATA_TEST_EU_AP_I_ENTID': idmap,
    'DUTCH_CUSTOMER_DATA_TEST_LIVE': 'FALSE',
    'DUTCH_CUSTOMER_DATA_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['DUTCH_CUSTOMER_DATA_TEST_EU_AP_I_ENTID']

  const live = 'TRUE' === env.DUTCH_CUSTOMER_DATA_TEST_LIVE

  if (live) {
    client = new DutchCustomerDataSDK(merge([
      {
      },
      extra
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
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
