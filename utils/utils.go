package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/adrien3d/go-sigfox-client/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
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
		logrus.Fatalln(resp.StatusCode, resp.Body, bodyResp)
		logrus.Fatalln(err.Error())
	}

	c.JSON(http.StatusOK, respType)
}
