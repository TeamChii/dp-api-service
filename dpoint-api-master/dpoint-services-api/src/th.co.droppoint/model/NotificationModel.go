package model

type NotiReq struct {
	Noti_id int `json:"noti_id"`
}
type NotiCategoryReq struct {
	Noti_category string `json:"noti_category"`
	Mc_id         int    `json:"mc_id"`
}
