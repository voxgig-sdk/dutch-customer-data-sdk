<?php
declare(strict_types=1);

// DutchCustomerData SDK utility: result_headers

class DutchCustomerDataResultHeaders
{
    public static function call(DutchCustomerDataContext $ctx): ?DutchCustomerDataResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
