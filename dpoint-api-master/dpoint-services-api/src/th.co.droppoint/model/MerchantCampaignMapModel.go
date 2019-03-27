package model

import (
	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	//"th.co.droppoint/entity"
)

type MerchantCampaignMapReq struct {
	Mc_id         int         `json:"mc_id"`
	Category_type string      `json:"category_type"`
	Paging        PagingModel `json:"paging"`
}

type MerchantCampaignMapReq2 struct {
	Mc_id         int    `json:"mc_id"`
	Campaign_id   int    `json:"campaign_id"`
	Category_type string `json:"category_type"`
}
type MerchantCampaignMapUpadteReq struct {
	Mc_id       int `json:"mc_id"`
	Campaign_id int `json:"campaign_id"`
}
type MerchantCampaignMapEntityResp struct {
	Mc_id           int                    `gorm:"column:mc_id" json:"mc_id"`
	Campaign_id     int                    `gorm:"column:campaign_id" json:"campaign_id"`
	CampaignEntity  *entity.CampaignEntity `gorm:"column:campaign_id;ForeignKey:campaign_id;AssociationForeignKey:campaign_id" json:"campaign"`
	Campaign_status string                 `gorm:"column:campaign_status" json:"campaign_status"`
	Category_type   string                 `gorm:"column:category_type" json:"category_type"`
	Create_by       string                 `gorm:"column:create_by" json:"create_by"`
	Create_date     *time.Time             `gorm:"column:create_date" json:"create_date"`
	Launch_date     *time.Time             `gorm:"column:launch_date" json:"launch_date"`
	Customer_count  int                    `gorm:"-" json:"customer_count"`
}

func (MerchantCampaignMapEntityResp) TableName() string {
	return "dp_tb_merchant_campaign"
}
