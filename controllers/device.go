package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models/devices"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
)

type DeviceController struct {
}

func NewDeviceController() DeviceController {
	return DeviceController{}
}

func (dc DeviceController) GetDevice(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("id")
	utils.SigfoxAPIRequest(c, url, new(devices.DeviceInformation))
}

func (dc DeviceController) GetDevices(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/"
	utils.SigfoxAPIRequest(c, url, new(devices.Devices))
}

func (dc DeviceController) GetDeviceMessages(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("id") + "/messages"
	utils.SigfoxAPIRequest(c, url, new(devices.DeviceMessages))
}

func (dc DeviceController) GetDeviceLocations(c *gin.Context) {
	url := "https://api.sigfox.com/v2/devices/" + c.Param("id") + "/locations"
	utils.SigfoxAPIRequest(c, url, new(devices.DeviceLocations))
}
