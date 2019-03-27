package entity

type MerchantImageMapEntity struct {
	Mc_id      int `gorm:"column:mc_id" json:"mc_id"`
	Content_id int `gorm:"column:content_id" json:"content_id"`
}

func (MerchantImageMapEntity) TableName() string {
	return "dp_mp_merchant_image"
}
