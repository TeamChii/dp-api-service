package model



type CRMReportRes struct {
	//All_Customer_Visit float64    `json:"all_customer_visit"`
	//New_Customer_Visit int    `json:"new_customer_visit"`
	//Repeating_Customer_Visit float64    `json:"repeating_customer_visit"`
	//Percent_Customer_Visit string    `json:"percent_customer_visit"`

	// not clear
	//Member_Card_Reward float64    `json:"payment_received_credits"`
	//Customer_Point_Reward float64    `json:"average_sales"`
	//Point_Card_Reward int    `json:"total_bills_receipts"`

	/*
	Customer_Amt_Point_Redeem int    `json:"customer_amt_point_redeem"`
	Point_Amt_Point_Redeem int    `json:"point_amt_point_redeem"`
	Item_Amt_Point_Redeem int    `json:"item_amt_point_redeem"`

	Customer_Amt_Coupon_Redeem int    `json:"customer_amt_coupon_redeem"`
	Coupon_Amt_Coupon_Redeem int    `json:"coupon_amt_coupon_redeem"`
	*/
	Give_member_cust_count int    `json:"give_member_cust_count"`
	Give_member_card_count int    `json:"give_member_card_count"`
	Give_member_amt int    `json:"give_member_amt"`

	Give_point_cust_count int    `json:"give_point_cust_count"`
	Give_point_card_count int    `json:"give_point_card_count"`
	Give_point_amt int    `json:"give_point_amt"`

	Give_coupon_cust_count int    `json:"give_coupon_cust_count"`
	Give_coupon_card_count int    `json:"give_coupon_card_count"`
	Give_coupon_amt int    `json:"give_coupon_amt"`

	Redeem_point_cust_count int    `json:"redeem_point_cust_count"`
	Redeem_point_card_count int    `json:"redeem_point_card_count"`
	Redeem_point_amt int    `json:"redeem_point_amt"`

	Redeem_coupon_cust_count int    `json:"redeem_coupon_cust_count"`
	Redeem_coupon_card_count int    `json:"redeem_coupon_card_count"`
	Redeem_coupon_amt int    `json:"redeem_coupon_amt"`

}


