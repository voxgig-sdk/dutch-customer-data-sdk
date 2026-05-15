<?php
declare(strict_types=1);

// DutchCustomerData SDK utility: prepare_headers

class DutchCustomerDataPrepareHeaders
{
    public static function call(DutchCustomerDataContext $ctx): array
    {
        $options = $ctx->client->options_map();
        $headers = \Voxgig\Struct\Struct::getprop($options, 'headers');
        if (!$headers) {
            return [];
        }
        $out = \Voxgig\Struct\Struct::clone($headers);
        return is_array($out) ? $out : [];
    }
}
