package controllers

import (
	"fmt"
	"github.com/adrien3d/go-sigfox-client/models"
	"github.com/adrien3d/go-sigfox-client/utils"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

type CoverageController struct {
}

func NewCoverageController() CoverageController {
	return CoverageController{}
}

func (cc CoverageController) GetCoveragePrediction(c *gin.Context) {
	var params models.CoverageQueryParams
	if c.ShouldBind(&params) == nil {
		if (math.Abs(params.Lat) < 85) && (math.Abs(params.Lng) < 180) {
			url := "https://api.sigfox.com/v2/coverages/global/predictions?lat=" + fmt.Sprintf("%f", params.Lat) +
				"&lng=" + fmt.Sprintf("%f", params.Lng)
			if params.Radius != 0 {
				url += "&radius=" + string(params.Radius)
			}
			if params.GroupId != "" {
				url += "&groupId=" + params.GroupId
			}
			utils.SigfoxAPIRequest(c, url, new(models.Coverage))
		}
	}
}

func (cc CoverageController) GetCoveragePredictions(c *gin.Context) {
	var params models.CoveragesRequest
	if c.BindJSON(&params) == nil {
		url := "https://api.sigfox.com/v2/coverages/global/predictions"
		utils.SigfoxAPIBodyRequest(c, url, params, new(models.CoveragesResponse))
	}
}

func (cc CoverageController) GetCoverageRedundancy(c *gin.Context) {
	var params models.CoverageQueryParams
	if c.ShouldBind(&params) == nil {
		if (math.Abs(params.Lat) < 85) && (math.Abs(params.Lng) < 180) {
			url := "https://api.sigfox.com/v2/coverages/operators/redundancy?lat=" + fmt.Sprintf("%f", params.Lat) +
				"&lng=" + fmt.Sprintf("%f", params.Lng)
			if params.Radius != 0 {
				url += "&radius=" + string(params.Radius)
			}
			if params.GroupId != "" {
				url += "&groupId=" + params.GroupId
			}
			utils.SigfoxAPIRequest(c, url, new(models.Redundancy))
		} else {
			c.JSON(http.StatusBadRequest, "Wrong lat or lng query parameter")
		}
	}
}
