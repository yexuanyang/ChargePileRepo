import service, {bonusService} from '@/utils/request'

// @Tags Car
// @Summary 创建Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Car true "创建Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carInfo/createCar [post]
export const createCar = (data) => {
  return service({
    url: '/carInfo/createCar',
    method: 'post',
    data
  })
}

export const createCar2 = (data) => {
  return bonusService({
    url: '/client/index/car',
    method: 'post',
    data
  })
}

// @Tags Car
// @Summary 删除Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Car true "删除Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carInfo/deleteCar [delete]
export const deleteCar = (data) => {
  return service({
    url: '/carInfo/deleteCar',
    method: 'delete',
    data
  })
}

// @Tags Car
// @Summary 删除Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carInfo/deleteCar [delete]
export const deleteCarByIds = (data) => {
  return service({
    url: '/carInfo/deleteCarByIds',
    method: 'delete',
    data
  })
}

// @Tags Car
// @Summary 更新Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Car true "更新Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carInfo/updateCar [put]
export const updateCar = (data) => {
  return service({
    url: '/carInfo/updateCar',
    method: 'put',
    data
  })
}


export const updateCar2 = (data) => {
  return bonusService({
    url: '/client/index/car',
    method: 'patch',
    data
  })
}

// @Tags Car
// @Summary 用id查询Car
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Car true "用id查询Car"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carInfo/findCar [get]
export const findCar = (params) => {
  return service({
    url: '/carInfo/findCar',
    method: 'get',
    params
  })
}

// @Tags Car
// @Summary 分页获取Car列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Car列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carInfo/getCarList [get]
export const getCarList = (params) => {
  return service({
    url: '/carInfo/getCarList',
    method: 'get',
    params
  })
}

export const getCarListByUserId = (params) => {
  return service({
    url: '/carInfo/getCarListByUserId',
    method: 'get',
    params
  })
}

export const getCarListByUserId2 = (params) => {
  return bonusService({
    url: '/client/index/car',
    method: 'get',
    params
  })
}