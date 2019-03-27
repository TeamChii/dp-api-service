package entity

type MiniposThaiPromptPayEntity struct {
	Thai_Prompt_Pay_Id 		int        `gorm:"column:thai_prompt_pay_id; primary_key" json:"thai_prompt_pay_id"`
	MC_Id      int        `gorm:"column:mc_id" json:"mc_id"`
	Minipos_Setting_Id      int        `gorm:"column:minipos_setting_id" json:"minipos_setting_id"`
	Prompt_Pay_Key        	string       `gorm:"column:prompt_pay_key" json:"prompt_pay_key"`
	Merchant_Name    		string     `gorm:"column:merchant_name" json:"merchant_name"`
	Payee_Name              string    `gorm:"column:payee_name" json:"payee_name"`
	Payee_Surename          string        `gorm:"column:payee_surename" json:"payee_surename"`
	Email         			string `gorm:"column:email" json:"email"`
	Merchant_Business_Type  string     `gorm:"column:merchant_business_type" json:"merchant_business_type"`
	Bank_Name     			string     `gorm:"column:bank_name" json:"bank_name"`
}

func (MiniposThaiPromptPayEntity) TableName() string {
	return "dp_ms_minipos_thai_prompt_pay"
}

