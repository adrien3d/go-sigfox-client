package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models/devices"
	"github.com/gin-gonic/gin"
)

type DeviceController struct {
}

func NewDeviceController() DeviceController {
	return DeviceController{}
}

func (dc DeviceController) GetDevices(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/"
	devices := devices.Devices{}
	SigfoxAPIRequest(c, url, devices)
}

func (dc DeviceController) GetDeviceMessages(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("sigfoxId") + "/messages"
	messages := devices.DeviceMessages{}
	SigfoxAPIRequest(c, url, messages)
}

func (dc DeviceController) GetDeviceLocations(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("sigfoxId") + "/locations"
	locations := devices.DeviceLocations{}
	SigfoxAPIRequest(c, url, locations)
}
