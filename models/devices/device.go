package devices

import "github.com/adrien3d/go-sigfox-client/models"

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type ComputedLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Radius    int32   `json:"radius"`
	Source    int8    `json:"source"`
}

type Group struct {
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Level int64  `json:"level,omitempty"`
}

type IdName struct {
	Id   string `json:"id"`
	Name string `json:"name,omitempty"`
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
	Group                  Group            `json:"group"`
	DeviceType             IdName           `json:"deviceType"`
	AverageSnr             int64            `json:"averageSnr,omitempty"`
	AverageRssi            int64            `json:"averageRssi,omitempty"`
	ActivationTime         int64            `json:"activationTime"`
	Contract               IdName           `json:"contract"`
	Token                  Token            `json:"token"`
	CreatedBy              string           `json:"createdBy"`
	LastEditionTime        int64            `json:"lastEditionTime"`
	LastEditedBy           string           `json:"lastEditedBy"`
	UnsubscriptionTime     int64            `json:"unsubscriptionTime"`
	CreationTime           int64            `json:"creationTime"`
	ModemCertificate       IdName           `json:"modemCertificate"`
	ProductCertificate     IdName           `json:"productCertificate"`
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
