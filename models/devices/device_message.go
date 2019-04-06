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
	Status int64  `json:"status"`
	Info   string `json:"info"`
	CbDef  string `json:"cbDef"`
	Time   int64  `json:"time"`
}

type ReceptionInfos struct {
	BaseStation   IdName       `json:"baseStation"`
	Delay         float64      `json:"delay"` // Not in Sigfox API Documentation
	Rssi          float64      `json:"rssi"`
	RssiRepeaters float64      `json:"rssiRepeaters"`
	Latitude      string       `json:"lat"` //Should be in int in Sigfox API
	Longitude     string       `json:"lng"` //Should be in int in Sigfox API
	Snr           float64      `json:"snr"`
	SnrRepeaters  float64      `json:"snrRepeaters"`
	Freq          float64      `json:"freq"`
	FreqRepeaters float64      `json:"freqRepeaters"`
	Rep           int64        `json:"rep"`
	Repetitions   []Repetition `json:"repetitions"`
	CbStatus      CbStatus     `json:"cbStatus"`
}

type DownlinkAnswerStatus struct {
	BaseStation  IdName  `json:"baseStation"`
	PlannedPower float64 `json:"plannedPower"`
	Data         string  `json:"data"`
	Operator     string  `json:"operator"`
	Country      string  `json:"country"`
}

type DeviceMessage struct {
	Device               IdName               `json:"device"`
	Time                 int64                `json:"time"`
	Data                 string               `json:"data"`
	RolloverCounter      int64                `json:"rolloverCounter"` // Not in Sigfox API Documentation
	AckRequired          bool                 `json:"ackRequired"`
	Lqi                  int64                `json:"lqi"`
	SequenceNbr          int64                `json:"seqNumber"`
	FramesNbr            int64                `json:"nbFrames"`
	CompLoc              []ComputedLocation   `json:"computedLocation"`
	RInfos               []ReceptionInfos     `json:"rinfos"`
	DownlinkAnswerStatus DownlinkAnswerStatus `json:"downlinkAnswerStatus"`
	Operator             string               `json:"operator"`
	Country              string               `json:"country"`
}

type DeviceMessages struct {
	Messages []DeviceMessage `json:"data" bson:"data"`
	Paging   models.NextURL  `json:"paging" bson:"paging"`
}
