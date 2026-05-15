-- ProjectName SDK exists test

local sdk = require("dutch-customer-data_sdk")

describe("DutchCustomerDataSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
