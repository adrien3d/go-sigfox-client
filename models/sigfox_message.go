package models

type SigfoxBaseStation struct {
	Id string `json:"id"`
}

type ReceptionInfos struct {
	SigfoxBaseStation SigfoxBaseStation `json:"baseStation"`
	Delay             float64           `json:"delay"`
	Latitude          string            `json:"lat"` //TODO: conversion from string to float64
	Longitude         string            `json:"lng"` //TODO: conversion from string to float64
}

type ComputedLocation struct {
	Latitude  float64 `json:"lat" bson:"lat" valid:"-"`
	Longitude float64 `json:"lng" bson:"lng" valid:"-"`
	Radius    int64   `json:"radius" bson:"radius" valid:"-"`
	Source    int64   `json:"source" bson:"source" valid:"-"`
}

type SigfoxDevice struct {
	Id string `json:"id"`
}

type SigfoxMessage struct {
	Device          SigfoxDevice       `json:"device" `
	Time            int64              `json:"time"`
	Data            string             `json:"data"`
	RolloverCounter int64              `json:"rolloverCounter"`
	SequenceNbr     int64              `json:"seqNumber"`
	RInfos          []ReceptionInfos   `json:"rinfos"`
	FramesNbr       int64              `json:"nbFrames"`
	Operator        string             `json:"operator"`
	Country         string             `json:"country"`
	CompLoc         []ComputedLocation `json:"computedLocation"`
	Lqi             int64              `json:"lqi"`
}

type NextURL struct {
	SigfoxNextURL string `json:"next" bson:"next"`
}

type APIMessages struct {
	Messages []SigfoxMessage `json:"data" bson:"data"`
	Paging   NextURL         `json:"paging" bson:"paging"`
}
