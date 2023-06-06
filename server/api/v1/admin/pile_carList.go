package admin

type PileCarRequest struct {
	Id uint `json:"id"` // 充电桩ID
}

// PileCarResponse 充电桩
type PileCarResponse struct {
	CarList []uint `json:"car_list"` // 汽车ID数组
}

// GetPileCarList 获取充电桩下等待队列中汽车列表
// @Tags ChargePile
// @Summary 获取充电桩下等待队列中汽车列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PileCarRequest true "获取充电桩下等待队列中汽车列表"
// @Success 200 {string} string PileCarResponse
// @Router /admin/manage_carList [post]

//func (PileManageApi) GetPileCarList(c *gin.Context) {
//	var cr PileCarRequest
//	err := c.ShouldBindJSON(&cr)
//	if err != nil {
//		response.FailWithMessage("参数绑定失败", c)
//		return
//	}
//	var chargePile system.ChargePile
//	var pileCarResponse PileCarResponse
//	err = global.GVA_DB.Model(&system.ChargePile{}).Preload("CarModel").Take(&chargePile, cr.Id).Error
//	if err != nil {
//		response.FailWithMessage("查数据库失败", c)
//		return
//	}
//	for _, car := range chargePile.CarModel {
//		pileCarResponse.CarList = append(pileCarResponse.CarList, car.ID)
//	}
//	response.OkWithData(pileCarResponse, c)
//}
