// Typed models for the DutchCustomerData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// EuApI is the typed data model for the eu_ap_i entity.
type EuApI struct {
	Buyer *string `json:"buyer,omitempty"`
	BuyerCountry *string `json:"buyer_country,omitempty"`
	ContractNature *string `json:"contract_nature,omitempty"`
	Html *string `json:"html,omitempty"`
	Id *string `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	NoticeType *string `json:"notice_type,omitempty"`
	OfficialLanguage *string `json:"official_language,omitempty"`
	Pdf *string `json:"pdf,omitempty"`
	PlaceOfPerformance *string `json:"place_of_performance,omitempty"`
	ProcedureType *string `json:"procedure_type,omitempty"`
	PublicationDate *string `json:"publication_date,omitempty"`
	Status *string `json:"status,omitempty"`
	Title *string `json:"title,omitempty"`
	Vat *map[string]any `json:"vat,omitempty"`
}

// EuApILoadMatch is the typed request payload for EuApI.LoadTyped.
type EuApILoadMatch struct {
	Buyer *string `json:"buyer,omitempty"`
	BuyerCountry *string `json:"buyer_country,omitempty"`
	ContractNature *string `json:"contract_nature,omitempty"`
	Html *string `json:"html,omitempty"`
	Id string `json:"id"`
	Link *string `json:"link,omitempty"`
	NoticeType *string `json:"notice_type,omitempty"`
	OfficialLanguage *string `json:"official_language,omitempty"`
	Pdf *string `json:"pdf,omitempty"`
	PlaceOfPerformance *string `json:"place_of_performance,omitempty"`
	ProcedureType *string `json:"procedure_type,omitempty"`
	PublicationDate *string `json:"publication_date,omitempty"`
	Status *string `json:"status,omitempty"`
	Title *string `json:"title,omitempty"`
	Vat *map[string]any `json:"vat,omitempty"`
}

// EuApIListMatch is the typed request payload for EuApI.ListTyped.
type EuApIListMatch struct {
	Buyer *string `json:"buyer,omitempty"`
	BuyerCountry *string `json:"buyer_country,omitempty"`
	ContractNature *string `json:"contract_nature,omitempty"`
	Html *string `json:"html,omitempty"`
	Id *string `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	NoticeType *string `json:"notice_type,omitempty"`
	OfficialLanguage *string `json:"official_language,omitempty"`
	Pdf *string `json:"pdf,omitempty"`
	PlaceOfPerformance *string `json:"place_of_performance,omitempty"`
	ProcedureType *string `json:"procedure_type,omitempty"`
	PublicationDate *string `json:"publication_date,omitempty"`
	Status *string `json:"status,omitempty"`
	Title *string `json:"title,omitempty"`
	Vat *map[string]any `json:"vat,omitempty"`
}

// GlobalApI is the typed data model for the global_ap_i entity.
type GlobalApI struct {
	Addition *string `json:"addition,omitempty"`
	Admin1 *string `json:"admin1,omitempty"`
	Admin2 *string `json:"admin2,omitempty"`
	Admin3 *string `json:"admin3,omitempty"`
	Bic *map[string]any `json:"bic,omitempty"`
	City *string `json:"city,omitempty"`
	Currency *map[string]any `json:"currency,omitempty"`
	Date *string `json:"date,omitempty"`
	Dns *map[string]any `json:"dns,omitempty"`
	Email *map[string]any `json:"email,omitempty"`
	Found *int `json:"found,omitempty"`
	Freeformaddress *string `json:"freeformaddress,omitempty"`
	FromCurrency *string `json:"from_currency,omitempty"`
	Iban *map[string]any `json:"iban,omitempty"`
	Ip *map[string]any `json:"ip,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
	Lei *map[string]any `json:"lei,omitempty"`
	Letter *string `json:"letter,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	Number *int `json:"number,omitempty"`
	Password *map[string]any `json:"password,omitempty"`
	Phone *map[string]any `json:"phone,omitempty"`
	Population *int `json:"population,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Province *string `json:"province,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Status *string `json:"status,omitempty"`
	Street *string `json:"street,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *map[string]any `json:"url,omitempty"`
	Webrank *map[string]any `json:"webrank,omitempty"`
}

// GlobalApILoadMatch is the typed request payload for GlobalApI.LoadTyped.
type GlobalApILoadMatch struct {
	Addition *string `json:"addition,omitempty"`
	Admin1 *string `json:"admin1,omitempty"`
	Admin2 *string `json:"admin2,omitempty"`
	Admin3 *string `json:"admin3,omitempty"`
	Bic *map[string]any `json:"bic,omitempty"`
	City *string `json:"city,omitempty"`
	Currency *map[string]any `json:"currency,omitempty"`
	Date *string `json:"date,omitempty"`
	Dns *map[string]any `json:"dns,omitempty"`
	Email *map[string]any `json:"email,omitempty"`
	Found *int `json:"found,omitempty"`
	Freeformaddress *string `json:"freeformaddress,omitempty"`
	FromCurrency *string `json:"from_currency,omitempty"`
	Iban *map[string]any `json:"iban,omitempty"`
	Ip *map[string]any `json:"ip,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
	Lei *map[string]any `json:"lei,omitempty"`
	Letter *string `json:"letter,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	Number *int `json:"number,omitempty"`
	Password *map[string]any `json:"password,omitempty"`
	Phone *map[string]any `json:"phone,omitempty"`
	Population *int `json:"population,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Province *string `json:"province,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Status *string `json:"status,omitempty"`
	Street *string `json:"street,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *map[string]any `json:"url,omitempty"`
	Webrank *map[string]any `json:"webrank,omitempty"`
}

// GlobalApIListMatch is the typed request payload for GlobalApI.ListTyped.
type GlobalApIListMatch struct {
	Addition *string `json:"addition,omitempty"`
	Admin1 *string `json:"admin1,omitempty"`
	Admin2 *string `json:"admin2,omitempty"`
	Admin3 *string `json:"admin3,omitempty"`
	Bic *map[string]any `json:"bic,omitempty"`
	City *string `json:"city,omitempty"`
	Currency *map[string]any `json:"currency,omitempty"`
	Date *string `json:"date,omitempty"`
	Dns *map[string]any `json:"dns,omitempty"`
	Email *map[string]any `json:"email,omitempty"`
	Found *int `json:"found,omitempty"`
	Freeformaddress *string `json:"freeformaddress,omitempty"`
	FromCurrency *string `json:"from_currency,omitempty"`
	Iban *map[string]any `json:"iban,omitempty"`
	Ip *map[string]any `json:"ip,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
	Lei *map[string]any `json:"lei,omitempty"`
	Letter *string `json:"letter,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	Number *int `json:"number,omitempty"`
	Password *map[string]any `json:"password,omitempty"`
	Phone *map[string]any `json:"phone,omitempty"`
	Population *int `json:"population,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Province *string `json:"province,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Status *string `json:"status,omitempty"`
	Street *string `json:"street,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *map[string]any `json:"url,omitempty"`
	Webrank *map[string]any `json:"webrank,omitempty"`
}

// GlobalApICreateData is the typed request payload for GlobalApI.CreateTyped.
type GlobalApICreateData struct {
	Addition *string `json:"addition,omitempty"`
	Admin1 *string `json:"admin1,omitempty"`
	Admin2 *string `json:"admin2,omitempty"`
	Admin3 *string `json:"admin3,omitempty"`
	Bic *map[string]any `json:"bic,omitempty"`
	City *string `json:"city,omitempty"`
	Currency *map[string]any `json:"currency,omitempty"`
	Date *string `json:"date,omitempty"`
	Dns *map[string]any `json:"dns,omitempty"`
	Email *map[string]any `json:"email,omitempty"`
	Found *int `json:"found,omitempty"`
	Freeformaddress *string `json:"freeformaddress,omitempty"`
	FromCurrency *string `json:"from_currency,omitempty"`
	Iban *map[string]any `json:"iban,omitempty"`
	Ip *map[string]any `json:"ip,omitempty"`
	Lat *float64 `json:"lat,omitempty"`
	Lei *map[string]any `json:"lei,omitempty"`
	Letter *string `json:"letter,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	Number *int `json:"number,omitempty"`
	Password *map[string]any `json:"password,omitempty"`
	Phone *map[string]any `json:"phone,omitempty"`
	Population *int `json:"population,omitempty"`
	Postcode *string `json:"postcode,omitempty"`
	Province *string `json:"province,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Status *string `json:"status,omitempty"`
	Street *string `json:"street,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *map[string]any `json:"url,omitempty"`
	Webrank *map[string]any `json:"webrank,omitempty"`
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

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
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

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
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
