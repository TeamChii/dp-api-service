package entity

import "time"

type MiniposTransactionRedeemDetailEntity struct {
	Mpos_Transaction_Id int        `gorm:"column:mpos_transaction_id; primary_key" json:"mpos_transaction_id"`
	Mpos_Receive_Id     int     `gorm:"column:mpos_receive_id" json:"mpos_receive_id"`
	Mc_Id               int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_Menu_Id        *int       `gorm:"column:mpos_menu_id" json:"mpos_menu_id"`
	Calculator_Keyin    string     `gorm:"column:calculator_keyin" json:"calculator_keyin"`
	Redeem_Amt        	int        `gorm:"column:redeem_amt" json:"redeem_amt"`
	Original_Price     	float64    `gorm:"column:original_price" json:"original_price"`
	Original_Item  		int        `gorm:"column:original_item" json:"original_item"`
	Sub_Total     		float64    `gorm:"column:sub_total" json:"sub_total"`
	Created_Date        *time.Time `gorm:"column:created_date" json:"created_date"`
	Created_By          string     `gorm:"column:created_by" json:"created_by"`
}






func (MiniposTransactionRedeemDetailEntity) TableName() string {
	return "dp_ts_minipos_transactoin_redeem_detail"
}
