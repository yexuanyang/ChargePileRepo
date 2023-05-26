package request

import (
	"time"
)

type ReportSearch struct {
	UserId  int       `json:"userId"`
	Date    time.Time `json:"date"`
	EndDate time.Time `json:"endDate"`
}

type ChargePileReportSearch struct {
	PileId  int       `json:"pileId"`
	Date    time.Time `json:"date"`
	EndDate time.Time `json:"endDate"`
}
