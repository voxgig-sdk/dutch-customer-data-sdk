package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewEuApIEntityFunc func(client *DutchCustomerDataSDK, entopts map[string]any) DutchCustomerDataEntity

var NewGlobalApIEntityFunc func(client *DutchCustomerDataSDK, entopts map[string]any) DutchCustomerDataEntity

var NewNetherlandsApIEntityFunc func(client *DutchCustomerDataSDK, entopts map[string]any) DutchCustomerDataEntity

