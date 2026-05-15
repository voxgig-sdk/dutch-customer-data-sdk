<?php
declare(strict_types=1);

// DutchCustomerData SDK utility: result_body

class DutchCustomerDataResultBody
{
    public static function call(DutchCustomerDataContext $ctx): ?DutchCustomerDataResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
