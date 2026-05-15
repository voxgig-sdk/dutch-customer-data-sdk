<?php
declare(strict_types=1);

// DutchCustomerData SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

DutchCustomerDataUtility::setRegistrar(function (DutchCustomerDataUtility $u): void {
    $u->clean = [DutchCustomerDataClean::class, 'call'];
    $u->done = [DutchCustomerDataDone::class, 'call'];
    $u->make_error = [DutchCustomerDataMakeError::class, 'call'];
    $u->feature_add = [DutchCustomerDataFeatureAdd::class, 'call'];
    $u->feature_hook = [DutchCustomerDataFeatureHook::class, 'call'];
    $u->feature_init = [DutchCustomerDataFeatureInit::class, 'call'];
    $u->fetcher = [DutchCustomerDataFetcher::class, 'call'];
    $u->make_fetch_def = [DutchCustomerDataMakeFetchDef::class, 'call'];
    $u->make_context = [DutchCustomerDataMakeContext::class, 'call'];
    $u->make_options = [DutchCustomerDataMakeOptions::class, 'call'];
    $u->make_request = [DutchCustomerDataMakeRequest::class, 'call'];
    $u->make_response = [DutchCustomerDataMakeResponse::class, 'call'];
    $u->make_result = [DutchCustomerDataMakeResult::class, 'call'];
    $u->make_point = [DutchCustomerDataMakePoint::class, 'call'];
    $u->make_spec = [DutchCustomerDataMakeSpec::class, 'call'];
    $u->make_url = [DutchCustomerDataMakeUrl::class, 'call'];
    $u->param = [DutchCustomerDataParam::class, 'call'];
    $u->prepare_auth = [DutchCustomerDataPrepareAuth::class, 'call'];
    $u->prepare_body = [DutchCustomerDataPrepareBody::class, 'call'];
    $u->prepare_headers = [DutchCustomerDataPrepareHeaders::class, 'call'];
    $u->prepare_method = [DutchCustomerDataPrepareMethod::class, 'call'];
    $u->prepare_params = [DutchCustomerDataPrepareParams::class, 'call'];
    $u->prepare_path = [DutchCustomerDataPreparePath::class, 'call'];
    $u->prepare_query = [DutchCustomerDataPrepareQuery::class, 'call'];
    $u->result_basic = [DutchCustomerDataResultBasic::class, 'call'];
    $u->result_body = [DutchCustomerDataResultBody::class, 'call'];
    $u->result_headers = [DutchCustomerDataResultHeaders::class, 'call'];
    $u->transform_request = [DutchCustomerDataTransformRequest::class, 'call'];
    $u->transform_response = [DutchCustomerDataTransformResponse::class, 'call'];
});
