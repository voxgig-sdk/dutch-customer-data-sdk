# DutchCustomerData SDK utility: make_context
require_relative '../core/context'
module DutchCustomerDataUtilities
  MakeContext = ->(ctxmap, basectx) {
    DutchCustomerDataContext.new(ctxmap, basectx)
  }
end
