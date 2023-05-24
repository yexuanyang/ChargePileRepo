package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type ReportApi struct {
}

var reportService = service.ServiceGroupApp.UserServiceGroup.ReportService

// GetDurationChargeKwh path: /report/getDurationTotalPrice
func (reportApi *ReportApi) GetDurationChargeKwh(c *gin.Context) {
	var report request.ReportSearch
	claims, err := utils.GetClaims(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if report.UserId == 0 {
		report.UserId = int(claims.BaseClaims.ID)
	}
	err = c.ShouldBindJSON(&report)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var total float64
	total, err = reportService.GetDurationTotalCharge(report)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"total": total}, c)
}

// GetDurationPrice path: /report/getDurationTotalPrice
func (reportApi *ReportApi) GetDurationPrice(c *gin.Context) {
	var report request.ReportSearch
	claims, err := utils.GetClaims(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if report.UserId == 0 {
		report.UserId = int(claims.BaseClaims.ID)
	}
	err = c.ShouldBindJSON(&report)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var total float64
	total, err = reportService.GetDurationTotalPrice(report)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"total": total}, c)
}
