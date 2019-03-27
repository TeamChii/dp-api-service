package service

import (
	"encoding/json"
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

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func PackageLoad(PackageReq model.PackageReq) map[string]interface{} {

	var order = PackageReq.OrderBy + " " + PackageReq.SortBy

	var count int
	var data []entity.PackageEntity
	conn := config.DBsql()
	//conn.Find(&data).Where("category_name = ?", searchObject.Category).Count(&count) // count record
	err := conn.Order(order).Find(&data).Count(&count).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func PackageAdd(PackageEntity entity.PackageEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)

	now := time.Now()
	PackageEntity.Create_by = utils.Decode(jwt.Raw).(string)
	PackageEntity.Create_date = &now
	err := config.DBsql().Create(&PackageEntity).Error

	return utils.ResDataAdd(err)

}

func PackageUpdate(PackageEntity entity.PackageEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	now := time.Now()
	PackageEntity.Update_by = utils.Decode(jwt.Raw).(string)
	PackageEntity.Update_date = &now

	err := config.DBsql().Model(&PackageEntity).Where("package_id = ?", PackageEntity.Package_id).Update(&PackageEntity).Error

	return utils.ResDataEdit(err)
}
func PackageDelete(data []entity.PackageEntity) map[string]string {

	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("package_id = ?", data[i].Package_id).
			Table("dp_ms_package").Count(&count).Delete(entity.PackageEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataDel(len(data), deleted)

}
