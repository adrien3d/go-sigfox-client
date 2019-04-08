package devices

import "github.com/adrien3d/go-sigfox-client/models"

type DeviceLocation struct {
	Time      int64   `json:"time"`
	Valid     bool    `json:"valid"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Radius    int32   `json:"radius"`
	Source    int8    `json:"source"`
}

type DeviceLocations struct {
	Locations []DeviceLocation `json:"data"`
	Paging    models.NextURL   `json:"paging"`
}
