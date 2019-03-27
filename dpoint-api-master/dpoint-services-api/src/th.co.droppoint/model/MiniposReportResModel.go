package model

import "github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"

type MiniposReportRes struct {
	Net_Sales                string  `json:"net_sales"`
	Total_Net_Sales          string  `json:"total_net_sales"`
	Payment_Received_Cash    string  `json:"payment_received_cash"`
	Payment_Received_Qr      string  `json:"payment_received_qr"`
	Payment_Received_Credits string  `json:"payment_received_credits"`
	Total_Bills_Receipts     int     `json:"total_bills_receipts"`
	Average_Sales            float64 `json:"average_sales"`

	Total_Discount_Redeem float64 `json:"total_discount_redeem"`
	//Discount_Value float64    `json:"discount_value"`
	Item_Discount_Value float64 `json:"item_discount_value"`
	Bill_Discount_Value float64 `json:"bill_discount_value"`
	Redeem_Value        float64 `json:"redeem_value"`
	Redeem_Item         int     `json:"redeem_item"`

	BestSellerMenuList []BestSellerMenuRes `json:"best_seller_list"`
	// Sales Traffic
	Sales_Traffice_List []SalesTrafficeRes `json:"sales_traffice_list"`

	Sales_Traffice_Peak_Time SalesTrafficeRes `json:"sales_traffice_peak_time"`
	Sales_Traffice_Lowest    SalesTrafficeRes `json:"sales_traffice_lowest"`
}
type BestSellerMenuRes struct {
	MenuCount   int                      `json:"menu_count"`
	MiniposMenu entity.MiniposMenuEntity `json:"menu_model"`
}
type SalesTrafficeRes struct {
	Sale_Time  string `json:"sale_time"`
	Sale_Count int    `json:"sale_count"`
}
