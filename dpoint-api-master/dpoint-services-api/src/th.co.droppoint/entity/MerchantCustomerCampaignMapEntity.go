package entity

import "time"

type MerchantCustomerCampaignMapEntity struct {
	Mc_id           int             `gorm:"column:mc_id" json:"mc_id"`
	MerchantEntity  *MerchantEntity `gorm:"column:mc_id;ForeignKey:mc_id;AssociationForeignKey:mc_id" json:"merchant"`
	Campaign_id     int             `gorm:"column:campaign_id" json:"campaign_id"`
	CampaignEntity  *CampaignEntity `gorm:"column:campaign_id;ForeignKey:campaign_id;AssociationForeignKey:campaign_id" json:"campaign"`
	Cust_id         int             `gorm:"column:cust_id" json:"cust_id"`
	CustomerEntity  *CustomerEntity `gorm:"column:cust_id;ForeignKey:cust_id;AssociationForeignKey:cust_id" json:"customer"`
	Message_detail  string          `gorm:"column:message_detail" json:"message_detail"`
	Category_type   *string         `gorm:"column:category_type" json:"category_type"`
	Send_status     string          `gorm:"column:send_status" json:"send_status"`
	Response_status string          `gorm:"column:response_status" json:"response_status"`
	Send_date       *time.Time      `gorm:"column:send_date" json:"send_date"`
	Response_date   *time.Time      `gorm:"column:response_date" json:"response_date"`
}

func (MerchantCustomerCampaignMapEntity) TableName() string {
	return "dp_mp_merchant_customer_campaign"
}
