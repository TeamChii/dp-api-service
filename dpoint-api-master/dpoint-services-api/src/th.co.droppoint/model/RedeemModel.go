package model

type RedeemReq struct {
	Mc_id  int         `json:"mc_id"`
	Paging PagingModel `json:"paging"`
}
type RedeemReqList struct {
	Mc_id     int             `json:"mc_id"`
	Cust_id   int             `json:"cust_id"`
	Container []ContainerList `json:"container"`
}
type ContainerList struct {
	Container_id int `json:"container_id"`
	Point_amt    int `json:"point_amt"`
	Container_Type_id int `json:"container_type_id"`
}
