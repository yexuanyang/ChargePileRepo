package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
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
