import service from '@/utils/request'

// @Tags ChargeStation
// @Summary 创建ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargeStation true "创建ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargeStation/createChargeStation [post]
export const createChargeStation = (data) => {
  return service({
    url: '/chargeStation/createChargeStation',
    method: 'post',
    data
  })
}

// @Tags ChargeStation
// @Summary 删除ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargeStation true "删除ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chargeStation/deleteChargeStation [delete]
export const deleteChargeStation = (data) => {
  return service({
    url: '/chargeStation/deleteChargeStation',
    method: 'delete',
    data
  })
}

// @Tags ChargeStation
// @Summary 删除ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chargeStation/deleteChargeStation [delete]
export const deleteChargeStationByIds = (data) => {
  return service({
    url: '/chargeStation/deleteChargeStationByIds',
    method: 'delete',
    data
  })
}

// @Tags ChargeStation
// @Summary 更新ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargeStation true "更新ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chargeStation/updateChargeStation [put]
export const updateChargeStation = (data) => {
  return service({
    url: '/chargeStation/updateChargeStation',
    method: 'put',
    data
  })
}

// @Tags ChargeStation
// @Summary 用id查询ChargeStation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ChargeStation true "用id查询ChargeStation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chargeStation/findChargeStation [get]
export const findChargeStation = (params) => {
  return service({
    url: '/chargeStation/findChargeStation',
    method: 'get',
    params
  })
}

// @Tags ChargeStation
// @Summary 分页获取ChargeStation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ChargeStation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargeStation/getChargeStationList [get]
export const getChargeStationList = (params) => {
  return service({
    url: '/chargeStation/getChargeStationList',
    method: 'get',
    params
  })
}
