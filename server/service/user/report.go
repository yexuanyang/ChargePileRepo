package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/response"
)

type ReportService struct {
}

func (reportService *ReportService) GetDurationTotalCharge(report request.ReportSearch) (total float64, err error) {
	db := global.GVA_DB.Model(&user.Order{}).Where("user_id = ?", report.UserId)
	db = db.Where("created_at BETWEEN ? AND ?", report.Date, report.EndDate)
	result := db.Select("SUM(kwh) as total").Find(&total)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return 0, nil
		}
	}
	return total, nil
}

func (reportService *ReportService) GetDurationTotalPrice(report request.ReportSearch) (total float64, err error) {
	db := global.GVA_DB.Model(&user.Order{}).Where("user_id = ?", report.UserId)
	db = db.Where("created_at BETWEEN ? AND ?", report.Date, report.EndDate)
	result := db.Select("SUM(price) as total").Find(&total)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return 0, nil
		}
	}
	return total, nil
}

func (reportService *ReportService) GetDurationChargeInfo(report request.ChargePileReportSearch) (total response.ChargePileInfoResponse, err error) {
	db := global.GVA_DB.Model(&system.ChargePile{}).Where("id = ?", report.PileId)
	db = db.Where("created_at BETWEEN ? AND ?", report.Date, report.EndDate)
	tx := db.Select("SUM(charge_count) as chargeCount,SUM(charge_time) as chargeTime,SUM(electricity) as chargeElectricity").First(&total)
	if tx.Error != nil {
		if tx.RowsAffected == 0 {
			return response.ChargePileInfoResponse{}, nil
		}
	}
	return total, nil
}
