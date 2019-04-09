package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
)

type GroupController struct {
}

func NewGroupController() GroupController {
	return GroupController{}
}

func (GroupController) GetGroup(c *gin.Context) {
	url := "https://api.sigfox.com/v2/groups/" + c.Param("id")
	utils.SigfoxAPIRequest(c, url, new(models.GroupInfo))
}

func (GroupController) GetGroups(c *gin.Context) {
	url := "https://api.sigfox.com/v2/groups/"
	utils.SigfoxAPIRequest(c, url, new(models.GroupInfos))
}
