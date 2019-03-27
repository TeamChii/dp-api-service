package entity

type MiniposItemRewardEntity struct {
	Mpos_Menu_Id         int        `gorm:"column:mpos_menu_id; primary_key" json:"mpos_menu_id"`
	Minipos_Setting_Id   int        `gorm:"column:minipos_setting_id; primary_key" json:"minipos_setting_id"`
	MiniposMenu  *MiniposMenuEntity  `gorm:"column:mpos_menu_id;ForeignKey:mpos_menu_id;AssociationForeignKey:mpos_menu_id" json:"miniposMenu"`
	Mc_Id int        `gorm:"column:mc_id; primary_key" json:"mc_id"`
	Reward_Item_Type     string     `gorm:"column:reward_item_type" json:"reward_item_type"`
	Ord     int     `gorm:"column:ord" json:"ord"`

}

func (MiniposItemRewardEntity) TableName() string {
	return "dp_ms_minipos_item_reward"
}
