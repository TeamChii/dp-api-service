package entity

import "time"

type RedeemEntity struct {
	Redeem_id       int              `gorm:"column:redeem_id; primary_key" json:"redeem_id"`
	Mc_id           int              `gorm:"column:mc_id" json:"mc_id"`
	MerchantEntity  *MerchantEntity  `gorm:"column:mc_id;ForeignKey:mc_id;AssociationForeignKey:mc_id" json:"merchant"`
	Cust_id         int              `gorm:"column:cust_id" json:"cust_id"`
	CustomerEntity  *CustomerEntity  `gorm:"column:cust_id;ForeignKey:cust_id;AssociationForeignKey:cust_id" json:"customer"`
	Container_id    int              `gorm:"column:container_id" json:"container_id"`
	ContainerEntity *ContainerEntity `gorm:"column:container_id;ForeignKey:container_id;AssociationForeignKey:container_id" json:"container"`
	Menu_id         *int             `gorm:"column:menu_id" json:"menu_id"`
	Point_amt       int              `gorm:"column:point_amt" json:"point_amt"`
	Create_date     *time.Time       `gorm:"column:create_date" json:"create_date"`
	Create_by       string           `gorm:"column:create_by" json:"create_by"`
}

func (RedeemEntity) TableName() string {
	return "dp_tb_redeem"
}
