package models

type Roles struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Path []IdName `json:"path"`
}

type Profile struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Group Group   `json:"group"`
	Roles []Roles `json:"roles"`
}

type Profiles struct {
	Data   []Profile `json:"data"`
	Paging NextURL   `json:"paging"`
}
