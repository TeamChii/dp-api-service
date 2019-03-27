package model

type PinReq struct {
	Pin     string `gorm:"json:"pin"`
	User_id int    `gorm:"json:"user_id"`
}
type PinAuthenReq struct {
	Pin        string `gorm:"json:"pin"`
	User_Phone string `gorm:"json:"user_phone"`
}
type UserAddReq struct {
	Mc_id      int    `json:"mc_id"`
	User_Phone string `json:"user_phone"`
	User_name  string `json:"user_name"`
	Role_id    string `json:"role_id"`
}
type UserLoadRoleReq struct {
	Mc_id   int    `json:"mc_id"`
	Role_id string `json:"role_id"`
}
type UserAuthenResp struct {
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
	Role_id         string `gorm:"column:role_id" json:"role_id"`
}

func (UserAuthenResp) TableName() string {
	return "dp_ms_user"
}
