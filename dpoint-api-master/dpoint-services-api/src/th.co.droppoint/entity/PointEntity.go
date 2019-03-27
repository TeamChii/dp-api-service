package entity

import "time"

type PointEntity struct {
	Point_id            int              `gorm:"column:point_id; primary_key" json:"point_id"`
	Mc_id               int              `gorm:"column:mc_id" json:"mc_id"`
	MerchantEntity      *MerchantEntity  `gorm:"column:mc_id;ForeignKey:mc_id;AssociationForeignKey:mc_id" json:"merchant"`
	Cust_id             int              `gorm:"column:cust_id" json:"cust_id"`
	CustomerEntity      *CustomerEntity  `gorm:"column:cust_id;ForeignKey:cust_id;AssociationForeignKey:cust_id" json:"customer"`
	Container_id        int              `gorm:"column:container_id" json:"container_id"`
	ContainerEntity     *ContainerEntity `gorm:"column:container_id;ForeignKey:container_id;AssociationForeignKey:container_id" json:"container"`
	Menu_id             int              `gorm:"column:menu_id" json:"menu_id"`
	Transfer_to_cust_id int              `gorm:"column:transfer_to_cust_id" json:"transfer_to_cust_id"`
	Point_amt           int              `gorm:"column:point_amt" json:"point_amt"`
	Create_by           string           `gorm:"column:create_by" json:"create_by"`
	Create_date         *time.Time       `gorm:"column:create_date" json:"create_date"`
	Expire_date         *time.Time       `gorm:"column:expire_date" json:"expire_date"`
	Expire_flag         string           `gorm:"column:expire_flag" json:"expire_flag"`

	Redeem_flag string     `gorm:"column:redeem_flag" json:"redeem_flag"`
	Redeem_date *time.Time `gorm:"column:redeem_date" json:"redeem_date"`
	Redeem_amt  int        `gorm:"column:redeem_amt" json:"redeem_amt"`
	Expire_amt  int        `gorm:"column:expire_amt" json:"expire_amt"`
}

func (PointEntity) TableName() string {
	return "dp_tb_point"
}
