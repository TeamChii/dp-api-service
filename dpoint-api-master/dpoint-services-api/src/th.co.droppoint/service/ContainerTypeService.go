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

/*
"th.co.droppoint/config"
"th.co.droppoint/entity"
"th.co.droppoint/model"
"th.co.droppoint/utils"
*/

func ContainerTypeSearch(searchObject model.SearchObjectModel) map[string]interface{} {
	//var accountModel model.CustomerModel

	var order = searchObject.Paging.OrderBy + " " + searchObject.Paging.SortBy
	var offset = searchObject.Paging.PageNo*searchObject.Paging.PageSize - searchObject.Paging.PageSize

	var count int
	var account []entity.ContainerTypeEntity
	conn := config.DBsql()
	conn.Find(&account).Count(&count) // count record
	err := conn.
		Offset(offset).
		Limit(searchObject.Paging.PageSize).
		Order(order).Find(&account).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(account)
	json.Unmarshal(b, &dataToJson)
	return utils.ResDataLoad(err, dataToJson, searchObject.Paging, count)
}
func ContainerTypeById(id string) map[string]interface{} {

	var data entity.ContainerTypeEntity
	check := config.DBsql().Where("container_type_id = ?", id).Find(&data).RecordNotFound()

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func AddContainerType(containerTypeMaster entity.ContainerTypeEntity) map[string]string {

	//	var resModel model.ResponseModel
	//id, _ := strconv.Atoi(utils.RandSeq(32))
	//ContainerTypeMaster.Mc_id = id
	//ContainerTypeMaster.Create_date = time.Now()

	err := config.DBsql().Create(&containerTypeMaster).Error
	return utils.ResDataAdd(err)

}

func UpdateContainerType(containerTypeMaster entity.ContainerTypeEntity) map[string]string {

	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	err := config.DBsql().Model(&containerTypeMaster).Where("container_type_id = ?", containerTypeMaster.Container_type_id).Update(&containerTypeMaster).Error

	return utils.ResDataEdit(err)
}
