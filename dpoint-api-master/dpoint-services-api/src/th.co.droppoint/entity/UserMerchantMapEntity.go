package entity

type UserMerchantMapEntity struct {
	User_id int    `gorm:"column:user_id" json:"user_id"`
	Mc_id   int    `gorm:"column:mc_id" json:"mc_id"`
	Role_id string `gorm:"column:role_id" json:"role_id"`
}

func (UserMerchantMapEntity) TableName() string {
	return "dp_mp_user_merchant"
}
