package model

type MerchantCustomerCampaignMapAddReq struct {
	Mc_id           int               `json:"mc_id"`
	Campaign_id     int               `json:"campaign_id"`
	Category_type   string            `json:"category_type"`
	Message_detail  string            `json:"message_detail"`
	ContainerDetail []ContainerDetail `gorm:"-" json:"container"`
}

type MerchantCustomerCampaignMapLoadCustReq struct {
	Mc_id       int         `json:"mc_id"`
	Campaign_id int         `json:"campaign_id"`
	Paging      PagingModel `json:"paging"`
}
