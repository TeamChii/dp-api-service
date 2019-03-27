package model

import (
	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	//"th.co.droppoint/entity"
)

type RequestPointReq struct {
	Mc_id         int         `gorm:"column:mc_id" json:"mc_id"`
	Category_type string      `gorm:"column:category_type" json:"category_type"`
	Reqeust_type  string      `gorm:"column:reqeust_type" json:"reqeust_type"`
	Paging        PagingModel `json:"paging"`
}
type RequestPointAddReq struct {
	Mc_id           int               `gorm:"column:mc_id" json:"mc_id"`
	Request_message string            `gorm:"column:request_message" json:"request_message"`
	Cust_id         int               `gorm:"column:cust_id" json:"cust_id"`
	ContainerDetail []ContainerDetail `gorm:"-" json:"container"`
}
type RequestPointAddReq2 struct {
	Mc_id           int               `gorm:"column:mc_id" json:"mc_id"`
	Request_message string            `gorm:"column:request_message" json:"request_message"`
	CustAr          []CustAr          `gorm:"-" json:"customer"`
	ContainerDetail []ContainerDetail `gorm:"-" json:"container"`
}
type CustAr struct {
	Cust_id          int `gorm:"column:cust_id" json:"cust_id"`
	Request_point_id int `gorm:"column:request_point_id" json:"request_point_id"`
}
type ContainerDetail struct {
	Container_id          int        `gorm:"column:container_id" json:"container_id"`
	Issue_date            *time.Time `gorm:"column:issue_date" json:"issue_date"`
	Expire_date           *time.Time `gorm:"column:expire_date" json:"expire_date"`
	Cust_tag              string     `gorm:"column:cust_tag" json:"cust_tag"`
	Cust_frg              string     `gorm:"column:cust_frg" json:"cust_frg"`
	Cust_status           string     `gorm:"column:cust_status" json:"cust_status"`
	Cust_first_visit_date *time.Time `gorm:"column:cust_first_visit_date" json:"cust_first_visit_date"`
	Cust_last_visit_date  *time.Time `gorm:"column:cust_last_visit_date" json:"cust_last_visit_date"`
	Create_by             string     `gorm:"column:create_by" json:"create_by"`
	Update_by             string     `gorm:"column:update_by" json:"update_by"`
	Create_date           *time.Time `gorm:"column:create_date" json:"create_date"`
	Update_date           *time.Time `gorm:"column:update_date" json:"update_date"`
	Container_type        string     `gorm:"-" json:"container_type"`
	Expire_mode           string     `gorm:"-" json:"expire_mode"`
	Expire_value          string     `gorm:"-" json:"expire_value"`
}
type RequestPointEntityReq struct {
	Request_point_id int                    `gorm:"column:request_point_id; primary_key" json:"request_point_id"`
	Mc_id            int                    `gorm:"column:mc_id" json:"mc_id"`
	Cust_id          int                    `gorm:"column:cust_id" json:"cust_id"`
	CustomerEntity   *entity.CustomerEntity `gorm:"-" json:"customer"`
	Container_id     int                    `gorm:"column:container_id" json:"container_id"`
	Request_status   string                 `gorm:"column:request_status" json:"request_status"`
	Request_date     *time.Time             `gorm:"column:request_date" json:"request_date"`
	Request_date_str string                 `gorm:"-" json:"request_date_str"`
	Reqeust_type     string                 `gorm:"column:reqeust_type" json:"reqeust_type"`
	Request_message  string                 `gorm:"column:request_message" json:"request_message"`
	Category_type    string                 `gorm:"column:category_type" json:"category_type"`
}

func (RequestPointEntityReq) TableName() string {
	return "dp_tb_request_point"
}
