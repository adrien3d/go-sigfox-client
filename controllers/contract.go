package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
)

type ContractController struct {
}

func NewContractController() ContractController {
	return ContractController{}
}
func (ContractController) GetContract(c *gin.Context) {
	url := "https://api.sigfox.com/v2/contract-infos/" + c.Param("id")
	utils.SigfoxAPIRequest(c, url, new(models.Contract))
}

func (ContractController) GetContracts(c *gin.Context) {
	url := "https://api.sigfox.com/v2/contract-infos/"
	utils.SigfoxAPIRequest(c, url, new(models.Contracts))
}
