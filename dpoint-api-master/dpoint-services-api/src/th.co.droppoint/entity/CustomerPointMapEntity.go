package entity

type CustomerPointMapEntity struct {
	Mc_id     int    `gorm:"column:mc_id" json:"mc_id"`
	Cust_id   int    `gorm:"column:cust_id" json:"cust_id"`
	Card_type string `gorm:"column:card_type" json:"card_type"`
	Point_amt int    `gorm:"column:point_amt" json:"point_amt"`
}

func (CustomerPointMapEntity) TableName() string {
	return "dp_mp_customer_point"
}
