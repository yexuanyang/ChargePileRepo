package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChargePileApi struct {
}

var chargePileService = service.ServiceGroupApp.AdminServiceGroup.ChargePileService

// CreateChargePile 创建ChargePile
// @Tags ChargePile
// @Summary 创建ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.ChargePile true "创建ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargePile/createChargePile [post]
func (chargePileApi *ChargePileApi) CreateChargePile(c *gin.Context) {
	var chargePile system.ChargePile
	err := c.ShouldBindJSON(&chargePile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PileType":  {utils.NotEmpty()},
		"StationId": {utils.NotEmpty()},
	}
	if err := utils.Verify(chargePile, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargePileService.CreateChargePile(&chargePile); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChargePile 删除ChargePile
// @Tags ChargePile
// @Summary 删除ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.ChargePile true "删除ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chargePile/deleteChargePile [delete]
func (chargePileApi *ChargePileApi) DeleteChargePile(c *gin.Context) {
	var chargePile system.ChargePile
	err := c.ShouldBindJSON(&chargePile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargePileService.DeleteChargePile(chargePile); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChargePileByIds 批量删除ChargePile
// @Tags ChargePile
// @Summary 批量删除ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /chargePile/deleteChargePileByIds [delete]
func (chargePileApi *ChargePileApi) DeleteChargePileByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargePileService.DeleteChargePileByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChargePile 更新ChargePile
// @Tags ChargePile
// @Summary 更新ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.ChargePile true "更新ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chargePile/updateChargePile [put]
func (chargePileApi *ChargePileApi) UpdateChargePile(c *gin.Context) {
	var chargePile system.ChargePile
	err := c.ShouldBindJSON(&chargePile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"PileType":  {utils.NotEmpty()},
		"StationId": {utils.NotEmpty()},
	}
	if err := utils.Verify(chargePile, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chargePileService.UpdateChargePile(chargePile); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChargePile 用id查询ChargePile
// @Tags ChargePile
// @Summary 用id查询ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.ChargePile true "用id查询ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chargePile/findChargePile [get]
func (chargePileApi *ChargePileApi) FindChargePile(c *gin.Context) {
	var chargePile system.ChargePile
	err := c.ShouldBindQuery(&chargePile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rechargePile, err := chargePileService.GetChargePile(chargePile.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rechargePile": rechargePile}, c)
	}
}

// GetChargePileList 分页获取ChargePile列表
// @Tags ChargePile
// @Summary 分页获取ChargePile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.ChargePileSearch true "分页获取ChargePile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargePile/getChargePileList [get]
func (chargePileApi *ChargePileApi) GetChargePileList(c *gin.Context) {
	var pageInfo adminReq.ChargePileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := chargePileService.GetChargePileInfoList(pageInfo); err != nil {
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
