package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/adrien3d/go-sigfox-client/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func SigfoxAPIRequest(c *gin.Context, apiUrl string, respType interface{}) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)

	key := config.RetrieveSigfoxAPIKey(c)
	fmt.Println(key)
	req.SetBasicAuth(key.SigfoxKey, key.SigfoxSecret)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(bodyResp, &respType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Wrong Sigfox API response, please check the provided credentials")
		panic(err.Error())
	}

	c.JSON(http.StatusOK, respType)
}
