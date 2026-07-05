<?php
declare(strict_types=1);

// Typed models for the DutchCustomerData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** EuApI entity data model. */
class EuApI
{
    public ?string $buyer = null;
    public ?string $buyer_country = null;
    public ?string $contract_nature = null;
    public ?string $html = null;
    public ?string $id = null;
    public ?string $link = null;
    public ?string $notice_type = null;
    public ?string $official_language = null;
    public ?string $pdf = null;
    public ?string $place_of_performance = null;
    public ?string $procedure_type = null;
    public ?string $publication_date = null;
    public ?string $status = null;
    public ?string $title = null;
    public ?array $vat = null;
}

/** Request payload for EuApI#load. */
class EuApILoadMatch
{
    public ?string $buyer = null;
    public ?string $buyer_country = null;
    public ?string $contract_nature = null;
    public ?string $html = null;
    public string $id;
    public ?string $link = null;
    public ?string $notice_type = null;
    public ?string $official_language = null;
    public ?string $pdf = null;
    public ?string $place_of_performance = null;
    public ?string $procedure_type = null;
    public ?string $publication_date = null;
    public ?string $status = null;
    public ?string $title = null;
    public ?array $vat = null;
}

/** Request payload for EuApI#list. */
class EuApIListMatch
{
    public ?string $buyer = null;
    public ?string $buyer_country = null;
    public ?string $contract_nature = null;
    public ?string $html = null;
    public ?string $id = null;
    public ?string $link = null;
    public ?string $notice_type = null;
    public ?string $official_language = null;
    public ?string $pdf = null;
    public ?string $place_of_performance = null;
    public ?string $procedure_type = null;
    public ?string $publication_date = null;
    public ?string $status = null;
    public ?string $title = null;
    public ?array $vat = null;
}

/** GlobalApI entity data model. */
class GlobalApI
{
    public ?string $addition = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?array $bic = null;
    public ?string $city = null;
    public ?array $currency = null;
    public ?string $date = null;
    public ?array $dns = null;
    public ?array $email = null;
    public ?int $found = null;
    public ?string $freeformaddress = null;
    public ?string $from_currency = null;
    public ?array $iban = null;
    public ?array $ip = null;
    public ?float $lat = null;
    public ?array $lei = null;
    public ?string $letter = null;
    public ?float $lon = null;
    public ?string $municipality = null;
    public ?int $number = null;
    public ?array $password = null;
    public ?array $phone = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?float $score = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?string $type = null;
    public ?array $url = null;
    public ?array $webrank = null;
}

/** Request payload for GlobalApI#load. */
class GlobalApILoadMatch
{
    public ?string $addition = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?array $bic = null;
    public ?string $city = null;
    public ?array $currency = null;
    public ?string $date = null;
    public ?array $dns = null;
    public ?array $email = null;
    public ?int $found = null;
    public ?string $freeformaddress = null;
    public ?string $from_currency = null;
    public ?array $iban = null;
    public ?array $ip = null;
    public ?float $lat = null;
    public ?array $lei = null;
    public ?string $letter = null;
    public ?float $lon = null;
    public ?string $municipality = null;
    public ?int $number = null;
    public ?array $password = null;
    public ?array $phone = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?float $score = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?string $type = null;
    public ?array $url = null;
    public ?array $webrank = null;
}

/** Request payload for GlobalApI#list. */
class GlobalApIListMatch
{
    public ?string $addition = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?array $bic = null;
    public ?string $city = null;
    public ?array $currency = null;
    public ?string $date = null;
    public ?array $dns = null;
    public ?array $email = null;
    public ?int $found = null;
    public ?string $freeformaddress = null;
    public ?string $from_currency = null;
    public ?array $iban = null;
    public ?array $ip = null;
    public ?float $lat = null;
    public ?array $lei = null;
    public ?string $letter = null;
    public ?float $lon = null;
    public ?string $municipality = null;
    public ?int $number = null;
    public ?array $password = null;
    public ?array $phone = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?float $score = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?string $type = null;
    public ?array $url = null;
    public ?array $webrank = null;
}

/** Request payload for GlobalApI#create. */
class GlobalApICreateData
{
    public ?string $addition = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?array $bic = null;
    public ?string $city = null;
    public ?array $currency = null;
    public ?string $date = null;
    public ?array $dns = null;
    public ?array $email = null;
    public ?int $found = null;
    public ?string $freeformaddress = null;
    public ?string $from_currency = null;
    public ?array $iban = null;
    public ?array $ip = null;
    public ?float $lat = null;
    public ?array $lei = null;
    public ?string $letter = null;
    public ?float $lon = null;
    public ?string $municipality = null;
    public ?int $number = null;
    public ?array $password = null;
    public ?array $phone = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?float $score = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?string $type = null;
    public ?array $url = null;
    public ?array $webrank = null;
}

/** NetherlandsApI entity data model. */
class NetherlandsApI
{
    public ?int $active = null;
    public ?string $addition = null;
    public ?string $city = null;
    public ?string $coc = null;
    public ?int $construction_year = null;
    public ?int $floor_area = null;
    public ?string $freeformaddress = null;
    public ?string $id = null;
    public ?float $lat = null;
    public ?string $letter = null;
    public ?float $lon = null;
    public ?string $municipality = null;
    public ?string $name = null;
    public ?string $number = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?string $purpose = null;
    public ?string $street = null;
    public ?string $type = null;
    public ?string $vestiging = null;
}

/** Request payload for NetherlandsApI#list. */
class NetherlandsApIListMatch
{
    public ?int $active = null;
    public ?string $addition = null;
    public ?string $city = null;
    public ?string $coc = null;
    public ?int $construction_year = null;
    public ?int $floor_area = null;
    public ?string $freeformaddress = null;
    public ?string $id = null;
    public ?float $lat = null;
    public ?string $letter = null;
    public ?float $lon = null;
    public ?string $municipality = null;
    public ?string $name = null;
    public ?string $number = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?string $purpose = null;
    public ?string $street = null;
    public ?string $type = null;
    public ?string $vestiging = null;
}

