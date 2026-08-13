# frozen_string_literal: true

# Typed models for the DutchCustomerData SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# EuApI entity data model.
#
# @!attribute [rw] active
#   @return [Integer, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] buyer
#   @return [String, nil]
#
# @!attribute [rw] buyer_country
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] contract_nature
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] html
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notice_type
#   @return [String, nil]
#
# @!attribute [rw] official_language
#   @return [String, nil]
#
# @!attribute [rw] pdf
#   @return [String, nil]
#
# @!attribute [rw] place_of_performance
#   @return [String, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] procedure_type
#   @return [String, nil]
#
# @!attribute [rw] publication_date
#   @return [String, nil]
#
# @!attribute [rw] response_date
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] vat
#   @return [String, nil]
EuApI = Struct.new(
  :active,
  :address,
  :buyer,
  :buyer_country,
  :city,
  :contract_nature,
  :country,
  :html,
  :id,
  :link,
  :name,
  :notice_type,
  :official_language,
  :pdf,
  :place_of_performance,
  :postcode,
  :procedure_type,
  :publication_date,
  :response_date,
  :title,
  :vat,
  keyword_init: true
)

# Request payload for EuApI#load.
#
# @!attribute [rw] active
#   @return [Integer, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] buyer
#   @return [String, nil]
#
# @!attribute [rw] buyer_country
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] contract_nature
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] html
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notice_type
#   @return [String, nil]
#
# @!attribute [rw] official_language
#   @return [String, nil]
#
# @!attribute [rw] pdf
#   @return [String, nil]
#
# @!attribute [rw] place_of_performance
#   @return [String, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] procedure_type
#   @return [String, nil]
#
# @!attribute [rw] publication_date
#   @return [String, nil]
#
# @!attribute [rw] response_date
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] vat
#   @return [String, nil]
EuApILoadMatch = Struct.new(
  :active,
  :address,
  :buyer,
  :buyer_country,
  :city,
  :contract_nature,
  :country,
  :html,
  :id,
  :link,
  :name,
  :notice_type,
  :official_language,
  :pdf,
  :place_of_performance,
  :postcode,
  :procedure_type,
  :publication_date,
  :response_date,
  :title,
  :vat,
  keyword_init: true
)

# Request payload for EuApI#list.
#
# @!attribute [rw] active
#   @return [Integer, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] buyer
#   @return [String, nil]
#
# @!attribute [rw] buyer_country
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] contract_nature
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] html
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notice_type
#   @return [String, nil]
#
# @!attribute [rw] official_language
#   @return [String, nil]
#
# @!attribute [rw] pdf
#   @return [String, nil]
#
# @!attribute [rw] place_of_performance
#   @return [String, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] procedure_type
#   @return [String, nil]
#
# @!attribute [rw] publication_date
#   @return [String, nil]
#
# @!attribute [rw] response_date
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] vat
#   @return [String, nil]
EuApIListMatch = Struct.new(
  :active,
  :address,
  :buyer,
  :buyer_country,
  :city,
  :contract_nature,
  :country,
  :html,
  :id,
  :link,
  :name,
  :notice_type,
  :official_language,
  :pdf,
  :place_of_performance,
  :postcode,
  :procedure_type,
  :publication_date,
  :response_date,
  :title,
  :vat,
  keyword_init: true
)

# GlobalApI entity data model.
#
# @!attribute [rw] addition
#   @return [String, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] admin1
#   @return [String, nil]
#
# @!attribute [rw] admin2
#   @return [String, nil]
#
# @!attribute [rw] admin3
#   @return [String, nil]
#
# @!attribute [rw] bank
#   @return [String, nil]
#
# @!attribute [rw] bic
#   @return [String, nil]
#
# @!attribute [rw] browser
#   @return [String, nil]
#
# @!attribute [rw] builtwith
#   @return [Integer, nil]
#
# @!attribute [rw] carrier
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] cloudflare
#   @return [Integer, nil]
#
# @!attribute [rw] commoncrawl
#   @return [Integer, nil]
#
# @!attribute [rw] content_length
#   @return [Integer, nil]
#
# @!attribute [rw] content_type
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] crux
#   @return [Integer, nil]
#
# @!attribute [rw] device_family
#   @return [String, nil]
#
# @!attribute [rw] device_name
#   @return [String, nil]
#
# @!attribute [rw] device_type
#   @return [String, nil]
#
# @!attribute [rw] disposable
#   @return [Integer, nil]
#
# @!attribute [rw] dns_a
#   @return [Array, nil]
#
# @!attribute [rw] dns_mx
#   @return [Array, nil]
#
# @!attribute [rw] dns_ns
#   @return [Array, nil]
#
# @!attribute [rw] dns_soa
#   @return [Array, nil]
#
# @!attribute [rw] dns_txt
#   @return [Array, nil]
#
# @!attribute [rw] dns_www_a
#   @return [Array, nil]
#
# @!attribute [rw] dnsserver
#   @return [String, nil]
#
# @!attribute [rw] domain
#   @return [String, nil]
#
# @!attribute [rw] domcop
#   @return [Integer, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] free
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] host
#   @return [String, nil]
#
# @!attribute [rw] host_type
#   @return [String, nil]
#
# @!attribute [rw] hostio
#   @return [Integer, nil]
#
# @!attribute [rw] http_code
#   @return [Integer, nil]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] iban_human
#   @return [String, nil]
#
# @!attribute [rw] int
#   @return [String, nil]
#
# @!attribute [rw] international
#   @return [String, nil]
#
# @!attribute [rw] ip
#   @return [String, nil]
#
# @!attribute [rw] ipint
#   @return [Integer, nil]
#
# @!attribute [rw] ismobile
#   @return [Integer, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [String, nil]
#
# @!attribute [rw] letter
#   @return [String, nil]
#
# @!attribute [rw] local_id
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [Float, nil]
#
# @!attribute [rw] mailserver
#   @return [String, nil]
#
# @!attribute [rw] majestic
#   @return [Integer, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] municipality
#   @return [String, nil]
#
# @!attribute [rw] mx_host
#   @return [String, nil]
#
# @!attribute [rw] mx_ip
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] national
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] ocid
#   @return [String, nil]
#
# @!attribute [rw] pagerank
#   @return [Integer, nil]
#
# @!attribute [rw] platform
#   @return [String, nil]
#
# @!attribute [rw] population
#   @return [Integer, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] province
#   @return [String, nil]
#
# @!attribute [rw] province_code
#   @return [String, nil]
#
# @!attribute [rw] redirect_count
#   @return [Integer, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] register_id
#   @return [String, nil]
#
# @!attribute [rw] renewal_date
#   @return [String, nil]
#
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] sepa
#   @return [Integer, nil]
#
# @!attribute [rw] spf
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] success
#   @return [Integer, nil]
#
# @!attribute [rw] swift
#   @return [Integer, nil]
#
# @!attribute [rw] tranco
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] umbrella
#   @return [Integer, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [String, nil]
#
# @!attribute [rw] user_agent
#   @return [String, nil]
#
# @!attribute [rw] valid
#   @return [Integer, nil]
#
# @!attribute [rw] verified
#   @return [Boolean, nil]
#
# @!attribute [rw] verified_checksum
#   @return [Boolean, nil]
#
# @!attribute [rw] webrank
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_email
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_format
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_password
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_phone
#   @return [Integer, nil]
GlobalApI = Struct.new(
  :addition,
  :address,
  :admin1,
  :admin2,
  :admin3,
  :bank,
  :bic,
  :browser,
  :builtwith,
  :carrier,
  :city,
  :cloudflare,
  :commoncrawl,
  :content_length,
  :content_type,
  :country,
  :country_code,
  :crux,
  :device_family,
  :device_name,
  :device_type,
  :disposable,
  :dns_a,
  :dns_mx,
  :dns_ns,
  :dns_soa,
  :dns_txt,
  :dns_www_a,
  :dnsserver,
  :domain,
  :domcop,
  :email,
  :found,
  :free,
  :freeformaddress,
  :host,
  :host_type,
  :hostio,
  :http_code,
  :iban,
  :iban_human,
  :int,
  :international,
  :ip,
  :ipint,
  :ismobile,
  :lat,
  :lei,
  :letter,
  :local_id,
  :lon,
  :mailserver,
  :majestic,
  :message,
  :municipality,
  :mx_host,
  :mx_ip,
  :name,
  :national,
  :number,
  :ocid,
  :pagerank,
  :platform,
  :population,
  :postcode,
  :province,
  :province_code,
  :redirect_count,
  :region,
  :register_id,
  :renewal_date,
  :score,
  :sepa,
  :spf,
  :status,
  :street,
  :success,
  :swift,
  :tranco,
  :type,
  :umbrella,
  :url,
  :user,
  :user_agent,
  :valid,
  :verified,
  :verified_checksum,
  :webrank,
  :wrong_email,
  :wrong_format,
  :wrong_password,
  :wrong_phone,
  keyword_init: true
)

# Request payload for GlobalApI#load.
#
# @!attribute [rw] addition
#   @return [String, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] admin1
#   @return [String, nil]
#
# @!attribute [rw] admin2
#   @return [String, nil]
#
# @!attribute [rw] admin3
#   @return [String, nil]
#
# @!attribute [rw] bank
#   @return [String, nil]
#
# @!attribute [rw] bic
#   @return [String, nil]
#
# @!attribute [rw] browser
#   @return [String, nil]
#
# @!attribute [rw] builtwith
#   @return [Integer, nil]
#
# @!attribute [rw] carrier
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] cloudflare
#   @return [Integer, nil]
#
# @!attribute [rw] commoncrawl
#   @return [Integer, nil]
#
# @!attribute [rw] content_length
#   @return [Integer, nil]
#
# @!attribute [rw] content_type
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] crux
#   @return [Integer, nil]
#
# @!attribute [rw] device_family
#   @return [String, nil]
#
# @!attribute [rw] device_name
#   @return [String, nil]
#
# @!attribute [rw] device_type
#   @return [String, nil]
#
# @!attribute [rw] disposable
#   @return [Integer, nil]
#
# @!attribute [rw] dns_a
#   @return [Array, nil]
#
# @!attribute [rw] dns_mx
#   @return [Array, nil]
#
# @!attribute [rw] dns_ns
#   @return [Array, nil]
#
# @!attribute [rw] dns_soa
#   @return [Array, nil]
#
# @!attribute [rw] dns_txt
#   @return [Array, nil]
#
# @!attribute [rw] dns_www_a
#   @return [Array, nil]
#
# @!attribute [rw] dnsserver
#   @return [String, nil]
#
# @!attribute [rw] domain
#   @return [String, nil]
#
# @!attribute [rw] domcop
#   @return [Integer, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] free
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] host
#   @return [String, nil]
#
# @!attribute [rw] host_type
#   @return [String, nil]
#
# @!attribute [rw] hostio
#   @return [Integer, nil]
#
# @!attribute [rw] http_code
#   @return [Integer, nil]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] iban_human
#   @return [String, nil]
#
# @!attribute [rw] int
#   @return [String, nil]
#
# @!attribute [rw] international
#   @return [String, nil]
#
# @!attribute [rw] ip
#   @return [String, nil]
#
# @!attribute [rw] ipint
#   @return [Integer, nil]
#
# @!attribute [rw] ismobile
#   @return [Integer, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [String, nil]
#
# @!attribute [rw] letter
#   @return [String, nil]
#
# @!attribute [rw] local_id
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [Float, nil]
#
# @!attribute [rw] mailserver
#   @return [String, nil]
#
# @!attribute [rw] majestic
#   @return [Integer, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] municipality
#   @return [String, nil]
#
# @!attribute [rw] mx_host
#   @return [String, nil]
#
# @!attribute [rw] mx_ip
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] national
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] ocid
#   @return [String, nil]
#
# @!attribute [rw] pagerank
#   @return [Integer, nil]
#
# @!attribute [rw] platform
#   @return [String, nil]
#
# @!attribute [rw] population
#   @return [Integer, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] province
#   @return [String, nil]
#
# @!attribute [rw] province_code
#   @return [String, nil]
#
# @!attribute [rw] redirect_count
#   @return [Integer, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] register_id
#   @return [String, nil]
#
# @!attribute [rw] renewal_date
#   @return [String, nil]
#
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] sepa
#   @return [Integer, nil]
#
# @!attribute [rw] spf
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] success
#   @return [Integer, nil]
#
# @!attribute [rw] swift
#   @return [Integer, nil]
#
# @!attribute [rw] tranco
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] umbrella
#   @return [Integer, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [String, nil]
#
# @!attribute [rw] user_agent
#   @return [String, nil]
#
# @!attribute [rw] valid
#   @return [Integer, nil]
#
# @!attribute [rw] verified
#   @return [Boolean, nil]
#
# @!attribute [rw] verified_checksum
#   @return [Boolean, nil]
#
# @!attribute [rw] webrank
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_email
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_format
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_password
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_phone
#   @return [Integer, nil]
GlobalApILoadMatch = Struct.new(
  :addition,
  :address,
  :admin1,
  :admin2,
  :admin3,
  :bank,
  :bic,
  :browser,
  :builtwith,
  :carrier,
  :city,
  :cloudflare,
  :commoncrawl,
  :content_length,
  :content_type,
  :country,
  :country_code,
  :crux,
  :device_family,
  :device_name,
  :device_type,
  :disposable,
  :dns_a,
  :dns_mx,
  :dns_ns,
  :dns_soa,
  :dns_txt,
  :dns_www_a,
  :dnsserver,
  :domain,
  :domcop,
  :email,
  :found,
  :free,
  :freeformaddress,
  :host,
  :host_type,
  :hostio,
  :http_code,
  :iban,
  :iban_human,
  :int,
  :international,
  :ip,
  :ipint,
  :ismobile,
  :lat,
  :lei,
  :letter,
  :local_id,
  :lon,
  :mailserver,
  :majestic,
  :message,
  :municipality,
  :mx_host,
  :mx_ip,
  :name,
  :national,
  :number,
  :ocid,
  :pagerank,
  :platform,
  :population,
  :postcode,
  :province,
  :province_code,
  :redirect_count,
  :region,
  :register_id,
  :renewal_date,
  :score,
  :sepa,
  :spf,
  :status,
  :street,
  :success,
  :swift,
  :tranco,
  :type,
  :umbrella,
  :url,
  :user,
  :user_agent,
  :valid,
  :verified,
  :verified_checksum,
  :webrank,
  :wrong_email,
  :wrong_format,
  :wrong_password,
  :wrong_phone,
  keyword_init: true
)

# Request payload for GlobalApI#list.
#
# @!attribute [rw] addition
#   @return [String, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] admin1
#   @return [String, nil]
#
# @!attribute [rw] admin2
#   @return [String, nil]
#
# @!attribute [rw] admin3
#   @return [String, nil]
#
# @!attribute [rw] bank
#   @return [String, nil]
#
# @!attribute [rw] bic
#   @return [String, nil]
#
# @!attribute [rw] browser
#   @return [String, nil]
#
# @!attribute [rw] builtwith
#   @return [Integer, nil]
#
# @!attribute [rw] carrier
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] cloudflare
#   @return [Integer, nil]
#
# @!attribute [rw] commoncrawl
#   @return [Integer, nil]
#
# @!attribute [rw] content_length
#   @return [Integer, nil]
#
# @!attribute [rw] content_type
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] crux
#   @return [Integer, nil]
#
# @!attribute [rw] device_family
#   @return [String, nil]
#
# @!attribute [rw] device_name
#   @return [String, nil]
#
# @!attribute [rw] device_type
#   @return [String, nil]
#
# @!attribute [rw] disposable
#   @return [Integer, nil]
#
# @!attribute [rw] dns_a
#   @return [Array, nil]
#
# @!attribute [rw] dns_mx
#   @return [Array, nil]
#
# @!attribute [rw] dns_ns
#   @return [Array, nil]
#
# @!attribute [rw] dns_soa
#   @return [Array, nil]
#
# @!attribute [rw] dns_txt
#   @return [Array, nil]
#
# @!attribute [rw] dns_www_a
#   @return [Array, nil]
#
# @!attribute [rw] dnsserver
#   @return [String, nil]
#
# @!attribute [rw] domain
#   @return [String, nil]
#
# @!attribute [rw] domcop
#   @return [Integer, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] free
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] host
#   @return [String, nil]
#
# @!attribute [rw] host_type
#   @return [String, nil]
#
# @!attribute [rw] hostio
#   @return [Integer, nil]
#
# @!attribute [rw] http_code
#   @return [Integer, nil]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] iban_human
#   @return [String, nil]
#
# @!attribute [rw] int
#   @return [String, nil]
#
# @!attribute [rw] international
#   @return [String, nil]
#
# @!attribute [rw] ip
#   @return [String, nil]
#
# @!attribute [rw] ipint
#   @return [Integer, nil]
#
# @!attribute [rw] ismobile
#   @return [Integer, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [String, nil]
#
# @!attribute [rw] letter
#   @return [String, nil]
#
# @!attribute [rw] local_id
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [Float, nil]
#
# @!attribute [rw] mailserver
#   @return [String, nil]
#
# @!attribute [rw] majestic
#   @return [Integer, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] municipality
#   @return [String, nil]
#
# @!attribute [rw] mx_host
#   @return [String, nil]
#
# @!attribute [rw] mx_ip
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] national
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] ocid
#   @return [String, nil]
#
# @!attribute [rw] pagerank
#   @return [Integer, nil]
#
# @!attribute [rw] platform
#   @return [String, nil]
#
# @!attribute [rw] population
#   @return [Integer, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] province
#   @return [String, nil]
#
# @!attribute [rw] province_code
#   @return [String, nil]
#
# @!attribute [rw] redirect_count
#   @return [Integer, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] register_id
#   @return [String, nil]
#
# @!attribute [rw] renewal_date
#   @return [String, nil]
#
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] sepa
#   @return [Integer, nil]
#
# @!attribute [rw] spf
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] success
#   @return [Integer, nil]
#
# @!attribute [rw] swift
#   @return [Integer, nil]
#
# @!attribute [rw] tranco
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] umbrella
#   @return [Integer, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [String, nil]
#
# @!attribute [rw] user_agent
#   @return [String, nil]
#
# @!attribute [rw] valid
#   @return [Integer, nil]
#
# @!attribute [rw] verified
#   @return [Boolean, nil]
#
# @!attribute [rw] verified_checksum
#   @return [Boolean, nil]
#
# @!attribute [rw] webrank
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_email
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_format
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_password
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_phone
#   @return [Integer, nil]
GlobalApIListMatch = Struct.new(
  :addition,
  :address,
  :admin1,
  :admin2,
  :admin3,
  :bank,
  :bic,
  :browser,
  :builtwith,
  :carrier,
  :city,
  :cloudflare,
  :commoncrawl,
  :content_length,
  :content_type,
  :country,
  :country_code,
  :crux,
  :device_family,
  :device_name,
  :device_type,
  :disposable,
  :dns_a,
  :dns_mx,
  :dns_ns,
  :dns_soa,
  :dns_txt,
  :dns_www_a,
  :dnsserver,
  :domain,
  :domcop,
  :email,
  :found,
  :free,
  :freeformaddress,
  :host,
  :host_type,
  :hostio,
  :http_code,
  :iban,
  :iban_human,
  :int,
  :international,
  :ip,
  :ipint,
  :ismobile,
  :lat,
  :lei,
  :letter,
  :local_id,
  :lon,
  :mailserver,
  :majestic,
  :message,
  :municipality,
  :mx_host,
  :mx_ip,
  :name,
  :national,
  :number,
  :ocid,
  :pagerank,
  :platform,
  :population,
  :postcode,
  :province,
  :province_code,
  :redirect_count,
  :region,
  :register_id,
  :renewal_date,
  :score,
  :sepa,
  :spf,
  :status,
  :street,
  :success,
  :swift,
  :tranco,
  :type,
  :umbrella,
  :url,
  :user,
  :user_agent,
  :valid,
  :verified,
  :verified_checksum,
  :webrank,
  :wrong_email,
  :wrong_format,
  :wrong_password,
  :wrong_phone,
  keyword_init: true
)

# Request payload for GlobalApI#create.
#
# @!attribute [rw] addition
#   @return [String, nil]
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] admin1
#   @return [String, nil]
#
# @!attribute [rw] admin2
#   @return [String, nil]
#
# @!attribute [rw] admin3
#   @return [String, nil]
#
# @!attribute [rw] bank
#   @return [String, nil]
#
# @!attribute [rw] bic
#   @return [String, nil]
#
# @!attribute [rw] browser
#   @return [String, nil]
#
# @!attribute [rw] builtwith
#   @return [Integer, nil]
#
# @!attribute [rw] carrier
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] cloudflare
#   @return [Integer, nil]
#
# @!attribute [rw] commoncrawl
#   @return [Integer, nil]
#
# @!attribute [rw] content_length
#   @return [Integer, nil]
#
# @!attribute [rw] content_type
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] crux
#   @return [Integer, nil]
#
# @!attribute [rw] device_family
#   @return [String, nil]
#
# @!attribute [rw] device_name
#   @return [String, nil]
#
# @!attribute [rw] device_type
#   @return [String, nil]
#
# @!attribute [rw] disposable
#   @return [Integer, nil]
#
# @!attribute [rw] dns_a
#   @return [Array, nil]
#
# @!attribute [rw] dns_mx
#   @return [Array, nil]
#
# @!attribute [rw] dns_ns
#   @return [Array, nil]
#
# @!attribute [rw] dns_soa
#   @return [Array, nil]
#
# @!attribute [rw] dns_txt
#   @return [Array, nil]
#
# @!attribute [rw] dns_www_a
#   @return [Array, nil]
#
# @!attribute [rw] dnsserver
#   @return [String, nil]
#
# @!attribute [rw] domain
#   @return [String, nil]
#
# @!attribute [rw] domcop
#   @return [Integer, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] free
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] host
#   @return [String, nil]
#
# @!attribute [rw] host_type
#   @return [String, nil]
#
# @!attribute [rw] hostio
#   @return [Integer, nil]
#
# @!attribute [rw] http_code
#   @return [Integer, nil]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] iban_human
#   @return [String, nil]
#
# @!attribute [rw] int
#   @return [String, nil]
#
# @!attribute [rw] international
#   @return [String, nil]
#
# @!attribute [rw] ip
#   @return [String, nil]
#
# @!attribute [rw] ipint
#   @return [Integer, nil]
#
# @!attribute [rw] ismobile
#   @return [Integer, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [String, nil]
#
# @!attribute [rw] letter
#   @return [String, nil]
#
# @!attribute [rw] local_id
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [Float, nil]
#
# @!attribute [rw] mailserver
#   @return [String, nil]
#
# @!attribute [rw] majestic
#   @return [Integer, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] municipality
#   @return [String, nil]
#
# @!attribute [rw] mx_host
#   @return [String, nil]
#
# @!attribute [rw] mx_ip
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] national
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] ocid
#   @return [String, nil]
#
# @!attribute [rw] pagerank
#   @return [Integer, nil]
#
# @!attribute [rw] platform
#   @return [String, nil]
#
# @!attribute [rw] population
#   @return [Integer, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] province
#   @return [String, nil]
#
# @!attribute [rw] province_code
#   @return [String, nil]
#
# @!attribute [rw] redirect_count
#   @return [Integer, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] register_id
#   @return [String, nil]
#
# @!attribute [rw] renewal_date
#   @return [String, nil]
#
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] sepa
#   @return [Integer, nil]
#
# @!attribute [rw] spf
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] success
#   @return [Integer, nil]
#
# @!attribute [rw] swift
#   @return [Integer, nil]
#
# @!attribute [rw] tranco
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] umbrella
#   @return [Integer, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [String, nil]
#
# @!attribute [rw] user_agent
#   @return [String, nil]
#
# @!attribute [rw] valid
#   @return [Integer, nil]
#
# @!attribute [rw] verified
#   @return [Boolean, nil]
#
# @!attribute [rw] verified_checksum
#   @return [Boolean, nil]
#
# @!attribute [rw] webrank
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_email
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_format
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_password
#   @return [Integer, nil]
#
# @!attribute [rw] wrong_phone
#   @return [Integer, nil]
GlobalApICreateData = Struct.new(
  :addition,
  :address,
  :admin1,
  :admin2,
  :admin3,
  :bank,
  :bic,
  :browser,
  :builtwith,
  :carrier,
  :city,
  :cloudflare,
  :commoncrawl,
  :content_length,
  :content_type,
  :country,
  :country_code,
  :crux,
  :device_family,
  :device_name,
  :device_type,
  :disposable,
  :dns_a,
  :dns_mx,
  :dns_ns,
  :dns_soa,
  :dns_txt,
  :dns_www_a,
  :dnsserver,
  :domain,
  :domcop,
  :email,
  :found,
  :free,
  :freeformaddress,
  :host,
  :host_type,
  :hostio,
  :http_code,
  :iban,
  :iban_human,
  :int,
  :international,
  :ip,
  :ipint,
  :ismobile,
  :lat,
  :lei,
  :letter,
  :local_id,
  :lon,
  :mailserver,
  :majestic,
  :message,
  :municipality,
  :mx_host,
  :mx_ip,
  :name,
  :national,
  :number,
  :ocid,
  :pagerank,
  :platform,
  :population,
  :postcode,
  :province,
  :province_code,
  :redirect_count,
  :region,
  :register_id,
  :renewal_date,
  :score,
  :sepa,
  :spf,
  :status,
  :street,
  :success,
  :swift,
  :tranco,
  :type,
  :umbrella,
  :url,
  :user,
  :user_agent,
  :valid,
  :verified,
  :verified_checksum,
  :webrank,
  :wrong_email,
  :wrong_format,
  :wrong_password,
  :wrong_phone,
  keyword_init: true
)

# NetherlandsApI entity data model.
#
# @!attribute [rw] active
#   @return [Integer, nil]
#
# @!attribute [rw] addition
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] coc
#   @return [String, nil]
#
# @!attribute [rw] construction_year
#   @return [Integer, nil]
#
# @!attribute [rw] floor_area
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] letter
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [Float, nil]
#
# @!attribute [rw] municipality
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [String, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] province
#   @return [String, nil]
#
# @!attribute [rw] province_code
#   @return [String, nil]
#
# @!attribute [rw] purpose
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] vestiging
#   @return [String, nil]
NetherlandsApI = Struct.new(
  :active,
  :addition,
  :city,
  :coc,
  :construction_year,
  :floor_area,
  :freeformaddress,
  :id,
  :lat,
  :letter,
  :lon,
  :municipality,
  :name,
  :number,
  :postcode,
  :province,
  :province_code,
  :purpose,
  :street,
  :type,
  :vestiging,
  keyword_init: true
)

# Request payload for NetherlandsApI#list.
#
# @!attribute [rw] active
#   @return [Integer, nil]
#
# @!attribute [rw] addition
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] coc
#   @return [String, nil]
#
# @!attribute [rw] construction_year
#   @return [Integer, nil]
#
# @!attribute [rw] floor_area
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] letter
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [Float, nil]
#
# @!attribute [rw] municipality
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [String, nil]
#
# @!attribute [rw] postcode
#   @return [String, nil]
#
# @!attribute [rw] province
#   @return [String, nil]
#
# @!attribute [rw] province_code
#   @return [String, nil]
#
# @!attribute [rw] purpose
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] vestiging
#   @return [String, nil]
NetherlandsApIListMatch = Struct.new(
  :active,
  :addition,
  :city,
  :coc,
  :construction_year,
  :floor_area,
  :freeformaddress,
  :id,
  :lat,
  :letter,
  :lon,
  :municipality,
  :name,
  :number,
  :postcode,
  :province,
  :province_code,
  :purpose,
  :street,
  :type,
  :vestiging,
  keyword_init: true
)

