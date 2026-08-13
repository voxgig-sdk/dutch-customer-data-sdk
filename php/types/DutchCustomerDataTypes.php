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
    public ?int $active = null;
    public ?string $address = null;
    public ?string $buyer = null;
    public ?string $buyer_country = null;
    public ?string $city = null;
    public ?string $contract_nature = null;
    public ?string $country = null;
    public ?string $html = null;
    public ?string $id = null;
    public ?string $link = null;
    public ?string $name = null;
    public ?string $notice_type = null;
    public ?string $official_language = null;
    public ?string $pdf = null;
    public ?string $place_of_performance = null;
    public ?string $postcode = null;
    public ?string $procedure_type = null;
    public ?string $publication_date = null;
    public ?string $response_date = null;
    public ?string $title = null;
    public ?string $vat = null;
}

/** Request payload for EuApI#load. */
class EuApILoadMatch
{
    public ?int $active = null;
    public ?string $address = null;
    public ?string $buyer = null;
    public ?string $buyer_country = null;
    public ?string $city = null;
    public ?string $contract_nature = null;
    public ?string $country = null;
    public ?string $html = null;
    public string $id;
    public ?string $link = null;
    public ?string $name = null;
    public ?string $notice_type = null;
    public ?string $official_language = null;
    public ?string $pdf = null;
    public ?string $place_of_performance = null;
    public ?string $postcode = null;
    public ?string $procedure_type = null;
    public ?string $publication_date = null;
    public ?string $response_date = null;
    public ?string $title = null;
    public ?string $vat = null;
}

/** Request payload for EuApI#list. */
class EuApIListMatch
{
    public ?int $active = null;
    public ?string $address = null;
    public ?string $buyer = null;
    public ?string $buyer_country = null;
    public ?string $city = null;
    public ?string $contract_nature = null;
    public ?string $country = null;
    public ?string $html = null;
    public ?string $id = null;
    public ?string $link = null;
    public ?string $name = null;
    public ?string $notice_type = null;
    public ?string $official_language = null;
    public ?string $pdf = null;
    public ?string $place_of_performance = null;
    public ?string $postcode = null;
    public ?string $procedure_type = null;
    public ?string $publication_date = null;
    public ?string $response_date = null;
    public ?string $title = null;
    public ?string $vat = null;
}

/** GlobalApI entity data model. */
class GlobalApI
{
    public ?string $addition = null;
    public ?string $address = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?string $bank = null;
    public ?string $bic = null;
    public ?string $browser = null;
    public ?int $builtwith = null;
    public ?string $carrier = null;
    public ?string $city = null;
    public ?int $cloudflare = null;
    public ?int $commoncrawl = null;
    public ?int $content_length = null;
    public ?string $content_type = null;
    public ?string $country = null;
    public ?string $country_code = null;
    public ?int $crux = null;
    public ?string $device_family = null;
    public ?string $device_name = null;
    public ?string $device_type = null;
    public ?int $disposable = null;
    public ?array $dns_a = null;
    public ?array $dns_mx = null;
    public ?array $dns_ns = null;
    public ?array $dns_soa = null;
    public ?array $dns_txt = null;
    public ?array $dns_www_a = null;
    public ?string $dnsserver = null;
    public ?string $domain = null;
    public ?int $domcop = null;
    public ?string $email = null;
    public ?int $found = null;
    public ?int $free = null;
    public ?string $freeformaddress = null;
    public ?string $host = null;
    public ?string $host_type = null;
    public ?int $hostio = null;
    public ?int $http_code = null;
    public ?string $iban = null;
    public ?string $iban_human = null;
    public ?string $int = null;
    public ?string $international = null;
    public ?string $ip = null;
    public ?int $ipint = null;
    public ?int $ismobile = null;
    public ?float $lat = null;
    public ?string $lei = null;
    public ?string $letter = null;
    public ?string $local_id = null;
    public ?float $lon = null;
    public ?string $mailserver = null;
    public ?int $majestic = null;
    public ?string $message = null;
    public ?string $municipality = null;
    public ?string $mx_host = null;
    public ?string $mx_ip = null;
    public ?string $name = null;
    public ?string $national = null;
    public ?int $number = null;
    public ?string $ocid = null;
    public ?int $pagerank = null;
    public ?string $platform = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?int $redirect_count = null;
    public ?string $region = null;
    public ?string $register_id = null;
    public ?string $renewal_date = null;
    public ?float $score = null;
    public ?int $sepa = null;
    public ?string $spf = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?int $success = null;
    public ?int $swift = null;
    public ?int $tranco = null;
    public ?string $type = null;
    public ?int $umbrella = null;
    public ?string $url = null;
    public ?string $user = null;
    public ?string $user_agent = null;
    public ?int $valid = null;
    public ?bool $verified = null;
    public ?bool $verified_checksum = null;
    public ?int $webrank = null;
    public ?int $wrong_email = null;
    public ?int $wrong_format = null;
    public ?int $wrong_password = null;
    public ?int $wrong_phone = null;
}

/** Request payload for GlobalApI#load. */
class GlobalApILoadMatch
{
    public ?string $addition = null;
    public ?string $address = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?string $bank = null;
    public ?string $bic = null;
    public ?string $browser = null;
    public ?int $builtwith = null;
    public ?string $carrier = null;
    public ?string $city = null;
    public ?int $cloudflare = null;
    public ?int $commoncrawl = null;
    public ?int $content_length = null;
    public ?string $content_type = null;
    public ?string $country = null;
    public ?string $country_code = null;
    public ?int $crux = null;
    public ?string $device_family = null;
    public ?string $device_name = null;
    public ?string $device_type = null;
    public ?int $disposable = null;
    public ?array $dns_a = null;
    public ?array $dns_mx = null;
    public ?array $dns_ns = null;
    public ?array $dns_soa = null;
    public ?array $dns_txt = null;
    public ?array $dns_www_a = null;
    public ?string $dnsserver = null;
    public ?string $domain = null;
    public ?int $domcop = null;
    public ?string $email = null;
    public ?int $found = null;
    public ?int $free = null;
    public ?string $freeformaddress = null;
    public ?string $host = null;
    public ?string $host_type = null;
    public ?int $hostio = null;
    public ?int $http_code = null;
    public ?string $iban = null;
    public ?string $iban_human = null;
    public ?string $int = null;
    public ?string $international = null;
    public ?string $ip = null;
    public ?int $ipint = null;
    public ?int $ismobile = null;
    public ?float $lat = null;
    public ?string $lei = null;
    public ?string $letter = null;
    public ?string $local_id = null;
    public ?float $lon = null;
    public ?string $mailserver = null;
    public ?int $majestic = null;
    public ?string $message = null;
    public ?string $municipality = null;
    public ?string $mx_host = null;
    public ?string $mx_ip = null;
    public ?string $name = null;
    public ?string $national = null;
    public ?int $number = null;
    public ?string $ocid = null;
    public ?int $pagerank = null;
    public ?string $platform = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?int $redirect_count = null;
    public ?string $region = null;
    public ?string $register_id = null;
    public ?string $renewal_date = null;
    public ?float $score = null;
    public ?int $sepa = null;
    public ?string $spf = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?int $success = null;
    public ?int $swift = null;
    public ?int $tranco = null;
    public ?string $type = null;
    public ?int $umbrella = null;
    public ?string $url = null;
    public ?string $user = null;
    public ?string $user_agent = null;
    public ?int $valid = null;
    public ?bool $verified = null;
    public ?bool $verified_checksum = null;
    public ?int $webrank = null;
    public ?int $wrong_email = null;
    public ?int $wrong_format = null;
    public ?int $wrong_password = null;
    public ?int $wrong_phone = null;
}

/** Request payload for GlobalApI#list. */
class GlobalApIListMatch
{
    public ?string $addition = null;
    public ?string $address = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?string $bank = null;
    public ?string $bic = null;
    public ?string $browser = null;
    public ?int $builtwith = null;
    public ?string $carrier = null;
    public ?string $city = null;
    public ?int $cloudflare = null;
    public ?int $commoncrawl = null;
    public ?int $content_length = null;
    public ?string $content_type = null;
    public ?string $country = null;
    public ?string $country_code = null;
    public ?int $crux = null;
    public ?string $device_family = null;
    public ?string $device_name = null;
    public ?string $device_type = null;
    public ?int $disposable = null;
    public ?array $dns_a = null;
    public ?array $dns_mx = null;
    public ?array $dns_ns = null;
    public ?array $dns_soa = null;
    public ?array $dns_txt = null;
    public ?array $dns_www_a = null;
    public ?string $dnsserver = null;
    public ?string $domain = null;
    public ?int $domcop = null;
    public ?string $email = null;
    public ?int $found = null;
    public ?int $free = null;
    public ?string $freeformaddress = null;
    public ?string $host = null;
    public ?string $host_type = null;
    public ?int $hostio = null;
    public ?int $http_code = null;
    public ?string $iban = null;
    public ?string $iban_human = null;
    public ?string $int = null;
    public ?string $international = null;
    public ?string $ip = null;
    public ?int $ipint = null;
    public ?int $ismobile = null;
    public ?float $lat = null;
    public ?string $lei = null;
    public ?string $letter = null;
    public ?string $local_id = null;
    public ?float $lon = null;
    public ?string $mailserver = null;
    public ?int $majestic = null;
    public ?string $message = null;
    public ?string $municipality = null;
    public ?string $mx_host = null;
    public ?string $mx_ip = null;
    public ?string $name = null;
    public ?string $national = null;
    public ?int $number = null;
    public ?string $ocid = null;
    public ?int $pagerank = null;
    public ?string $platform = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?int $redirect_count = null;
    public ?string $region = null;
    public ?string $register_id = null;
    public ?string $renewal_date = null;
    public ?float $score = null;
    public ?int $sepa = null;
    public ?string $spf = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?int $success = null;
    public ?int $swift = null;
    public ?int $tranco = null;
    public ?string $type = null;
    public ?int $umbrella = null;
    public ?string $url = null;
    public ?string $user = null;
    public ?string $user_agent = null;
    public ?int $valid = null;
    public ?bool $verified = null;
    public ?bool $verified_checksum = null;
    public ?int $webrank = null;
    public ?int $wrong_email = null;
    public ?int $wrong_format = null;
    public ?int $wrong_password = null;
    public ?int $wrong_phone = null;
}

/** Request payload for GlobalApI#create. */
class GlobalApICreateData
{
    public ?string $addition = null;
    public ?string $address = null;
    public ?string $admin1 = null;
    public ?string $admin2 = null;
    public ?string $admin3 = null;
    public ?string $bank = null;
    public ?string $bic = null;
    public ?string $browser = null;
    public ?int $builtwith = null;
    public ?string $carrier = null;
    public ?string $city = null;
    public ?int $cloudflare = null;
    public ?int $commoncrawl = null;
    public ?int $content_length = null;
    public ?string $content_type = null;
    public ?string $country = null;
    public ?string $country_code = null;
    public ?int $crux = null;
    public ?string $device_family = null;
    public ?string $device_name = null;
    public ?string $device_type = null;
    public ?int $disposable = null;
    public ?array $dns_a = null;
    public ?array $dns_mx = null;
    public ?array $dns_ns = null;
    public ?array $dns_soa = null;
    public ?array $dns_txt = null;
    public ?array $dns_www_a = null;
    public ?string $dnsserver = null;
    public ?string $domain = null;
    public ?int $domcop = null;
    public ?string $email = null;
    public ?int $found = null;
    public ?int $free = null;
    public ?string $freeformaddress = null;
    public ?string $host = null;
    public ?string $host_type = null;
    public ?int $hostio = null;
    public ?int $http_code = null;
    public ?string $iban = null;
    public ?string $iban_human = null;
    public ?string $int = null;
    public ?string $international = null;
    public ?string $ip = null;
    public ?int $ipint = null;
    public ?int $ismobile = null;
    public ?float $lat = null;
    public ?string $lei = null;
    public ?string $letter = null;
    public ?string $local_id = null;
    public ?float $lon = null;
    public ?string $mailserver = null;
    public ?int $majestic = null;
    public ?string $message = null;
    public ?string $municipality = null;
    public ?string $mx_host = null;
    public ?string $mx_ip = null;
    public ?string $name = null;
    public ?string $national = null;
    public ?int $number = null;
    public ?string $ocid = null;
    public ?int $pagerank = null;
    public ?string $platform = null;
    public ?int $population = null;
    public ?string $postcode = null;
    public ?string $province = null;
    public ?string $province_code = null;
    public ?int $redirect_count = null;
    public ?string $region = null;
    public ?string $register_id = null;
    public ?string $renewal_date = null;
    public ?float $score = null;
    public ?int $sepa = null;
    public ?string $spf = null;
    public ?string $status = null;
    public ?string $street = null;
    public ?int $success = null;
    public ?int $swift = null;
    public ?int $tranco = null;
    public ?string $type = null;
    public ?int $umbrella = null;
    public ?string $url = null;
    public ?string $user = null;
    public ?string $user_agent = null;
    public ?int $valid = null;
    public ?bool $verified = null;
    public ?bool $verified_checksum = null;
    public ?int $webrank = null;
    public ?int $wrong_email = null;
    public ?int $wrong_format = null;
    public ?int $wrong_password = null;
    public ?int $wrong_phone = null;
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

