package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewEuApIEntityFunc func(client *DutchCustomerDataSDK, entopts map[string]any) DutchCustomerDataEntity

var NewGlobalApIEntityFunc func(client *DutchCustomerDataSDK, entopts map[string]any) DutchCustomerDataEntity

var NewNetherlandsApIEntityFunc func(client *DutchCustomerDataSDK, entopts map[string]any) DutchCustomerDataEntity

