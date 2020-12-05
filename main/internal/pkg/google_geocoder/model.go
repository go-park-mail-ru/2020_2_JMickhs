package google_geocoder

type Response struct {
	Status  string                     `json:"status"`
	Results []GoogleGeoResponseResults `json:"results"`
}

type GoogleGeoResponseResults struct {
	AddressComponents []AddressComponents `json:"address_components"`
	FormattedAddress  string              `json:"formatted_address"`
	Geometry          Geometry            `json:"geometry"`
}

type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Geometry struct {
	Location     GeoCoordinates            `json:"location"`
	LocationType string                    `json:"location_type"`
	Viewport     map[string]GeoCoordinates `json:"viewport"`
}

type GeoCoordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}
