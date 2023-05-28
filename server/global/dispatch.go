package global

import (
	"modernc.org/libc/limits"
	"time"
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
	lowStart    = 23  // 谷时开始时间
	lowEnd      = 7   // 谷时结束时间
)

type Order struct {
	GVA_MODEL
	UserId      int       `json:"userId" form:"userId" gorm:"column:user_id;comment:订单所属的用户id;"`
	CarId       string    `json:"carId" form:"carId" gorm:"column:car_id;comment:充电的车牌号;size:20;"`
	ChargeType  string    `json:"chargeType" form:"chargeType" gorm:"column:charge_type;type:enum('快充','慢充','其他');comment:充电类型;"`
	ChargeCost  float64   `json:"chargeCost" form:"chargeCost" gorm:"column:charge_cost;comment:;size:22;"`
	Kwh         float64   `json:"kwh" form:"kwh" gorm:"column:kwh;comment:;size:22;"`
	Time        float64   `json:"time" form:"time" gorm:"column:time;comment:;"`
	PileId      int       `json:"pileId" form:"pileId" gorm:"column:pile_id;comment:;size:20;"`
	ServiceCost float64   `json:"serviceCost" form:"serviceCost" gorm:"column:service_cost;comment:;size:22;"`
	StartedAt   time.Time `json:"startedAt" form:"startedAt" gorm:"column:started_at;comment:;"`
	StopAt      time.Time `json:"stopAt" form:"stopAt" gorm:"column:stop_at;comment:;"`
	TotalCost   float64   `json:"totalCost" form:"totalCost" gorm:"column:total_cost;comment:;size:22;"`
}

func (chargeStation *ChargeStation) Dispatch(order Order) (retOrder Order) {
	// 获取等待区的车辆数目
	waitingNum := len(chargeStation.Waiting.Cars)
	// 等待区满
	if waitingNum >= WaitingAreaSize {
		return
	}
	// 获取当前订单的车辆信息
	currentCar := Car{CarId: order.CarId, Energy: order.Kwh, Mode: Mode[order.ChargeType]}

	// 订单请求的是快充
	if currentCar.Mode == 0 {
		// 获取充电时间
		currentCar.ChargeTime = currentCar.Energy / FastPower
		order.Time = currentCar.ChargeTime
		order.ServiceCost = currentCar.Energy * ServiceCostRate

		// 获得充电队列有空闲位置的快充充电桩列表
		freeFastChargePileList, err := chargeStation.getFreeFastChargePileList()
		if err != nil {
			panic(err)
			return
		}
		// 没有空闲的快充充电桩
		if len(freeFastChargePileList) == 0 {
			chargeStation.Waiting.Cars = append(chargeStation.Waiting.Cars, currentCar)
		} else { //有空闲充电桩，将车辆调度到等待时间最短的空间充电桩
			// 得到等待时间最短的充电桩
			currentChargePile, pileId := getMinWaitingTime(freeFastChargePileList)
			order.PileId = pileId

			// 把车辆加入充电桩等待队列
			currentChargePile.Cars[currentChargePile.Length] = currentCar
			currentChargePile.Length += 1
			// 获取车辆的开始充电时间
			currentTime := time.Now().Unix()
			order.StartedAt = time.Unix(currentTime+int64(currentChargePile.WaitingTime*3600), 0)
			order.StopAt = time.Unix(order.StartedAt.Unix()+int64(currentCar.ChargeTime*3600), 0)
			order.ChargeCost = calculateFee(order.StartedAt, order.StopAt, FastPower)
			order.TotalCost = order.ChargeCost + order.ServiceCost
			// 更新充电桩的waitingTime
			currentChargePile.WaitingTime += currentCar.ChargeTime

			retOrder = Order(order)
			return retOrder
		}

	} else { //订单请求的是慢充
		currentCar.ChargeTime = currentCar.Energy / TrickPower
		order.Time = currentCar.ChargeTime
		order.ServiceCost = currentCar.Energy * ServiceCostRate

		// 获得充电队列有空闲位置的慢充充电桩列表
		freeFastChargePileList, err := chargeStation.getFreeTrickChargePileList()
		if err != nil {
			panic(err)
			return
		}
		// 没有空闲的快充充电桩
		if len(freeFastChargePileList) == 0 {
			chargeStation.Waiting.Cars = append(chargeStation.Waiting.Cars, currentCar)
		} else { //有空闲充电桩，将车辆调度到等待时间最短的空间充电桩
			// 得到等待时间最短的充电桩
			currentChargePile, pileId := getMinWaitingTime(freeFastChargePileList)
			order.PileId = pileId

			// 把车辆加入充电桩等待队列
			currentChargePile.Cars[currentChargePile.Length] = currentCar
			currentChargePile.Length += 1
			// 获取车辆的开始充电时间
			currentTime := time.Now().Unix()
			order.StartedAt = time.Unix(currentTime+int64(currentChargePile.WaitingTime*3600), 0)
			order.StopAt = time.Unix(order.StartedAt.Unix()+int64(currentCar.ChargeTime*3600), 0)
			order.ChargeCost = calculateFee(order.StartedAt, order.StopAt, TrickPower)
			order.TotalCost = order.ChargeCost + order.ServiceCost
			// 更新充电桩的waitingTime
			currentChargePile.WaitingTime += currentCar.ChargeTime

			retOrder = Order(order)
			return retOrder
		}
	}
	return Order{}
}

// getFreeFastChargePileList 获取充电站内充电队列有空闲的所有快充充电桩
func (chargeStation *ChargeStation) getFreeFastChargePileList() (chargePileList []*ChargePile, err error) {
	for i := 0; i < len(chargeStation.FastChargePiles); i++ {
		if chargeStation.FastChargePiles[i].Length < ChargingQueueLen {
			chargePileList = append(chargePileList, &chargeStation.FastChargePiles[i])
		}
	}
	return chargePileList, nil
}

// getFreeTrickChargePileList 获取充电站内充电队列有空闲的所有慢充充电桩
func (chargeStation *ChargeStation) getFreeTrickChargePileList() (chargePileList []*ChargePile, err error) {
	for i := 0; i < len(chargeStation.TrickleChargingPiles); i++ {
		if chargeStation.FastChargePiles[i].Length < ChargingQueueLen {
			chargePileList = append(chargePileList, &chargeStation.TrickleChargingPiles[i])
		}
	}
	return chargePileList, nil
}

// getMinWaitingTime 获得充电桩队列中的等待时间最短的充电桩
func getMinWaitingTime(chargePileList []*ChargePile) (chargePile *ChargePile, minIndex int) {
	min := float64(limits.INT_MAX)
	for index, item := range chargePileList {
		if item.WaitingTime < min {
			min = item.WaitingTime
			minIndex = index
			chargePile = item
		}
	}
	return chargePile, minIndex
}

func calculateFee(start, end time.Time, power float64) float64 {
	var fee float64
	subHours := end.Sub(start).Hours() // 计算充电时间(小时)
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
