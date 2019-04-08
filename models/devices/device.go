package devices

import "github.com/adrien3d/go-sigfox-client/models"

type Location struct {
	Latitude  float64 `json:"lat" form:"lat"`
	Longitude float64 `json:"lng" form:"lng"`
}

type ComputedLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Radius    int32   `json:"radius"`
	Source    int8    `json:"source"`
}

type Token struct {
	State              int64  `json:"state"`
	DetailMessage      string `json:"detailMessage"`
	End                int64  `json:"end"`
	UnsubscriptionTime int64  `json:"unsubscriptionTime,omitempty"`
	FreeMessages       int64  `json:"freeMessages,omitempty"`
	FreeMessagesSent   int64  `json:"freeMessagesSent,omitempty"`
}

type DeviceInformation struct {
	Id                     string           `json:"id"`
	Name                   string           `json:"name"`
	SequenceNumber         int64            `json:"sequenceNumber"`
	TrashSequenceNumber    int64            `json:"trashSequenceNumber"`
	LastCom                int64            `json:"lastCom"`
	Pac                    string           `json:"pac"`
	Location               Location         `json:"location"`
	LastComputedLocation   ComputedLocation `json:"lastComputedLocation"`
	Group                  models.Group     `json:"group"`
	DeviceType             models.IdName    `json:"deviceType"`
	AverageSnr             int64            `json:"averageSnr,omitempty"`
	AverageRssi            int64            `json:"averageRssi,omitempty"`
	ActivationTime         int64            `json:"activationTime"`
	Contract               models.IdName    `json:"contract"`
	Token                  Token            `json:"token"`
	CreatedBy              string           `json:"createdBy"`
	LastEditionTime        int64            `json:"lastEditionTime"`
	LastEditedBy           string           `json:"lastEditedBy"`
	UnsubscriptionTime     int64            `json:"unsubscriptionTime"`
	CreationTime           int64            `json:"creationTime"`
	ModemCertificate       models.IdName    `json:"modemCertificate"`
	ProductCertificate     models.IdName    `json:"productCertificate"`
	State                  int64            `json:"state"`
	ComState               int64            `json:"comState"`
	AutomaticRenewal       bool             `json:"automaticRenewal"`
	AutomaticRenewalStatus int64            `json:"automaticRenewalStatus"`
	Activable              bool             `json:"activable"`
}

type Devices struct {
	DeviceInfos []DeviceInformation `json:"data"`
	Paging      models.NextURL      `json:"paging"`
}
