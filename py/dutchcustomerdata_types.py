# Typed models for the DutchCustomerData SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class EuApI:
    buyer: Optional[str] = None
    buyer_country: Optional[str] = None
    contract_nature: Optional[str] = None
    html: Optional[str] = None
    id: Optional[str] = None
    link: Optional[str] = None
    notice_type: Optional[str] = None
    official_language: Optional[str] = None
    pdf: Optional[str] = None
    place_of_performance: Optional[str] = None
    procedure_type: Optional[str] = None
    publication_date: Optional[str] = None
    status: Optional[str] = None
    title: Optional[str] = None
    vat: Optional[dict] = None


@dataclass
class EuApILoadMatch:
    buyer: Optional[str] = None
    buyer_country: Optional[str] = None
    contract_nature: Optional[str] = None
    html: Optional[str] = None
    id: Optional[str] = None
    link: Optional[str] = None
    notice_type: Optional[str] = None
    official_language: Optional[str] = None
    pdf: Optional[str] = None
    place_of_performance: Optional[str] = None
    procedure_type: Optional[str] = None
    publication_date: Optional[str] = None
    status: Optional[str] = None
    title: Optional[str] = None
    vat: Optional[dict] = None


@dataclass
class EuApIListMatch:
    buyer: Optional[str] = None
    buyer_country: Optional[str] = None
    contract_nature: Optional[str] = None
    html: Optional[str] = None
    id: Optional[str] = None
    link: Optional[str] = None
    notice_type: Optional[str] = None
    official_language: Optional[str] = None
    pdf: Optional[str] = None
    place_of_performance: Optional[str] = None
    procedure_type: Optional[str] = None
    publication_date: Optional[str] = None
    status: Optional[str] = None
    title: Optional[str] = None
    vat: Optional[dict] = None


@dataclass
class GlobalApI:
    addition: Optional[str] = None
    admin1: Optional[str] = None
    admin2: Optional[str] = None
    admin3: Optional[str] = None
    bic: Optional[dict] = None
    city: Optional[str] = None
    currency: Optional[dict] = None
    date: Optional[str] = None
    dns: Optional[dict] = None
    email: Optional[dict] = None
    found: Optional[int] = None
    freeformaddress: Optional[str] = None
    from_currency: Optional[str] = None
    iban: Optional[dict] = None
    ip: Optional[dict] = None
    lat: Optional[float] = None
    lei: Optional[dict] = None
    letter: Optional[str] = None
    lon: Optional[float] = None
    municipality: Optional[str] = None
    number: Optional[int] = None
    password: Optional[dict] = None
    phone: Optional[dict] = None
    population: Optional[int] = None
    postcode: Optional[str] = None
    province: Optional[str] = None
    province_code: Optional[str] = None
    score: Optional[float] = None
    status: Optional[str] = None
    street: Optional[str] = None
    type: Optional[str] = None
    url: Optional[dict] = None
    webrank: Optional[dict] = None


@dataclass
class GlobalApILoadMatch:
    addition: Optional[str] = None
    admin1: Optional[str] = None
    admin2: Optional[str] = None
    admin3: Optional[str] = None
    bic: Optional[dict] = None
    city: Optional[str] = None
    currency: Optional[dict] = None
    date: Optional[str] = None
    dns: Optional[dict] = None
    email: Optional[dict] = None
    found: Optional[int] = None
    freeformaddress: Optional[str] = None
    from_currency: Optional[str] = None
    iban: Optional[dict] = None
    ip: Optional[dict] = None
    lat: Optional[float] = None
    lei: Optional[dict] = None
    letter: Optional[str] = None
    lon: Optional[float] = None
    municipality: Optional[str] = None
    number: Optional[int] = None
    password: Optional[dict] = None
    phone: Optional[dict] = None
    population: Optional[int] = None
    postcode: Optional[str] = None
    province: Optional[str] = None
    province_code: Optional[str] = None
    score: Optional[float] = None
    status: Optional[str] = None
    street: Optional[str] = None
    type: Optional[str] = None
    url: Optional[dict] = None
    webrank: Optional[dict] = None


@dataclass
class GlobalApIListMatch:
    addition: Optional[str] = None
    admin1: Optional[str] = None
    admin2: Optional[str] = None
    admin3: Optional[str] = None
    bic: Optional[dict] = None
    city: Optional[str] = None
    currency: Optional[dict] = None
    date: Optional[str] = None
    dns: Optional[dict] = None
    email: Optional[dict] = None
    found: Optional[int] = None
    freeformaddress: Optional[str] = None
    from_currency: Optional[str] = None
    iban: Optional[dict] = None
    ip: Optional[dict] = None
    lat: Optional[float] = None
    lei: Optional[dict] = None
    letter: Optional[str] = None
    lon: Optional[float] = None
    municipality: Optional[str] = None
    number: Optional[int] = None
    password: Optional[dict] = None
    phone: Optional[dict] = None
    population: Optional[int] = None
    postcode: Optional[str] = None
    province: Optional[str] = None
    province_code: Optional[str] = None
    score: Optional[float] = None
    status: Optional[str] = None
    street: Optional[str] = None
    type: Optional[str] = None
    url: Optional[dict] = None
    webrank: Optional[dict] = None


@dataclass
class GlobalApICreateData:
    addition: Optional[str] = None
    admin1: Optional[str] = None
    admin2: Optional[str] = None
    admin3: Optional[str] = None
    bic: Optional[dict] = None
    city: Optional[str] = None
    currency: Optional[dict] = None
    date: Optional[str] = None
    dns: Optional[dict] = None
    email: Optional[dict] = None
    found: Optional[int] = None
    freeformaddress: Optional[str] = None
    from_currency: Optional[str] = None
    iban: Optional[dict] = None
    ip: Optional[dict] = None
    lat: Optional[float] = None
    lei: Optional[dict] = None
    letter: Optional[str] = None
    lon: Optional[float] = None
    municipality: Optional[str] = None
    number: Optional[int] = None
    password: Optional[dict] = None
    phone: Optional[dict] = None
    population: Optional[int] = None
    postcode: Optional[str] = None
    province: Optional[str] = None
    province_code: Optional[str] = None
    score: Optional[float] = None
    status: Optional[str] = None
    street: Optional[str] = None
    type: Optional[str] = None
    url: Optional[dict] = None
    webrank: Optional[dict] = None


@dataclass
class NetherlandsApI:
    active: Optional[int] = None
    addition: Optional[str] = None
    city: Optional[str] = None
    coc: Optional[str] = None
    construction_year: Optional[int] = None
    floor_area: Optional[int] = None
    freeformaddress: Optional[str] = None
    id: Optional[str] = None
    lat: Optional[float] = None
    letter: Optional[str] = None
    lon: Optional[float] = None
    municipality: Optional[str] = None
    name: Optional[str] = None
    number: Optional[str] = None
    postcode: Optional[str] = None
    province: Optional[str] = None
    province_code: Optional[str] = None
    purpose: Optional[str] = None
    street: Optional[str] = None
    type: Optional[str] = None
    vestiging: Optional[str] = None


@dataclass
class NetherlandsApIListMatch:
    active: Optional[int] = None
    addition: Optional[str] = None
    city: Optional[str] = None
    coc: Optional[str] = None
    construction_year: Optional[int] = None
    floor_area: Optional[int] = None
    freeformaddress: Optional[str] = None
    id: Optional[str] = None
    lat: Optional[float] = None
    letter: Optional[str] = None
    lon: Optional[float] = None
    municipality: Optional[str] = None
    name: Optional[str] = None
    number: Optional[str] = None
    postcode: Optional[str] = None
    province: Optional[str] = None
    province_code: Optional[str] = None
    purpose: Optional[str] = None
    street: Optional[str] = None
    type: Optional[str] = None
    vestiging: Optional[str] = None

