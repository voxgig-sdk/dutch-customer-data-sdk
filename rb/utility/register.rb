# DutchCustomerData SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

DutchCustomerDataUtility.registrar = ->(u) {
  u.clean = DutchCustomerDataUtilities::Clean
  u.done = DutchCustomerDataUtilities::Done
  u.make_error = DutchCustomerDataUtilities::MakeError
  u.feature_add = DutchCustomerDataUtilities::FeatureAdd
  u.feature_hook = DutchCustomerDataUtilities::FeatureHook
  u.feature_init = DutchCustomerDataUtilities::FeatureInit
  u.fetcher = DutchCustomerDataUtilities::Fetcher
  u.make_fetch_def = DutchCustomerDataUtilities::MakeFetchDef
  u.make_context = DutchCustomerDataUtilities::MakeContext
  u.make_options = DutchCustomerDataUtilities::MakeOptions
  u.make_request = DutchCustomerDataUtilities::MakeRequest
  u.make_response = DutchCustomerDataUtilities::MakeResponse
  u.make_result = DutchCustomerDataUtilities::MakeResult
  u.make_point = DutchCustomerDataUtilities::MakePoint
  u.make_spec = DutchCustomerDataUtilities::MakeSpec
  u.make_url = DutchCustomerDataUtilities::MakeUrl
  u.param = DutchCustomerDataUtilities::Param
  u.prepare_auth = DutchCustomerDataUtilities::PrepareAuth
  u.prepare_body = DutchCustomerDataUtilities::PrepareBody
  u.prepare_headers = DutchCustomerDataUtilities::PrepareHeaders
  u.prepare_method = DutchCustomerDataUtilities::PrepareMethod
  u.prepare_params = DutchCustomerDataUtilities::PrepareParams
  u.prepare_path = DutchCustomerDataUtilities::PreparePath
  u.prepare_query = DutchCustomerDataUtilities::PrepareQuery
  u.result_basic = DutchCustomerDataUtilities::ResultBasic
  u.result_body = DutchCustomerDataUtilities::ResultBody
  u.result_headers = DutchCustomerDataUtilities::ResultHeaders
  u.transform_request = DutchCustomerDataUtilities::TransformRequest
  u.transform_response = DutchCustomerDataUtilities::TransformResponse
}
