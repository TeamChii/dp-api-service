package entity

import "time"

type DeviceMerchantMapEntity struct {
	Device_uid      string     `gorm:"column:device_uid" json:"device_uid"`
	Mc_id           int        `gorm:"column:mc_id" json:"mc_id"`
	Device_name     string     `gorm:"column:device_name" json:"device_name"`
	Device_detail   string     `gorm:"column:device_detail" json:"device_detail"`
	Create_by       string     `gorm:"column:create_by" json:"create_by"`
	Create_date     *time.Time `gorm:"column:create_date" json:"create_date"`
	Create_date_str string     `gorm:"-" json:"create_date_str"`
	Update_by       string     `gorm:"column:update_by" json:"update_by"`
	Update_date     *time.Time `gorm:"column:update_date" json:"update_date"`
}

func (DeviceMerchantMapEntity) TableName() string {
	return "dp_mp_device_merchant"
}
