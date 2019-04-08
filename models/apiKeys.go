package models

type APIKey struct {
	Name         string `json:"name"`
	DeviceTypeId string `json:"deviceTypeId"`
	SigfoxId     string `json:"sigfoxId"`
	SigfoxKey    string `json:"sigfoxKey"`
	SigfoxSecret string `json:"sigfoxSecret"`
}

type APIKeys struct {
	Keys []APIKey `json:"data"`
}
