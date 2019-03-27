package entity

import (
	"time"
)

type MiniposMenuCategoryEntity struct {
	Mpos_menu_category_id     int        `gorm:"column:mpos_menu_category_id; primary_key" json:"mpos_menu_category_id"`
	Mc_id                     int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_menu_category_name   string     `gorm:"column:mpos_menu_category_name" json:"mpos_menu_category_name"`
	Mpos_menu_category_detail string     `gorm:"column:mpos_menu_category_detail" json:"mpos_menu_category_detail"`
	Create_date               *time.Time `gorm:"column:create_date" json:"create_date"`
	Create_by                 string     `gorm:"column:create_by" json:"create_by"`
	Update_date               *time.Time `gorm:"column:update_date" json:"update_date"`
	Update_by                 string     `gorm:"column:update_by" json:"update_by"`
	Active                    string     `gorm:"column:active" json:"active"`
	Ord                       int 		 `gorm:"column:ord" json:"ord"`
}

func (MiniposMenuCategoryEntity) TableName() string {
	return "dp_ms_minipos_menu_category"
}
