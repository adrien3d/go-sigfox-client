package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

func (uc UserController) GetUser(c *gin.Context) {
	url := "https://api.sigfox.com/v2/api-users/" + c.Param("userId")
	SigfoxAPIRequest(c, url, new(models.User))
}

func (uc UserController) GetUsers(c *gin.Context) {
	url := "https://api.sigfox.com/v2/api-users/"
	SigfoxAPIRequest(c, url, new(models.Users))
}
