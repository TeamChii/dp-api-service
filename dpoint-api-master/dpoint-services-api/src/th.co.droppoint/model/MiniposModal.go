package model

import "github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"

type MenuReq struct {
	Mc_id                 int         `json:"mc_id"`
	Mpos_menu_category_id *int        `json:"mpos_menu_category_id"`
	Paging                PagingModel `json:"paging"`
}
type TransactionReceiveReq struct {
	Transaction_Purchase []Transaction_Purchase `json:"transaction_purchase"`
	Transaction_Discount []Transaction_Discount `json:"transaction_discount"`
	Transaction_Redeem   []Transaction_Redeem   `json:"transaction_redeem"`
	Receive              Receive                `json:"receive"`
}
type Transaction_Purchase struct {
	Mc_id            int     `json:"mc_id"`
	Mpos_Menu_Id     *int    `json:"mpos_menu_id"`
	Calculator_Keyin string  `json:"calculator_keyin"`
	Price            float64 `json:"price"`
	Amt              int     `json:"amt"`
}
type Transaction_Discount struct {
	Mc_id            int     `json:"mc_id"`
	Mpos_Menu_Id     *int    `json:"mpos_menu_id"`
	Calculator_Keyin string  `json:"calculator_keyin"`
	Discount_Type    string  `json:"discount_type"`
	Discount_Amt     float64 `json:"discount_amt"`
	Remaining_Price  float64 `json:"remaining_price"`
	Remaining_Item   int     `json:"remaining_item"`
	Sub_Total        float64 `json:"sub_total"`
}
type Transaction_Redeem struct {
	Mc_id            int     `json:"mc_id"`
	Mpos_Menu_Id     *int    `json:"mpos_menu_id"`
	Calculator_Keyin string  `json:"calculator_keyin"`
	Redeem_Amt       int     `json:"redeem_amt"`
	Original_Price   float64 `json:"original_price"`
	Original_Item    int     `json:"original_item"`
	Sub_Total        float64 `json:"sub_total"`
}
type Receive struct {
	Mc_id                  int      `json:"mc_id"`
	Mpos_payment_method_id int      `json:"mpos_payment_method_id"`
	Total_charge_amt       float64  `json:"total_charge_amt"`
	Receive_amt            float64  `json:"receive_amt"`
	Change_amt             float64  `json:"change_amt"`
	Cust_Mobile_No         string   `json:"cust_mobile_no"`
	Discount_Amt           *float64 `json:"discount_amt"`
	Discount_Type          string   `json:"discount_type"`
}
type MiniPOSMenuCategoryReq struct {
	Transaction_Purchase []Transaction_Purchase `json:"transaction_purchase"`
	Receive              Receive                `json:"receive"`
}

type GivePointRewardModel struct {
	MiniposPointReward entity.MiniposPointRewardEntity  `json:"minipos_point_reward"`
	MenuItems          []entity.MiniposItemRewardEntity `json:"menu_items"`
}

type MiniposReceiptHistoryReq struct {
	Customer_Mobile_No string `json:"customer_mobile_no"`
	MC_Id              int    `json:"mc_id"`
	Mpos_Receive_Id    int    `json:"mpos_receive_id"`
	Status_Flag        string `json:"status_flag"`
}
