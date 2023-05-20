package ctype

import "encoding/json"

type ChargingType int

const (
	FastCharging ChargingType = 1
	SlowCharging ChargingType = 2
)

func (r ChargingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r ChargingType) String() string {
	var str string
	switch r {
	case FastCharging:
		str = "快充桩"
	case SlowCharging:
		str = "慢充桩"
	default:
		str = "其他"
	}
	return str
}
