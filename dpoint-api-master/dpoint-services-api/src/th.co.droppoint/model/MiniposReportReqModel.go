package model


type MiniposReportReq struct {
	From_date                 string         `json:"from_date"`
	To_date string        `json:"to_date"`
	MC_Id int         `json:"mc_id"`
	Filter_Type string        `json:"filter_type"`
}
