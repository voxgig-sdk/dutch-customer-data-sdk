# DutchCustomerData SDK utility: feature_add
module DutchCustomerDataUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
