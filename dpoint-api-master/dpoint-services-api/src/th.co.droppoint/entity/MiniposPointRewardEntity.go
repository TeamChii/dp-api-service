package entity

type MiniposPointRewardEntity struct {
	Minipos_Point_reward_Id int        `gorm:"column:minipos_point_reward_id; primary_key" json:"minipos_point_reward_id"`
	MC_Id   int `gorm:"column:mc_id" json:"mc_id"`
	Minipos_Setting_Id      int        `gorm:"column:minipos_setting_id" json:"minipos_setting_id"`
	Container_Id       		int     `gorm:"column:container_id" json:"container_id"`
	Reward_Type     		string     `gorm:"column:reward_type" json:"reward_type"`
	Reward_Amount_Amt 		*float64     `gorm:"column:reward_amount_amt" json:"reward_amount_amt"`
	Reward_Amount_Point     *int `gorm:"column:reward_amount_point" json:"reward_amount_point"`
	Reward_Item_Type        string     `gorm:"column:reward_item_type" json:"reward_item_type"`

}

func (MiniposPointRewardEntity) TableName() string {
	return "dp_ms_minipos_point_reward"
}

