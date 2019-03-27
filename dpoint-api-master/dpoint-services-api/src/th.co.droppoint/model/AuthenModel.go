package model

type AuthenReq struct {
	Mc_phone string `json:"mc_phone"`
}
type AuthenSecReq struct {
	Pin        string `json:"pin"`
	Device_uid string `json:"device_uid"`
}
