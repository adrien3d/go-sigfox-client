package models

type Parameters struct {
	Duration int64 `json:"duration"`
	Nb       int64 `json:"nb"`
	Level    int64 `json:"level"`
}

type Option struct {
	Id         string     `json:"id"`
	Parameters Parameters `json:"parameters"`
}

type Territory struct {
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Type  int64  `json:"type,omitempty"`
	Level int64  `json:"level,omitempty"`
}

type Contract struct {
	Id                     string      `json:"id"`
	Name                   string      `json:"name"`
	Group                  Group       `json:"group"`
	MaxUplinkFrames        int64       `json:"maxUplinkFrames"`
	MaxDownlinkFrames      int64       `json:"maxDownlinkFrames"`
	MaxTokens              int64       `json:"maxTokens"`
	ActivationEndTime      int64       `json:"activationEndTime"`
	CommunicationEndTime   int64       `json:"communicationEndTime"`
	Bidir                  bool        `json:"bidir"`
	HighPriorityDownlink   bool        `json:"highPriorityDownlink"`
	AutomaticRenewal       bool        `json:"automaticRenewal"`
	RenewalDuration        int64       `json:"renewalDuration"`
	Options                []Option    `json:"options"`
	ContractId             string      `json:"contractId"`
	UserId                 string      `json:"userId"`
	Order                  IdName      `json:"order"`
	PricingModel           int64       `json:"pricingModel"`
	CreatedBy              string      `json:"createdBy"`
	LastEditionTime        int64       `json:"lastEditionTime"`
	LastEditedBy           string      `json:"lastEditedBy"`
	StartTime              int64       `json:"startTime"`
	Timezone               string      `json:"timezone"`
	SubscriptionPlan       int64       `json:"subscriptionPlan"`
	TokenDuration          int64       `json:"tokenDuration"`
	BlacklistedTerritories []Territory `json:"blacklistedTerritories"`
	TokensInUse            int64       `json:"tokensInUse"`
	TokensUsed             int64       `json:"tokensUsed"`
	NextContractInfo       IdName      `json:"nextContractInfo"`
}

type Contracts struct {
	Contract []Contract `json:"data"`
	Paging   NextURL    `json:"paging"`
}
