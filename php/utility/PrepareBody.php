<?php
declare(strict_types=1);

// DutchCustomerData SDK utility: prepare_body

class DutchCustomerDataPrepareBody
{
    public static function call(DutchCustomerDataContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
