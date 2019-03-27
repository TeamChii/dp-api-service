package entity

import "time"

type MiniposTransactionDetailEntity struct {
	Mpos_transaction_id int        `gorm:"column:mpos_transaction_id; primary_key" json:"mpos_transaction_id"`
	Mc_id               int        `gorm:"column:mc_id" json:"mc_id"`
	Mpos_menu_id        *int       `gorm:"column:mpos_menu_id" json:"mpos_menu_id"`
	Calculator_keyin    string     `gorm:"column:calculator_keyin" json:"calculator_keyin"`
	Price               float64    `gorm:"column:price" json:"price"`
	Amt                 int        `gorm:"column:amt" json:"amt"`
	Create_date         *time.Time `gorm:"column:create_date" json:"create_date"`
	Create_by           string     `gorm:"column:create_by" json:"create_by"`

	Mpos_receive_id     int     `gorm:"column:mpos_receive_id" json:"mpos_receive_id"`
}

func (MiniposTransactionDetailEntity) TableName() string {
	return "dp_ts_minipos_transaction_detail"
}
