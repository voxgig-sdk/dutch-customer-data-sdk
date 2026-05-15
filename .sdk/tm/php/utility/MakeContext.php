<?php
declare(strict_types=1);

// DutchCustomerData SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class DutchCustomerDataMakeContext
{
    public static function call(array $ctxmap, ?DutchCustomerDataContext $basectx): DutchCustomerDataContext
    {
        return new DutchCustomerDataContext($ctxmap, $basectx);
    }
}
