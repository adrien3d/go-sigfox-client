package models

type UserRoles struct {
	Group   Group  `json:"group"`
	Profile IdName `json:"profile"`
}

type User struct {
	FirstName     string      `json:"firstName"`
	LastName      string      `json:"lastName"`
	Timezone      string      `json:"timezone"`
	ID            string      `json:"id"`
	Email         string      `json:"email"`
	Password      string      `json:"password"`
	Locked        bool        `json:"locked"`
	CreationTime  int         `json:"creationTime"`
	LastLoginTime int         `json:"lastLoginTime"`
	UserRoles     []UserRoles `json:"userRoles"`
}

type Users struct {
	Data   []User  `json:"data"`
	Paging NextURL `json:"paging"`
}
