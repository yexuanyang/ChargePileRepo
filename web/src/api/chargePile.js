import service, {bonusService} from '@/utils/request'

// @Tags ChargePile
// @Summary 创建ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargePile true "创建ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargePile/createChargePile [post]
export const createChargePile = (data) => {
  return service({
    url: '/chargePile/createChargePile',
    method: 'post',
    data
  })
}

// @Tags ChargePile
// @Summary 删除ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargePile true "删除ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chargePile/deleteChargePile [delete]
export const deleteChargePile = (data) => {
  return service({
    url: '/chargePile/deleteChargePile',
    method: 'delete',
    data
  })
}

// @Tags ChargePile
// @Summary 删除ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chargePile/deleteChargePile [delete]
export const deleteChargePileByIds = (data) => {
  return service({
    url: '/chargePile/deleteChargePileByIds',
    method: 'delete',
    data
  })
}

// @Tags ChargePile
// @Summary 更新ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargePile true "更新ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chargePile/updateChargePile [put]
export const updateChargePile = (data) => {
  return service({
    url: '/chargePile/updateChargePile',
    method: 'put',
    data
  })
}

// @Tags ChargePile
// @Summary 用id查询ChargePile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ChargePile true "用id查询ChargePile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chargePile/findChargePile [get]
export const findChargePile = (params) => {
  return service({
    url: '/chargePile/findChargePile',
    method: 'get',
    params
  })
}

// @Tags ChargePile
// @Summary 分页获取ChargePile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ChargePile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chargePile/getChargePileList [get]
export const getChargePileList = (params) => {
  return service({
    url: '/chargePile/getChargePileList',
    method: 'get',
    params
  })
}

export const UpdateChargePileByIds = (data) => {
  return service({
    url: '/admin/manage_isOpen',
    method: 'post',
    data
  })
}

export const UpdateChargePileByIds2 = (data) => {
  return service({
    url: '/admin/index/manage',
    method: 'post',
    data
  })
}

export const getChargeStationInfo = () => {
  return bonusService({
    url: '/admin/index/manage',
    method: 'get',
  })
}
