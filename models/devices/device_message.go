package devices

import "github.com/adrien3d/go-sigfox-client/models"

type BaseStation struct {
	Id string `json:"id"`
}

type ReceptionInfos struct {
	BaseStation BaseStation `json:"baseStation"`
	Delay       float64     `json:"delay"`
	Latitude    string      `json:"lat"` //TODO: conversion from string to float64
	Longitude   string      `json:"lng"` //TODO: conversion from string to float64
}

type ComputedLocation struct {
	Latitude  float64 `json:"lat" bson:"lat" valid:"-"`
	Longitude float64 `json:"lng" bson:"lng" valid:"-"`
	Radius    int64   `json:"radius" bson:"radius" valid:"-"`
	Source    int64   `json:"source" bson:"source" valid:"-"`
}

type DeviceId struct {
	Id string `json:"id"`
}

type DeviceMessage struct {
	Device          DeviceId           `json:"device"`
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

type DeviceMessages struct {
	Messages []DeviceMessage `json:"data" bson:"data"`
	Paging   models.NextURL  `json:"paging" bson:"paging"`
}
