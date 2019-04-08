package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models/devices"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
)

type DeviceTypeController struct {
}

func NewDeviceTypeController() DeviceTypeController {
	return DeviceTypeController{}
}

func (DeviceTypeController) GetDeviceType(c *gin.Context) {
	url := "https://api.sigfox.com/v2/device-types/" + c.Param("id")
	utils.SigfoxAPIRequest(c, url, new(devices.DeviceType))
}

func (DeviceTypeController) GetDeviceTypes(c *gin.Context) {
	url := "https://api.sigfox.com/v2/device-types/"
	utils.SigfoxAPIRequest(c, url, new(devices.DeviceTypes))
}
