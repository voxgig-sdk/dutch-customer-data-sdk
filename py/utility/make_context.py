# DutchCustomerData SDK utility: make_context

from core.context import DutchCustomerDataContext


def make_context_util(ctxmap, basectx):
    return DutchCustomerDataContext(ctxmap, basectx)
