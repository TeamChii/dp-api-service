package entity

type SystemParamEntity struct {
	Key_code      string `gorm:"column:key_code" json:"key_code"`
	Key_value     string `gorm:"column:key_value" json:"key_value"`
	Key_ref       string `gorm:"column:key_ref" json:"key_ref"`
	Category_name string `gorm:"column:category_name" json:"category_name"`
	Key_extra     string `gorm:"column:key_extra" json:"key_extra"`
	Ord           int    `gorm:"column:ord" json:"ord"`
	Active        string `gorm:"column:active" json:"active"`
	Image_ref     string `gorm:"column:image_ref" json:"image_ref"`
}

func (SystemParamEntity) TableName() string {
	return "dp_ms_system_param"
}
