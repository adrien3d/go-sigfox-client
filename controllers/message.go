package controllers

import (
	"encoding/json"
	"github.com/adrien3d/go-sigfox-client/config"
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type MessageController struct {
}

func NewMessageController() MessageController {
	return MessageController{}
}

func (mc MessageController) GetDeviceMessages(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.sigfox.com/v2/devices/" + c.Param("sigfoxId") + "/messages", nil)

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

	var messages = new(models.APIMessages)
	err = json.Unmarshal(bodyResp, &messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Wrong Sigfox API response, please check the provided credentials")
		panic(err.Error())
	}

	c.JSON(http.StatusOK, messages)
}
