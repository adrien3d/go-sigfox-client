package models

type SigfoxLocation struct {
	Time      int64   `json:"time"`
	Valid     bool    `json:"valid"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Radius    int32   `json:"radius"`
	Source    int8    `json:"source"`
}

type APILocations struct {
	Locations []SigfoxLocation `json:"data"`
	Paging    NextURL          `json:"paging"`
}
