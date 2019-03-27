package service

import (
	"encoding/json"
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
)

func SystemParamSearch(searchObject model.SearchObjectSystemParamModel) map[string]interface{} {

	var order = searchObject.Paging.OrderBy + " " + searchObject.Paging.SortBy
	var offset = searchObject.Paging.PageNo*searchObject.Paging.PageSize - searchObject.Paging.PageSize

	var count int
	var data []entity.SystemParamEntity
	conn := config.DBsql()
	//conn.Find(&data).Where("category_name = ?", searchObject.Category).Count(&count) // count record
	err := conn.
		Where("category_name = ? AND key_code LIKE ?", searchObject.Category, "%"+searchObject.SearchString+"%").
		Offset(offset).
		Limit(searchObject.Paging.PageSize).
		Order(order).Find(&data).Count(&count).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)
	return utils.ResDataLoad(err, dataToJson, searchObject.Paging, count)
}
func SystemParamAdd(SystemParamEntity entity.SystemParamEntity) map[string]string {

	//	var resModel model.ResponseModel
	//id, _ := strconv.Atoi(utils.RandSeq(32))
	//RedeemMaster.Mc_id = id
	//RedeemMaster.Create_date = time.Now()

	err := config.DBsql().Create(&SystemParamEntity).Error

	return utils.ResDataAdd(err)

}

func SystemParamUpdate(SystemParamEntity entity.SystemParamEntity) map[string]string {

	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	err := config.DBsql().Model(&SystemParamEntity).Where("key_code = ?", SystemParamEntity.Key_code).Update(&SystemParamEntity).Error

	return utils.ResDataEdit(err)
}
func SystemParamDelete(data []entity.SystemParamEntity) map[string]string {

	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("key_code = ?", data[i].Key_code).
			Table("dp_ms_system_param").Count(&count).Update(entity.SystemParamEntity{Active: "0"})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataDel(len(data), deleted)

}
