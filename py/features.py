# DutchCustomerData SDK feature factory

from feature.base_feature import DutchCustomerDataBaseFeature
from feature.test_feature import DutchCustomerDataTestFeature


def _make_feature(name):
    features = {
        "base": lambda: DutchCustomerDataBaseFeature(),
        "test": lambda: DutchCustomerDataTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
