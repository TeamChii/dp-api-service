package entity

import "time"

type MerchantEntity struct {
	Mc_id             int        `gorm:"column:mc_id; primary_key" json:"mc_id"`
	Mc_name           string     `gorm:"column:mc_name" json:"mc_name"`
	Mc_email          string     `gorm:"column:mc_email" json:"mc_email"`
	Mc_email_ext      string     `gorm:"column:mc_email_ext" json:"mc_email_ext"`
	Mc_facebook_id    string     `gorm:"column:mc_facebook_id" json:"mc_facebook_id"`
	Mc_line_id        string     `gorm:"column:mc_line_id" json:"mc_line_id"`
	Mc_twitter_id     string     `gorm:"column:mc_twitter_id" json:"mc_twitter_id"`
	Mc_ig_id          string     `gorm:"column:mc_ig_id" json:"mc_ig_id"`
	Mc_google_id      string     `gorm:"column:mc_google_id" json:"mc_google_id"`
	Mc_phone          string     `gorm:"column:mc_phone" json:"mc_phone"`
	Mc_phone_ext      string     `gorm:"column:mc_phone_ext" json:"mc_phone_ext"`
	Mc_mob            string     `gorm:"column:mc_mob" json:"mc_mob"`
	Mc_mob_ext        string     `gorm:"column:mc_mob_ext" json:"mc_mob_ext"`
	Mc_region         string     `gorm:"column:mc_region" json:"mc_region"`
	Mc_language       string     `gorm:"column:mc_language" json:"mc_language"`
	Mc_country        string     `gorm:"column:mc_country" json:"mc_country"`
	Mc_address_1      string     `gorm:"column:mc_address_1" json:"mc_address_1"`
	Mc_address_2      string     `gorm:"column:mc_address_2" json:"mc_address_2"`
	Mc_key_contact    string     `gorm:"column:mc_key_contact" json:"mc_key_contact"`
	Mc_lat            float64    `gorm:"column:mc_lat" json:"mc_lat"`
	Mc_long           float64    `gorm:"column:mc_long" json:"mc_long"`
	Mc_status         string     `gorm:"column:mc_status" json:"mc_status"`
	Mc_business_cat   string     `gorm:"column:mc_business_cat" json:"mc_business_cat"`
	Mc_business_size  string     `gorm:"column:mc_business_size" json:"mc_business_size"`
	Mc_sale_volume    int        `gorm:"column:mc_sale_volume" json:"mc_sale_volume"`
	Mc_currency       string     `gorm:"column:mc_currency" json:"mc_currency"`
	Mc_is_head_office string     `gorm:"column:mc_is_head_office" json:"mc_is_head_office"`
	Mc_ref            int        `gorm:"column:mc_ref" json:"mc_ref"`
	Create_by         string     `gorm:"column:create_by" json:"create_by"`
	Update_by         string     `gorm:"column:update_by" json:"update_by"`
	Create_date       *time.Time `gorm:"column:create_date" json:"create_date"`
	Update_date       *time.Time `gorm:"column:update_date" json:"update_date"`
	Mc_tax            string     `gorm:"column:mc_tax" json:"mc_tax"`
	Mc_detail         string     `gorm:"column:mc_detail" json:"mc_detail"`
	Package_id        int        `gorm:"column:package_id" json:"package_id"`
	//Image_ref         string     `gorm:"column:image_ref" json:"image_ref"`
	//Image_ref_id      int        `gorm:"column:image_ref_id" json:"image_ref_id"`
	//MerchantImageMapEntity MerchantImageMapEntity `gorm:"-" json:"image"`
	Logo_ref    string `gorm:"column:logo_ref" json:"logo_ref"`
	Logo_ref_id int    `gorm:"column:logo_ref_id" json:"logo_ref_id"`

	Mc_province string `gorm:"column:mc_province" json:"mc_province"`
	Mc_district string `gorm:"column:mc_district" json:"mc_district"`
	Mc_postcode string `gorm:"column:mc_postcode" json:"mc_postcode"`
	Mc_credit        int        `gorm:"column:mc_credit" json:"mc_credit"`
	Country_code        string        `gorm:"column:country_code" json:"country_code"`
	Mc_group_id        *int        `gorm:"column:mc_group_id" json:"mc_group_id"`


}

func (MerchantEntity) TableName() string {
	return "dp_ms_merchant"
}
