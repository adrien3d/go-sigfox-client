package config

import (
	"encoding/json"
	"fmt"
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func RetrieveSigfoxAPIKey(c *gin.Context) models.APIKey {
	var ret models.APIKey
	apiKeys := ExtractSigfoxAPIKey(GetString(c, "API_KEYS_FILENAME"))
	for _, key := range apiKeys.Keys {
		if key.SigfoxId == c.Param("sigfoxId") {
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
	utils.CheckErr(err)

	fmt.Println(data)

	return data
}
