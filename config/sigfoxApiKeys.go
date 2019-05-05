package config

import (
	"encoding/json"
	"fmt"
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func RetrieveSigfoxAPIKey(c *gin.Context) models.APIKey {
	var ret models.APIKey
	apiKeys := ExtractSigfoxAPIKey(GetString(c, "API_KEYS_FILENAME"))
	id := c.Param("id")

	for _, key := range apiKeys.Keys {
		for _, sigId := range key.SigfoxId {
			if sigId == id {
				ret = key
			}
		}
		if key.DeviceTypeId == id {
			ret = key
		}
	}

	return ret
}

func ExtractSigfoxAPIKey(fileName string) models.APIKeys {
	var data models.APIKeys

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	return data
}
