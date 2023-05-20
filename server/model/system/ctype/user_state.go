package ctype

import "encoding/json"

type State int

const (
	ChargingUser State = 1
	WaitingUser  State = 2
	QueuingUser  State = 3
)

func (r State) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r State) String() string {
	var str string
	switch r {
	case ChargingUser:
		str = "充电区充电中"
	case WaitingUser:
		str = "充电区等待中"
	case QueuingUser:
		str = "排队中"
	default:
		str = "其他"
	}
	return str
}
