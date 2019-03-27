package model

type MerchantImageMapResp struct {
	Content_id   int    `gorm:"column:content_id" json:"content_id"`
	Content_path string `gorm:"column:content_path" json:"content_path"`
}
