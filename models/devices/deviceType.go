package devices

import "github.com/adrien3d/go-sigfox-client/models"

type DeviceType struct {
	Name                string        `json:"name"`
	KeepAlive           int           `json:"keepAlive"`
	AlertEmail          string        `json:"alertEmail"`
	ID                  string        `json:"id"`
	Description         string        `json:"description"`
	DownlinkMode        int           `json:"downlinkMode"`
	DownlinkDataString  string        `json:"downlinkDataString"`
	PayloadType         int           `json:"payloadType"`
	PayloadConfig       string        `json:"payloadConfig"`
	Group               models.Group  `json:"group"`
	Contract            models.IdName `json:"contract"`
	GeolocPayloadConfig models.IdName `json:"geolocPayloadConfig"`
	CreationTime        int           `json:"creationTime"`
	CreatedBy           string        `json:"createdBy"`
	LastEditionTime     int           `json:"lastEditionTime"`
	LastEditedBy        string        `json:"lastEditedBy"`
	AutomaticRenewal    bool          `json:"automaticRenewal"`
}

type DeviceTypes struct {
	Data   []DeviceType   `json:"data"`
	Paging models.NextURL `json:"paging"`
}
