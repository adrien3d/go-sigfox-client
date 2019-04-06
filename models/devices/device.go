package devices

import "github.com/adrien3d/go-sigfox-client/models"

type DeviceInformations struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	SequenceNumber      int64  `json:"sequenceNumber"`
	TrashSequenceNumber int64  `json:"trashSequenceNumber"`
	LastCom             int64  `json:"lastCom"`
	Pac                 string `json:"pac"`
}

type Devices struct {
	DeviceInfos []DeviceInformations `json:"data"`
	Paging      models.NextURL       `json:"paging"`
}

/*
{
  "data": [
    {
      "id": "00FF",
      "name": "Device 1",
      "sequenceNumber": 5,
      "trashSequenceNumber": 1,
      "lastCom": 1462801589979,
      "pac": "585CB3B42AC98BD4",
      "location": {
        "lat": 48.8585715,
        "lng": 2.2922923
      },
      "lastComputedLocation": {
        "lat": 42,
        "lng": 1.4,
        "radius": 4000,
        "source": 2
      },
      "group": {
        "id": "572f1204017975032d8ec1dd",
        "name": "Group 1",
        "type": "0",
        "level": 1
      },
      "deviceType": {
        "id": "57309548171c857460043085",
        "name": "Device type 1"
      },
      "averageSnr": 26.1,
      "averageRssi": -124.84,
      "activationTime": 1462801512365,
      "contract": {
        "id": "573095b7171c857460043086",
        "name": "First contract"
      },
      "token": {
        "state": 0,
        "detailMessage": "Active",
        "end": 1503619200000,
        "unsubscriptionTime": 1503602300000,
        "freeMessages": 1,
        "freeMessagesSent": 9
      },
      "createdBy": "57986ea2e89a8e255c31ddf7",
      "lastEditionTime": 1487065492780,
      "lastEditedBy": "57986ea2e89a8e255c31ddf7",
      "unsubscriptionTime": 1513619200000,
      "creationTime": 1462801032158,
      "modemCertificate": {
        "id": "57309674171c857460043087",
        "name": "M_0004_AC59_02"
      },
      "productCertificate": {
        "id": "53aaad5ae4b0ba69e6d561b9",
        "name": "P_0003_71CF_01"
      },
      "state": 0,
      "comState": 1,
      "automaticRenewal": true,
      "automaticRenewalStatus": 0,
      "activable": true
    }
  ],
  "paging": {
    "next": "string"
  }
}
*/
