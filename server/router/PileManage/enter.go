package PileManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	PileRouter
}

type PileRouter struct {
}

func (s *PileRouter) InitPileRouter(Router *gin.RouterGroup) {
	pileRouter := Router.Group("admin").Use(middleware.OperationRecord())
	pileRouterApi := v1.ApiGroupApp.PileManageApiGroup.PileRouter
	{
		pileRouter.GET("manage", pileRouterApi.GetPileList)
	}
}
