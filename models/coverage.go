package models

type Location struct {
	Lat float64 `json:"lat" form:"lat"`
	Lng float64 `json:"lng" form:"lng"`
}

type CoveragesRequest struct {
	Locations []Location `json:"locations"`
	Radius    int        `json:"radius,omitempty"`
	GroupID   string     `json:"groupId,omitempty"`
}

type CoveragesResponse struct {
	Data []struct {
		Location
		Coverage
	} `json:"data"`
}

type CoverageQueryParams struct {
	Lat     float64 `json:"lat" form:"lat"`
	Lng     float64 `json:"lng" form:"lng"`
	Radius  int     `json:"radius,omitempty" form:"radius,omitempty"`
	GroupId string  `json:"groupId,omitempty" form:"groupId,omitempty"`
}

type Coverage struct {
	LocationCovered bool  `json:"locationCovered"`
	Margins         []int `json:"margins"`
}

type Redundancy struct {
	Redundancy int `json:"redundancy"`
}
