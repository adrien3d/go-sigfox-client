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

func (dc DeviceController) GetDevice(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("id")
	SigfoxAPIRequest(c, url, new(devices.DeviceInformation))
}

func (dc DeviceController) GetDevices(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/"
	SigfoxAPIRequest(c, url, new(devices.Devices))
}

func (dc DeviceController) GetDeviceMessages(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("id") + "/messages"
	SigfoxAPIRequest(c, url, new(devices.DeviceMessages))
}

func (dc DeviceController) GetDeviceLocations(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("id") + "/locations"
	SigfoxAPIRequest(c, url, new(devices.DeviceLocations))
}
