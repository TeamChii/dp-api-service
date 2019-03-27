package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
		"th.co.droppoint/service/authentication"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func LoadCategoryMenu(MenuReq model.MenuReq) map[string]interface{} {
	var data []entity.MiniposMenuCategoryEntity
	err := config.DBsql().
		Where("mc_id = ?", MenuReq.Mc_id).
		Find(&data).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}

func LoadMenuByCategory(MenuReq model.MenuReq) map[string]interface{} {
	var data []entity.MiniposMenuEntity
	var count int
	var order = MenuReq.Paging.OrderBy + " " + MenuReq.Paging.SortBy
	var offset = MenuReq.Paging.PageNo*MenuReq.Paging.PageSize - MenuReq.Paging.PageSize
	var err error
	if MenuReq.Mpos_menu_category_id == nil {
		config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id is null ", MenuReq.Mc_id).
			Find(&data).Count(&count)

		err = config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id is null ", MenuReq.Mc_id).
			Offset(offset).
			Limit(MenuReq.Paging.PageSize).
			Order(order).Find(&data).Error
	} else {
		config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id = ?", MenuReq.Mc_id, MenuReq.Mpos_menu_category_id).
			Find(&data).Count(&count)

		err = config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id = ?", MenuReq.Mc_id, MenuReq.Mpos_menu_category_id).
			Offset(offset).
			Limit(MenuReq.Paging.PageSize).
			Order(order).Find(&data).Error
	}

	for index := 0; index < len(data); index++ {
		data[index].Mpos_menu_image_ref = utils.CONTENT_URL + "/" + data[index].Mpos_menu_image_ref + ""
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, MenuReq.Paging, count)
}
func LoadActiveMenuByCategory(MenuReq model.MenuReq) map[string]interface{} {
	var data []entity.MiniposMenuEntity
	var count int
	var order = MenuReq.Paging.OrderBy + " " + MenuReq.Paging.SortBy
	var offset = MenuReq.Paging.PageNo*MenuReq.Paging.PageSize - MenuReq.Paging.PageSize
	var err error
	if MenuReq.Mpos_menu_category_id == nil {
		config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id is null and dp_ms_minipos_menu.active = '1' ", MenuReq.Mc_id).
			Find(&data).Count(&count)

		err = config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id is null and dp_ms_minipos_menu.active = '1' ", MenuReq.Mc_id).
			Offset(offset).
			Limit(MenuReq.Paging.PageSize).
			Order(order).Find(&data).Error
	} else {
		config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id = ? and dp_ms_minipos_menu.active = '1' ", MenuReq.Mc_id, MenuReq.Mpos_menu_category_id).
			Find(&data).Count(&count)

		err = config.DBsql().
			Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
			Where("dp_ms_minipos_menu.mc_id = ? and mpos_menu_category_id = ? and dp_ms_minipos_menu.active = '1' ", MenuReq.Mc_id, MenuReq.Mpos_menu_category_id).
			Offset(offset).
			Limit(MenuReq.Paging.PageSize).
			Order(order).Find(&data).Error
	}

	for index := 0; index < len(data); index++ {
		data[index].Mpos_menu_image_ref = utils.CONTENT_URL + "/" + data[index].Mpos_menu_image_ref + ""
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, MenuReq.Paging, count)
}
func LoadAllMenu(MenuReq model.MenuReq) map[string]interface{} {
	var data []entity.MiniposMenuEntity
	var count int
	var order = MenuReq.Paging.OrderBy + " " + MenuReq.Paging.SortBy
	var offset = MenuReq.Paging.PageNo*MenuReq.Paging.PageSize - MenuReq.Paging.PageSize
	var err error

	config.DBsql().
		Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
		Where("dp_ms_minipos_menu.mc_id = ? ", MenuReq.Mc_id).
		Find(&data).Count(&count)

	err = config.DBsql().
		Joins("left join dp_mp_minipos_menu_merchant on dp_mp_minipos_menu_merchant.mpos_menu_id = dp_ms_minipos_menu.mpos_menu_id").
		Where("dp_ms_minipos_menu.mc_id = ? ", MenuReq.Mc_id).
		Offset(offset).
		Limit(MenuReq.Paging.PageSize).
		Order(order).Find(&data).Error

	for index := 0; index < len(data); index++ {
		data[index].Mpos_menu_image_ref = utils.CONTENT_URL + "/" + data[index].Mpos_menu_image_ref + ""
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, MenuReq.Paging, count)
}
func AddTransactionAndReceive(TransactionReceiveReq model.TransactionReceiveReq, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)

	var err error
	var err2 error

	now := time.Now()
	createBy := utils.Decode(jwt.Raw).(string)

	var MiniposReceiveEntity entity.MiniposReceiveEntity
	trx_No := ""
	if err == nil {

		MiniposReceiveEntity.Mc_id = TransactionReceiveReq.Receive.Mc_id
		MiniposReceiveEntity.Mpos_payment_method_id = TransactionReceiveReq.Receive.Mpos_payment_method_id
		MiniposReceiveEntity.Total_charge_amt = TransactionReceiveReq.Receive.Total_charge_amt
		MiniposReceiveEntity.Receive_amt = TransactionReceiveReq.Receive.Receive_amt
		MiniposReceiveEntity.Change_amt = TransactionReceiveReq.Receive.Change_amt
		MiniposReceiveEntity.Create_date = &now
		MiniposReceiveEntity.Create_by = createBy
		trx_No = genTrxNo(MiniposReceiveEntity.Mc_id)
		MiniposReceiveEntity.Cust_Mobile_No = TransactionReceiveReq.Receive.Cust_Mobile_No
		MiniposReceiveEntity.Trx_No = trx_No
		MiniposReceiveEntity.Status_Flag = "A"
		MiniposReceiveEntity.Discount_Amt = TransactionReceiveReq.Receive.Discount_Amt
		MiniposReceiveEntity.Discount_Type = TransactionReceiveReq.Receive.Discount_Type
		err2 = config.DBsql().Create(&MiniposReceiveEntity).Error

	}
	for index := 0; index < len(TransactionReceiveReq.Transaction_Purchase); index++ {
		var MiniposTransactionDetailEntity entity.MiniposTransactionDetailEntity
		menu_id := TransactionReceiveReq.Transaction_Purchase[index].Mpos_Menu_Id
		MiniposTransactionDetailEntity.Mc_id = TransactionReceiveReq.Transaction_Purchase[index].Mc_id
		MiniposTransactionDetailEntity.Mpos_menu_id = menu_id
		MiniposTransactionDetailEntity.Calculator_keyin = TransactionReceiveReq.Transaction_Purchase[index].Calculator_Keyin
		MiniposTransactionDetailEntity.Price = TransactionReceiveReq.Transaction_Purchase[index].Price
		MiniposTransactionDetailEntity.Amt = TransactionReceiveReq.Transaction_Purchase[index].Amt
		MiniposTransactionDetailEntity.Create_date = &now
		MiniposTransactionDetailEntity.Create_by = createBy

		MiniposTransactionDetailEntity.Mpos_receive_id = MiniposReceiveEntity.Mpos_receive_id
		err = config.DBsql().Create(&MiniposTransactionDetailEntity).Error

	}
	for index := 0; index < len(TransactionReceiveReq.Transaction_Discount); index++ {
		var MiniposTransactionDistcountDetailEntity entity.MiniposTransactionDistcountDetailEntity
		menu_id := TransactionReceiveReq.Transaction_Discount[index].Mpos_Menu_Id
		MiniposTransactionDistcountDetailEntity.Mpos_Receive_Id = MiniposReceiveEntity.Mpos_receive_id
		MiniposTransactionDistcountDetailEntity.Mc_Id = TransactionReceiveReq.Transaction_Discount[index].Mc_id
		MiniposTransactionDistcountDetailEntity.Mpos_Menu_Id = menu_id
		MiniposTransactionDistcountDetailEntity.Calculator_Keyin = TransactionReceiveReq.Transaction_Discount[index].Calculator_Keyin
		MiniposTransactionDistcountDetailEntity.Discount_Type = TransactionReceiveReq.Transaction_Discount[index].Discount_Type
		MiniposTransactionDistcountDetailEntity.Discount_Amt = TransactionReceiveReq.Transaction_Discount[index].Discount_Amt
		MiniposTransactionDistcountDetailEntity.Remaining_Price = TransactionReceiveReq.Transaction_Discount[index].Remaining_Price
		MiniposTransactionDistcountDetailEntity.Remaining_Item = TransactionReceiveReq.Transaction_Discount[index].Remaining_Item
		MiniposTransactionDistcountDetailEntity.Sub_Total = TransactionReceiveReq.Transaction_Discount[index].Sub_Total
		MiniposTransactionDistcountDetailEntity.Created_Date = &now
		MiniposTransactionDistcountDetailEntity.Created_By = createBy

		err = config.DBsql().Create(&MiniposTransactionDistcountDetailEntity).Error

	}
	for index := 0; index < len(TransactionReceiveReq.Transaction_Redeem); index++ {
		var MiniposTransactionRedeemDetailEntity entity.MiniposTransactionRedeemDetailEntity
		menu_id := TransactionReceiveReq.Transaction_Redeem[index].Mpos_Menu_Id
		MiniposTransactionRedeemDetailEntity.Mpos_Receive_Id = MiniposReceiveEntity.Mpos_receive_id
		MiniposTransactionRedeemDetailEntity.Mc_Id = TransactionReceiveReq.Transaction_Redeem[index].Mc_id
		MiniposTransactionRedeemDetailEntity.Mpos_Menu_Id = menu_id
		MiniposTransactionRedeemDetailEntity.Calculator_Keyin = TransactionReceiveReq.Transaction_Redeem[index].Calculator_Keyin
		MiniposTransactionRedeemDetailEntity.Redeem_Amt = TransactionReceiveReq.Transaction_Redeem[index].Redeem_Amt
		MiniposTransactionRedeemDetailEntity.Original_Price = TransactionReceiveReq.Transaction_Redeem[index].Original_Price
		MiniposTransactionRedeemDetailEntity.Original_Item = TransactionReceiveReq.Transaction_Redeem[index].Original_Item
		MiniposTransactionRedeemDetailEntity.Sub_Total = TransactionReceiveReq.Transaction_Redeem[index].Sub_Total
		MiniposTransactionRedeemDetailEntity.Created_Date = &now
		MiniposTransactionRedeemDetailEntity.Created_By = createBy

		err = config.DBsql().Create(&MiniposTransactionRedeemDetailEntity).Error

	}
	member_point := ""
	if TransactionReceiveReq.Receive.Cust_Mobile_No != "" {
		member_point = pointLink(TransactionReceiveReq, ctx)
	}

	return utils.ResDataReceipt(err2, trx_No, member_point)

}
func pointLink(transaction_Purchase model.TransactionReceiveReq, ctx iris.Context) string {

	mcId := transaction_Purchase.Receive.Mc_id
	//total_amount := transaction_Purchase.Receive.Total_charge_amt
	cust_mobile := transaction_Purchase.Receive.Cust_Mobile_No
	member_point := ""
	type PointLinkData struct {
		Loyalty_Point_Link_flag string  `json:"loyalty_point_link_flag"`
		Reward_Type             string  `json:"reward_type"`
		Reward_Amount_Amt       float64 `json:"reward_amount_amt"`
		Reward_Amount_Point     int     `json:"reward_amount_point"`
		Container_Id            int     `json:"container_id"`
		Reward_Item_Type        string  `json:"reward_item_type"`
	}
	pointLinkData := PointLinkData{}
	sql := `select 
				msetting.loyalty_point_link_flag,
				mreward.reward_type,
				mreward.reward_amount_amt,
				mreward.reward_amount_point ,
				mreward.container_id,
				mreward.reward_item_type
				FROM
  					 dp_ms_minipos_point_reward mreward left join dp_ms_minipos_setting msetting
					on mreward.minipos_setting_id=msetting.minipos_setting_id 
				where 
					mreward.mc_id= ?  `

	config.DBsql().Raw(sql, mcId).
		Scan(&pointLinkData)
	if pointLinkData.Loyalty_Point_Link_flag == "Y" &&
		pointLinkData.Reward_Type == "A" && pointLinkData.Reward_Amount_Amt != 0 &&
		pointLinkData.Reward_Amount_Point != 0 {

		quantity := 0
		total_amount := 0.0
		reward_item_type := pointLinkData.Reward_Item_Type
		// A=All Items , O=Only Items , E=Exclude Items
		for index := 0; index < len(transaction_Purchase.Transaction_Purchase); index++ {
			menu_id := transaction_Purchase.Transaction_Purchase[index].Mpos_Menu_Id
			if menu_id != nil {
				type CountResult struct {
					Count int
				}
				var countResult CountResult
				if reward_item_type != "A" {
					condition := ""
					if reward_item_type == "E" {
						condition = " not "
					}
					sql_item := `select count(1) as count from dp_ms_minipos_item_reward
							 where ` + condition + ` exists(
							 select * from dp_ms_minipos_item_reward
								 where mpos_menu_id = ? 
								 and mc_id = ? and reward_item_type= ? 
								 )`
					config.DBsql().Raw(sql_item, menu_id, mcId, reward_item_type).
						Scan(&countResult)
					if countResult.Count > 0 {
						amt := transaction_Purchase.Transaction_Purchase[index].Amt
						price := transaction_Purchase.Transaction_Purchase[index].Price
						quantity = quantity + amt
						total_amount = total_amount + (price * float64(amt))
					}
				} else {
					quantity = 1
					amt := transaction_Purchase.Transaction_Purchase[index].Amt
					price := transaction_Purchase.Transaction_Purchase[index].Price
					total_amount = total_amount + (price * float64(amt))
				}
			} else {
				quantity = 1
				amt := transaction_Purchase.Transaction_Purchase[index].Amt
				price := transaction_Purchase.Transaction_Purchase[index].Price
				total_amount = total_amount + (price * float64(amt))
			}
		}

		if quantity > 0 {
			point := int(total_amount / pointLinkData.Reward_Amount_Amt)
			point_multply := int(pointLinkData.Reward_Amount_Point)
			//println(point)
			if point > 0 {
				pointMaster := model.PointEntityReq{}

				var data []model.CustomerResp
				pointMaster.Container_id = pointLinkData.Container_Id
				pointMaster.Expire_flag = ""
				pointMaster.Mc_id = mcId
				pointMaster.Point_amt = point * point_multply
				pointMaster.Transfer_to_cust_id = 0
				pointMaster.Cust_mobile = cust_mobile
				pointMaster.Status = "1"
				config.DBsql().Where("cust_mobile = ?", cust_mobile).Find(&data)
				if len(data) > 0 {
					pointMaster.Cust_id = data[0].Cust_id
					pointMaster.Status = "3"
				}
				AddPoint(pointMaster, ctx)
				member_point = strconv.Itoa(pointMaster.Point_amt)
			}
		}

		/*
			{
				"container_id": 68,
				"expire_flag": "",
				"cust_id": 22,
				"mc_id": 54,
				"menu_id": 0,
				"point_amt": 10,
				"transfer_to_cust_id": 0,
				"status": "3",
				"cust_mobile": "0828442963"
			}
		*/

	} else if pointLinkData.Loyalty_Point_Link_flag == "Y" &&
		pointLinkData.Reward_Type == "Q" && pointLinkData.Reward_Amount_Amt != 0 && pointLinkData.Reward_Amount_Point != 0 {
		quantity := 0
		reward_item_type := pointLinkData.Reward_Item_Type
		// A=All Items , O=Only Items , E=Exclude Items

		for index := 0; index < len(transaction_Purchase.Transaction_Purchase); index++ {
			menu_id := transaction_Purchase.Transaction_Purchase[index].Mpos_Menu_Id
			if menu_id != nil {
				type CountResult struct {
					Count int
				}
				var countResult CountResult
				if reward_item_type != "A" {
					condition := ""
					if reward_item_type == "E" {
						condition = " not "
					}
					sql_item := `select count(1) as count from dp_ms_minipos_item_reward
							 where ` + condition + ` exists(
							 select * from dp_ms_minipos_item_reward
								 where mpos_menu_id = ? 
								 and mc_id = ? and reward_item_type = ? 
								 )`
					config.DBsql().Raw(sql_item, menu_id, mcId, reward_item_type).
						Scan(&countResult)
					if countResult.Count > 0 {
						amt := transaction_Purchase.Transaction_Purchase[index].Amt
						quantity = quantity + amt
					}
				} else {
					amt := transaction_Purchase.Transaction_Purchase[index].Amt
					quantity = quantity + amt
				}
			}
		}

		if quantity > 0 {
			point := int(quantity / int(pointLinkData.Reward_Amount_Amt))
			point_multply := int(pointLinkData.Reward_Amount_Point)
			if point > 0 {
				pointMaster := model.PointEntityReq{}
				var data []model.CustomerResp
				pointMaster.Container_id = pointLinkData.Container_Id
				pointMaster.Expire_flag = ""
				pointMaster.Mc_id = mcId
				pointMaster.Point_amt = point * point_multply
				pointMaster.Transfer_to_cust_id = 0
				pointMaster.Cust_mobile = cust_mobile
				pointMaster.Status = "1"
				config.DBsql().Where("cust_mobile = ?", cust_mobile).Find(&data)
				if len(data) > 0 {
					pointMaster.Cust_id = data[0].Cust_id
					pointMaster.Status = "3"
				}
				AddPoint(pointMaster, ctx)
				member_point = strconv.Itoa(pointMaster.Point_amt)
			}
		}
	}
	return member_point
}
func genTrxNo(mcId int) string {
	mcId_padded := fmt.Sprintf("%05d", mcId)
	token := utils.RandString(8)
	create_time := time.Now()
	year, month, day := time.Now().Date()
	month_to_int := int(month)

	year_str := strconv.Itoa(year)
	month_str := fmt.Sprintf("%02d", month_to_int)
	day_str := fmt.Sprintf("%02d", day)

	hour_str := fmt.Sprintf("%02d", create_time.Hour())
	minute_str := fmt.Sprintf("%02d", create_time.Minute())
	return mcId_padded + "-" + year_str + month_str + day_str + "-" + hour_str + minute_str + "-" + strings.ToUpper(token)
}
func GetMiniPOSCategoryById(id string) map[string]interface{} {

	var data entity.MiniposMenuCategoryEntity
	check := config.DBsql().Where("mpos_menu_category_id = ?", id).Find(&data).RecordNotFound()

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func AddMiniPOSCategory(MiniposMenuCategoryEntity entity.MiniposMenuCategoryEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	createBy := utils.Decode(jwt.Raw).(string)
	now := time.Now()
	MiniposMenuCategoryEntity.Create_by = createBy
	MiniposMenuCategoryEntity.Create_date = &now
	err := config.DBsql().Create(&MiniposMenuCategoryEntity).Error
	return utils.ResDataAdd(err)
}
func UpdateMiniPOSCategory(MiniposMenuCategoryEntity entity.MiniposMenuCategoryEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	now := time.Now()
	MiniposMenuCategoryEntity.Update_by = utils.Decode(jwt.Raw).(string)
	MiniposMenuCategoryEntity.Update_date = &now

	err := config.DBsql().Model(&MiniposMenuCategoryEntity).Where("mpos_menu_category_id = ? AND mc_id = ? ",
		MiniposMenuCategoryEntity.Mpos_menu_category_id, MiniposMenuCategoryEntity.Mc_id).Update(&MiniposMenuCategoryEntity).Error

	return utils.ResDataEdit(err)
}
func DeleteMiniPOSCategory(data []entity.MiniposMenuCategoryEntity) map[string]string {
	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("mpos_menu_category_id = ? AND mc_id = ?", data[i].Mpos_menu_category_id, data[i].Mc_id).
			Table("dp_ms_minipos_menu_category").Count(&count).Delete(entity.DeviceMerchantMapEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataDel(len(data), deleted)
}
func GetMiniposMenuById(id string) map[string]interface{} {
	var data entity.MiniposMenuEntity
	check := config.DBsql().Where("mpos_menu_id = ?", id).Find(&data).RecordNotFound()

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func AddMiniposMenu(MiniposMenuEntity entity.MiniposMenuEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	createBy := utils.Decode(jwt.Raw).(string)
	now := time.Now()
	MiniposMenuEntity.Create_by = createBy
	MiniposMenuEntity.Create_date = &now
	err := config.DBsql().Create(&MiniposMenuEntity).Error
	return utils.ResDataAdd(err)
}
func UpdateMiniposMenu(MiniposMenuEntity entity.MiniposMenuEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	now := time.Now()
	MiniposMenuEntity.Update_by = utils.Decode(jwt.Raw).(string)
	MiniposMenuEntity.Update_date = &now

	err := config.DBsql().Model(&MiniposMenuEntity).Where("mpos_menu_id = ? AND mc_id = ? ",
		MiniposMenuEntity.Mpos_menu_id, MiniposMenuEntity.Mc_id).Update(&MiniposMenuEntity).Error

	return utils.ResDataEdit(err)
}
func DeleteMiniposMenu(data []entity.MiniposMenuEntity) map[string]string {
	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("mpos_menu_id = ? AND mc_id = ?", data[i].Mpos_menu_id, data[i].Mc_id).
			Table("dp_ms_minipos_menu").Count(&count).Delete(entity.MiniposMenuEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataDel(len(data), deleted)
}

func AddMiniposMappingMenu(data []entity.MiniposMenuEntity) map[string]string {
	var count int
	var added = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("mpos_menu_id = ? AND mc_id = ?", data[i].Mpos_menu_id, data[i].Mc_id).
			Table("dp_ms_minipos_menu").Count(&count).Update("mpos_menu_category_id", data[i].Mpos_menu_category_id) //Delete(entity.MiniposMenuEntity{})
		if count != 0 {
			added = added + 1
		}
	}

	return utils.ResDataMapping(len(data), added)
}

func DeleteMiniposMappingMenu(data []entity.MiniposMenuEntity) map[string]string {
	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("mpos_menu_id = ? AND mc_id = ?", data[i].Mpos_menu_id, data[i].Mc_id).
			Table("dp_ms_minipos_menu").Count(&count).Update("mpos_menu_category_id", nil) //Delete(entity.MiniposMenuEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataRemvoveMapping(len(data), deleted)
}

func SearchReceiptHistory(MiniposReceiptHistoryReq model.MiniposReceiptHistoryReq) map[string]interface{} {
	type DataResp struct {
		Mpos_Receive_Id  int        `json:"mpos_receive_id"`
		Hh24             string     `json:"hh24"`
		Hh_time          string     `json:"hh_time"`
		Trx_no           string     `json:"trx_no"`
		Total_charge_amt float64    `json:"total_charge_amt"`
		Status_flag      string     `json:"status_flag"`
		Cust_mobile_no   string     `json:"cust_mobile_no"`
		Mc_Id            int        `json:"mc_id"`
		Create_date      *time.Time `json:"create_date"`
	}

	type Items struct {
		History_date string     `json:"history_date"`
		DataRespList []DataResp `json:"items"`
	}

	year, month, day := time.Now().Date()
	month_to_int := int(month)

	year_str := strconv.Itoa(year)
	month_str := fmt.Sprintf("%02d", month_to_int)
	day_str := fmt.Sprintf("%02d", day)
	day_filter := year_str + "-" + month_str + "-" + day_str
	// '2018-11-14 00:00:00'
	//'2018-11-14 23:59:50'
	start_time := day_filter + " 00:00:00"
	finish_time := day_filter + " 23:59:50"
	dataresp := []DataResp{}

	sql := `SELECT
			 mpos_receive_id,
			 TO_CHAR(create_date, 'HH24') as hh24 ,
			 TO_CHAR(create_date, 'HH24:MI:SS') as hh_time ,
			 trx_no ,
 			 total_charge_amt,
			 status_flag,
             cust_mobile_no,
			 create_date
		FROM
  			dp_ts_minipos_receive
		WHERE
		    mc_id = ? 
 			AND
       		create_date >= ?
   			AND
			create_date <=  ?  `
	if len(MiniposReceiptHistoryReq.Customer_Mobile_No) > 0 {
		sql = sql + " AND cust_mobile_no = '" + MiniposReceiptHistoryReq.Customer_Mobile_No + "' "
	}
	sql = sql + " order by create_date desc "
	err := config.DBsql().Raw(sql, MiniposReceiptHistoryReq.MC_Id, start_time, finish_time).
		Scan(&dataresp).Error
	hh24 := ""
	dataList := []DataResp{}
	items := []Items{}
	//items := []Items;
	size := len(dataresp)
	if size > 0 {
		hh24 = dataresp[0].Hh24
		for i := 0; i < size; i++ {
			if dataresp[i].Hh24 != hh24 {
				item := Items{}
				history_date := hh24 + ":00 - " + hh24 + ":59"
				item.History_date = history_date
				item.DataRespList = dataList
				items = append(items, item)

				dataList = nil
				dataList = append(dataList, dataresp[i])
				hh24 = dataresp[i].Hh24
			} else {
				dataList = append(dataList, dataresp[i])
			}
		}
		if len(dataList) > 0 {
			item := Items{}
			history_date := hh24 + ":00 - " + hh24 + ":59"
			item.History_date = history_date
			item.DataRespList = dataList
			items = append(items, item)
		}

	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(items)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}

func UpdateReceiptStatus(MiniposReceiptHistoryReq model.MiniposReceiptHistoryReq, ctx iris.Context) map[string]string {
	/*
		jwt := ctx.Values().Get("jwt").(*jwt.Token)
		now := time.Now()

		MiniposReceiveEntity := entity.MiniposReceiveEntity{}
		MiniposReceiveEntity.Update_by = utils.Decode(jwt.Raw).(string)
		MiniposReceiveEntity.Update_date = &now
	*/
	err := config.DBsql().Model(&entity.MiniposReceiveEntity{}).Where("mpos_receive_id = ? ",
		MiniposReceiptHistoryReq.Mpos_Receive_Id).Update(
		map[string]interface{}{"status_flag": MiniposReceiptHistoryReq.Status_Flag}).Error

	return utils.ResDataEdit(err)
}
func GetMiniPOSReport(miniposReportReq model.MiniposReportReq) map[string]interface{} {
	var data model.MiniposReportRes

	check := false
	mc_id := strconv.Itoa(miniposReportReq.MC_Id)
	from_date := miniposReportReq.From_date
	to_date := miniposReportReq.To_date
	from_date_array := strings.Split(from_date, "/")
	to_date_array := strings.Split(to_date, "/")

	start_day := from_date_array[0] + "-" + from_date_array[1] + "-" + from_date_array[2]
	finish_day := to_date_array[0] + "-" + to_date_array[1] + "-" + to_date_array[2]

	start_time := start_day + " 00:00:00"
	finish_time := finish_day + " 23:59:50"

	sql := `
select
        coalesce(total_bill.net_sales,0)+
        coalesce(redeem_amount.discount_amt,0)+
        coalesce(discount_amount.discount_amt,0)+
        coalesce(discount_percent.discount_amt,0) net_sales

       ,coalesce(total_bill.net_sales,0) sales
       ,coalesce(total_bill.net_sales,0) -
        (coalesce(bill_discount_amount.discount_amt,0)+
         coalesce(bill_discount_percent.discount_amt,0))
       total_net_sales
       ,CASE WHEN coalesce(total_bill.total_bills_receipts,0)=0 then 0
            WHEN  coalesce(total_bill.total_bills_receipts,0)>0 THEN
           --   (coalesce(total_bill.net_sales,0) +
        -- coalesce(redeem_amount.discount_amt,0)+
        -- coalesce(discount_amount.discount_amt,0)+
       -- coalesce(discount_percent.discount_amt,0))/coalesce(total_bill.total_bills_receipts,0)
		(coalesce(total_bill.net_sales,0) -
        coalesce(bill_discount_amount.discount_amt,0)+
        coalesce(bill_discount_percent.discount_amt,0))/coalesce(total_bill.total_bills_receipts,0)
           -- ELSE 'other'
       END average_sales

       ,coalesce(cash.cash_amt,0)
        -- +coalesce(item_disc_cash.item_discount_amt,0)+
        -- coalesce(item_percent_disc_cash.item_discount_amt,0)
		payment_received_cash

       ,coalesce(qr.qr_amt,0)
        -- +coalesce(item_disc_qr.item_discount_amt,0)+
        -- coalesce(item_percent_disc_qr.item_discount_amt,0) 
		payment_received_qr

       ,coalesce(credit.credit_amt,0)
        -- +coalesce(item_disc_credit.item_discount_amt,0)+
        -- coalesce(item_percent_disc_credit.item_discount_amt,0) 
		payment_received_credits

       ,coalesce(item_disc_cash.item_discount_amt,0)+
        coalesce(item_percent_disc_cash.item_discount_amt,0) item_discount_amt_cash

       ,coalesce(item_disc_qr.item_discount_amt,0)+
        coalesce(item_percent_disc_qr.item_discount_amt,0) item_discount_amt_qr

       ,coalesce(item_disc_credit.item_discount_amt,0)+
        coalesce(item_percent_disc_credit.item_discount_amt,0) item_discount_amt_credit

       ,coalesce(total_bill.total_bills_receipts,0) total_bills_receipts
       -- ,all_amt.total_bill x
       ,coalesce(discount_amount.discount_amt,0) discount_amount_value
       ,coalesce(discount_percent.discount_amt,0) discount_percent_value


      ,coalesce(redeem_amount.discount_amt,0) redeem_value
      ,coalesce(redeem_amount.redeem_count,0) redeem_item

       ,coalesce(bill_discount_amount.discount_amt,0) bill_disc_amt
       ,coalesce(bill_discount_percent.discount_amt,0) bill_disc_percent


       ,coalesce(bill_discount_amount.discount_amt,0)+coalesce(bill_discount_percent.discount_amt,0) as bill_discount_value

       ,coalesce(discount_amount.discount_amt,0)+
       coalesce(discount_percent.discount_amt,0) item_discount_value

       ,coalesce(discount_amount.discount_amt,0)+
        coalesce(discount_percent.discount_amt,0)+
        coalesce(redeem_amount.discount_amt,0)+
        coalesce(bill_discount_amount.discount_amt,0)+
        coalesce(bill_discount_percent.discount_amt,0) total_discount_redeem
from
 (
  SELECT
   -- to_char(sum(total_charge_amt),'FM999,999,999,990D00') as net_sales
    count(*) as total_bills_receipts
    ,sum(total_charge_amt) as net_sales
		FROM dp_ts_minipos_receive
    where mc_id=` + mc_id + ` and status_flag = 'A'
	   and create_date >= '` + start_time + `'
      and
			create_date <=  '` + finish_time + `'
   ) total_bill ,
 (
  SELECT
   -- to_char(sum(total_charge_amt),'FM999,999,999,990D00') as cash_amt
   sum(total_charge_amt) as cash_amt
		FROM dp_ts_minipos_receive
		where mc_id=` + mc_id + ` and status_flag = 'A' and mpos_payment_method_id=1
		and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
) cash,
(
  SELECT
    -- to_char(sum(total_charge_amt),'FM999,999,999,990D00') as qr_amt
    sum(total_charge_amt) as qr_amt
		FROM dp_ts_minipos_receive
		where mc_id=` + mc_id + ` and status_flag = 'A' and mpos_payment_method_id=2
		and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
) qr,
(
  SELECT
   --  to_char(sum(total_charge_amt),'FM999,999,999,990D00') as credit_amt
   sum(total_charge_amt) as credit_amt
		FROM dp_ts_minipos_receive
		where mc_id=` + mc_id + ` and status_flag = 'A' and mpos_payment_method_id=3
		and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
) credit ,
 (
 select
 sum(mtdd.discount_amt) as item_discount_amt
 from dp_ts_minipos_transactoin_discount_detail mtdd left join
  dp_ts_minipos_receive mr on mr.mpos_receive_id = mtdd.mpos_receive_id
where mtdd.mc_id=` + mc_id + ` and status_flag = 'A' 
and mtdd.discount_type='A'
and mr.mpos_payment_method_id=1
   and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
	 ) item_disc_cash,
     (
 select
 sum(mtdd.discount_amt) as item_discount_amt
 from dp_ts_minipos_transactoin_discount_detail mtdd left join
  dp_ts_minipos_receive mr on mr.mpos_receive_id = mtdd.mpos_receive_id
where mtdd.mc_id=` + mc_id + ` and status_flag = 'A' 
and mtdd.discount_type='A'
and mr.mpos_payment_method_id=2
   and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
	 ) item_disc_qr,
     (
 select
 sum(mtdd.discount_amt) as item_discount_amt
 from dp_ts_minipos_transactoin_discount_detail mtdd left join
  dp_ts_minipos_receive mr on mr.mpos_receive_id = mtdd.mpos_receive_id
where mtdd.mc_id=` + mc_id + ` and status_flag = 'A' 
and mtdd.discount_type='A'
and mr.mpos_payment_method_id=3
   and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
	 ) item_disc_credit,
     (
 select
 sum(((mtdd.remaining_price*mtdd.remaining_item)*mtdd.discount_amt)/100) as item_discount_amt
 from dp_ts_minipos_transactoin_discount_detail mtdd left join
  dp_ts_minipos_receive mr on mr.mpos_receive_id = mtdd.mpos_receive_id
where mtdd.mc_id=` + mc_id + ` and status_flag = 'A' 
and mtdd.discount_type='C'
and mr.mpos_payment_method_id=1
   and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
	 ) item_percent_disc_cash,
     (
 select
 sum(((mtdd.remaining_price*mtdd.remaining_item)*mtdd.discount_amt)/100) as item_discount_amt
 from dp_ts_minipos_transactoin_discount_detail mtdd left join
  dp_ts_minipos_receive mr on mr.mpos_receive_id = mtdd.mpos_receive_id
where mtdd.mc_id=` + mc_id + ` and status_flag = 'A' 
and mtdd.discount_type='C'
and mr.mpos_payment_method_id=2
   and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
	 ) item_percent_disc_qr,
     (
 select
 sum(((mtdd.remaining_price*mtdd.remaining_item)*mtdd.discount_amt)/100) as item_discount_amt
 from dp_ts_minipos_transactoin_discount_detail mtdd left join
  dp_ts_minipos_receive mr on mr.mpos_receive_id = mtdd.mpos_receive_id
where mtdd.mc_id=` + mc_id + ` and status_flag = 'A' 
and mtdd.discount_type='C'
and mr.mpos_payment_method_id=3
   and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
	 ) item_percent_disc_credit,
( SELECT
    sum(discount_amt) as discount_amt
		FROM dp_ts_minipos_transactoin_discount_detail
		where mc_id=` + mc_id + `
		and discount_type='A'
		and
			created_date >= '` + start_time + `'
		AND
			created_date <= '` + finish_time + `'
)discount_amount,
( SELECT
    sum(((remaining_price*remaining_item)*discount_amt)/100) as discount_amt
		FROM dp_ts_minipos_transactoin_discount_detail
		where mc_id=` + mc_id + `
		and discount_type='C'
		and
			created_date >= '` + start_time + `'
		AND
			created_date <= '` + finish_time + `'
)discount_percent,
( SELECT
    sum(redeem_amt) as redeem_count
    ,sum((redeem_amt*original_price)) as discount_amt
		FROM dp_ts_minipos_transactoin_redeem_detail
		where mc_id=` + mc_id + `
		and
			created_date >= '` + start_time + `'
		AND
			created_date <= '` + finish_time + `'
)redeem_amount,
     ( SELECT
    sum(discount_amt) as discount_amt
		FROM dp_ts_minipos_receive
		where mc_id=` + mc_id + ` and status_flag = 'A' 
		and discount_type='A'
		and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
)bill_discount_amount,
( SELECT
    sum((total_charge_amt*discount_amt)/100) as discount_amt
		FROM dp_ts_minipos_receive
		where mc_id=` + mc_id + ` and status_flag = 'A' 
		and discount_type='C'
		and
			create_date >= '` + start_time + `'
		AND
			create_date <= '` + finish_time + `'
)bill_discount_percent
			`
	/*
		err := config.DBsql().Raw(sql,mc_id,start_time,finish_time).
			Scan(&data).Error
		if(err!=nil){
			println(err)
		}
	*/
	err := config.DBsql().Raw(sql).
		Scan(&data).Error
	if err != nil {
		println(err)
	}
	type MenuResp struct {
		MPOS_menu_id    int `json:"mpos_menu_id"`
		MPOS_menu_count int `json:"mpos_menu_count"`
	}
	var menuResp []MenuResp
	sql = `
		select * from (
									 select sum(amt) mpos_menu_count ,mpos_menu_id
									 from dp_ts_minipos_transaction_detail
									 where mc_id =` + mc_id + `
										and
										create_date >= '` + start_time + `'
										AND
										create_date <= '` + finish_time + `'
										 and mpos_menu_id is not null
									 group by mpos_menu_id
								 )x order by  x.mpos_menu_count desc limit 10
	`
	err = config.DBsql().Raw(sql).
		Scan(&menuResp).Error
	if err != nil {
		println(err)
	}
	//BestSellerMenuList []BestSellerMenuRes `json:"bestSeller_list"`
	var bestSellerMenuList []model.BestSellerMenuRes
	for index := 0; index < len(menuResp); index++ {
		var bestSellerMenuRes model.BestSellerMenuRes
		var miniposMenu entity.MiniposMenuEntity
		config.DBsql().Where("mpos_menu_id = ? ", menuResp[index].MPOS_menu_id).Find(&miniposMenu)
		miniposMenu.Mpos_menu_image_ref = utils.CONTENT_URL + "/" + miniposMenu.Mpos_menu_image_ref + ""
		bestSellerMenuRes.MiniposMenu = miniposMenu
		bestSellerMenuRes.MenuCount = menuResp[index].MPOS_menu_count
		bestSellerMenuList = append(bestSellerMenuList, bestSellerMenuRes)
	}
	data.BestSellerMenuList = bestSellerMenuList

	type SalesTrafficResp struct {
		Count_time00_06 int `json:"count_time00_06"`
		Count_time06_09 int `json:"count_time06_09"`
		Count_time10_12 int `json:"count_time10_12"`
		Count_time13_15 int `json:"count_time13_15"`
		Count_time16_18 int `json:"count_time16_18"`
		Count_time19_21 int `json:"count_time19_21"`
		Count_time22_23 int `json:"count_time22_23"`
	}
	var salesTrafficResp SalesTrafficResp
	sql = `
	select * from (
			  				select
									count(mpos_receive_id) as count_time00_06
								from (
									SELECT *
										FROM dp_ts_minipos_receive
										WHERE mc_id = ` + mc_id + `
										and status_flag = 'A'
										AND
											create_date >= '` + start_day + ` 00:00:00'
										AND
											create_date <= '` + finish_day + ` 23:59:59'
								)x where
							      TO_CHAR(x.create_date,'HH24:MI:SS')>='00:00:00'
								  AND
							      TO_CHAR(x.create_date,'HH24:MI:SS')<='05:59:59'
							) time00,(
							  select
									count(mpos_receive_id) as count_time06_09
								from (
									SELECT *
										FROM dp_ts_minipos_receive
										WHERE mc_id = ` + mc_id + `
										and status_flag = 'A'
										AND
											create_date >= '` + start_day + ` 00:00:00'
										AND
											create_date <= '` + finish_day + ` 23:59:59'
								)x where
							      TO_CHAR(x.create_date,'HH24:MI:SS')>='06:00:00'
								  AND
							      TO_CHAR(x.create_date,'HH24:MI:SS')<='09:59:59'
							) time1,
              (
                select
									count(mpos_receive_id) as count_time10_12
								from (
									SELECT *
										FROM dp_ts_minipos_receive
										WHERE mc_id = ` + mc_id + `
										and status_flag = 'A'
										AND
											create_date >= '` + start_day + ` 00:00:00'
										AND
											create_date <= '` + finish_day + ` 23:59:59'
								)x where
							      TO_CHAR(x.create_date,'HH24:MI:SS')>='10:00:00'
								  AND
							      TO_CHAR(x.create_date,'HH24:MI:SS')<='12:59:59'
							) time2,
              (
                select
									count(mpos_receive_id) as count_time13_15
								from (
									SELECT *
										FROM dp_ts_minipos_receive
										WHERE mc_id = ` + mc_id + `
										and status_flag = 'A'
										AND
											create_date >= '` + start_day + ` 00:00:00'
										AND
											create_date <= '` + finish_day + ` 23:59:59'
								)x where
							      TO_CHAR(x.create_date,'HH24:MI:SS')>='13:00:00'
								  AND
							      TO_CHAR(x.create_date,'HH24:MI:SS')<='15:59:59'
							) time3,
              (
                select
									count(mpos_receive_id) as count_time16_18
								from (
									SELECT *
										FROM dp_ts_minipos_receive
										WHERE mc_id = ` + mc_id + `
										and status_flag = 'A'
										AND
											create_date >= '` + start_day + ` 00:00:00'
										AND
											create_date <= '` + finish_day + ` 23:59:59'
								)x where
							      TO_CHAR(x.create_date,'HH24:MI:SS')>='16:00:00'
								  AND
							      TO_CHAR(x.create_date,'HH24:MI:SS')<='18:59:59'
							) time4,
              (
                select
									count(mpos_receive_id) as count_time19_21
								from (
									SELECT *
										FROM dp_ts_minipos_receive
										WHERE mc_id = ` + mc_id + `
										and status_flag = 'A'
										AND
											create_date >= '` + start_day + ` 00:00:00'
										AND
											create_date <= '` + finish_day + ` 23:59:59'
								)x where
							      TO_CHAR(x.create_date,'HH24:MI:SS')>='19:00:00'
								  AND
							      TO_CHAR(x.create_date,'HH24:MI:SS')<='21:59:59'
							) time5,
              (
                select
									count(mpos_receive_id) as count_time22_23
								from (
									SELECT *
										FROM dp_ts_minipos_receive
										WHERE mc_id = ` + mc_id + `
										and status_flag = 'A'
										AND
											create_date >= '` + start_day + ` 00:00:00'
										AND
											create_date <= '` + finish_day + ` 23:59:59'
								)x where
							      TO_CHAR(x.create_date,'HH24:MI:SS')>='22:00:00'
								  AND
							      TO_CHAR(x.create_date,'HH24:MI:SS')<='23:59:59'
							) time6
	`
	err = config.DBsql().Raw(sql).
		Scan(&salesTrafficResp).Error
	if err != nil {
		println(err)
	}
	var sales_Traffice_List []model.SalesTrafficeRes

	salesTrafficeRes_time_00_06 := model.SalesTrafficeRes{}
	salesTrafficeRes_time_00_06.Sale_Time = "00:00 - 05:59"
	salesTrafficeRes_time_00_06.Sale_Count = salesTrafficResp.Count_time00_06

	salesTrafficeRes_time_06_09 := model.SalesTrafficeRes{}
	salesTrafficeRes_time_06_09.Sale_Time = "06:00 - 09:59"
	salesTrafficeRes_time_06_09.Sale_Count = salesTrafficResp.Count_time06_09

	salesTrafficeRes_time_10_12 := model.SalesTrafficeRes{}
	salesTrafficeRes_time_10_12.Sale_Time = "10:00 - 12:59"
	salesTrafficeRes_time_10_12.Sale_Count = salesTrafficResp.Count_time10_12

	salesTrafficeRes_time_13_15 := model.SalesTrafficeRes{}
	salesTrafficeRes_time_13_15.Sale_Time = "13:00 - 15:59"
	salesTrafficeRes_time_13_15.Sale_Count = salesTrafficResp.Count_time13_15

	salesTrafficeRes_time_16_18 := model.SalesTrafficeRes{}
	salesTrafficeRes_time_16_18.Sale_Time = "16:00 - 18:59"
	salesTrafficeRes_time_16_18.Sale_Count = salesTrafficResp.Count_time16_18

	salesTrafficeRes_time_19_21 := model.SalesTrafficeRes{}
	salesTrafficeRes_time_19_21.Sale_Time = "19:00 - 21:59"
	salesTrafficeRes_time_19_21.Sale_Count = salesTrafficResp.Count_time19_21

	salesTrafficeRes_time_22_23 := model.SalesTrafficeRes{}
	salesTrafficeRes_time_22_23.Sale_Time = "22:00 - 23:59"
	salesTrafficeRes_time_22_23.Sale_Count = salesTrafficResp.Count_time22_23

	sales_Traffice_List = append(sales_Traffice_List, salesTrafficeRes_time_00_06)
	sales_Traffice_List = append(sales_Traffice_List, salesTrafficeRes_time_06_09)
	sales_Traffice_List = append(sales_Traffice_List, salesTrafficeRes_time_10_12)
	sales_Traffice_List = append(sales_Traffice_List, salesTrafficeRes_time_13_15)
	sales_Traffice_List = append(sales_Traffice_List, salesTrafficeRes_time_16_18)
	sales_Traffice_List = append(sales_Traffice_List, salesTrafficeRes_time_19_21)
	sales_Traffice_List = append(sales_Traffice_List, salesTrafficeRes_time_22_23)

	peak_time := ""
	peak_count := 0

	lowest_time := ""
	lowest_count := 0
	for index := 0; index < len(sales_Traffice_List); index++ {
		salesTraffice := sales_Traffice_List[index]
		saleCount := salesTraffice.Sale_Count
		saleTime := salesTraffice.Sale_Time
		if index == 0 {
			peak_time = saleTime
			peak_count = saleCount
			lowest_time = saleTime
			lowest_count = saleCount
		}
		if saleCount > peak_count {
			peak_count = saleCount
			peak_time = saleTime
		}

		if saleCount < lowest_count {
			lowest_count = saleCount
			lowest_time = saleTime
		}
	}

	sales_traffice_peak_time := model.SalesTrafficeRes{}
	sales_traffice_peak_time.Sale_Count = peak_count
	sales_traffice_peak_time.Sale_Time = peak_time

	sales_traffice_lowest_time := model.SalesTrafficeRes{}
	sales_traffice_lowest_time.Sale_Count = lowest_count
	sales_traffice_lowest_time.Sale_Time = lowest_time

	data.Sales_Traffice_List = sales_Traffice_List
	data.Sales_Traffice_Peak_Time = sales_traffice_peak_time
	data.Sales_Traffice_Lowest = sales_traffice_lowest_time
	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)
	//return utils.ResDataLoadNonPagin(err, dataToJson)
	return utils.ResDataLoadById(check, dataToJson)
}
