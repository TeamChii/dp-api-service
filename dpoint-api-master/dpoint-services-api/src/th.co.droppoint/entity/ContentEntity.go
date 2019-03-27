package entity

import "time"

type ContentEntity struct {
	Content_id     int        `gorm:"column:content_id; primary_key" json:"content_id"`
	Content_type   string     `gorm:"column:content_type" json:"content_type"`
	File_name      string     `gorm:"column:file_name" json:"file_name"`
	Content_root   string     `gorm:"column:content_root" json:"content_root"`
	Content_path   string     `gorm:"column:content_path" json:"content_path"`
	Create_by      string     `gorm:"column:create_by" json:"create_by"`
	Update_by      string     `gorm:"column:update_by" json:"update_by"`
	Create_date    *time.Time `gorm:"column:create_date" json:"create_date"`
	Update_date    *time.Time `gorm:"column:update_date" json:"update_date"`
	CreatedDateStr string     `gorm:"-" json:"createdDateStr"`
	UpdatedDateStr string     `gorm:"-" json:"updatedDateStr"`
}

func (ContentEntity) TableName() string {
	return "dp_tb_content"
}
