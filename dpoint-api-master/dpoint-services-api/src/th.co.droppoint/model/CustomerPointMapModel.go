package model

type CustomerPointMapReq struct {
	Mc_id   int         `json:"mc_id"`
	Cust_id int         `json:"cust_id"`
	Paging  PagingModel `json:"paging"`
}
