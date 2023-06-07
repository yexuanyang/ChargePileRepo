package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.UserServiceGroup.OrderService

// CreateOrder 创建Order
// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
// /client/index/order [post]
func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var order system.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.StopAt = time.Now().Add(time.Hour * 24)
	order.StartedAt = time.Now()
	if order.UserId == 0 {
		order.UserId = int(utils.GetUserID(c))
	}
	if order.StationId == 0 {
		// 默认充电站是第一个充电站
		order.StationId = 1
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.State = "WAITING"

	verify := utils.Rules{
		"user_id":    {utils.NotEmpty()},
		"carId":      {utils.NotEmpty()},
		"chargeType": {utils.NotEmpty()},
		"kwh":        {utils.NotEmpty()},
	}
	if err := utils.Verify(order, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order.Kwh = 0
	if err := orderService.CreateOrder(&order); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		car := GetCarInfoByOrder(order)
		err = ChargeStations[order.StationId-1].Waiting.Enqueue(car)
		CarNum[order.StationId-1][car.Mode] += 1
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除Order
// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	var req systemReq.OrderDelete
	var order system.Order
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uintId, _ := strconv.ParseUint(req.Id, 10, 64)
	order.ID = uint(uintId)

	order, err = orderService.GetOrder(order.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	car := GetCarInfoByOrder(order)
	CarNum[order.StationId-1][car.Mode] -= 1
	stationId := order.StationId - 1
	if order.State == "WAITING" {
		// 将汽车从WAITING中移除,订单在数据库中也删除
		err = ChargeStations[stationId].Waiting.Delete(car)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else if order.State == "CHARGING" {
		// 该汽车正在充电
		FreePileMutex[stationId].Lock()
		defer FreePileMutex[stationId].Unlock()
		FreePile[stationId][order.PileId] += 1 // 空闲位置+1

		// 向指定的充电桩线程发送提前结束充电请求
		IsInterrupt[stationId][order.PileId] = true
		// 数据库中的订单不需要删除
		// 向前端返回成功即可
		response.OkWithMessage("删除成功", c)
		return
	} else if order.State == "DISPATCHED " {
		currentPile := ChargeStations[stationId].ChargePiles[order.PileId]
		// 汽车不在充电，数据库中的订单直接删除
		currentPile.mu.Lock() // 获取当前充电桩的充电队列的锁
		defer currentPile.mu.Unlock()

		currentPile.Cars = append(currentPile.Cars[:1], currentPile.Cars[2:]...) // 将汽车移除

		FreePileMutex[stationId].Lock()
		defer FreePileMutex[stationId].Unlock()
		FreePile[stationId][order.PileId] += 1 // 空闲位置+1
	}

	if err := orderService.DeleteOrder(order); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除Order
// @Tags Order
// @Summary 批量删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /order/deleteOrderByIds [delete]
func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.DeleteOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新Order
// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var req systemReq.OrderUpdate
	err := c.ShouldBindJSON(&req)
	id, _ := strconv.ParseUint(req.Id, 10, 32)
	var order system.Order
	order.ID = uint(id)
	order.ChargeType = req.Mode
	order.ApplyKwh = req.ApplyKwh

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims, err := utils.GetClaims(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if order.UserId == 0 {
		order.UserId = int(claims.BaseClaims.ID)
	}
	if order.StationId == 0 {
		order.StationId = 1
	}
	updateCar := GetCarInfoByOrder(order)
	err = ChargeStations[order.StationId-1].Waiting.Update(updateCar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"mode":      {utils.NotEmpty()},
		"apply_kwh": {utils.NotEmpty()},
	}
	if err := utils.Verify(order, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.UpdateOrder(order); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询Order
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	var order system.Order
	err := c.ShouldBindQuery(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorder, err := orderService.GetOrder(order.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// GetOrderList 分页获取Order列表
// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.OrderSearch true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo systemReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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

func (orderApi *OrderApi) GetOrderListByUserId(c *gin.Context) {
	claims, err := utils.GetClaims(c)
	var pageInfo systemReq.OrderSearch
	err = c.ShouldBindQuery(&pageInfo)
	pageInfo.Order.UserId = int(claims.BaseClaims.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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

// GetOrderListByUserId2 /client/index/order
func (orderApi *OrderApi) GetOrderListByUserId2(c *gin.Context) {
	var search systemReq.OrderSearch2
	err := c.ShouldBindQuery(&search)
	search.Order.UserId = int(utils.GetUserID(c))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList2(search); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		res := make([]systemRes.OrderResponse, 0, total)
		for _, item := range list {
			cur := systemRes.OrderResponse{}
			cur.Mode = QueuePrefix[Mode[item.ChargeType]]
			cur.ID = strconv.Itoa(int(item.ID))
			cur.UserID = strconv.Itoa(item.UserId)
			cur.CarID = item.CarId
			cur.ApplyKwh = item.ApplyKwh
			cur.ChargeKwh = item.Kwh
			cur.State = item.State
			cur.ChargePrice = item.ChargeCost
			cur.ServicePrice = item.ServiceCost
			cur.Fee = item.TotalCost
			cur.CreateTime = item.CreatedAt.String()
			cur.FinishTime = item.StopAt.String()
			cur.DispatchTime = "xx-xx-xx xx:xx:xx"
			cur.ChargeID = strconv.Itoa(item.PileId)
			cur.StartTime = item.StartedAt.String()
			if search.Mode == "HISTORY" {
				cur.FrontCars = 0
			} else {
				cur.FrontCars, err = orderService.GetFrontCars(item)
				if err != nil {
					response.FailWithMessage(err.Error(), c)
					return
				}
			}
			res = append(res, cur)
		}
		response.OkWithData(res, c)
		return
	}
}

func (orderApi *OrderApi) GetUnFinishedOrderNumber(c *gin.Context) {
	var pageInfo systemReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var total int64
	pageInfo.State = "WAITING"
	if _, total1, err := orderService.GetOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		total += total1
	}
	pageInfo.State = "CHARGING"
	if _, total2, err := orderService.GetOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		total += total2
	}
	response.OkWithData(total, c)

}
