package PileManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

func (PileRouter) PileListView(c *gin.Context) {
	var cr system.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	//list, count, _ := common.ComList(&models.BannerModel{}, common.Option{
	//	PageInfo: cr,
	//	Debug:    true,
	//})
	//
	//res.OKWithList(list, count, c)
}
