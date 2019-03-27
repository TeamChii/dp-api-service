package entity

type ContainerTypeEntity struct {
	Container_type_id   int    `gorm:"column:container_type_id; primary_key" json:"container_type_id"`
	Container_type_name string `gorm:"column:container_type_name" json:"container_type_name"`
	Container_type_code string `gorm:"column:container_type_code" json:"container_type_code"`
	Ref_id              string `gorm:"column:ref_id" json:"ref_id"`
}

func (ContainerTypeEntity) TableName() string {
	return "dp_ms_container_type"
}
