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
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level int64  `json:"level"`
}

type IdName struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Token struct {
	State              int64  `json:"state"`
	DetailMessage      string `json:"detailMessage"`
	End                int64  `json:"end"`
	UnsubscriptionTime int64  `json:"unsubscriptionTime"`
	FreeMessages       int64  `json:"freeMessages"`
	FreeMessagesSent   int64  `json:"freeMessagesSent"`
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
	AverageSnr             int64            `json:"averageSnr"`
	AverageRssi            int64            `json:"averageRssi"`
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
