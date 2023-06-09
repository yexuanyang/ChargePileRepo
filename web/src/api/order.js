import service, {bonusService} from '@/utils/request'

// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
export const createOrder = (data) => {
  return service({
    url: '/order/createOrder',
    method: 'post',
    data
  })
}

export const createOrder2 = (data) => {
  return bonusService({
    url: '/client/index/order',
    method: 'post',
    data
  })
}

// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
export const deleteOrder = (data) => {
  return service({
    url: '/order/deleteOrder',
    method: 'delete',
    data
  })
}

export const deleteOrder2 = (data) => {
  return bonusService({
    url: '/client/index/order',
    method: 'delete',
    data
  })
}

// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
export const deleteOrderByIds = (data) => {
  return service({
    url: '/order/deleteOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
export const updateOrder = (data) => {
  return service({
    url: '/order/updateOrder',
    method: 'put',
    data
  })
}

export const updateOrder2 = (data) => {
  return bonusService({
    url: '/client/index/order',
    method: 'patch',
    data
  })
}

// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
export const findOrder = (params) => {
  return service({
    url: '/order/findOrder',
    method: 'get',
    params
  })
}

// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
export const getOrderList = (params) => {
  return service({
    url: '/order/getOrderList',
    method: 'get',
    params
  })
}

export const getUnFinishedOrderNumber = (params) => {
  return service({
    url: '/order/getUnFinishedOrderNumber',
    method: 'get',
    params
  })
}

export const getOrderListByUserId2 = (params) => {
  return bonusService({
    url: '/client/index/order',
    method: 'get',
    params
  })
}

export const getOrderListByUserId = (params) => {
  return service({
    url: '/order/getOrderListByUserId',
    method: 'get',
    params
  })
}