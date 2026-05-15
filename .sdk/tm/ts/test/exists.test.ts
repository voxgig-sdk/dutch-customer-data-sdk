
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { DutchCustomerDataSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await DutchCustomerDataSDK.test()
    equal(null !== testsdk, true)
  })

})
