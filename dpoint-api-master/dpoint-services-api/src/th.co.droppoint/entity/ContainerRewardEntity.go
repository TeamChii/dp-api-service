package entity

type ContainerRewardEntity struct {
	Container_id    int              `gorm:"column:container_id; primary_key" json:"container_id"`
	ContainerEntity *ContainerEntity `gorm:"column:container_id;ForeignKey:container_id;AssociationForeignKey:container_id" json:"container"`
	Point_amt       int              `gorm:"column:point_amt" json:"point_amt"`
	Reward_detail   string           `gorm:"column:reward_detail" json:"reward_detail"`
}

func (ContainerRewardEntity) TableName() string {
	return "dp_ms_container_reward"
}
