package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
)

type ApiUserController struct {
}

func NewApiUserController() ApiUserController {
	return ApiUserController{}
}

func (ApiUserController) GetApiUser(c *gin.Context) {
	url := "https://api.sigfox.com/v2/api-users/" + c.Param("id")
	utils.SigfoxAPIRequest(c, url, new(models.ApiUser))
}

func (ApiUserController) GetApiUsers(c *gin.Context) {
	url := "https://api.sigfox.com/v2/api-users/"
	utils.SigfoxAPIRequest(c, url, new(models.ApiUsers))
}

func (ApiUserController) CreateApiUser(c *gin.Context) {
	var params models.NewApiUserRequest
	if c.BindJSON(&params) == nil {
		url := "https://api.sigfox.com/v2/api-users/"
		utils.SigfoxAPIBodyRequest(c, url, params, new(models.NewApiUserResponse))
	}
}
