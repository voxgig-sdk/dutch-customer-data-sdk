package voxgigdutchcustomerdatasdk

import (
	"github.com/voxgig-sdk/dutch-customer-data-sdk/core"
	"github.com/voxgig-sdk/dutch-customer-data-sdk/entity"
	"github.com/voxgig-sdk/dutch-customer-data-sdk/feature"
	_ "github.com/voxgig-sdk/dutch-customer-data-sdk/utility"
)

// Type aliases preserve external API.
type DutchCustomerDataSDK = core.DutchCustomerDataSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type DutchCustomerDataEntity = core.DutchCustomerDataEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type DutchCustomerDataError = core.DutchCustomerDataError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewEuApIEntityFunc = func(client *core.DutchCustomerDataSDK, entopts map[string]any) core.DutchCustomerDataEntity {
		return entity.NewEuApIEntity(client, entopts)
	}
	core.NewGlobalApIEntityFunc = func(client *core.DutchCustomerDataSDK, entopts map[string]any) core.DutchCustomerDataEntity {
		return entity.NewGlobalApIEntity(client, entopts)
	}
	core.NewNetherlandsApIEntityFunc = func(client *core.DutchCustomerDataSDK, entopts map[string]any) core.DutchCustomerDataEntity {
		return entity.NewNetherlandsApIEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewDutchCustomerDataSDK = core.NewDutchCustomerDataSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
