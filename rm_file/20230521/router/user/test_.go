package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

// InitTestRouter 初始化 Test 路由信息
func (s *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	Test22Router := Router.Group("Test22").Use(middleware.OperationRecord())
	Test22RouterWithoutRecord := Router.Group("Test22")
	var Test22Api = v1.ApiGroupApp.UserApiGroup.TestApi
	{
		Test22Router.POST("createTest", Test22Api.CreateTest)             // 新建Test
		Test22Router.DELETE("deleteTest", Test22Api.DeleteTest)           // 删除Test
		Test22Router.DELETE("deleteTestByIds", Test22Api.DeleteTestByIds) // 批量删除Test
		Test22Router.PUT("updateTest", Test22Api.UpdateTest)              // 更新Test
	}
	{
		Test22RouterWithoutRecord.GET("findTest", Test22Api.FindTest)       // 根据ID获取Test
		Test22RouterWithoutRecord.GET("getTestList", Test22Api.GetTestList) // 获取Test列表
	}
}
