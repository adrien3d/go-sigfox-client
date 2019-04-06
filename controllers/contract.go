package controllers

import (
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/gin-gonic/gin"
)

type ContractController struct {
}

func NewContractController() ContractController {
	return ContractController{}
}
func (cc ContractController) GetContract(c *gin.Context) {
	url := "https://api.sigfox.com/v2/contract-infos/" + c.Param("id")
	SigfoxAPIRequest(c, url, new(models.Contract))
}

func (cc ContractController) GetContracts(c *gin.Context) {
	url := "https://api.sigfox.com/v2/contract-infos/"
	SigfoxAPIRequest(c, url, new(models.Contracts))
}
