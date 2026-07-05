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
# @!attribute [rw] buyer
#   @return [String, nil]
#
# @!attribute [rw] buyer_country
#   @return [String, nil]
#
# @!attribute [rw] contract_nature
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
# @!attribute [rw] procedure_type
#   @return [String, nil]
#
# @!attribute [rw] publication_date
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] vat
#   @return [Hash, nil]
EuApI = Struct.new(
  :buyer,
  :buyer_country,
  :contract_nature,
  :html,
  :id,
  :link,
  :notice_type,
  :official_language,
  :pdf,
  :place_of_performance,
  :procedure_type,
  :publication_date,
  :status,
  :title,
  :vat,
  keyword_init: true
)

# Request payload for EuApI#load.
#
# @!attribute [rw] buyer
#   @return [String, nil]
#
# @!attribute [rw] buyer_country
#   @return [String, nil]
#
# @!attribute [rw] contract_nature
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
# @!attribute [rw] procedure_type
#   @return [String, nil]
#
# @!attribute [rw] publication_date
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] vat
#   @return [Hash, nil]
EuApILoadMatch = Struct.new(
  :buyer,
  :buyer_country,
  :contract_nature,
  :html,
  :id,
  :link,
  :notice_type,
  :official_language,
  :pdf,
  :place_of_performance,
  :procedure_type,
  :publication_date,
  :status,
  :title,
  :vat,
  keyword_init: true
)

# Request payload for EuApI#list.
#
# @!attribute [rw] buyer
#   @return [String, nil]
#
# @!attribute [rw] buyer_country
#   @return [String, nil]
#
# @!attribute [rw] contract_nature
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
# @!attribute [rw] procedure_type
#   @return [String, nil]
#
# @!attribute [rw] publication_date
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] vat
#   @return [Hash, nil]
EuApIListMatch = Struct.new(
  :buyer,
  :buyer_country,
  :contract_nature,
  :html,
  :id,
  :link,
  :notice_type,
  :official_language,
  :pdf,
  :place_of_performance,
  :procedure_type,
  :publication_date,
  :status,
  :title,
  :vat,
  keyword_init: true
)

# GlobalApI entity data model.
#
# @!attribute [rw] addition
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
# @!attribute [rw] bic
#   @return [Hash, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [Hash, nil]
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] dns
#   @return [Hash, nil]
#
# @!attribute [rw] email
#   @return [Hash, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] from_currency
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [Hash, nil]
#
# @!attribute [rw] ip
#   @return [Hash, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [Hash, nil]
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
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] password
#   @return [Hash, nil]
#
# @!attribute [rw] phone
#   @return [Hash, nil]
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
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [Hash, nil]
#
# @!attribute [rw] webrank
#   @return [Hash, nil]
GlobalApI = Struct.new(
  :addition,
  :admin1,
  :admin2,
  :admin3,
  :bic,
  :city,
  :currency,
  :date,
  :dns,
  :email,
  :found,
  :freeformaddress,
  :from_currency,
  :iban,
  :ip,
  :lat,
  :lei,
  :letter,
  :lon,
  :municipality,
  :number,
  :password,
  :phone,
  :population,
  :postcode,
  :province,
  :province_code,
  :score,
  :status,
  :street,
  :type,
  :url,
  :webrank,
  keyword_init: true
)

# Request payload for GlobalApI#load.
#
# @!attribute [rw] addition
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
# @!attribute [rw] bic
#   @return [Hash, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [Hash, nil]
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] dns
#   @return [Hash, nil]
#
# @!attribute [rw] email
#   @return [Hash, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] from_currency
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [Hash, nil]
#
# @!attribute [rw] ip
#   @return [Hash, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [Hash, nil]
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
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] password
#   @return [Hash, nil]
#
# @!attribute [rw] phone
#   @return [Hash, nil]
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
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [Hash, nil]
#
# @!attribute [rw] webrank
#   @return [Hash, nil]
GlobalApILoadMatch = Struct.new(
  :addition,
  :admin1,
  :admin2,
  :admin3,
  :bic,
  :city,
  :currency,
  :date,
  :dns,
  :email,
  :found,
  :freeformaddress,
  :from_currency,
  :iban,
  :ip,
  :lat,
  :lei,
  :letter,
  :lon,
  :municipality,
  :number,
  :password,
  :phone,
  :population,
  :postcode,
  :province,
  :province_code,
  :score,
  :status,
  :street,
  :type,
  :url,
  :webrank,
  keyword_init: true
)

# Request payload for GlobalApI#list.
#
# @!attribute [rw] addition
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
# @!attribute [rw] bic
#   @return [Hash, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [Hash, nil]
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] dns
#   @return [Hash, nil]
#
# @!attribute [rw] email
#   @return [Hash, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] from_currency
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [Hash, nil]
#
# @!attribute [rw] ip
#   @return [Hash, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [Hash, nil]
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
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] password
#   @return [Hash, nil]
#
# @!attribute [rw] phone
#   @return [Hash, nil]
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
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [Hash, nil]
#
# @!attribute [rw] webrank
#   @return [Hash, nil]
GlobalApIListMatch = Struct.new(
  :addition,
  :admin1,
  :admin2,
  :admin3,
  :bic,
  :city,
  :currency,
  :date,
  :dns,
  :email,
  :found,
  :freeformaddress,
  :from_currency,
  :iban,
  :ip,
  :lat,
  :lei,
  :letter,
  :lon,
  :municipality,
  :number,
  :password,
  :phone,
  :population,
  :postcode,
  :province,
  :province_code,
  :score,
  :status,
  :street,
  :type,
  :url,
  :webrank,
  keyword_init: true
)

# Request payload for GlobalApI#create.
#
# @!attribute [rw] addition
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
# @!attribute [rw] bic
#   @return [Hash, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [Hash, nil]
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] dns
#   @return [Hash, nil]
#
# @!attribute [rw] email
#   @return [Hash, nil]
#
# @!attribute [rw] found
#   @return [Integer, nil]
#
# @!attribute [rw] freeformaddress
#   @return [String, nil]
#
# @!attribute [rw] from_currency
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [Hash, nil]
#
# @!attribute [rw] ip
#   @return [Hash, nil]
#
# @!attribute [rw] lat
#   @return [Float, nil]
#
# @!attribute [rw] lei
#   @return [Hash, nil]
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
# @!attribute [rw] number
#   @return [Integer, nil]
#
# @!attribute [rw] password
#   @return [Hash, nil]
#
# @!attribute [rw] phone
#   @return [Hash, nil]
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
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] street
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [Hash, nil]
#
# @!attribute [rw] webrank
#   @return [Hash, nil]
GlobalApICreateData = Struct.new(
  :addition,
  :admin1,
  :admin2,
  :admin3,
  :bic,
  :city,
  :currency,
  :date,
  :dns,
  :email,
  :found,
  :freeformaddress,
  :from_currency,
  :iban,
  :ip,
  :lat,
  :lei,
  :letter,
  :lon,
  :municipality,
  :number,
  :password,
  :phone,
  :population,
  :postcode,
  :province,
  :province_code,
  :score,
  :status,
  :street,
  :type,
  :url,
  :webrank,
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

