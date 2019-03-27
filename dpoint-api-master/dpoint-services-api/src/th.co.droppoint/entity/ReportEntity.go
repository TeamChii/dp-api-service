package entity

import "time"

type ReportEntity struct {
	Report_id       int        `gorm:"column:report_id" json:"report_id"`
	Mc_id           int        `gorm:"column:mc_id" json:"mc_id"`
	Report_category string     `gorm:"column:report_category" json:"report_category"`
	Customer_id     int        `gorm:"column:customer_id" json:"customer_id"`
	Mobile_no       string     `gorm:"column:mobile_no" json:"mobile_no"`
	Amt             int        `gorm:"column:amt" json:"amt"`
	Unit            string     `gorm:"column:unit" json:"unit"`
	Report_detail   string     `gorm:"column:report_detail" json:"report_detail"`
	Date_date       *time.Time `gorm:"column:date_date" json:"date_date"`
	Create_by       string     `gorm:"column:create_by" json:"create_by"`
	Create_date     *time.Time `gorm:"column:create_date" json:"create_date"`
}

func (ReportEntity) TableName() string {
	return "dp_rp_report"
}
