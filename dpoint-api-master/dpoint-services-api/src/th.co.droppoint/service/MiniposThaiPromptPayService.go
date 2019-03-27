package service

import (
	"encoding/json"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetMiniposThaiPromptPayById(id string) map[string]interface{} {
	var data entity.MiniposThaiPromptPayEntity
	check := config.DBsql().Where("mc_id = ?", id).First(&data).RecordNotFound()

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func AddMiniposThaiPromptPay(MiniposThaiPromptPayEntity entity.MiniposThaiPromptPayEntity, ctx iris.Context) map[string]string {
	/*
		jwt := ctx.Values().Get("jwt").(*jwt.Token)
		createBy := utils.Decode(jwt.Raw).(string)
		now := time.Now()
		MiniposThaiPromptPayEntity.Create_by = createBy
		MiniposMenuCategoryEntity.Create_date = &now;
	*/
	err := config.DBsql().Create(&MiniposThaiPromptPayEntity).Error
	return utils.ResDataAdd(err)
}
func UpdateMiniposThaiPromptPay(MiniposThaiPromptPayEntity entity.MiniposThaiPromptPayEntity, ctx iris.Context) map[string]string {
	/*
		jwt := ctx.Values().Get("jwt").(*jwt.Token)
		now := time.Now()
		MiniposThaiPromptPayEntity.Update_by = utils.Decode(jwt.Raw).(string)
		MiniposThaiPromptPayEntity.Update_date = &now
	*/
	err := config.DBsql().Model(&MiniposThaiPromptPayEntity).Where("Thai_Prompt_Pay_Id = ? AND minipos_setting_id = ? ",
		MiniposThaiPromptPayEntity.Thai_Prompt_Pay_Id, MiniposThaiPromptPayEntity.Minipos_Setting_Id).Update(&MiniposThaiPromptPayEntity).Error

	return utils.ResDataEdit(err)
}
