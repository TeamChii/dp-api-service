package entity

import "time"

type RequestPointEntity struct {
	Request_point_id int        `gorm:"column:request_point_id; primary_key" json:"request_point_id"`
	Mc_id            int        `gorm:"column:mc_id" json:"mc_id"`
	Cust_id          int        `gorm:"column:cust_id" json:"cust_id"`
	Container_id     int        `gorm:"column:container_id" json:"container_id"`
	Request_status   string     `gorm:"column:request_status" json:"request_status"`
	Request_date     *time.Time `gorm:"column:request_date" json:"request_date"`
	Reqeust_type     string     `gorm:"column:reqeust_type" json:"reqeust_type"`
	Request_message  string     `gorm:"column:request_message" json:"request_message"`
	Category_type    string     `gorm:"column:category_type" json:"category_type"`
}

func (RequestPointEntity) TableName() string {
	return "dp_tb_request_point"
}
