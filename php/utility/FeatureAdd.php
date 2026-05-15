<?php
declare(strict_types=1);

// DutchCustomerData SDK utility: feature_add

class DutchCustomerDataFeatureAdd
{
    public static function call(DutchCustomerDataContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
