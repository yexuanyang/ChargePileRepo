package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type ChargeApi struct{}

func (chargeApi *ChargeApi) GetChargePileList(c *gin.Context) {
	list, total, err := chargePileService.GetChargePileList()
	if err != nil {
		return
	}
	response.OkWithDetailed(response.ListResult{List: list, Total: total}, "获取成功", c)
}

func (chargeApi *ChargeApi) CreateChargePile(c *gin.Context) {
}
