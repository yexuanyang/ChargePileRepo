package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestApi struct {
}

var Test22Service = service.ServiceGroupApp.UserServiceGroup.TestService

// CreateTest 创建Test
// @Tags Test
// @Summary 创建Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Test true "创建Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Test22/createTest [post]
func (Test22Api *TestApi) CreateTest(c *gin.Context) {
	var Test22 user.Test
	err := c.ShouldBindJSON(&Test22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := Test22Service.CreateTest(&Test22); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTest 删除Test
// @Tags Test
// @Summary 删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Test true "删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Test22/deleteTest [delete]
func (Test22Api *TestApi) DeleteTest(c *gin.Context) {
	var Test22 user.Test
	err := c.ShouldBindJSON(&Test22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := Test22Service.DeleteTest(Test22); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestByIds 批量删除Test
// @Tags Test
// @Summary 批量删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Test22/deleteTestByIds [delete]
func (Test22Api *TestApi) DeleteTestByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := Test22Service.DeleteTestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTest 更新Test
// @Tags Test
// @Summary 更新Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Test true "更新Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Test22/updateTest [put]
func (Test22Api *TestApi) UpdateTest(c *gin.Context) {
	var Test22 user.Test
	err := c.ShouldBindJSON(&Test22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := Test22Service.UpdateTest(Test22); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTest 用id查询Test
// @Tags Test
// @Summary 用id查询Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.Test true "用id查询Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Test22/findTest [get]
func (Test22Api *TestApi) FindTest(c *gin.Context) {
	var Test22 user.Test
	err := c.ShouldBindQuery(&Test22)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reTest22, err := Test22Service.GetTest(Test22.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTest22": reTest22}, c)
	}
}

// GetTestList 分页获取Test列表
// @Tags Test
// @Summary 分页获取Test列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.TestSearch true "分页获取Test列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Test22/getTestList [get]
func (Test22Api *TestApi) GetTestList(c *gin.Context) {
	var pageInfo userReq.TestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := Test22Service.GetTestInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
