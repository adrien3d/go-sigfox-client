package models

type IdName struct {
	Id   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type Group struct {
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Level int64  `json:"level,omitempty"`
}

type NextURL struct {
	SigfoxNextURL string `json:"next" bson:"next"`
}
