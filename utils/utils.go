package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/adrien3d/go-sigfox-client/config"
	"github.com/adrien3d/go-sigfox-client/models/devices"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

func CheckErr(e error) {
	if e != nil {
		logrus.Errorln(e)
		panic(e)
	}
}

func SigfoxAPIRequest(c *gin.Context, apiUrl string, respType interface{}) {
	req, _ := http.NewRequest("GET", apiUrl, nil)

	handleRequest(c, req, respType)
}

func SigfoxAPIBodyRequest(c *gin.Context, apiUrl string, reqBody interface{}, respType interface{}) {
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonBody))

	handleRequest(c, req, respType)
}

func handleRequest(c *gin.Context, req *http.Request, respType interface{}) {
	client := &http.Client{}
	key := config.RetrieveSigfoxAPIKey(c)
	fmt.Println(key)
	req.SetBasicAuth(key.SigfoxKey, key.SigfoxSecret)

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalln(err)
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusNoContent) {
		logrus.Warnln(resp.StatusCode, resp.Body)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)

	err = json.Unmarshal(bodyResp, &respType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Wrong Sigfox API response, please check the provided credentials")
		logrus.Warnln(resp.StatusCode, resp.Body, bodyResp)
		logrus.Warnln(err.Error())
	}

	switch reflect.ValueOf(respType).Elem().Interface().(type) {
	case devices.Devices:
		fmt.Println("Devices detected")
		var devs = devices.Devices{}
		CheckErr(json.Unmarshal(bodyResp, &devs))
		for _, device := range devs.DeviceInfos {
			unitTimeInRFC3339 := time.Unix(device.LastCom/1000, 0).Format(time.RFC850) // converts utc time to RFC3339 format
			fmt.Println(unitTimeInRFC3339, device.Location, device.LastComputedLocation)
		}
	case devices.DeviceInformation:
		fmt.Println("Device detected")
	case devices.DeviceMessages:
		var mess = devices.DeviceMessages{}
		CheckErr(json.Unmarshal(bodyResp, &mess))
		for i, mes := range mess.Messages {
			fmt.Println(i, mes.CompLoc, mes.RInfos)
		}
	}

	c.JSON(http.StatusOK, respType)
}
