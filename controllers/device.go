package controllers

import (
	"encoding/json"
	"github.com/adrien3d/go-sigfox-client/config"
	"github.com/adrien3d/go-sigfox-client/models/devices"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type DeviceController struct {
}

func NewDeviceController() DeviceController {
	return DeviceController{}
}

func (dc DeviceController) GetDeviceMessages(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.sigfox.com/v2/devices/"+c.Param("sigfoxId")+"/messages", nil)

	key := config.RetrieveSigfoxAPIKey(c)
	req.SetBasicAuth(key.SigfoxKey, key.SigfoxSecret)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var messages = new(devices.DeviceMessages)
	err = json.Unmarshal(bodyResp, &messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Wrong Sigfox API response, please check the provided credentials")
		panic(err.Error())
	}

	c.JSON(http.StatusOK, messages)
}

func (dc DeviceController) GetDeviceLocations(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.sigfox.com/v2/devices/"+c.Param("sigfoxId")+"/locations", nil)

	key := config.RetrieveSigfoxAPIKey(c)
	req.SetBasicAuth(key.SigfoxKey, key.SigfoxSecret)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var locations = new(devices.DeviceLocations)
	err = json.Unmarshal(bodyResp, &locations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Wrong Sigfox API response, please check the provided credentials")
		panic(err.Error())
	}

	c.JSON(http.StatusOK, locations)
}
