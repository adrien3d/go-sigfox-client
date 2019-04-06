package models

type User struct {
	Name         string   `json:"name"`
	Timezone     string   `json:"timezone"`
	Group        Group    `json:"group"`
	CreationTime int64    `json:"creationTime"`
	Id           string   `json:"id"`
	AccessToken  string   `json:"accessToken"`
	Profiles     []IdName `json:"profiles"`
}

type Users struct {
	Users  []User  `json:"data"`
	Paging NextURL `json:"paging"`
}
