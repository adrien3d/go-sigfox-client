package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
}

func NewProfileController() ProfileController {
	return ProfileController{}
}

func (ProfileController) GetProfiles(c *gin.Context) {
	url := "https://api.sigfox.com/v2/profiles?groupId=" + c.Param("id")
	utils.SigfoxAPIRequest(c, url, new(models.Profiles))
}
