// Typed models for the DutchCustomerData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface EuApI {
  buyer?: string
  buyer_country?: string
  contract_nature?: string
  html?: string
  id?: string
  link?: string
  notice_type?: string
  official_language?: string
  pdf?: string
  place_of_performance?: string
  procedure_type?: string
  publication_date?: string
  status?: string
  title?: string
  vat?: Record<string, any>
}

export type EuApILoadMatch = Partial<EuApI>

export type EuApIListMatch = Partial<EuApI>

export interface GlobalApI {
  addition?: string
  admin1?: string
  admin2?: string
  admin3?: string
  bic?: Record<string, any>
  city?: string
  currency?: Record<string, any>
  date?: string
  dns?: Record<string, any>
  email?: Record<string, any>
  found?: number
  freeformaddress?: string
  from_currency?: string
  iban?: Record<string, any>
  ip?: Record<string, any>
  lat?: number
  lei?: Record<string, any>
  letter?: string
  lon?: number
  municipality?: string
  number?: number
  password?: Record<string, any>
  phone?: Record<string, any>
  population?: number
  postcode?: string
  province?: string
  province_code?: string
  score?: number
  status?: string
  street?: string
  type?: string
  url?: Record<string, any>
  webrank?: Record<string, any>
}

export type GlobalApILoadMatch = Partial<GlobalApI>

export type GlobalApIListMatch = Partial<GlobalApI>

export type GlobalApICreateData = Partial<GlobalApI>

export interface NetherlandsApI {
  active?: number
  addition?: string
  city?: string
  coc?: string
  construction_year?: number
  floor_area?: number
  freeformaddress?: string
  id?: string
  lat?: number
  letter?: string
  lon?: number
  municipality?: string
  name?: string
  number?: string
  postcode?: string
  province?: string
  province_code?: string
  purpose?: string
  street?: string
  type?: string
  vestiging?: string
}

export type NetherlandsApIListMatch = Partial<NetherlandsApI>

