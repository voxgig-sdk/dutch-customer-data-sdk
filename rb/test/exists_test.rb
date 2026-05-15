# DutchCustomerData SDK exists test

require "minitest/autorun"
require_relative "../DutchCustomerData_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = DutchCustomerDataSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
