package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("client")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	chargeApi := v1.ApiGroupApp.UserApiGroup.ChargeApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	adminBaseRouter := Router.Group("admin")
	{
		adminBaseRouter.POST("login", baseApi.AdminLogin)
		adminBaseRouter.GET("index/manage", chargeApi.GetChargePileList2)  // 获得充电桩列表
		adminBaseRouter.POST("index/manage", chargeApi.ChangeChargePile)   //开关充电桩
		adminBaseRouter.GET("index/report", chargeApi.GetChargePileReport) //查看单个充电桩报表信息
		adminBaseRouter.GET("index/charge", chargeApi.GetChargePileInfo)   //查询充电桩服务的车辆信息
	}
	return baseRouter
}
