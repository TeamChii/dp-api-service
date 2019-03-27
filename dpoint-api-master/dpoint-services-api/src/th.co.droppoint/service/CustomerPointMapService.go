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

func CustomerPointMapByIdMcCust(CustomerPointMapMaster model.CustomerPointMapReq) map[string]interface{} {

	var data []entity.CustomerPointMapEntity

	var count int
	err := config.DBsql().Where("mc_id = ? AND cust_id = ?", CustomerPointMapMaster.Mc_id, CustomerPointMapMaster.Cust_id).
		Find(&data).Count(&count).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func AddCustomerPointMap(CustomerPointMapMaster entity.CustomerPointMapEntity) map[string]string {

	//	var resModel model.ResponseModel
	//id, _ := strconv.Atoi(utils.RandSeq(32))
	//CustomerPointMapMaster.Mc_id = id
	//CustomerPointMapMaster.Create_date = time.Now()

	err := config.DBsql().Create(&CustomerPointMapMaster).Error
	return utils.ResDataAdd(err)

}

func UpdateCustomerPointMap(CustomerPointMapMaster entity.CustomerPointMapEntity) map[string]string {

	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	err := config.DBsql().Model(&CustomerPointMapMaster).Where("mc_id = ? AND cust_id = ? AND card_type = ?",
		CustomerPointMapMaster.Mc_id, CustomerPointMapMaster.Cust_id, CustomerPointMapMaster.Card_type).
		Update(&CustomerPointMapMaster).Error

	return utils.ResDataEdit(err)
}
