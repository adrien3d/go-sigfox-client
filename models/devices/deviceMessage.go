package devices

import "github.com/adrien3d/go-sigfox-client/models"

type Repetition struct {
	Nseq     int64   `json:"nseq"`
	Rssi     float64 `json:"rssi"`
	Snr      float64 `json:"snr"`
	Freq     float64 `json:"freq"`
	Repeated bool    `json:"repeated"`
}

type CbStatus struct {
	Status int64  `json:"status,omitempty"`
	Info   string `json:"info,omitempty"`
	CbDef  string `json:"cbDef,omitempty"`
	Time   int64  `json:"time,omitempty"`
}

type ReceptionInfos struct {
	BaseStation   models.IdName `json:"baseStation"`
	Delay         float64       `json:"delay"` // Not in Sigfox API Documentation
	Rssi          float64       `json:"rssi,omitempty"`
	RssiRepeaters float64       `json:"rssiRepeaters,omitempty"`
	Latitude      string        `json:"lat"` //Should be in int in Sigfox API
	Longitude     string        `json:"lng"` //Should be in int in Sigfox API
	Snr           float64       `json:"snr,omitempty"`
	SnrRepeaters  float64       `json:"snrRepeaters,omitempty"`
	Freq          float64       `json:"freq,omitempty"`
	FreqRepeaters float64       `json:"freqRepeaters,omitempty"`
	Rep           int64         `json:"rep,omitempty"`
	Repetitions   []Repetition  `json:"repetitions,omitempty"`
	CbStatus      CbStatus      `json:"cbStatus,omitempty"`
}

type DownlinkAnswerStatus struct {
	BaseStation  models.IdName `json:"baseStation,omitempty"`
	PlannedPower float64       `json:"plannedPower,omitempty"`
	Data         string        `json:"data,omitempty"`
	Operator     string        `json:"operator,omitempty"`
	Country      string        `json:"country,omitempty"`
}

type DeviceMessage struct {
	Device               models.IdName        `json:"device"`
	Time                 int64                `json:"time"`
	Data                 string               `json:"data"`
	RolloverCounter      int64                `json:"rolloverCounter"` // Not in Sigfox API Documentation
	AckRequired          bool                 `json:"ackRequired"`
	Lqi                  int64                `json:"lqi"`
	SequenceNbr          int64                `json:"seqNumber"`
	FramesNbr            int64                `json:"nbFrames"`
	CompLoc              []ComputedLocation   `json:"computedLocation"`
	RInfos               []ReceptionInfos     `json:"rinfos"`
	DownlinkAnswerStatus DownlinkAnswerStatus `json:"downlinkAnswerStatus,omitempty"`
	Operator             string               `json:"operator"`
	Country              string               `json:"country"`
}

type DeviceMessages struct {
	Messages []DeviceMessage `json:"data" bson:"data"`
	Paging   models.NextURL  `json:"paging" bson:"paging"`
}
