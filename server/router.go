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
	devices := router.Group("/devices")
	{
		deviceController := controllers.NewDeviceController()
		devices.GET("/:id/info", deviceController.GetDevice)
		devices.GET("/:id/infos", deviceController.GetDevices)
		devices.GET("/:id/messages", deviceController.GetDeviceMessages)
		devices.GET("/:id/locations", deviceController.GetDeviceLocations)
	}

	users := router.Group("/users")
	{
		userController := controllers.NewUserController()
		users.GET("/:id/info", userController.GetUser)
		users.GET("/:id/infos", userController.GetUsers)
	}
}
