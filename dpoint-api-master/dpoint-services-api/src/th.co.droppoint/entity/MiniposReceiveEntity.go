package entity

import "time"

type MiniposReceiveEntity struct {
	Mpos_receive_id        int        `gorm:"column:mpos_receive_id; primary_key" json:"mpos_receive_id"`
	Mc_id                  int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_payment_method_id int        `gorm:"column:mpos_payment_method_id" json:"mpos_payment_method_id"`
	Total_charge_amt       float64    `gorm:"column:total_charge_amt" json:"total_charge_amt"`
	Receive_amt            float64    `gorm:"column:receive_amt" json:"receive_amt"`
	Change_amt             float64    `gorm:"column:change_amt" json:"change_amt"`
	Create_date            *time.Time `gorm:"column:create_date" json:"create_date"`
	Create_by              string     `gorm:"column:create_by" json:"create_by"`

	Cust_Mobile_No  	   string     `gorm:"column:cust_mobile_no" json:"cust_mobile_no"`
	Cust_Country_Code  	   string     `gorm:"column:cust_country_code" json:"cust_country_code"`
	Trx_No  			   string     `gorm:"column:trx_no" json:"trx_no"`
	Discount_Amt 		   *float64    `gorm:"column:discount_amt" json:"discount_amt"`
	Discount_Type  			string    `gorm:"column:discount_type" json:"discount_type"`
	Cust_Id  				*int      `gorm:"column:cust_id" json:"cust_id"`
	Status_Flag  			string      `gorm:"column:status_flag" json:"status_flag"`
}

func (MiniposReceiveEntity) TableName() string {
	return "dp_ts_minipos_receive"
}
