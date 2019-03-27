package entity

type MiniposSettingEntity struct {
	Minipos_Setting_Id int        `gorm:"column:minipos_setting_id; primary_key" json:"minipos_setting_id"`
	Mc_id               int        `gorm:"column:mc_id" json:"mc_id"`
	Currency_Code       string      `gorm:"column:currency_code" json:"currency_code"`
	Vat_Amt    			*float64     `gorm:"column:vat_amt" json:"vat_amt"`
	Service_Charge_Amt  *float64    `gorm:"column:service_charge_amt" json:"service_charge_amt"`
	Thai_Prompt_Pay_Flag string        `gorm:"column:thai_prompt_pay_flag" json:"thai_prompt_pay_flag"`
	Line_Pay_Flag        string `gorm:"column:line_pay_flag" json:"line_pay_flag"`
	Ali_Pay_Flag         string     `gorm:"column:ali_pay_flag" json:"ali_pay_flag"`
	Loyalty_Point_Link_Flag string     `gorm:"column:loyalty_point_link_flag" json:"loyalty_point_link_flag"`
	Point_Reward_Flag       string     `gorm:"column:point_reward_flag" json:"point_reward_flag"`
}

func (MiniposSettingEntity) TableName() string {
	return "dp_ms_minipos_setting"
}

