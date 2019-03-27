package entity

import "time"

type PackageEntity struct {
	Package_id        int        `gorm:"column:package_id; primary_key" json:"package_id"`
	Package_name      string     `gorm:"column:package_name" json:"package_name"`
	Max_container_amt int        `gorm:"column:max_container_amt" json:"max_container_amt"`
	Create_by         string     `gorm:"column:create_by" json:"create_by"`
	Update_by         string     `gorm:"column:update_by" json:"update_by"`
	Create_date       *time.Time `gorm:"column:create_date" json:"create_date"`
	Update_date       *time.Time `gorm:"column:update_date" json:"update_date"`
	Package_desc      string     `gorm:"column:package_desc" json:"package_desc"`
	Color             string     `gorm:"column:color" json:"color"`
}

func (PackageEntity) TableName() string {
	return "dp_ms_package"
}
