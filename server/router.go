package server

import (
	"github.com/adrien3d/go-sigfox-client/config"
	"github.com/adrien3d/go-sigfox-client/controllers"
	"github.com/spf13/viper"
	"net/http"

	"github.com/adrien3d/go-sigfox-client/middlewares"
	"github.com/gin-gonic/gin"
	"time"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "You successfully reached Sigfox API client."})
}

func ConfigMiddleware(viper *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.ToContext(c, config.New(viper))
		c.Next()
	}
}

func (a *API) SetupRouter() {
	router := a.Router

	router.Use(middlewares.CorsMiddleware(middlewares.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Requested-With, Accept, Token",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Use(ConfigMiddleware(a.Config))

	router.GET("/", Index)

	users := router.Group("/api-users")
	{
		apiUserController := controllers.NewApiUserController()
		users.GET("/:id/info", apiUserController.GetApiUser)
		users.GET("/:id/infos", apiUserController.GetApiUsers)
	}

	contracts := router.Group("/contracts")
	{
		contractController := controllers.NewContractController()
		contracts.GET("/:id/info", contractController.GetContract)
		contracts.GET("/:id/infos", contractController.GetContracts)
	}

	coverages := router.Group("/coverage")
	{
		coverageController := controllers.NewCoverageController()
		coverages.GET("/prediction", coverageController.GetCoveragePrediction)
		coverages.GET("/predictions", coverageController.GetCoveragePredictions)
		coverages.GET("/redundancy", coverageController.GetCoverageRedundancy)
	}

	devices := router.Group("/devices")
	{
		deviceController := controllers.NewDeviceController()
		devices.GET("/:id/info", deviceController.GetDevice)
		devices.GET("/:id/infos", deviceController.GetDevices)
		devices.GET("/:id/messages", deviceController.GetDeviceMessages)
		devices.GET("/:id/locations", deviceController.GetDeviceLocations)
	}

	deviceTypes := router.Group("device-types")
	{
		devicetypeController := controllers.NewDeviceTypeController()
		deviceTypes.GET("/:id/info", devicetypeController.GetDeviceType)
		deviceTypes.GET("/:id/infos", devicetypeController.GetDeviceTypes)
	}

	groups := router.Group("/groups")
	{
		groupController := controllers.NewGroupController()
		groups.GET("/:id/info", groupController.GetGroup)
		groups.GET("/:id/infos", groupController.GetGroups)
	}

	profiles := router.Group("/profiles")
	{
		profileController := controllers.NewProfileController()
		profiles.GET("/:id/infos", profileController.GetProfiles)
	}
}
