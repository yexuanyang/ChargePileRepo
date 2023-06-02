package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/user"
	"go.uber.org/zap"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
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
// @name                        x-token
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
	var wg sync.WaitGroup
	//var ctx context.Context
	wg.Add(user.ChargeStationNumber + user.FastChargingPileNum + user.TrickleChargingPileNum)
	for i := range user.ChargeStations {
		//go func() {
		//	for {
		//		select {
		//		case <-ctx.Done():
		//				return
		//		default:
		//			user.ChargeStations[i].Waiting.DispatchCar(&user.ChargeStations[i], &wg)
		//		}
		//	}
		//}()
		go user.ChargeStations[i].Waiting.DispatchCar(&user.ChargeStations[i], &wg)
		for j := range user.ChargeStations[i].ChargePiles {
			go user.ChargeStations[i].ChargePiles[j].Charging(&user.ChargeStations[i], &wg)
		}
	}
	core.RunWindowsServer()
	wg.Wait()
}
