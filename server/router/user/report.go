package user

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReportRouter struct{}

func (reportRouter *ReportRouter) InitReportRouter(Router *gin.RouterGroup) {
	reportRouterGroup := Router.Group("report").Use(middleware.OperationRecord())
	//reportRouterGroupWithoutRecord := Router.Group("report")
	var reportApi = v1.ApiGroupApp.UserApiGroup.ReportApi
	{
		reportRouterGroup.POST("getDurationTotalCharge", reportApi.GetDurationChargeKwh)
		reportRouterGroup.POST("getDurationTotalPrice", reportApi.GetDurationPrice)
	}
}
