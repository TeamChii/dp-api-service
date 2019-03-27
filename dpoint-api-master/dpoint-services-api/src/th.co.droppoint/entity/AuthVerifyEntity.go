package entity

import "time"

type AuthVerifyEntity struct {
	Auth_token   string     `gorm:"column:auth_token; primary_key:true" json:"auth_token"`
	Auth_type    string     `gorm:"column:auth_type" json:"auth_type"`
	Auth_code    string     `gorm:"column:auth_code" json:"auth_code"`
	Expire_time  int        `gorm:"column:expire_time" json:"expire_time"`
	Created_time *time.Time `gorm:"column:created_time" json:"created_time"`
	//CreatedDateStr     string `gorm:"-" json:"createdDateStr"`
}

func (AuthVerifyEntity) TableName() string {
	return "dp_tb_auth_verify"
}
