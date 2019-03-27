package entity

type UserEntity struct {
	User_id         int    `gorm:"column:user_id; primary_key" json:"user_id"`
	User_name       string `gorm:"column:user_name" json:"user_name"`
	User_first_name string `gorm:"column:user_first_name" json:"user_first_name"`
	User_last_name  string `gorm:"column:user_last_name" json:"user_last_name"`
	User_phone      string `gorm:"column:user_phone" json:"user_phone"`
	Pin             string `gorm:"column:pin" json:"pin"`
	Image_ref       string `gorm:"column:image_ref" json:"image_ref"`
	Image_ref_id    int    `gorm:"column:image_ref_id" json:"image_ref_id"`
	User_email      string `gorm:"column:user_email" json:"user_email"`
	Staff_pin       string `gorm:"column:staff_pin" json:"staff_pin"`
}

func (UserEntity) TableName() string {
	return "dp_ms_user"
}
