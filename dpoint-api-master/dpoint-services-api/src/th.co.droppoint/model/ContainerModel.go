package model

import (
	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	//"th.co.droppoint/entity"
)

type ContainerReq struct {
	Mc_id          int         `json:"mc_id"`
	Contaienr_type string      `json:"contaienr_type"`
	Active_Status  string      `json:"active_status"`
	Paging         PagingModel `json:"paging"`
}

type ContainerReq2 struct {
	Mc_id          int         `json:"mc_id"`
	Cust_id        int         `json:"cust_id"`
	Contaienr_type string      `json:"contaienr_type"`
	Paging         PagingModel `json:"paging"`
}
type ContainerByIdReq struct {
	Mc_id        int `json:"mc_id"`
	Container_id int `gorm:"-" json:"container_id"`
	Cust_id      int `json:"cust_id"`
}
type ContainerEntityReq struct {
	Container_id              int                            `gorm:"column:container_id; primary_key" json:"container_id"`
	Container_type_id         int                            `gorm:"column:container_type_id" json:"container_type_id"`
	ContainerTypeEntity       *entity.ContainerTypeEntity    `gorm:"column:container_type_id;ForeignKey:container_type_id;AssociationForeignKey:container_type_id" json:"containerType"`
	Mc_id                     int                            `gorm:"column:mc_id" json:"mc_id"`
	Cust_id                   int                            `gorm:"-" json:"cust_id"`
	Layout_template           string                         `gorm:"column:layout_template" json:"layout_template"`
	Layout_color              string                         `gorm:"column:layout_color" json:"layout_color"`
	Container_subject         string                         `gorm:"column:container_subject" json:"container_subject"`
	Container_detail          string                         `gorm:"column:container_detail" json:"container_detail"`
	Image_ref                 string                         `gorm:"column:image_ref" json:"image_ref"`
	Image_ref_id              int                            `gorm:"column:image_ref_id" json:"image_ref_id"`
	Container_term_conditions string                         `gorm:"column:container_term_conditions" json:"container_term_conditions"`
	Expire_mode               string                         `gorm:"column:expire_mode" json:"expire_mode"`
	Point_amt                 int                            `gorm:"-" json:"point_amt"`
	ContainerRewardEntity     []entity.ContainerRewardEntity `gorm:"-" json:"reward"`
	Total_point               int                            `gorm:"-" json:"total_point"`
	Expire_value              string                         `gorm:"expire_value" json:"expire_value"`
	Allow_borrow              string                         `gorm:"column:allow_borrow" json:"allow_borrow"`
	Transferable              string                         `gorm:"column:transferable" json:"transferable"`
	Create_by                 string                         `gorm:"column:create_by" json:"create_by"`
	Create_date               *time.Time                     `gorm:"column:create_date" json:"create_date"`
	Create_date_str           string                         `gorm:"-" json:"create_date_str"`
	Update_by                 string                         `gorm:"column:update_by" json:"update_by"`
	Update_date               *time.Time                     `gorm:"column:update_date" json:"update_date"`
	Update_date_str           string                         `gorm:"-" json:"update_date_str"`
	Customer_Flag             string                         `gorm:"column:customer_flag" json:"customer_flag"`
	Active_Status             string                         `gorm:"column:active_status" json:"active_status"`

	Card_Usage_Every string `gorm:"column:card_usage_every" json:"card_usage_every"`
	Card_Usage_Once  string `gorm:"column:card_usage_once" json:"card_usage_once"`
	ItemTime         int    `gorm:"column:item_time" json:"item_time"`

	Expire_flag string `gorm:"-" json:"expire_flag"`
}

func (ContainerEntityReq) TableName() string {
	return "dp_ms_container"
}
