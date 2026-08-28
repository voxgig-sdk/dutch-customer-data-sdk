// Typed models for the DutchCustomerData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface EuApI {
  active?: number
  address?: string
  buyer?: string
  buyer_country?: string
  city?: string
  contract_nature?: string
  country?: string
  html?: string
  id?: string
  link?: string
  name?: string
  notice_type?: string
  official_language?: string
  pdf?: string
  place_of_performance?: string
  postcode?: string
  procedure_type?: string
  publication_date?: string
  response_date?: string
  title?: string
  vat?: string
}

export interface EuApILoadMatch {
  vat: string
}

export interface EuApIListMatch {
  q: string
}

export interface GlobalApI {
  addition?: string
  address?: string
  admin1?: string
  admin2?: string
  admin3?: string
  bank?: string
  bic?: string
  browser?: string
  builtwith?: number
  carrier?: string
  city?: string
  cloudflare?: number
  commoncrawl?: number
  content_length?: number
  content_type?: string
  country?: string
  country_code?: string
  crux?: number
  device_family?: string
  device_name?: string
  device_type?: string
  disposable?: number
  dns_a?: any[]
  dns_mx?: any[]
  dns_ns?: any[]
  dns_soa?: any[]
  dns_txt?: any[]
  dns_www_a?: any[]
  dnsserver?: string
  domain?: string
  domcop?: number
  email?: string
  found?: number
  free?: number
  freeformaddress?: string
  host?: string
  host_type?: string
  hostio?: number
  http_code?: number
  iban?: string
  iban_human?: string
  int?: string
  international?: string
  ip?: string
  ipint?: number
  ismobile?: number
  lat?: number
  lei?: string
  letter?: string
  local_id?: string
  lon?: number
  mailserver?: string
  majestic?: number
  message?: string
  municipality?: string
  mx_host?: string
  mx_ip?: string
  name?: string
  national?: string
  number?: number
  ocid?: string
  pagerank?: number
  platform?: string
  population?: number
  postcode?: string
  province?: string
  province_code?: string
  redirect_count?: number
  region?: string
  register_id?: string
  renewal_date?: string
  score?: number
  sepa?: number
  spf?: string
  status?: string
  street?: string
  success?: number
  swift?: number
  tranco?: number
  type?: string
  umbrella?: number
  url?: string
  user?: string
  user_agent?: string
  valid?: number
  verified?: boolean
  verified_checksum?: boolean
  webrank?: number
  wrong_email?: number
  wrong_format?: number
  wrong_password?: number
  wrong_phone?: number
}

export interface GlobalApILoadMatch {
  bic?: string
  lei?: string
  local_id?: string
}

export interface GlobalApIListMatch {
  city?: string
  country_code: string
  full?: number
  postcode?: string
  q?: string
}

export interface GlobalApICreateData {
  addition?: string
  address?: string
  admin1?: string
  admin2?: string
  admin3?: string
  bank?: string
  bic?: string
  browser?: string
  builtwith?: number
  carrier?: string
  city?: string
  cloudflare?: number
  commoncrawl?: number
  content_length?: number
  content_type?: string
  country?: string
  country_code?: string
  crux?: number
  device_family?: string
  device_name?: string
  device_type?: string
  disposable?: number
  dns_a?: any[]
  dns_mx?: any[]
  dns_ns?: any[]
  dns_soa?: any[]
  dns_txt?: any[]
  dns_www_a?: any[]
  dnsserver?: string
  domain?: string
  domcop?: number
  email?: string
  found?: number
  free?: number
  freeformaddress?: string
  host?: string
  host_type?: string
  hostio?: number
  http_code?: number
  iban?: string
  iban_human?: string
  int?: string
  international?: string
  ip?: string
  ipint?: number
  ismobile?: number
  lat?: number
  lei?: string
  letter?: string
  local_id?: string
  lon?: number
  mailserver?: string
  majestic?: number
  message?: string
  municipality?: string
  mx_host?: string
  mx_ip?: string
  name?: string
  national?: string
  number?: number
  ocid?: string
  pagerank?: number
  platform?: string
  population?: number
  postcode?: string
  province?: string
  province_code?: string
  redirect_count?: number
  region?: string
  register_id?: string
  renewal_date?: string
  score?: number
  sepa?: number
  spf?: string
  status?: string
  street?: string
  success?: number
  swift?: number
  tranco?: number
  type?: string
  umbrella?: number
  url?: string
  user?: string
  user_agent?: string
  valid?: number
  verified?: boolean
  verified_checksum?: boolean
  webrank?: number
  wrong_email?: number
  wrong_format?: number
  wrong_password?: number
  wrong_phone?: number
}

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

export interface NetherlandsApIListMatch {
  number: string
  postcode: string
  suffix?: string
}

