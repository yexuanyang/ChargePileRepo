package user

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"math"
	"modernc.org/libc/limits"
	"strconv"
	"sync"
	"time"
)

var (
	// Mode 充电模式
	Mode = map[string]int{"F": 0, "T": 1}
	// Power 充电功率
	Power = []float64{30.0, 7.0}
	// ChargeStations 充电站集合
	ChargeStations []ChargeStation
	// FreePile 充电桩空闲位置数量，FreePile[1][1] = 2表示第2个充电站第2个充电桩有2个空闲位置
	FreePile      []map[int]int
	FreePileMutex []sync.Mutex
	// IsInterrupt 用来处理前端发送的取消订单操作,IsInterrupt[0][1] = true 表示第1个充电站的第2个充电桩需要提前中断充电
	IsInterrupt []map[int]bool
	// CarNum WAITING汽车数量 CarNum[1][1] = 2表示充电站2的WAITING中有2辆慢充汽车 CarNum[1][0] = 3 表示充电站2的WAITING有3辆快充汽车
	CarNum [][]int
	// QueuePrefix 队列号前缀，根据充电类型区分
	QueuePrefix = []string{0: "F", 1: "T"}
	// IsOpen 处理充电桩关闭的操作,IsOpen[0][2] = true表示第一个充电站的第三个充电桩开启
	IsOpen           []map[int]bool
	WaitingBlockBusy = false // WAITING是否忙（是否正在给故障队列调度）,忙的时候不允许向里面加数据

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
	// WaitingAreaSize WAITING大小
	WaitingAreaSize = 10
	// ChargingQueueLen 充电队列长度
	ChargingQueueLen = 2
	// ServiceCostRate 服务费
	ServiceCostRate = 0.8
)

type ChargeStation struct {
	StationId   int
	ChargePiles map[int]*ChargePile // pileId <-> chargePile(数据库中的pileId和充电桩对象的一一对应)
	Waiting     *WaitingBlock
}

type ChargePile struct {
	StationId   int
	PileId      int // 存储在数据库中的id
	Id          int //存储在充电站中的id位置
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

	FreePile = make([]map[int]int, number)
	IsInterrupt = make([]map[int]bool, number)
	IsOpen = make([]map[int]bool, number)
	for i := 0; i < int(number); i++ {
		FreePile[i] = make(map[int]int)
		IsInterrupt[i] = make(map[int]bool)
		IsOpen[i] = make(map[int]bool)
	}
	FreePileMutex = make([]sync.Mutex, number)
	CarNum = make([][]int, number)
	for i := range CarNum {
		// 列标0表示快充，列标1表示慢充
		CarNum[i] = make([]int, 2)
	}

	// 初始化充电站中充电桩列表
	for i := range ChargeStations {
		// 初始充电桩
		ChargeStations[i].ChargePiles = make(map[int]*ChargePile)
		// 初始WAITING车辆
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
		// 充电桩的Id存储在充电站内的下标
		pile.Id = len(ChargeStations[stationIndex].ChargePiles)
		// pileId存储在数据库中充电桩的id
		pile.PileId = int(chargePile.ID)
		// id存储在充电站中的下标
		fmt.Println(pile.PileId)
		pile.StationId = int(stationIndex)
		pile.Mode = Mode[chargePile.PileType]
		// 初始化充电桩充电队列内的汽车
		pile.Cars = make([]Car, 0, ChargingQueueLen)

		// 充电桩加入对应的充电站
		ChargeStations[stationIndex].ChargePiles[int(chargePile.ID)] = &pile

		// 初始化充电桩的充电队列空闲数量,FreePile[stationIndex]是对应充电站下的空闲列表
		FreePile[stationIndex][int(chargePile.ID)] = ChargingQueueLen

		// 每加入一个充电桩都需要增加一个新的IsInterrupt标志
		IsInterrupt[stationIndex][int(chargePile.ID)] = false

		// 不管充电桩开没开都要初始化一个线程
		// 没开的充电桩初始化的时候IsOpen是false
		IsOpen[stationIndex][int(chargePile.ID)] = chargePile.IsOpen

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
		// 查看是否有车在充电队列，没有车就空闲
		// 充电桩如果没有开启，那么等待队列不会向该充电桩的充电队列增加汽车
		// 没车，等一段时间之后再次查询是否有车
		if len(chargePile.Cars) == 0 {
			time.Sleep(3 * time.Second)
			continue
		}

		car := chargePile.Cars[0]
		chargeTime, _ := time.ParseDuration(strconv.FormatFloat(car.ChargeTime, 'f', 5, 64) + "h")
		fmt.Printf("充电桩 %v 开始充电，充电汽车 %v\n", chargePile.PileId, car.CarId)
		global.GVA_DB.Model(&system.Order{}).Where("car_id = ? AND state <> 'CHARGING'", car.CarId).Update("state", "CHARGING")
		startTime := time.Now()
		// 通过随眠模拟充电,写在循环里方便实现删除订单时的提前结束充电
		var i int
		for i = 1; i <= int(chargeTime.Seconds()/20); i++ {
			time.Sleep(time.Second - time.Millisecond)
			if IsInterrupt[station.StationId][chargePile.PileId] {
				IsInterrupt[station.StationId][chargePile.PileId] = false
				break
			}
		}
		// 实际充电时间
		realChargeTime := time.Now().Sub(startTime) * 20

		// 修改充电桩的等待时间
		chargePile.WaitingTime -= car.ChargeTime
		// 计算充电所使用的钱
		endTime := startTime.Add(realChargeTime)
		chargeCost := calculateFee(startTime, endTime, Power[chargePile.Mode])
		serviceCost := realChargeTime.Hours() * ServiceCostRate

		fmt.Printf("充电桩 %v 结束充电，充电汽车 %v %v, 充电费用 %v\n", chargePile.PileId, car.CarId, realChargeTime.String(), chargeCost)
		// 更新数据库订单中汽车开始充电时间、结束充电时间和充电费用、总充电时间
		var currentOrder system.Order
		for i := 0; i < 5; i++ {
			tx := global.GVA_DB.Model(&system.Order{}).Where("car_id = ? AND state <> 'FINISHED'", car.CarId).First(&currentOrder)
			if tx.RowsAffected != 0 {
				break
			}
		}
		currentOrder.Time += realChargeTime.Hours()
		currentOrder.Kwh += currentOrder.Time * Power[car.Mode]
		currentOrder.StartedAt = startTime
		currentOrder.StopAt = endTime
		currentOrder.TotalCost += serviceCost + chargeCost
		currentOrder.ServiceCost += serviceCost
		currentOrder.ChargeCost += chargeCost
		currentOrder.PileId = chargePile.PileId
		var cp system.ChargePile
		tx := global.GVA_DB.First(&cp, "id = ?", chargePile.PileId)
		if tx.RowsAffected != 0 {
			cp.ChargeCost += currentOrder.ChargeCost
			cp.Electricity += currentOrder.Kwh
			cp.ChargeCount += 1
			cp.ServiceCost += currentOrder.ServiceCost
			cp.TotalCost += currentOrder.TotalCost
			cp.ChargeTime += realChargeTime.Hours()
			global.GVA_DB.Save(&cp)
		} else {
			global.GVA_LOG.Error("订单结束时向充电桩中写入数据时出错")
		}
		if i != int(chargeTime.Seconds()/20)+1 {
			// 提前中断，有两种情况，一种是删除订单，一种是充电桩故障
			// 删除订单需要把state设置为finished

			// 充电桩故障
			// 退出充电的时候要采用故障调度，直接将WAITING拷贝出来，然后将故障队列的车放到WAITING，同时将WAITING上锁不允许车辆加入，等到故障队列的车都调度完毕之后将原来的WAITING复制回来并解锁
			if !IsOpen[station.StationId][chargePile.PileId] {
				// 把第一个正在充电的车移除
				chargePile.Dequeue()
				// 计算这段时间这辆车充电的信息，赋值给新的车对象
				newCar := car
				// 剩余的充电时间减少
				newCar.ChargeTime -= realChargeTime.Hours()
				// 剩余的充电量减少
				newCar.Energy = newCar.ChargeTime * Power[car.Mode]

				// 操作数据库，把正在充电的订单状态改为WAITING,订单的充电桩号改成0
				currentOrder.State = "WAITING"
				currentOrder.PileId = 0
				global.GVA_DB.Save(&currentOrder)
				// 如果队列里面有第二辆车，需要把第二个订单的车的状态也更新成WAITING
				if len(chargePile.Cars) > 0 {
					var o system.Order
					global.GVA_DB.Model(&system.Order{}).Where("car_id = ?", chargePile.Cars[0].CarId).First(&o)
					o.State = "WAITING"
					o.PileId = 0
					global.GVA_DB.Save(&o)
				}

				fmt.Println(fmt.Sprintf("充电桩 %d 故障", chargePile.PileId))

				// WAITING不让新汽车进入，优先调度故障充电桩队列的车
				WaitingBlockBusy = true
				wBlock := station.Waiting.Cars
				// 重置充电桩的等待时间，方便故障恢复的时候充电桩正常工作
				chargePile.WaitingTime = 0
				// 把故障充电桩充电队列内的车都加入等待区，优先重新调度
				station.Waiting.Cars = chargePile.Cars
				// 充电队列车清零
				chargePile.Cars = nil
				station.Waiting.Cars = append(station.Waiting.Cars, newCar)
				fmt.Println(station.Waiting.Cars)
				for {
					if len(station.Waiting.Cars) == 0 {
						break
					}
					time.Sleep(5 * time.Second)
					fmt.Println("故障队列还在调度，故障队列：")
					fmt.Println(station.Waiting)
				}
				station.Waiting.Cars = wBlock
				fmt.Println(station.Waiting.Cars)
				// 解除限制，新汽车可以加入WAITING
				WaitingBlockBusy = false

			} else {
				//充电桩没有故障，中断是因为取消订单，订单状态变成已完成
				currentOrder.State = "FINISHED"
			}
		} else {
			// 正常充完电
			FreePileMutex[station.StationId].Lock()
			// 移除一辆汽车（正在充电的汽车将会被移除队列，如果充电桩故障，需要重新将新的车辆信息加入WAITING）
			chargePile.Dequeue()
			// 队列空闲数+1
			FreePile[station.StationId][chargePile.PileId] += 1
			fmt.Printf("充电桩 %d 空闲数量 %d\n", chargePile.PileId, FreePile[station.StationId][chargePile.PileId])
			FreePileMutex[station.StationId].Unlock()

			currentOrder.State = "FINISHED"
			fmt.Println("数据库修改完成，car_id = " + car.CarId + " CHARGING到FINISHED")
		}
		global.GVA_DB.Save(&currentOrder)
	}

}

// Dequeue 从WAITING汽车队列中取出一辆汽车
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

// Enqueue 将汽车加入WAITING
func (waitingBlock *WaitingBlock) Enqueue(car Car) error {
	if WaitingBlockBusy {
		return errors.New("系统繁忙，请稍后再试")
	}
	waitingBlock.mu.Lock()
	defer waitingBlock.mu.Unlock()
	if len(waitingBlock.Cars) < WaitingAreaSize {
		waitingBlock.Cars = append(waitingBlock.Cars, car)
		return nil
	} else {
		return errors.New("该充电站WAITING已满，请稍后再试或更换充电站")
	}
}

// Delete 从WAITING删除一辆汽车
func (waitingBlock *WaitingBlock) Delete(car Car) error {
	waitingBlock.mu.Lock()
	defer waitingBlock.mu.Unlock()
	for i := range waitingBlock.Cars {
		if car.CarId == waitingBlock.Cars[i].CarId {
			waitingBlock.Cars = append(waitingBlock.Cars[:i], waitingBlock.Cars[i+1:]...)
			return nil
		}
	}
	return errors.New("该车辆不在WAITING中")
}

// Update 更新WAITING的车辆信息(用户修改了充电请求，重新把用户加入队列)
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
	return errors.New("WAITING没有该汽车(汽车已经开始充电或者汽车车牌被修改)")
}

// DispatchCar WAITING线程执行的程序,WAITING的汽车持续向充电汽车请求加入汽车
func (waitingBlock *WaitingBlock) DispatchCar(station *ChargeStation) {
	for {
		var minIndex int
		var min float64
		if len(waitingBlock.Cars) == 0 {
			time.Sleep(5 * time.Second)
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
			// 从WAITING中取出汽车
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

				// 访问临界资源，给临界资源上锁
				FreePileMutex[station.StationId].Lock()
				if minIndex != limits.INT_MAX && FreePile[station.StationId][minIndex] > 0 {
					// 更新数据库里面的订单信息
					// 修改订单在数据库里面的状态，从WAITING变成DISPATCHED
					CarNum[currentCar.Mode][station.StationId] -= 1

					var currentOrder system.Order
					for i := 0; i < 5; i++ {
						tx := global.GVA_DB.Model(&system.Order{}).Where("car_id = ? AND state <> 'FINISHED'", currentCar.CarId).First(&currentOrder)
						if tx.RowsAffected != 0 {
							break
						}
						time.Sleep(1 * time.Second)
					}
					currentOrder.State = "DISPATCHED"
					currentOrder.PileId = minIndex // pileId存储的是数据库充电桩表中的充电桩下标
					global.GVA_DB.Save(&currentOrder)
					fmt.Println("修改数据库信息，WAITING->DISPATCHED car_id = " + currentCar.CarId)

					// 增加充电桩队列的等待时间
					station.ChargePiles[minIndex].WaitingTime += currentCar.ChargeTime

					fmt.Printf("充电桩%d 剩余空闲 %d\n", minIndex, FreePile[station.StationId][minIndex])
					FreePile[station.StationId][minIndex] -= 1
					fmt.Printf("从WAITING取出一辆汽车加入充电桩%d\n", minIndex)
					station.ChargePiles[minIndex].Enqueue(currentCar)
					waitingBlock.Dequeue()
					fmt.Printf("充电桩%d 剩余空闲 %d\n", minIndex, FreePile[station.StationId][minIndex])
				}
				FreePileMutex[station.StationId].Unlock()
			}
		}
		// 每5s请求一次
		time.Sleep(5 * time.Second)
	}
}

// GetCarInfoByOrder 从订单中获取车辆信息
func GetCarInfoByOrder(order system.Order) (car Car) {
	car.ChargeTime = order.ApplyKwh / Power[Mode[order.ChargeType]]
	car.Mode = Mode[order.ChargeType]

	car.QueueId = QueuePrefix[car.Mode] + strconv.Itoa(CarNum[order.StationId-1][car.Mode])
	car.Energy = order.ApplyKwh
	if order.CarId != "" {
		car.CarId = order.CarId
	} else {
		var carId string
		global.GVA_DB.Model(&system.Order{}).Select("car_id").First(&carId, "id = ?", order.ID)
		car.CarId = carId
	}

	return car
}
