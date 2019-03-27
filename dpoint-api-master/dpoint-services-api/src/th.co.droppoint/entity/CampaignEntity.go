package entity

type CampaignEntity struct {
	Campaign_id     int    `gorm:"column:campaign_id; primary_key" json:"campaign_id"`
	Campaign_name   string `gorm:"column:campaign_name" json:"campaign_name"`
	Campaign_code   string `gorm:"column:campaign_code" json:"campaign_code"`
	Campagin_detail string `gorm:"column:campagin_detail" json:"campagin_detail"`
	Campagin_color  string `gorm:"column:campagin_color" json:"campagin_color"`
}

func (CampaignEntity) TableName() string {
	return "dp_ms_campaign"
}
