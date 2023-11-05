package entity

type Place struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location GeoPoint `json:"location"`
}

type GeoPoint struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}
