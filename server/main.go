package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/user"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	// 初始化调度信息
	user.InitStation()
	for i := range user.ChargeStations {
		go user.ChargeStations[i].Waiting.DispatchCar(&user.ChargeStations[i])
		for j := range user.ChargeStations[i].ChargePiles {
			go user.ChargeStations[i].ChargePiles[j].Charging(&user.ChargeStations[i])
		}
	}

	core.RunWindowsServer()
}
