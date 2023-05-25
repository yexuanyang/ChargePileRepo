package user

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/response"
)

type ReportService struct {
}

// GetDurationTotalCharge 获得一段时间的充电总度数
func (reportService *ReportService) GetDurationTotalCharge(report request.ReportSearch) (total float64, err error) {
	db := global.GVA_DB.Model(&user.Order{}).Where("user_id = ?", report.UserId)
	db = db.Where("created_at BETWEEN ? AND ?", report.Date, report.EndDate)
	result := db.Select("CAST(SUM(kwh) as DECIMAL(10,2)) as total").Find(&total)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return 0, nil
		}
	}
	return total, nil
}

// GetDurationTotalPrice 获得一段时间的充电总金额
func (reportService *ReportService) GetDurationTotalPrice(report request.ReportSearch) (total float64, err error) {
	db := global.GVA_DB.Model(&user.Order{}).Where("user_id = ?", report.UserId)
	db = db.Where("created_at BETWEEN ? AND ?", report.Date, report.EndDate)
	result := db.Select("CAST(SUM(total_cost) as DECIMAL(10,2)) as total").Find(&total)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return 0, nil
		}
	}
	return total, nil
}

// GetDurationChargeInfo 获得充电桩一段时间的充电信息
func (reportService *ReportService) GetDurationChargeInfo(report request.ChargePileReportSearch) (total response.ChargePileInfoResponse, err error) {
	db := global.GVA_DB.Model(&system.ChargePile{}).Where("id = ?", report.PileId)
	db = db.Where("created_at BETWEEN ? AND ?", report.Date, report.EndDate)
	db = db.Select("id as pileId,SUM(charge_count) as chargeCount,CAST(SUM(charge_time) as DECIMAL(10,2)) as chargeTime,CAST(SUM(electricity) as DECIMAL(10,2))as chargeElectricity")
	tx := db.Group("id").First(&total)
	if tx.Error != nil {
		if tx.RowsAffected == 0 {
			return response.ChargePileInfoResponse{}, errors.New("没有找到符合条件的记录")
		}
	}
	return total, nil
}

// GetDurationReportInfo 获得一段时间的订单报表,按照日期来分组
// 返回的结构体中含有日期、总充电量、总金额
func (reportService *ReportService) GetDurationReportInfo(report request.ReportSearch) (res []response.OrderReportResponse, err error) {
	const selectSQL = "DATE_FORMAT( created_at, '%Y-%m-%d' ) AS date,CAST(SUM( kwh ) AS DECIMAL(10,2)) AS totalKwh,CAST(SUM( total_cost ) AS DECIMAL(10,2)) AS totalCost"
	db := global.GVA_DB.Model(&user.Order{}).Select(selectSQL)
	db = db.Where("created_at BETWEEN ? AND ?", report.Date, report.EndDate).Where("user_id = ?", report.UserId).Group("DATE_FORMAT( created_at, '%Y-%m-%d' )")
	tx := db.Order("DATE_FORMAT( created_at, '%Y-%m-%d' ) DESC").Debug().Find(&res)
	if tx.Error != nil {
		if tx.RowsAffected == 0 {
			return nil, errors.New("没有符合条件的记录")
		}
	}
	fmt.Println(res)
	return res, nil
}
