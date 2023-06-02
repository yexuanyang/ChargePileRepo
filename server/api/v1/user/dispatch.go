package user

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"math"
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
	Mode = map[string]int{"快充": 0, "慢充": 1}
	// Power 充电功率
	Power = []float64{30.0, 7.0}
	// ChargeStations 充电站集合
	ChargeStations []ChargeStation
	// FreePile 充电桩空闲位置数量，FreePile[1][1] = 2表示第2个充电站第2个充电桩有2个空闲位置
	FreePile      [][]int
	FreePileMutex []sync.Mutex
	// IsInterrupt 用来处理前端发送的取消订单操作,IsInterrupt[0][1] = true 表示第1个充电站的第2个充电桩需要提前中断充电
	IsInterrupt [][]bool
	// CarNum 等待区汽车数量 CarNum[1][1] = 2表示充电站2的等待区中有2辆慢充汽车 CarNum[1][0] = 3 表示充电站2的等待区有3辆快充汽车
	CarNum [][]int
	// QueuePrefix 队列号前缀，根据充电类型区分
	QueuePrefix = []string{0: "F", 1: "T"}
	// IsOpen 处理充电桩关闭的操作,IsOpen[0][2] = true表示第一个充电站的第三个充电桩开启
	IsOpen [][]bool

	chargeStationService = service.ServiceGroupApp.AdminServiceGroup.ChargeStationService
	chargePileService    = service.ServiceGroupApp.AdminServiceGroup.ChargePileService
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
	// ServiceCostRate 服务费
	ServiceCostRate = 0.8
)

type ChargeStation struct {
	StationId   int
	ChargePiles []*ChargePile
	Waiting     *WaitingBlock
}

type ChargePile struct {
	StationId   int
	PileId      int // 存储在充电站中的id位置，从0开始编号
	Id          int //存储在数据库中的id位置
	WaitingTime float64
	Mode        int
	Cars        []Car
	mu          sync.Mutex
}

type Car struct {
	CarId      string
	ChargeTime float64
	Mode       int
	Energy     float64
	QueueId    string // 存储本车的排队号码，排队号码显示充电模式和排队号,快充F，慢充T
}

type WaitingBlock struct {
	StationId int
	Cars      []Car
	mu        sync.Mutex
}

func InitStation() {
	// 获取充电站的数量
	_, number, _ := chargeStationService.GetChargeStationInfoList(request.ChargeStationSearch{})
	// 初始化充电站切片,数据库中充电站从1开始编号,程序里充电站从0开始编号
	ChargeStations = make([]ChargeStation, number)
	for i := range ChargeStations {
		ChargeStations[i].StationId = i
	}

	// 初始化充电桩充电队列空闲数量列表
	FreePile = make([][]int, number)
	FreePileMutex = make([]sync.Mutex, number)
	CarNum = make([][]int, number)
	for i := range CarNum {
		// 列标0表示快充，列标1表示慢充
		CarNum[i] = make([]int, 2)
	}
	IsInterrupt = make([][]bool, number)
	IsOpen = make([][]bool, number)

	// 初始化充电站中充电桩列表
	for i := range ChargeStations {
		// 初始充电桩
		ChargeStations[i].ChargePiles = make([]*ChargePile, 0)
		// 初始等待区车辆
		ChargeStations[i].Waiting = &WaitingBlock{}
		ChargeStations[i].Waiting.Cars = make([]Car, 0, WaitingAreaSize)
		ChargeStations[i].Waiting.StationId = i
	}

	// 获取数据库充电桩列表,列表按照充电站编号升序排列
	chargePileList, _, _ := chargePileService.GetChargePileInfoList(request.ChargePileSearch{})
	// 遍历充电桩列表，从中获取初始化充电桩的信息
	for _, chargePile := range chargePileList {

		stationIndex := chargePile.StationID - 1
		pile := ChargePile{}
		// 充电桩的PileId存储在充电站内的下标
		pile.PileId = len(ChargeStations[stationIndex].ChargePiles)
		// id存储在数据库中充电桩的id
		pile.Id = int(chargePile.ID)
		fmt.Println(pile.PileId)
		pile.StationId = int(stationIndex)
		pile.Mode = Mode[chargePile.PileType]
		// 初始化充电桩充电队列内的汽车
		pile.Cars = make([]Car, 0, ChargingQueueLen)

		// 充电桩加入对应的充电站
		ChargeStations[stationIndex].ChargePiles = append(ChargeStations[stationIndex].ChargePiles, &pile)

		// 初始化充电桩的充电队列空闲数量,FreePile[stationIndex]是对应充电站下的空闲列表
		FreePile[stationIndex] = append(FreePile[stationIndex], ChargingQueueLen)

		// 每加入一个充电桩都需要增加一个新的IsInterrupt标志
		IsInterrupt[stationIndex] = append(IsInterrupt[stationIndex], false)

		// 不管充电桩开没开都要初始化一个线程
		// 没开的充电桩初始化的时候IsOpen是false
		if chargePile.IsOpen {
			IsOpen[stationIndex] = append(IsOpen[stationIndex], true)
		} else {
			IsOpen[stationIndex] = append(IsOpen[stationIndex], false)
		}

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
	restHour := math.Mod(end.Sub(start).Minutes(), 60.0) / 60.0
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
func (chargePile *ChargePile) Charging(station *ChargeStation) {
	for {
		// 查看是否有车在充电队列，没有车就空闲（不论充电桩是否开启）
		// 没车，等一段时间之后再次查询是否有车
		if len(chargePile.Cars) == 0 {
			time.Sleep(3 * time.Second)
			continue
		}

		// 充电桩关闭（或者充电桩故障），提前退出充电
		// 退出充电的时候要采用故障调度，直接将等待区拷贝出来，然后将故障队列的车放到等待区，同时将等待区上锁不允许车辆加入，等到故障队列的车都调度完毕之后将原来的等待区复制回来并解锁
		if !IsOpen[station.StationId][chargePile.PileId] {
			station.Waiting.mu.Lock()
			wBlock := station.Waiting
			station.Waiting.Cars = chargePile.Cars
			for {
				if len(station.Waiting.Cars) == 0 {
					break
				}
			}
			station.Waiting = wBlock
			station.Waiting.mu.Unlock()
			// 故障队列调度完毕，充电桩进入阻塞，每5s查询一次充电桩是否开启
			for {
				if IsOpen[station.StationId][chargePile.PileId] {
					break
				}
				time.Sleep(5 * time.Second)
			}
		}

		car := chargePile.Cars[0]
		chargeTime, _ := time.ParseDuration(strconv.FormatFloat(car.ChargeTime, 'f', 5, 64) + "h")
		fmt.Printf("充电桩 %v 开始充电，充电汽车 %v\n", chargePile.PileId, car.CarId)
		startTime := time.Now()
		// 通过随眠模拟充电,写在循环里方便实现删除订单时的提前结束充电
		for i := 1; i <= int(chargeTime.Seconds()); i++ {
			time.Sleep(time.Second - time.Millisecond)
			if IsInterrupt[station.StationId][chargePile.PileId] {
				IsInterrupt[station.StationId][chargePile.PileId] = false
				break
			}
		}
		// 实际充电时间
		realChargeTime := time.Now().Sub(startTime)
		// 充电完毕
		FreePileMutex[station.StationId].Lock()
		chargePile.Dequeue()                                // 移除一辆汽车
		FreePile[station.StationId][chargePile.PileId] += 1 // 队列空闲数+1
		FreePileMutex[station.StationId].Unlock()
		// 修改充电桩的等待时间
		chargePile.WaitingTime -= car.ChargeTime
		// 计算充电所使用的钱
		endTime := startTime.Add(realChargeTime)
		chargeCost := calculateFee(startTime, endTime, Power[chargePile.Mode])
		serviceCost := realChargeTime.Hours() * ServiceCostRate

		fmt.Printf("充电桩 %v 结束充电，充电汽车 %v %v, 充电费用 %v\n", chargePile.PileId, car.CarId, realChargeTime.String(), chargeCost)
		// 更新数据库订单中汽车开始充电时间、结束充电时间和充电费用
		var currentOrder user.Order
		for i := 0; i < 5; i++ {
			tx := global.GVA_DB.Model(&user.Order{}).Where("car_id = ? AND state <> '已完成'", car.CarId).First(&currentOrder)
			if tx.RowsAffected != 0 {
				break
			}
		}
		currentOrder.Time = realChargeTime.Hours()
		currentOrder.Kwh = currentOrder.Time * Power[car.Mode]
		currentOrder.StartedAt = startTime
		currentOrder.StopAt = endTime
		currentOrder.TotalCost = serviceCost + chargeCost
		currentOrder.ServiceCost = serviceCost
		currentOrder.ChargeCost = chargeCost
		currentOrder.State = "已完成"
		global.GVA_DB.Save(&currentOrder)
		fmt.Println("数据库修改完成，car_id = " + car.CarId + " 队列区到已完成")
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
	hasLock := waitingBlock.mu.TryLock()
	if !hasLock {
		return errors.New("系统繁忙，请稍后再试")
	}
	defer waitingBlock.mu.Unlock()
	if len(waitingBlock.Cars) < WaitingAreaSize {
		waitingBlock.Cars = append(waitingBlock.Cars, car)
		return nil
	} else {
		return errors.New("该充电站等待区已满，请稍后再试或更换充电站")
	}
}

// Delete 从等待区删除一辆汽车
func (waitingBlock *WaitingBlock) Delete(car Car) error {
	waitingBlock.mu.Lock()
	defer waitingBlock.mu.Unlock()
	for i := range waitingBlock.Cars {
		if car.CarId == waitingBlock.Cars[i].CarId {
			waitingBlock.Cars = append(waitingBlock.Cars[:i], waitingBlock.Cars[i+1:]...)
			return nil
		}
	}
	return errors.New("该车辆不在等待区中")
}

// Update 更新等待区的车辆信息(用户修改了充电请求，重新把用户加入队列)
func (waitingBlock *WaitingBlock) Update(car Car) error {
	waitingBlock.mu.Lock()
	defer waitingBlock.mu.Unlock()
	for index, curCar := range waitingBlock.Cars {
		if curCar.CarId == car.CarId {
			if curCar.Energy != car.Energy {
				// 只修改充电量
				waitingBlock.Cars[index].Energy = car.Energy
				waitingBlock.Cars[index].ChargeTime = car.Energy / Power[car.Mode]
				return nil
			} else if curCar.Mode != car.Mode {
				// 只修改充电模式,重新生成排队号
				waitingBlock.Cars = append(waitingBlock.Cars[:index], waitingBlock.Cars[index+1:]...)
				car.QueueId = QueuePrefix[car.Mode] + strconv.Itoa(CarNum[waitingBlock.StationId][car.Mode])

				waitingBlock.Cars = append(waitingBlock.Cars, car)
				return nil
			} else {
				return errors.New("只允许修改充电量或者充电模式")
			}
		}
	}
	return errors.New("等待区没有该汽车(汽车已经开始充电或者汽车车牌被修改)")
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
		for _, num := range FreePile[station.StationId] {
			if num != 0 {
				hasFree = true
				break
			}
		}
		if hasFree {
			// 从等待区中取出汽车
			for _, currentCar := range waitingBlock.Cars {
				fmt.Println(currentCar)
				minIndex = limits.INT_MAX
				min = float64(limits.INT_MAX)
				// 判断应该加入哪一个充电桩
				for i, pile := range station.ChargePiles {
					// 满足三个条件：充电模式匹配、充电时间最短、充电桩开启
					if pile.WaitingTime < min && pile.Mode == currentCar.Mode && IsOpen[station.StationId][pile.PileId] {
						min = pile.WaitingTime
						minIndex = i
					}
				}

				if minIndex != limits.INT_MAX && FreePile[station.StationId][minIndex] > 0 {
					// 更新数据库里面的订单信息
					// 修改订单在数据库里面的状态，从等待区变成队列区
					CarNum[currentCar.Mode][station.StationId] -= 1

					var currentOrder user.Order
					for i := 0; i < 5; i++ {
						tx := global.GVA_DB.Model(&user.Order{}).Where("car_id = ? AND state = '等待区'", currentCar.CarId).First(&currentOrder)
						if tx.RowsAffected != 0 {
							break
						}
						time.Sleep(1 * time.Second)
					}
					currentOrder.State = "队列区"
					currentOrder.PileId = minIndex // pileId存储的是该站中的充电桩下标
					currentOrder.Time = currentCar.ChargeTime
					global.GVA_DB.Save(&currentOrder)
					fmt.Println("修改数据库信息，等待区->队列区 car_id = " + currentCar.CarId)

					// 增加充电桩队列的等待时间
					station.ChargePiles[minIndex].WaitingTime += currentCar.ChargeTime

					// 把空闲数量-1，加锁防止冲突;将汽车从等待区取出加入充电队列
					FreePileMutex[station.StationId].Lock()
					fmt.Printf("充电桩%d 剩余空闲 %d\n", minIndex, FreePile[station.StationId][minIndex])
					FreePile[station.StationId][minIndex] -= 1
					fmt.Printf("从等待区取出一辆汽车加入充电桩%d\n", minIndex)
					station.ChargePiles[minIndex].Enqueue(waitingBlock.Dequeue())
					fmt.Printf("充电桩%d 剩余空闲 %d\n", minIndex, FreePile[station.StationId][minIndex])
					FreePileMutex[station.StationId].Unlock()

				}
			}
		}
		// 每5s请求一次
		time.Sleep(5 * time.Second)
	}
}

// GetCarInfoByOrder 从订单中获取车辆信息
func GetCarInfoByOrder(order user.Order) (car Car) {
	car.ChargeTime = order.Kwh / Power[Mode[order.ChargeType]]
	car.Mode = Mode[order.ChargeType]
	CarNum[order.StationId-1][car.Mode] += 1
	car.QueueId = QueuePrefix[car.Mode] + strconv.Itoa(CarNum[order.StationId-1][car.Mode])
	car.Energy = order.Kwh
	car.CarId = order.CarId
	return car
}
