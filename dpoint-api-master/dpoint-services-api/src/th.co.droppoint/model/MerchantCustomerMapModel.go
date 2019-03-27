package model

import (
	"time"
)

type MerchantCustomerMapReq struct {
	Mc_id          int         `json:"mc_id"`
	Cust_id        int         `json:"cust_id"`
	Contaienr_type string      `json:"contaienr_type"`
	Paging         PagingModel `json:"paging"`
}

type MerchantCustomerMapEntity struct {
	Mc_id                 int                 `gorm:"column:mc_id" json:"mc_id"`
	Cust_id               int                 `gorm:"column:cust_id" json:"cust_id"`
	Container_id          int                 `gorm:"column:container_id" json:"container_id"`
	ContainerEntityReq    *ContainerEntityReq `gorm:"column:container_id;ForeignKey:container_id;AssociationForeignKey:container_id" json:"container"`
	Issue_date            *time.Time          `gorm:"column:issue_date" json:"issue_date"`
	Expire_date           *time.Time          `gorm:"column:expire_date" json:"expire_date"`
	Cust_tag              string              `gorm:"column:cust_tag" json:"cust_tag"`
	Cust_frg              string              `gorm:"column:cust_frg" json:"cust_frg"`
	Cust_status           string              `gorm:"column:cust_status" json:"cust_status"`
	Cust_first_visit_date *time.Time          `gorm:"column:cust_first_visit_date" json:"cust_first_visit_date"`
	Cust_last_visit_date  *time.Time          `gorm:"column:cust_last_visit_date" json:"cust_last_visit_date"`
	Create_by             string              `gorm:"column:create_by" json:"create_by"`
	Update_by             string              `gorm:"column:update_by" json:"update_by"`
	Create_date           *time.Time          `gorm:"column:create_date" json:"create_date"`
	Update_date           *time.Time          `gorm:"column:update_date" json:"update_date"`
}

func (MerchantCustomerMapEntity) TableName() string {
	return "dp_mp_merchant_customer"
}
