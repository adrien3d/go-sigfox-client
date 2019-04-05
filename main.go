package main

import (
	"github.com/adrien3d/go-sigfox-client/server"
	"github.com/adrien3d/go-sigfox-client/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	api := &server.API{Router: gin.Default(), Config: viper.New()}

	// Configuration setup
	err := api.SetupViper()
	if err != nil {
		panic(err)
	}

	// Router setup
	api.SetupRouter()

	err = api.Router.Run(api.Config.GetString("host_address"))

	utils.CheckErr(err)
}
