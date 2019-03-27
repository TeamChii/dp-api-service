package entity

import "time"

type ContainerEntity struct {
	Container_id              int                  `gorm:"column:container_id; primary_key" json:"container_id"`
	Container_type_id         int                  `gorm:"column:container_type_id" json:"container_type_id"`
	ContainerTypeEntity       *ContainerTypeEntity `gorm:"column:container_type_id;ForeignKey:container_type_id;AssociationForeignKey:container_type_id" json:"containerType"`
	Mc_id                     int                  `gorm:"column:mc_id" json:"mc_id"`
	MerchantEntity            *MerchantEntity      `gorm:"column:mc_id;ForeignKey:mc_id;AssociationForeignKey:mc_id" json:"merchant"`
	Layout_template           string               `gorm:"column:layout_template" json:"layout_template"`
	Layout_color              string               `gorm:"column:layout_color" json:"layout_color"`
	Container_subject         string               `gorm:"column:container_subject" json:"container_subject"`
	Container_detail          string               `gorm:"column:container_detail" json:"container_detail"`
	Image_ref                 string               `gorm:"column:image_ref" json:"image_ref"`
	Image_ref_id              int                  `gorm:"column:image_ref_id" json:"image_ref_id"`
	Container_term_conditions string               `gorm:"column:container_term_conditions" json:"container_term_conditions"`
	Expire_mode               string               `gorm:"column:expire_mode" json:"expire_mode"`
	Expire_value              string               `gorm:"column:expire_value" json:"expire_value"`
	Allow_borrow              string               `gorm:"column:allow_borrow" json:"allow_borrow"`
	Transferable              string               `gorm:"column:transferable" json:"transferable"`
	Create_by                 string               `gorm:"column:create_by" json:"create_by"`
	Create_date               *time.Time           `gorm:"column:create_date" json:"create_date"`
	Update_by                 string               `gorm:"column:update_by" json:"update_by"`
	Update_date               *time.Time           `gorm:"column:update_date" json:"update_date"`
	Customer_Flag             string           		`gorm:"column:customer_flag" json:"customer_flag"`
	Active_Status             string           		`gorm:"column:active_status" json:"active_status"`

	Card_Usage_Every          string                         `gorm:"column:card_usage_every" json:"card_usage_every"`
	Card_Usage_Once           string                         `gorm:"column:card_usage_once" json:"card_usage_once"`
	ItemTime                  int                         	 `gorm:"column:item_time" json:"item_time"`

	ContainerRewardEntity     []ContainerRewardEntity `gorm:"-" json:"reward"`

}

func (ContainerEntity) TableName() string {
	return "dp_ms_container"
}
