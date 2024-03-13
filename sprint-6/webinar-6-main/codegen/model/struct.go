package smartyapi

//go:generate easyjson --all struct.go

type SmartyStreetsAPI []struct {
	InputIndex           int        `json:"input_index,omitempty"`
	CandidateIndex       int        `json:"candidate_index,omitempty"`
	DeliveryLine1        string     `json:"delivery_line_1,omitempty"`
	LastLine             string     `json:"last_line,omitempty"`
	DeliveryPointBarcode string     `json:"delivery_point_barcode,omitempty"`
	Components           Components `json:"components,omitempty"`
	Metadata             Metadata   `json:"metadata,omitempty"`
	Analysis             Analysis   `json:"analysis,omitempty"`
}

type Components struct {
	PrimaryNumber           string `json:"primary_number,omitempty"`
	StreetPredirection      string `json:"street_predirection,omitempty"`
	StreetName              string `json:"street_name,omitempty"`
	StreetSuffix            string `json:"street_suffix,omitempty"`
	CityName                string `json:"city_name,omitempty"`
	StateAbbreviation       string `json:"state_abbreviation,omitempty"`
	Zipcode                 string `json:"zipcode,omitempty"`
	Plus4Code               string `json:"plus_4_code,omitempty"`
	DeliveryPoint           string `json:"delivery_point,omitempty"`
	DeliveryPointCheckDigit string `json:"delivery_point_check_digit,omitempty"`
}
type Metadata struct {
	RecordType            string  `json:"record_type,omitempty"`
	ZipType               string  `json:"zip_type,omitempty"`
	CountyFips            string  `json:"county_fips,omitempty"`
	CountyName            string  `json:"county_name,omitempty"`
	CarrierRoute          string  `json:"carrier_route,omitempty"`
	CongressionalDistrict string  `json:"congressional_district,omitempty"`
	Rdi                   string  `json:"rdi,omitempty"`
	ElotSequence          string  `json:"elot_sequence,omitempty"`
	ElotSort              string  `json:"elot_sort,omitempty"`
	Latitude              float64 `json:"latitude,omitempty"`
	Longitude             float64 `json:"longitude,omitempty"`
	Precision             string  `json:"precision,omitempty"`
	TimeZone              string  `json:"time_zone,omitempty"`
	UtcOffset             int     `json:"utc_offset,omitempty"`
	Dst                   bool    `json:"dst,omitempty"`
}

type Analysis struct {
	DpvMatchCode string `json:"dpv_match_code,omitempty"`
	DpvFootnotes string `json:"dpv_footnotes,omitempty"`
	DpvCmra      string `json:"dpv_cmra,omitempty"`
	DpvVacant    string `json:"dpv_vacant,omitempty"`
	Active       string `json:"active,omitempty"`
}
