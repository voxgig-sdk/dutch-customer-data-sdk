// Typed models for the DutchCustomerData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/dutch-customer-data-sdk/go/core"
)

// EuApI is the typed data model for the eu_ap_i entity.
type EuApI struct {
	Active *int `json:"active,omitempty"`
	Address *string `json:"address,omitempty"`
	Buyer *string `json:"buyer,omitempty"`
	BuyerCountry *string `json:"buyer_country,omitempty"`
	City *string `json:"city,omitempty"`
	ContractNature *string `json:"contract_nature,omitempty"`
	Country *string `json:"country,omitempty"`
	Html *string `json:"html,omitempty"`
	Id *string `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
	NoticeType *string `json:"notice_type,omitempty"`
	OfficialLanguage *string `json:"official_language,omitempty"`
	Pdf *string `json:"pdf,omitempty"`
	PlaceOfPerformance *string `json:"place_of_performance,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	ProcedureType *string `json:"procedure_type,omitempty"`
	PublicationDate *string `json:"publication_date,omitempty"`
	ResponseDate *string `json:"response_date,omitempty"`
	Title *string `json:"title,omitempty"`
	Vat *string `json:"vat,omitempty"`
}

// EuApILoadMatch is the typed request payload for EuApI.LoadTyped.
type EuApILoadMatch struct {
	Vat string `json:"vat"`
}

// EuApIListMatch is the typed request payload for EuApI.ListTyped.
type EuApIListMatch struct {
	Q string `json:"q"`
}

// GlobalApI is the typed data model for the global_ap_i entity.
type GlobalApI struct {
	Addition *string `json:"addition,omitempty"`
	Address *string `json:"address,omitempty"`
	Admin1 *string `json:"admin1,omitempty"`
	Admin2 *string `json:"admin2,omitempty"`
	Admin3 *string `json:"admin3,omitempty"`
	Bank *string `json:"bank,omitempty"`
	Bic *string `json:"bic,omitempty"`
	Browser *string `json:"browser,omitempty"`
	Builtwith *int `json:"builtwith,omitempty"`
	Carrier *string `json:"carrier,omitempty"`
	City *string `json:"city,omitempty"`
	Cloudflare *int `json:"cloudflare,omitempty"`
	Commoncrawl *int `json:"commoncrawl,omitempty"`
	ContentLength *int `json:"content_length,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Country *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Crux *int `json:"crux,omitempty"`
	DeviceFamily *string `json:"device_family,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	DeviceType *string `json:"device_type,omitempty"`
	Disposable *int `json:"disposable,omitempty"`
	DnsA *[]any `json:"dns_a,omitempty"`
	DnsMx *[]any `json:"dns_mx,omitempty"`
	DnsNs *[]any `json:"dns_ns,omitempty"`
	DnsSoa *[]any `json:"dns_soa,omitempty"`
	DnsTxt *[]any `json:"dns_txt,omitempty"`
	DnsWwwA *[]any `json:"dns_www_a,omitempty"`
	Dnsserver *string `json:"dnsserver,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Domcop *int `json:"domcop,omitempty"`
	Email *string `json:"email,omitempty"`
	Found *int `json:"found,omitempty"`
	Free *int `json:"free,omitempty"`
	Freeformaddress *string `json:"freeformaddress,omitempty"`
	Host *string `json:"host,omitempty"`
	HostType *string `json:"host_type,omitempty"`
	Hostio *int `json:"hostio,omitempty"`
	HttpCode *int `json:"http_code,omitempty"`
	Iban *string `json:"iban,omitempty"`
	IbanHuman *string `json:"iban_human,omitempty"`
	Int *string `json:"int,omitempty"`
	International *string `json:"international,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Ipint *int `json:"ipint,omitempty"`
	Ismobile *int `json:"ismobile,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
	Lei *string `json:"lei,omitempty"`
	Letter *string `json:"letter,omitempty"`
	LocalId *string `json:"local_id,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
	Mailserver *string `json:"mailserver,omitempty"`
	Majestic *int `json:"majestic,omitempty"`
	Message *string `json:"message,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	MxHost *string `json:"mx_host,omitempty"`
	MxIp *string `json:"mx_ip,omitempty"`
	Name *string `json:"name,omitempty"`
	National *string `json:"national,omitempty"`
	Number *int `json:"number,omitempty"`
	Ocid *string `json:"ocid,omitempty"`
	Pagerank *int `json:"pagerank,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Population *int `json:"population,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Province *string `json:"province,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	RedirectCount *int `json:"redirect_count,omitempty"`
	Region *string `json:"region,omitempty"`
	RegisterId *string `json:"register_id,omitempty"`
	RenewalDate *string `json:"renewal_date,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Sepa *int `json:"sepa,omitempty"`
	Spf *string `json:"spf,omitempty"`
	Status *string `json:"status,omitempty"`
	Street *string `json:"street,omitempty"`
	Success *int `json:"success,omitempty"`
	Swift *int `json:"swift,omitempty"`
	Tranco *int `json:"tranco,omitempty"`
	Type *string `json:"type,omitempty"`
	Umbrella *int `json:"umbrella,omitempty"`
	Url *string `json:"url,omitempty"`
	User *string `json:"user,omitempty"`
	UserAgent *string `json:"user_agent,omitempty"`
	Valid *int `json:"valid,omitempty"`
	Verified *bool `json:"verified,omitempty"`
	VerifiedChecksum *bool `json:"verified_checksum,omitempty"`
	Webrank *int `json:"webrank,omitempty"`
	WrongEmail *int `json:"wrong_email,omitempty"`
	WrongFormat *int `json:"wrong_format,omitempty"`
	WrongPassword *int `json:"wrong_password,omitempty"`
	WrongPhone *int `json:"wrong_phone,omitempty"`
}

// GlobalApILoadMatch is the typed request payload for GlobalApI.LoadTyped.
type GlobalApILoadMatch struct {
	Bic *string `json:"bic,omitempty"`
	Lei *string `json:"lei,omitempty"`
	LocalId *string `json:"local_id,omitempty"`
}

// GlobalApIListMatch is the typed request payload for GlobalApI.ListTyped.
type GlobalApIListMatch struct {
	City *string `json:"city,omitempty"`
	CountryCode string `json:"country_code"`
	Full *int `json:"full,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Q *string `json:"q,omitempty"`
}

// GlobalApICreateData is the typed request payload for GlobalApI.CreateTyped.
type GlobalApICreateData struct {
	Addition *string `json:"addition,omitempty"`
	Address *string `json:"address,omitempty"`
	Admin1 *string `json:"admin1,omitempty"`
	Admin2 *string `json:"admin2,omitempty"`
	Admin3 *string `json:"admin3,omitempty"`
	Bank *string `json:"bank,omitempty"`
	Bic *string `json:"bic,omitempty"`
	Browser *string `json:"browser,omitempty"`
	Builtwith *int `json:"builtwith,omitempty"`
	Carrier *string `json:"carrier,omitempty"`
	City *string `json:"city,omitempty"`
	Cloudflare *int `json:"cloudflare,omitempty"`
	Commoncrawl *int `json:"commoncrawl,omitempty"`
	ContentLength *int `json:"content_length,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Country *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Crux *int `json:"crux,omitempty"`
	DeviceFamily *string `json:"device_family,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	DeviceType *string `json:"device_type,omitempty"`
	Disposable *int `json:"disposable,omitempty"`
	DnsA *[]any `json:"dns_a,omitempty"`
	DnsMx *[]any `json:"dns_mx,omitempty"`
	DnsNs *[]any `json:"dns_ns,omitempty"`
	DnsSoa *[]any `json:"dns_soa,omitempty"`
	DnsTxt *[]any `json:"dns_txt,omitempty"`
	DnsWwwA *[]any `json:"dns_www_a,omitempty"`
	Dnsserver *string `json:"dnsserver,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Domcop *int `json:"domcop,omitempty"`
	Email *string `json:"email,omitempty"`
	Found *int `json:"found,omitempty"`
	Free *int `json:"free,omitempty"`
	Freeformaddress *string `json:"freeformaddress,omitempty"`
	Host *string `json:"host,omitempty"`
	HostType *string `json:"host_type,omitempty"`
	Hostio *int `json:"hostio,omitempty"`
	HttpCode *int `json:"http_code,omitempty"`
	Iban *string `json:"iban,omitempty"`
	IbanHuman *string `json:"iban_human,omitempty"`
	Int *string `json:"int,omitempty"`
	International *string `json:"international,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Ipint *int `json:"ipint,omitempty"`
	Ismobile *int `json:"ismobile,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
	Lei *string `json:"lei,omitempty"`
	Letter *string `json:"letter,omitempty"`
	LocalId *string `json:"local_id,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
	Mailserver *string `json:"mailserver,omitempty"`
	Majestic *int `json:"majestic,omitempty"`
	Message *string `json:"message,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	MxHost *string `json:"mx_host,omitempty"`
	MxIp *string `json:"mx_ip,omitempty"`
	Name *string `json:"name,omitempty"`
	National *string `json:"national,omitempty"`
	Number *int `json:"number,omitempty"`
	Ocid *string `json:"ocid,omitempty"`
	Pagerank *int `json:"pagerank,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Population *int `json:"population,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Province *string `json:"province,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	RedirectCount *int `json:"redirect_count,omitempty"`
	Region *string `json:"region,omitempty"`
	RegisterId *string `json:"register_id,omitempty"`
	RenewalDate *string `json:"renewal_date,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Sepa *int `json:"sepa,omitempty"`
	Spf *string `json:"spf,omitempty"`
	Status *string `json:"status,omitempty"`
	Street *string `json:"street,omitempty"`
	Success *int `json:"success,omitempty"`
	Swift *int `json:"swift,omitempty"`
	Tranco *int `json:"tranco,omitempty"`
	Type *string `json:"type,omitempty"`
	Umbrella *int `json:"umbrella,omitempty"`
	Url *string `json:"url,omitempty"`
	User *string `json:"user,omitempty"`
	UserAgent *string `json:"user_agent,omitempty"`
	Valid *int `json:"valid,omitempty"`
	Verified *bool `json:"verified,omitempty"`
	VerifiedChecksum *bool `json:"verified_checksum,omitempty"`
	Webrank *int `json:"webrank,omitempty"`
	WrongEmail *int `json:"wrong_email,omitempty"`
	WrongFormat *int `json:"wrong_format,omitempty"`
	WrongPassword *int `json:"wrong_password,omitempty"`
	WrongPhone *int `json:"wrong_phone,omitempty"`
}

// NetherlandsApI is the typed data model for the netherlands_ap_i entity.
type NetherlandsApI struct {
	Active *int `json:"active,omitempty"`
	Addition *string `json:"addition,omitempty"`
	City *string `json:"city,omitempty"`
	Coc *string `json:"coc,omitempty"`
	ConstructionYear *int `json:"construction_year,omitempty"`
	FloorArea *int `json:"floor_area,omitempty"`
	Freeformaddress *string `json:"freeformaddress,omitempty"`
	Id *string `json:"id,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
	Letter *string `json:"letter,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	Name *string `json:"name,omitempty"`
	Number *string `json:"number,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Province *string `json:"province,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	Street *string `json:"street,omitempty"`
	Type *string `json:"type,omitempty"`
	Vestiging *string `json:"vestiging,omitempty"`
}

// NetherlandsApIListMatch is the typed request payload for NetherlandsApI.ListTyped.
type NetherlandsApIListMatch struct {
	Number string `json:"number"`
	Postcode string `json:"postcode"`
	Suffix *string `json:"suffix,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
