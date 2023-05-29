package user

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"modernc.org/libc/limits"
	"strconv"
	"sync"
	"time"
)

var (
	// ErrorDispatch 出现故障时调度算法
	ErrorDispatch = map[int]string{0: "优先级调度", 1: "时间顺序调度"}
	// Dispatch 调度算法
	Dispatch = map[int]string{0: "先来先服务调度", 1: "singleDispatch", 2: "batchDispatch"}
	// Mode 充电模式
	Mode = map[string]int{"fast": 0, "slow": 1}
	// Power 充电功率
	Power = []float64{30.0, 7.0}
	// ChargeStations 充电站集合
	ChargeStations []ChargeStation
	// FreePile 充电桩空闲位置数量，FreePile[1] = 2表示第二个充电桩有两个空闲位置；FreePile[0]=3表示第一个充电桩有三个空闲位置
	FreePile      = make([]int, (FastChargingPileNum+TrickleChargingPileNum)*ChargingQueueLen)
	FreePileMutex sync.Mutex
)

const (
	peakPrice   = 1.0 // 峰时单价
	normalPrice = 0.7 // 平时单价
	lowPrice    = 0.4 // 谷时单价
	peakStart   = 10  // 峰时开始时间
	peakEnd     = 15  // 峰时结束时间
	peakStart2  = 18  // 峰时开始时间
	peakEnd2    = 21  // 峰时结束时间
	normalStart = 7   // 平时开始时间
	normalEnd   = 23  // 平时结束时间
	// ChargeStationNumber 充电站的数量
	ChargeStationNumber = 2
	// WaitingAreaSize 等待区大小
	WaitingAreaSize = 6
	// ChargingQueueLen 充电队列长度
	ChargingQueueLen = 2
	// FastChargingPileNum 快充充电桩数量
	FastChargingPileNum = 4
	// TrickleChargingPileNum 慢充充电桩数量
	TrickleChargingPileNum = 2
	// ServiceCostRate 服务费
	ServiceCostRate = 0.8
)

func InitStation() {
	ChargeStations = make([]ChargeStation, ChargeStationNumber)
	for i := range ChargeStations {
		// 初始充电桩
		ChargeStations[i].ChargePiles = make([]ChargePile, FastChargingPileNum+TrickleChargingPileNum)
		for k := range ChargeStations[i].ChargePiles {
			// 前面是快充，后面是慢充
			if k < FastChargingPileNum {
				ChargeStations[i].ChargePiles[k].Mode = Mode["fast"]
			} else {
				ChargeStations[i].ChargePiles[k].Mode = Mode["slow"]
			}
			ChargeStations[i].ChargePiles[k].Cars = make([]Car, 0, ChargingQueueLen)
			ChargeStations[i].ChargePiles[k].PileId = k
			// 初始化空闲的充电桩
			FreePile[k] = ChargingQueueLen
		}
		// 初始等待区车辆
		ChargeStations[i].Waiting.Cars = make([]Car, 0, WaitingAreaSize)
	}
}

func calculateFee(start, end time.Time, power float64) float64 {
	var fee float64
	subHours := end.Sub(start).Hours() // 计算充电时间(小时)
	// 计算小时所用的费用
	for i := 0; i < int(subHours); i++ {
		currTime := start.Add(time.Duration(i) * time.Hour)
		if isPeakTime(currTime) {
			fee += peakPrice * power
		} else if isNormalTime(currTime) {
			fee += normalPrice * power
		} else {
			fee += lowPrice * power
		}
	}
	// 计算剩下的分钟所用的费用
	restHour := (end.Sub(start).Minutes() - subHours*60) / 60.0
	if isPeakTime(end) {
		fee += peakPrice * power * restHour
	} else if isNormalTime(end) {
		fee += normalPrice * power * restHour
	} else {
		fee += lowPrice * power * restHour
	}
	return fee
}

// 判断是否是峰时
func isPeakTime(t time.Time) bool {
	hour := t.Hour()
	return (hour >= peakStart && hour < peakEnd) || (hour >= peakStart2 && hour < peakEnd2)
}

// 判断是否是平时
func isNormalTime(t time.Time) bool {
	hour := t.Hour()
	return (hour >= normalStart && hour < normalEnd) && !isPeakTime(t)
}

type ChargeStation struct {
	ChargePiles []ChargePile
	Waiting     WaitingBlock
}

type ChargePile struct {
	PileId      int
	WaitingTime float64
	Mode        int
	Cars        []Car
	FinishTime  time.Time
	mu          sync.Mutex
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
	mu   sync.Mutex
}

// Enqueue 往充电桩的充电队列中加入汽车
func (chargePile *ChargePile) Enqueue(car Car) {
	chargePile.mu.Lock()
	defer chargePile.mu.Unlock()
	chargePile.Cars = append(chargePile.Cars, car)
}

// Dequeue 从充电桩的充电队列中取出汽车
func (chargePile *ChargePile) Dequeue() (car Car) {
	chargePile.mu.Lock()
	defer chargePile.mu.Unlock()
	if len(chargePile.Cars) == 0 {
		return Car{}
	}
	car = chargePile.Cars[0]
	chargePile.Cars = chargePile.Cars[1:]
	return car
}

// Charging 充电桩线程运行的程序，从充电队列中取出汽车开始充电
func (chargePile *ChargePile) Charging() {
	for {
		// 从充电队列中取出汽车
		car := chargePile.Dequeue()
		if car.CarId == "" {
			// 充电队列中没有汽车
			continue
		}
		chargeTime, _ := time.ParseDuration(strconv.FormatFloat(car.ChargeTime, 'f', 5, 64) + "h")
		fmt.Printf("充电桩 %v 开始充电，充电汽车 %v\n", chargePile.PileId, car.CarId)
		// 更新数据库订单中汽车开始充电时间
		startTime := time.Now()
		global.GVA_DB.Model(&user.Order{}).Where("car_id = ? AND state = '队列区'", car.CarId).Update("started_at", startTime)
		time.Sleep(chargeTime)
		// 充电完毕，向等待区发送空闲信号
		FreePileMutex.Lock()
		FreePile[chargePile.PileId] += 1
		FreePileMutex.Unlock()
		// 修改充电桩的等待时间
		chargePile.WaitingTime -= car.ChargeTime
		// 计算充电所使用的钱
		endTime := startTime.Add(chargeTime)
		chargeCost := calculateFee(startTime, endTime, Power[chargePile.Mode])
		serviceCost := car.ChargeTime * ServiceCostRate

		fmt.Printf("充电桩 %v 开始充电，充电汽车 %v, 充电费用 %v\n", chargePile.PileId, chargeTime.String(), chargeCost)
		// 更新数据库订单中汽车结束充电时间和充电费用
		global.GVA_DB.Model(&user.Order{}).Where("car_id = ? AND state = '队列区'", car.CarId).Updates(map[string]interface{}{
			"stop_at": endTime, "total_cost": serviceCost + chargeCost, "service_cost": serviceCost,
			"charge_cost": chargeCost, "state": "已完成",
		})
	}
}

// Dequeue 从等待区汽车队列中取出一辆汽车
func (waitingBlock *WaitingBlock) Dequeue() (car Car) {
	waitingBlock.mu.Lock()
	defer waitingBlock.mu.Unlock()
	if len(waitingBlock.Cars) == 0 {
		return Car{}
	}
	car = waitingBlock.Cars[0]
	waitingBlock.Cars = waitingBlock.Cars[1:]
	return car
}

// Enqueue 将汽车加入等待区
func (waitingBlock *WaitingBlock) Enqueue(car Car) error {
	waitingBlock.mu.Lock()
	defer waitingBlock.mu.Unlock()
	if len(waitingBlock.Cars) < WaitingAreaSize {
		waitingBlock.Cars = append(waitingBlock.Cars, car)
		return nil
	} else {
		return errors.New("该充电站等待区已满，请稍后再试或更换充电站")
	}
}

// Delete 从等待区删除一辆汽车
func (waitingBlock *WaitingBlock) Delete(car Car) {
	waitingBlock.mu.Lock()
	defer waitingBlock.mu.Unlock()
	for i := range waitingBlock.Cars {
		if car.CarId == waitingBlock.Cars[i].CarId {
			waitingBlock.Cars = append(waitingBlock.Cars[:i], waitingBlock.Cars[i+1:]...)
		}
	}
}

// Update 更新等待区的车辆信息(用户修改了充电请求，重新把用户加入队列)
func (waitingBlock *WaitingBlock) Update(car Car) error {
	waitingBlock.Delete(car)
	err := waitingBlock.Enqueue(car)
	return err
}

// DispatchCar 等待区线程执行的程序,等待区的汽车持续向充电汽车请求加入汽车
func (waitingBlock *WaitingBlock) DispatchCar(station *ChargeStation) {
	for {
		var minIndex int
		var min float64
		if len(waitingBlock.Cars) == 0 {
			continue
		}
		var hasFree = false
		// 判断是否有空闲的充电桩
		for pile := range FreePile {
			if pile != 0 {
				hasFree = true
				break
			}
		}
		if hasFree {
			// 从等待区中取出第一辆汽车的信息
			currentCar := waitingBlock.Cars[0]
			minIndex = limits.INT_MAX
			min = float64(limits.INT_MAX)
			// 判断应该加入哪一个充电桩
			if currentCar.Mode == Mode["fast"] {
				for i := 0; i < FastChargingPileNum; i++ {
					if station.ChargePiles[i].WaitingTime < min {
						min = station.ChargePiles[i].WaitingTime
						minIndex = i
					}
				}
			} else {
				for i := FastChargingPileNum; i < FastChargingPileNum+TrickleChargingPileNum; i++ {
					if station.ChargePiles[i].WaitingTime < min {
						min = station.ChargePiles[i].WaitingTime
						minIndex = i
					}
				}
			}
			if FreePile[minIndex] > 0 {
				fmt.Printf("充电桩%d 剩余空闲 %d\n", minIndex, FreePile[minIndex])
				// 充电桩空闲，将汽车取出等待区加入到充电队列
				station.ChargePiles[minIndex].Enqueue(waitingBlock.Dequeue())
				fmt.Printf("从等待区取出一辆汽车加入充电桩%d\n", minIndex)
				// 增加充电桩队列的等待时间
				station.ChargePiles[minIndex].WaitingTime += currentCar.ChargeTime

				// 把空闲数量-1，加锁防止冲突
				FreePileMutex.Lock()
				FreePile[minIndex] -= 1
				FreePileMutex.Unlock()
				fmt.Printf("充电桩%d 剩余空闲 %d\n", minIndex, FreePile[minIndex])
				// 修改订单在数据库里面的状态，从等待区变成队列区
				// 更新数据库里面的订单信息
				global.GVA_DB.Model(&user.Order{}).Where("car_id = ? AND state = '等待区'", currentCar.CarId).Updates(map[string]interface{}{"state": "队列区",
					"pile_id": minIndex, "time": currentCar.ChargeTime})
			}
		}
	}
}

// GetCarInfoByOrder 从订单中获取车辆信息
func GetCarInfoByOrder(order user.Order) (car Car) {
	if order.ChargeType == "快充" {
		car.ChargeTime = order.Kwh / Power[Mode["fast"]]
		car.Mode = Mode["fast"]
	} else {
		car.ChargeTime = order.Kwh / Power[Mode["slow"]]
		car.Mode = Mode["slow"]
	}
	car.Energy = order.Kwh
	car.CarId = order.CarId
	return car
}
