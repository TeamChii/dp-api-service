package entity

import "time"

type MerchantCampaignMapEntity struct {
	Mc_id           int             `gorm:"column:mc_id" json:"mc_id"`
	Campaign_id     int             `gorm:"column:campaign_id" json:"campaign_id"`
	CampaignEntity  *CampaignEntity `gorm:"column:campaign_id;ForeignKey:campaign_id;AssociationForeignKey:campaign_id" json:"campaign"`
	Campaign_status string          `gorm:"column:campaign_status" json:"campaign_status"`
	Category_type   *string         `gorm:"column:category_type" json:"category_type"`
	Create_by       string          `gorm:"column:create_by" json:"create_by"`
	Create_date     *time.Time      `gorm:"column:create_date" json:"create_date"`
	Launch_date     *time.Time      `gorm:"column:launch_date" json:"launch_date"`
}

func (MerchantCampaignMapEntity) TableName() string {
	return "dp_tb_merchant_campaign"
}
