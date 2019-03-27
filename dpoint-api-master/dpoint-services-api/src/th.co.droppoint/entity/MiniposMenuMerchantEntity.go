package entity

import "time"

type MiniposMenuMerchantEntity struct {
	Mc_id            int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_menu_id     int        `gorm:"column:mpos_menu_id" json:"mpos_menu_id"`
	Mpos_menu_detail string     `gorm:"column:mpos_menu_detail" json:"mpos_menu_detail"`
	Create_date      *time.Time `gorm:"column:create_date" json:"create_date"`
	Create_by        string     `gorm:"column:create_by" json:"create_by"`
	Update_date      *time.Time `gorm:"column:update_date" json:"update_date"`
	Update_by        string     `gorm:"column:update_by" json:"update_by"`
	Active           string     `gorm:"column:active" json:"active"`
}

func (MiniposMenuMerchantEntity) TableName() string {
	return "dp_mp_minipos_menu_merchant"
}
