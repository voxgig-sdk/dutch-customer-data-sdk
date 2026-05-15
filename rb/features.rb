# DutchCustomerData SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module DutchCustomerDataFeatures
  def self.make_feature(name)
    case name
    when "base"
      DutchCustomerDataBaseFeature.new
    when "test"
      DutchCustomerDataTestFeature.new
    else
      DutchCustomerDataBaseFeature.new
    end
  end
end
