package models

type GroupInfo struct {
	ID                    string   `json:"id"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Type                  int      `json:"type"`
	Timezone              string   `json:"timezone"`
	NameCI                string   `json:"nameCI"`
	Path                  []IdName `json:"path"`
	CurrentPrototypeCount int      `json:"currentPrototypeCount"`
	NetworkOperatorID     string   `json:"networkOperatorId,omitempty"`
	CountryISOAlpha3      string   `json:"countryISOAlpha3,omitempty"`
	Billable              bool     `json:"billable,omitempty"`
	TechnicalEmail        string   `json:"technicalEmail,omitempty"`
	MaxPrototypeAllowed   int      `json:"maxPrototypeAllowed,omitempty"`
}

type GroupInfos struct {
	GroupInfos []GroupInfo `json:"data"`
	Paging     NextURL     `json:"paging"`
}
