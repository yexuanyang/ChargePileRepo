package global

import (
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
	// 出现故障时调度算法
	ErrorDispatch = map[int]string{0: "priorityDispatch", 1: "timeDispatch"}
	// 调度算法
	Dispatch = map[int]string{0: "default", 1: "singleDispatch", 2: "batchDispatch"}
	// 充电模式
	Mode           = map[string]int{"fast": 0, "slow": 1}
	ChargeStations []ChargeStation
)

const (
	// 充电站的数量
	ChargeStationNumber = 2
	// 等待区大小
	WaitingAreaSize = 6
	// 充电队列长度
	ChargingQueueLen = 2
	// 快充充电桩数量
	FastChargingPileNum = 4
	// 慢充充电桩数量
	TrickleChargingPileNum = 2
	// 快充功率
	FastPower = 30
	// 慢充功率
	TrickPower = 7
	// 服务费
	ServiceCostRate = 0.8
)

type ChargeStation struct {
	FastChargePiles      [FastChargingPileNum]ChargePile
	TrickleChargingPiles [TrickleChargingPileNum]ChargePile
	Waiting              WaitingBlock
}

type ChargePile struct {
	WaitingTime float64
	Mode        int
	Cars        [ChargingQueueLen]Car
	Length      int
	FinishTime  time.Time
}

type Car struct {
	CarId      string
	ChargeTime float64
	Mode       int
	Energy     float64
	QueueId    int
}

type WaitingBlock struct {
	Cars []Car
}

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
