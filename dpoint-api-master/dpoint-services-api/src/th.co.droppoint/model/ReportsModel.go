package model

type ReportReq struct {
	Mc_id int    `json:"mc_id"`
	Type  string `json:"type"`
}
type ReportDateReq struct {
	Mc_id           int    `json:"mc_id"`
	Start_month     string `json:"start_month"`
	End_month       string `json:"end_month"`
	Report_category string `json:"report_category"`
}
type ReportEnityModel struct {
	Mc_id           int    `json:"mc_id"`
	Report_category string `json:"report_category"`
	Count           int    `json:"count"`
	Unit            string `json:"unit"`
}
type ReportEnityModel2 struct {
	Mc_id           int    `json:"mc_id"`
	Report_category string `json:"report_category"`
	Name            string `json:"name"`
}
type ReportResp struct {
	New_customers     *ReportEnityModel `json:"new_customers"`
	Point_to_customer *ReportEnityModel `json:"point_to_customer"`
	Redeemed_cutomers *ReportEnityModel `json:"redeemed_cutomers"`
	Point_giving      *ReportEnityModel `json:"point_giving"`
	Special_request   *ReportEnityModel `json:"special_request"`
	Who_view          *ReportEnityModel `json:"who_view"`
}
type ReportMainResp struct {
	Day   *ReportResp `json:"day"`
	Week  *ReportResp `json:"week"`
	Month *ReportResp `json:"month"`
}
