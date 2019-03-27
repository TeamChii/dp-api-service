package entity

import "time"

type MiniposPaymentMethodEntity struct {
	Mpos_payment_method_id         int        `gorm:"column:mpos_payment_method_id; primary_key" json:"mpos_payment_method_id"`
	Mc_id                          int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_payment_method_name       string     `gorm:"column:mpos_payment_method_name" json:"mpos_payment_method_name"`
	Mpos_payment_method_detail     string     `gorm:"column:mpos_payment_method_detail" json:"mpos_payment_method_detail"`
	Mpos_payment_method_qrcode_ref string     `gorm:"column:mpos_payment_method_qrcode_ref" json:"mpos_payment_method_qrcode_ref"`
	Create_date                    *time.Time `gorm:"column:create_date" json:"create_date"`
	Create_by                      string     `gorm:"column:create_by" json:"create_by"`
	Update_date                    *time.Time `gorm:"column:update_date" json:"update_date"`
	Update_by                      string     `gorm:"column:update_by" json:"update_by"`
	Active                         string     `gorm:"column:active" json:"active"`
}

func (MiniposPaymentMethodEntity) TableName() string {
	return "dp_ms_minipos_payment_method"
}
