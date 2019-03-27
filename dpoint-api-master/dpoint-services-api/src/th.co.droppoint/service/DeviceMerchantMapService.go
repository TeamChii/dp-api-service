package service

import (
	"encoding/json"
	"time"

	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"

		"th.co.droppoint/utils"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func LoadDevice(id string) map[string]interface{} {
	var data []entity.DeviceMerchantMapEntity
	err := config.DBsql().
		Where(`mc_id = ?`, id).
		Find(&data).Error

	for index := 0; index < len(data); index++ {
		if data[index].Create_date.Year() != 1 {
			notidate := data[index].Create_date.Format("02/01/2006")
			data[index].Create_date_str = notidate
		}
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func CheckUidDevice(DeviceMerchantMapEntity entity.DeviceMerchantMapEntity) map[string]interface{} {
	var data entity.DeviceMerchantMapEntity
	check := config.DBsql().
		Where(`device_uid = ? AND mc_id = ?`, DeviceMerchantMapEntity.Device_uid, DeviceMerchantMapEntity.Mc_id).
		Find(&data).RecordNotFound()

	if data.Device_uid != "" && data.Create_date.Year() != 1 {
		notidate := data.Create_date.Format("02/01/2006")
		data.Create_date_str = notidate
	}

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	if check == false {
		return map[string]interface{}{
			"statusCode":  "S",
			"status":      true,
			"data":        dataToJson,
			"messageDesc": "Found uid"}
	} else {
		return map[string]interface{}{
			"statusCode":  "E",
			"status":      false,
			"data":        dataToJson,
			"messageDesc": "Not Found uid"}
	}

}
func AddDevice(DeviceMerchantMapEntity entity.DeviceMerchantMapEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	now := time.Now()
	DeviceMerchantMapEntity.Create_by = utils.Decode(jwt.Raw).(string)
	DeviceMerchantMapEntity.Create_date = &now

	err := config.DBsql().Create(&DeviceMerchantMapEntity).Error

	return utils.ResDataAdd(err)

}
func UpdateDevice(DeviceMerchantMapEntity entity.DeviceMerchantMapEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	now := time.Now()
	DeviceMerchantMapEntity.Update_by = utils.Decode(jwt.Raw).(string)
	DeviceMerchantMapEntity.Update_date = &now

	err := config.DBsql().Model(&DeviceMerchantMapEntity).Where("mc_id = ? AND device_uid = ?",
		DeviceMerchantMapEntity.Mc_id, DeviceMerchantMapEntity.Device_uid).Update(&DeviceMerchantMapEntity).Error

	return utils.ResDataEdit(err)

}
func DeleteDevice(data []entity.DeviceMerchantMapEntity) map[string]string {

	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("device_uid = ? AND mc_id = ?", data[i].Device_uid, data[i].Mc_id).
			Table("dp_mp_device_merchant").Count(&count).Delete(entity.DeviceMerchantMapEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataDel(len(data), deleted)

}

/*func UpdateDevice(DeviceMerchantMapEntity entity.DeviceMerchantMapEntity) map[string]string {

	err := config.DBsql().Model(&DeviceMerchantMapEntity).Where("user_id = ?", UserMaster.User_id).Update(&UserMaster).Error

	return utils.ResDataEdit(err)
}*/
