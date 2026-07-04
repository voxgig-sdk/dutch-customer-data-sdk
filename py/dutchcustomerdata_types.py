# Typed models for the DutchCustomerData SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class EuApI(TypedDict, total=False):
    buyer: str
    buyer_country: str
    contract_nature: str
    html: str
    id: str
    link: str
    notice_type: str
    official_language: str
    pdf: str
    place_of_performance: str
    procedure_type: str
    publication_date: str
    status: str
    title: str
    vat: dict


class EuApILoadMatch(TypedDict, total=False):
    buyer: str
    buyer_country: str
    contract_nature: str
    html: str
    id: str
    link: str
    notice_type: str
    official_language: str
    pdf: str
    place_of_performance: str
    procedure_type: str
    publication_date: str
    status: str
    title: str
    vat: dict


class EuApIListMatch(TypedDict, total=False):
    buyer: str
    buyer_country: str
    contract_nature: str
    html: str
    id: str
    link: str
    notice_type: str
    official_language: str
    pdf: str
    place_of_performance: str
    procedure_type: str
    publication_date: str
    status: str
    title: str
    vat: dict


class GlobalApI(TypedDict, total=False):
    addition: str
    admin1: str
    admin2: str
    admin3: str
    bic: dict
    city: str
    currency: dict
    date: str
    dns: dict
    email: dict
    found: int
    freeformaddress: str
    from_currency: str
    iban: dict
    ip: dict
    lat: float
    lei: dict
    letter: str
    lon: float
    municipality: str
    number: int
    password: dict
    phone: dict
    population: int
    postcode: str
    province: str
    province_code: str
    score: float
    status: str
    street: str
    type: str
    url: dict
    webrank: dict


class GlobalApILoadMatch(TypedDict, total=False):
    addition: str
    admin1: str
    admin2: str
    admin3: str
    bic: dict
    city: str
    currency: dict
    date: str
    dns: dict
    email: dict
    found: int
    freeformaddress: str
    from_currency: str
    iban: dict
    ip: dict
    lat: float
    lei: dict
    letter: str
    lon: float
    municipality: str
    number: int
    password: dict
    phone: dict
    population: int
    postcode: str
    province: str
    province_code: str
    score: float
    status: str
    street: str
    type: str
    url: dict
    webrank: dict


class GlobalApIListMatch(TypedDict, total=False):
    addition: str
    admin1: str
    admin2: str
    admin3: str
    bic: dict
    city: str
    currency: dict
    date: str
    dns: dict
    email: dict
    found: int
    freeformaddress: str
    from_currency: str
    iban: dict
    ip: dict
    lat: float
    lei: dict
    letter: str
    lon: float
    municipality: str
    number: int
    password: dict
    phone: dict
    population: int
    postcode: str
    province: str
    province_code: str
    score: float
    status: str
    street: str
    type: str
    url: dict
    webrank: dict


class GlobalApICreateData(TypedDict, total=False):
    addition: str
    admin1: str
    admin2: str
    admin3: str
    bic: dict
    city: str
    currency: dict
    date: str
    dns: dict
    email: dict
    found: int
    freeformaddress: str
    from_currency: str
    iban: dict
    ip: dict
    lat: float
    lei: dict
    letter: str
    lon: float
    municipality: str
    number: int
    password: dict
    phone: dict
    population: int
    postcode: str
    province: str
    province_code: str
    score: float
    status: str
    street: str
    type: str
    url: dict
    webrank: dict


class NetherlandsApI(TypedDict, total=False):
    active: int
    addition: str
    city: str
    coc: str
    construction_year: int
    floor_area: int
    freeformaddress: str
    id: str
    lat: float
    letter: str
    lon: float
    municipality: str
    name: str
    number: str
    postcode: str
    province: str
    province_code: str
    purpose: str
    street: str
    type: str
    vestiging: str


class NetherlandsApIListMatch(TypedDict, total=False):
    active: int
    addition: str
    city: str
    coc: str
    construction_year: int
    floor_area: int
    freeformaddress: str
    id: str
    lat: float
    letter: str
    lon: float
    municipality: str
    name: str
    number: str
    postcode: str
    province: str
    province_code: str
    purpose: str
    street: str
    type: str
    vestiging: str
