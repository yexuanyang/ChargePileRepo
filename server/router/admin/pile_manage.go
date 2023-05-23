package admin

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PileManageRouter struct {
}

func (s *PileManageRouter) InitPileRouter(Router *gin.RouterGroup) {
	pileRouter := Router.Group("admin").Use(middleware.OperationRecord())
	pileRouterApi := v1.ApiGroupApp.AdminApiGroup.PileManageApi
	{
		pileRouter.GET("manage", pileRouterApi.GetPileList)
		pileRouter.POST("manage_carList", pileRouterApi.GetPileCarList)
	}
}
