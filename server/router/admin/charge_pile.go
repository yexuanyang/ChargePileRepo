package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChargePileRouter struct {
}

// InitChargePileRouter 初始化 ChargePile 路由信息
func (s *ChargePileRouter) InitChargePileRouter(Router *gin.RouterGroup) {
	chargePileRouter := Router.Group("chargePile").Use(middleware.OperationRecord())
	chargePileRouterWithoutRecord := Router.Group("chargePile")
	var chargePileApi = v1.ApiGroupApp.AdminApiGroup.ChargePileApi
	{
		chargePileRouter.POST("createChargePile", chargePileApi.CreateChargePile)             // 新建ChargePile
		chargePileRouter.DELETE("deleteChargePile", chargePileApi.DeleteChargePile)           // 删除ChargePile
		chargePileRouter.DELETE("deleteChargePileByIds", chargePileApi.DeleteChargePileByIds) // 批量删除ChargePile
		chargePileRouter.PUT("updateChargePile", chargePileApi.UpdateChargePile)              // 更新ChargePile
	}
	{
		chargePileRouterWithoutRecord.GET("findChargePile", chargePileApi.FindChargePile)       // 根据ID获取ChargePile
		chargePileRouterWithoutRecord.GET("getChargePileList", chargePileApi.GetChargePileList) // 获取ChargePile列表
	}
}
