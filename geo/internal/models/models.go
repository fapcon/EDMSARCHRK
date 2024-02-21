package models

type Data struct {
	Email string
	Phone string
}

type AddressSearch []AddressSearchElement

type AddressSearchEl struct {
	Result string `json:"result"`
	GeoLat string `json:"lat"`
	GeoLon string `json:"lon"`
}

type AddressSearchElement struct {
	Source               string      `json:"source"`
	Result               string      `json:"result"`
	PostalCode           string      `json:"postal_code"`
	Country              string      `json:"country"`
	CountryISOCode       string      `json:"country_iso_code"`
	FederalDistrict      string      `json:"federal_district"`
	RegionFiasID         string      `json:"region_fias_id"`
	RegionKladrID        string      `json:"region_kladr_id"`
	RegionISOCode        string      `json:"region_iso_code"`
	RegionWithType       string      `json:"region_with_type"`
	RegionType           string      `json:"region_type"`
	RegionTypeFull       string      `json:"region_type_full"`
	Region               string      `json:"region"`
	AreaFiasID           interface{} `json:"area_fias_id"`
	AreaKladrID          interface{} `json:"area_kladr_id"`
	AreaWithType         interface{} `json:"area_with_type"`
	AreaType             interface{} `json:"area_type"`
	AreaTypeFull         interface{} `json:"area_type_full"`
	Area                 interface{} `json:"area"`
	CityFiasID           interface{} `json:"city_fias_id"`
	CityKladrID          interface{} `json:"city_kladr_id"`
	CityWithType         interface{} `json:"city_with_type"`
	CityType             interface{} `json:"city_type"`
	CityTypeFull         interface{} `json:"city_type_full"`
	City                 interface{} `json:"city"`
	CityArea             string      `json:"city_area"`
	CityDistrictFiasID   interface{} `json:"city_district_fias_id"`
	CityDistrictKladrID  interface{} `json:"city_district_kladr_id"`
	CityDistrictWithType string      `json:"city_district_with_type"`
	CityDistrictType     string      `json:"city_district_type"`
	CityDistrictTypeFull string      `json:"city_district_type_full"`
	CityDistrict         string      `json:"city_district"`
	SettlementFiasID     interface{} `json:"settlement_fias_id"`
	SettlementKladrID    interface{} `json:"settlement_kladr_id"`
	SettlementWithType   interface{} `json:"settlement_with_type"`
	SettlementType       interface{} `json:"settlement_type"`
	SettlementTypeFull   interface{} `json:"settlement_type_full"`
	Settlement           interface{} `json:"settlement"`
	StreetFiasID         string      `json:"street_fias_id"`
	StreetKladrID        string      `json:"street_kladr_id"`
	StreetWithType       string      `json:"street_with_type"`
	StreetType           string      `json:"street_type"`
	StreetTypeFull       string      `json:"street_type_full"`
	Street               string      `json:"street"`
	SteadFiasID          interface{} `json:"stead_fias_id"`
	SteadKladrID         interface{} `json:"stead_kladr_id"`
	SteadCadnum          interface{} `json:"stead_cadnum"`
	SteadType            interface{} `json:"stead_type"`
	SteadTypeFull        interface{} `json:"stead_type_full"`
	Stead                interface{} `json:"stead"`
	HouseFiasID          string      `json:"house_fias_id"`
	HouseKladrID         string      `json:"house_kladr_id"`
	HouseCadnum          string      `json:"house_cadnum"`
	HouseType            string      `json:"house_type"`
	HouseTypeFull        string      `json:"house_type_full"`
	House                string      `json:"house"`
	BlockType            interface{} `json:"block_type"`
	BlockTypeFull        interface{} `json:"block_type_full"`
	Block                interface{} `json:"block"`
	Entrance             interface{} `json:"entrance"`
	Floor                interface{} `json:"floor"`
	FlatFiasID           string      `json:"flat_fias_id"`
	FlatCadnum           string      `json:"flat_cadnum"`
	FlatType             string      `json:"flat_type"`
	FlatTypeFull         string      `json:"flat_type_full"`
	Flat                 string      `json:"flat"`
	FlatArea             string      `json:"flat_area"`
	SquareMeterPrice     string      `json:"square_meter_price"`
	FlatPrice            string      `json:"flat_price"`
	PostalBox            interface{} `json:"postal_box"`
	FiasID               string      `json:"fias_id"`
	FiasCode             string      `json:"fias_code"`
	FiasLevel            string      `json:"fias_level"`
	FiasActualityState   string      `json:"fias_actuality_state"`
	KladrID              string      `json:"kladr_id"`
	CapitalMarker        string      `json:"capital_marker"`
	Okato                string      `json:"okato"`
	Oktmo                string      `json:"oktmo"`
	TaxOffice            string      `json:"tax_office"`
	TaxOfficeLegal       string      `json:"tax_office_legal"`
	Timezone             string      `json:"timezone"`
	GeoLat               string      `json:"geo_lat"`
	GeoLon               string      `json:"geo_lon"`
	BeltwayHit           string      `json:"beltway_hit"`
	BeltwayDistance      interface{} `json:"beltway_distance"`
	QcGeo                string      `json:"qc_geo"`
	QcComplete           int64       `json:"qc_complete"`
	QcHouse              int64       `json:"qc_house"`
	Qc                   int64       `json:"qc"`
	UnparsedParts        interface{} `json:"unparsed_parts"`
	Metro                []Metro     `json:"metro"`
}

type Metro struct {
	Distance float64 `json:"distance"`
	Line     string  `json:"line"`
	Name     string  `json:"name"`
}

type GeocodeResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Value             string               `json:"value"`
	UnrestrictedValue string               `json:"unrestricted_value"`
	Data              AddressSearchElement `json:"data"`
}

type ResponseGeocode struct {
	Value             string            `json:"value"`
	UnrestrictedValue string            `json:"unrestricted_value"`
	Data              []AddressSearchEl `json:"addresses"`
}

type ResponseAddress struct {
	Addresses []AddressSearchEl `json:"addresses"`
}

type SearchRequest struct {
	Query string `json:"query"`
}

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
