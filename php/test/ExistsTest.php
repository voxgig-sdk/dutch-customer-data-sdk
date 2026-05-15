<?php
declare(strict_types=1);

// DutchCustomerData SDK exists test

require_once __DIR__ . '/../dutchcustomerdata_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = DutchCustomerDataSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
