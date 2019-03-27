package service

import (
	"encoding/json"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetMiniPOSSettingById(id string) map[string]interface{} {
	var data entity.MiniposSettingEntity
	check := config.DBsql().Where("mc_id = ?", id).First(&data).RecordNotFound()

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func AddMiniPOSSetting(MiniposSettingEntity entity.MiniposSettingEntity, ctx iris.Context) map[string]string {
	/*
		jwt := ctx.Values().Get("jwt").(*jwt.Token)
		createBy := utils.Decode(jwt.Raw).(string)
		now := time.Now()
		MiniposSettingEntity.Create_by = createBy
		MiniposMenuCategoryEntity.Create_date = &now;
	*/
	err := config.DBsql().Create(&MiniposSettingEntity).Error
	return utils.ResDataAdd(err)
}
func UpdateMiniPOSSetting(MiniposSettingEntity entity.MiniposSettingEntity, ctx iris.Context) map[string]string {
	/*
		jwt := ctx.Values().Get("jwt").(*jwt.Token)
		now := time.Now()
		MiniposSettingEntity.Update_by = utils.Decode(jwt.Raw).(string)
		MiniposSettingEntity.Update_date = &now
	*/
	err := config.DBsql().Model(&MiniposSettingEntity).Where("minipos_setting_id = ? AND mc_id = ? ",
		MiniposSettingEntity.Minipos_Setting_Id, MiniposSettingEntity.Mc_id).Update(&MiniposSettingEntity).Error

	return utils.ResDataEdit(err)
}
