package entity

import "time"

type MiniposMenuEntity struct {
	Mpos_menu_id          int        `gorm:"column:mpos_menu_id; primary_key" json:"mpos_menu_id"`
	Mc_id                 int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_menu_set_name    string     `gorm:"column:mpos_menu_set_name" json:"mpos_menu_set_name"`
	Mpos_menu_name        string     `gorm:"column:mpos_menu_name" json:"mpos_menu_name"`
	Mpos_menu_detail      string     `gorm:"column:mpos_menu_detail" json:"mpos_menu_detail"`
	Mpos_menu_category_id *int        `gorm:"column:mpos_menu_category_id" json:"mpos_menu_category_id"`
	Mpos_menu_price       *float64    `gorm:"column:mpos_menu_price" json:"mpos_menu_price"`
	Mpos_menu_currency    string     `gorm:"column:mpos_menu_currency" json:"mpos_menu_currency"`
	Create_date           *time.Time `gorm:"column:create_date" json:"create_date"`
	Create_by             string     `gorm:"column:create_by" json:"create_by"`
	Update_date           *time.Time `gorm:"column:update_date" json:"update_date"`
	Update_by             string     `gorm:"column:update_by" json:"update_by"`
	Active                string     `gorm:"column:active" json:"active"`
	Mpos_menu_image_ref   string     `gorm:"column:mpos_menu_image_ref" json:"mpos_menu_image_ref"`

	Ord  				  int        `gorm:"column:ord" json:"ord"`
	Allowed_Change		   string     `gorm:"column:allowed_change" json:"allowed_change"`
	Allowed_Quantity_Change 	 string     `gorm:"column:allowed_quantity_change" json:"allowed_quantity_change"`
	Allowed_Price_Chage  string     `gorm:"column:allowed_price_chage" json:"allowed_price_chage"`
	Multiple_Price_Flag  string     `gorm:"column:multiple_price_flag" json:"multiple_price_flag"`
	Price1_Amt			*float64    `gorm:"column:price1_amt" json:"price1_amt"`
	Price1_Detail		 string     `gorm:"column:price1_detail" json:"price1_detail"`
	Price2_Amt			*float64    `gorm:"column:price2_amt" json:"price2_amt"`
	Price2_Detail		 string     `gorm:"column:price2_detail" json:"price2_detail"`
	Price3_Amt			*float64    `gorm:"column:price3_amt" json:"price3_amt"`
	Price3_Detail		 string     `gorm:"column:price3_detail" json:"price3_detail"`
	Price4_Amt			*float64    `gorm:"column:price4_amt" json:"price4_amt"`
	Price4_Detail		 string     `gorm:"column:price4_detail" json:"price4_detail"`
	Price5_Amt			*float64    `gorm:"column:price5_amt" json:"price5_amt"`
	Price5_Detail		 string     `gorm:"column:price5_detail" json:"price5_detail"`
	Price1_Flag		 	string     `gorm:"column:price1_flag" json:"price1_flag"`
	Price2_Flag			 string     `gorm:"column:price2_flag" json:"price2_flag"`
	Price3_Flag		 	string     `gorm:"column:price3_flag" json:"price3_flag"`
	Price4_Flag			 string     `gorm:"column:price4_flag" json:"price4_flag"`
	Price5_Flag			 string     `gorm:"column:price5_flag" json:"price5_flag"`
}

func (MiniposMenuEntity) TableName() string {
	return "dp_ms_minipos_menu"
}
