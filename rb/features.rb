# DutchCustomerData SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module DutchCustomerDataFeatures
  def self.make_feature(name)
    case name
    when "base"
      DutchCustomerDataBaseFeature.new
    when "ratelimit"
      DutchCustomerDataRatelimitFeature.new
    when "retry"
      DutchCustomerDataRetryFeature.new
    when "test"
      DutchCustomerDataTestFeature.new
    when "timeout"
      DutchCustomerDataTimeoutFeature.new
    else
      DutchCustomerDataBaseFeature.new
    end
  end
end
