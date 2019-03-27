package model

import (
	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	//"th.co.droppoint/entity"
)

type CustomerModel struct {
	Data []entity.CustomerEntity `json:"data"`
}
type CustomerReq struct {
	Cust_mobile string `json:"cust_mobile"`
	Mc_id       int    `json:"mc_id"`
}

type CustomerResp struct {
	Cust_id          int        `gorm:"column:cust_id; primary_key" json:"cust_id"`
	Cust_name        string     `gorm:"column:cust_name" json:"cust_name"`
	Cust_alt_name    string     `gorm:"column:cust_alt_name" json:"cust_alt_name"`
	Cust_language    int64      `gorm:"column:cust_language" json:"cust_language"`
	Lead_by          string     `gorm:"column:lead_by" json:"lead_by"`
	Cust_status      int        `gorm:"column:cust_status" json:"cust_status"`
	Cust_dob         *time.Time `gorm:"column:cust_dob" json:"cust_dob"`
	Cust_sex         int        `gorm:"column:cust_sex" json:"cust_sex"`
	Cust_region      int        `gorm:"column:cust_region" json:"cust_region"`
	Cust_mobile      string     `gorm:"column:cust_mobile" json:"cust_mobile"`
	Cust_mobile_ext  string     `gorm:"column:cust_mobile_ext" json:"cust_mobile_ext"`
	Cust_email       string     `gorm:"column:cust_email" json:"cust_email"`
	Cust_email_ext   string     `gorm:"column:cust_email_ext" json:"cust_email_ext"`
	Cust_facebook_id string     `gorm:"column:cust_facebook_id" json:"cust_facebook_id"`
	Cust_line_id     string     `gorm:"column:cust_line_id" json:"cust_line_id"`
	Cust_twitter_id  string     `gorm:"column:cust_twitter_id" json:"cust_twitter_id"`
	Cust_google_id   string     `gorm:"column:cust_google_id" json:"cust_google_id"`
	Create_by        string     `gorm:"column:create_by" json:"create_by"`
	Update_by        string     `gorm:"column:update_by" json:"update_by"`
	Create_date      *time.Time `gorm:"column:create_date" json:"create_date"`
	Update_date      *time.Time `gorm:"column:update_date" json:"update_date"`
	Image_ref        string     `gorm:"column:image_ref" json:"image_ref"`
	Image_ref_id     int        `gorm:"column:image_ref_id" json:"image_ref_id"`
	Country_code     string     `gorm:"column:country_code" json:"country_code"`
	Last_action_date string     `gorm:"-" json:"last_purchase"`
}

func (CustomerResp) TableName() string {
	return "dp_ms_customer"
}
