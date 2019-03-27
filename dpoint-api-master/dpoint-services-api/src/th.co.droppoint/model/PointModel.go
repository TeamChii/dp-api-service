package model

import "time"

type PointReq struct {
	Mc_id  int         `json:"mc_id"`
	Paging PagingModel `json:"paging"`
}
type PointCheckMobileReq struct {
	Mc_id       int    `json:"mc_id"`
	Cust_mobile string `json:"cust_mobile"`
}
type PointEntityReq struct {
	Point_id            int        `gorm:"column:point_id; primary_key" json:"point_id"`
	Mc_id               int        `gorm:"column:mc_id" json:"mc_id"`
	Cust_id             int        `gorm:"column:cust_id" json:"cust_id"`
	Container_id        int        `gorm:"column:container_id" json:"container_id"`
	Menu_id             int        `gorm:"column:menu_id" json:"menu_id"`
	Transfer_to_cust_id int        `gorm:"column:transfer_to_cust_id" json:"transfer_to_cust_id"`
	Point_amt           int        `gorm:"column:point_amt" json:"point_amt"`
	Create_by           string     `gorm:"column:create_by" json:"create_by"`
	Create_date         *time.Time `gorm:"column:create_date" json:"create_date"`
	Expire_date         string     `gorm:"column:expire_date" json:"expire_date"`
	Expire_flag         string     `gorm:"column:expire_flag" json:"expire_flag"`
	Status              string     `gorm:"-" json:"status"`
	Cust_mobile         string     `gorm:"-" json:"cust_mobile"`
}
