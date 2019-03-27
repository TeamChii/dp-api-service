package entity

import "time"

type SumCustomerPointEntity struct {
	Cust_id                   int        `gorm:"column:cust_id" json:"cust_id"`
	Mc_id                     int        `gorm:"column:mc_id" json:"mc_id"`
	Container_id              int        `gorm:"column:container_id" json:"container_id"`
	Current_point_amt         *int       `gorm:"column:current_point_amt" json:"current_point_amt"`
	Archived_point_amt        int        `gorm:"column:archived_point_amt" json:"archived_point_amt"`
	Archived_transfer_amt     int        `gorm:"column:archived_transfer_amt" json:"archived_transfer_amt"`
	Archived_redeem_amt       int        `gorm:"column:archived_redeem_amt" json:"archived_redeem_amt"`
	Last_archived_point_date  *time.Time `gorm:"column:last_archived_point_date" json:"last_archived_point_date"`
	Last_archived_redeem_date *time.Time `gorm:"column:last_archived_redeem_date" json:"last_archived_redeem_date"`
	Created_by                string     `gorm:"column:created_by" json:"created_by"`
	Created_date              *time.Time `gorm:"column:created_date" json:"created_date"`
	Last_action_date          *time.Time `gorm:"column:last_action_date" json:"last_action_date"`

	Expired_point_amt          int `gorm:"column:expired_point_amt" json:"expired_point_amt"`
	Archived_expired_point_amt int `gorm:"column:archived_expired_point_amt" json:"archived_expired_point_amt"`
}

func (SumCustomerPointEntity) TableName() string {
	return "dp_sum_customer_point"
}
