package model

type UserMerchantMapEntityResp struct {
	User_id            int                 `gorm:"column:user_id" json:"user_id"`
	Mc_id              int                 `gorm:"column:mc_id" json:"mc_id"`
	MerchantEntityResp *MerchantEntityResp `gorm:"column:mc_id;ForeignKey:mc_id;AssociationForeignKey:mc_id" json:"merchant"`
	Role_id            string              `gorm:"column:role_id" json:"role_id"`
}

func (UserMerchantMapEntityResp) TableName() string {
	return "dp_mp_user_merchant"
}

type UserMerchantMapEntityResp2 struct {
	User_id         int    `json:"user_id"`
	Mc_id           int    `json:"mc_id"`
	User_name       string `json:"user_name"`
	User_first_name string `json:"user_first_name"`
	User_last_name  string `json:"user_last_name"`
	User_phone      string `json:"user_phone"`
	Image_ref       string `json:"image_ref"`
	Image_ref_id    int    `json:"image_ref_id"`
	Role_id         string `json:"role_id"`
	Staff_pin       string `json:"staff_pin"`
}
