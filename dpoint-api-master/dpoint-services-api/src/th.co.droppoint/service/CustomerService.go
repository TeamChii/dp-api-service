package service

import (
	"encoding/json"

	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func GetCustomerByMobile(customerObj model.CustomerReq) map[string]interface{} {

	var data []model.CustomerResp
	err := config.DBsql().Where("cust_mobile = ?", customerObj.Cust_mobile).Find(&data).Error
	if len(data) > 0 {
		data[0].Image_ref = utils.CONTENT_URL + "/" + data[0].Image_ref + ""

		var lastPurchase []entity.SumCustomerPointEntity
		config.DBsql().Where("cust_id = ?", data[0].Cust_id).Order("last_action_date ASC").Find(&lastPurchase)

		if len(lastPurchase) < 1 {
			data[0].Last_action_date = ""
		} else {
			data[0].Last_action_date = lastPurchase[0].Last_action_date.Format("02/01/2006")

		}
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func GetCustomerByMobileMcId(customerObj model.CustomerReq) map[string]interface{} {

	var data []model.CustomerResp
	err := config.DBsql().Where("cust_mobile = ?", customerObj.Cust_mobile).Find(&data).Error

	if len(data) > 0 {
		data[0].Image_ref = utils.CONTENT_URL + "/" + data[0].Image_ref + ""
		var mc_group entity.MerchantEntity

		config.DBsql().
			Where("mc_id = ? ", customerObj.Mc_id).Find(&mc_group)
		var lastPurchase []entity.SumCustomerPointEntity
		config.DBsql().Where("cust_id = ? AND mc_id = ?", data[0].Cust_id, mc_group.Mc_group_id).Order("last_action_date ASC").Find(&lastPurchase)

		if len(lastPurchase) < 1 {
			data[0].Last_action_date = ""
		} else {
			data[0].Last_action_date = lastPurchase[0].Last_action_date.Format("02/01/2006")

		}

	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func GetCustomerById(id string) map[string]interface{} {

	var data model.CustomerResp
	check := config.DBsql().Where("cust_id = ?", id).Find(&data).RecordNotFound()

	data.Image_ref = utils.CONTENT_URL + "/" + data.Image_ref + ""

	var lastPurchase []entity.SumCustomerPointEntity
	config.DBsql().Where("cust_id = ?", id).Order("last_action_date ASC").Find(&lastPurchase)

	if len(lastPurchase) < 1 {
		data.Last_action_date = ""
	} else {
		data.Last_action_date = lastPurchase[0].Last_action_date.Format("02/01/2006")

	}

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func UpdateCustomer(customerMaster entity.CustomerEntity) map[string]string {

	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	err := config.DBsql().Model(&customerMaster).Where("cust_id = ?", customerMaster.Cust_id).Update(&customerMaster).Error

	return utils.ResDataEdit(err)
}

/*func CustomerSearch(searchObject model.SearchObjectModel) map[string]interface{} {
	log.Info("into CustomerList")
	//var accountModel model.CustomerModel

	var order = searchObject.Paging.OrderBy + " " + searchObject.Paging.SortBy
	var offset = searchObject.Paging.PageNo*searchObject.Paging.PageSize - searchObject.Paging.PageSize

	var count int
	var account []entity.CustomerEntity
	conn := config.DBsql()
	conn.Find(&account).Count(&count) // count record
	err := conn.
		Offset(offset).
		Limit(searchObject.Paging.PageSize).
		Order(order).Find(&account).Error

	for i := 0; i < len(account); i++ {
		if account[i].Create_date.Year() != 1 {
			account[i].Create_date_str = account[i].Create_date.Format("02/01/2006 15:04:05")
		}
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(account)
	json.Unmarshal(b, &dataToJson)
	return utils.ResDataLoad(err, dataToJson, searchObject.Paging, count)
}

func AddCustomer(customerMaster entity.CustomerEntity) map[string]string {

	//	var resModel model.ResponseModel
	customerMaster.KeyId = utils.RandSeq(40)
	customerMaster.Create_date = time.Now()

	err := config.DBsql().Create(&customerMaster).Error
	return utils.ResDataAdd(err)

}


func DeleteCustomer(data []entity.CustomerEntity) map[string]string {

	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("key_id = ?", data[i].KeyId).
			Table("customer_master").Count(&count).Delete(entity.CustomerEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataDel(len(data), deleted)

}*/
