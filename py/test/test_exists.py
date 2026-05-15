# ProjectName SDK exists test

import pytest
from dutchcustomerdata_sdk import DutchCustomerDataSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = DutchCustomerDataSDK.test(None, None)
        assert testsdk is not None
