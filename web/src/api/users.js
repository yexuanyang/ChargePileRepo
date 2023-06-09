import service from '@/utils/request'

// @Tags Users
// @Summary 创建Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "创建Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userInfo/createUsers [post]
export const createUsers = (data) => {
  return service({
    url: '/userInfo/createUsers',
    method: 'post',
    data
  })
}

// @Tags Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userInfo/deleteUsers [delete]
export const deleteUsers = (data) => {
  return service({
    url: '/userInfo/deleteUsers',
    method: 'delete',
    data
  })
}

// @Tags Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userInfo/deleteUsers [delete]
export const deleteUsersByIds = (data) => {
  return service({
    url: '/userInfo/deleteUsersByIds',
    method: 'delete',
    data
  })
}

// @Tags Users
// @Summary 更新Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "更新Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userInfo/updateUsers [put]
export const updateUsers = (data) => {
  return service({
    url: '/userInfo/updateUsers',
    method: 'put',
    data
  })
}

// @Tags Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userInfo/findUsers [get]
export const findUsers = (params) => {
  return service({
    url: '/userInfo/findUsers',
    method: 'get',
    params
  })
}

// @Tags Users
// @Summary 分页获取Users列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Users列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userInfo/getUsersList [get]
export const getUsersList = (params) => {
  return service({
    url: '/userInfo/getUsersList',
    method: 'get',
    params
  })
}
