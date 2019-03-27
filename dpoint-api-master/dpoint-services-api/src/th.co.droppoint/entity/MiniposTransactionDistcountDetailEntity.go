package entity

import "time"

type MiniposTransactionDistcountDetailEntity struct {
	Mpos_Transaction_Id int        `gorm:"column:mpos_transaction_id; primary_key" json:"mpos_transaction_id"`
	Mpos_Receive_Id     int     `gorm:"column:mpos_receive_id" json:"mpos_receive_id"`
	Mc_Id               int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_Menu_Id        *int       `gorm:"column:mpos_menu_id" json:"mpos_menu_id"`
	Calculator_Keyin    string     `gorm:"column:calculator_keyin" json:"calculator_keyin"`
	Discount_Type    	string     `gorm:"column:discount_type" json:"discount_type"`
	Discount_Amt        float64     `gorm:"column:discount_amt" json:"discount_amt"`
	Remaining_Price     float64    `gorm:"column:remaining_price" json:"remaining_price"`
	Remaining_Item  	int        `gorm:"column:remaining_item" json:"remaining_item"`
	Sub_Total     		float64    `gorm:"column:sub_total" json:"sub_total"`
	Created_Date        *time.Time `gorm:"column:created_date" json:"created_date"`
	Created_By          string     `gorm:"column:created_by" json:"created_by"`

}

func (MiniposTransactionDistcountDetailEntity) TableName() string {
	return "dp_ts_minipos_transactoin_discount_detail"
}
