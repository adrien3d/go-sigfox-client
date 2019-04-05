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

	// Simple group: sigfox
	router.GET("/", Index)
	sigfox := router.Group("/sigfox")
	{
		messageController := controllers.NewMessageController()
		sigfox.GET("/:sigfoxId/messages", messageController.GetDeviceMessages)
		locationController := controllers.NewLocationController()
		sigfox.GET("/:sigfoxId/locations", locationController.GetDeviceLocations)
	}
}
