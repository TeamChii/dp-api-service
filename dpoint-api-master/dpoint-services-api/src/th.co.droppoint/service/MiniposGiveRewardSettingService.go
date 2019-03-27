package service

import (
	"encoding/json"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetGiveRewardSettingById(id string) map[string]interface{} {
	var data entity.MiniposPointRewardEntity

	check := config.DBsql().Where("mc_id = ?", id).First(&data).RecordNotFound()

	var dataRes model.GivePointRewardModel

	var menuItems []entity.MiniposItemRewardEntity
	config.DBsql().Where("mc_id = ?", id).Preload("MiniposMenu").Find(&menuItems)

	dataRes.MiniposPointReward = data
	dataRes.MenuItems = menuItems

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(dataRes)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func AddGiveRewardSetting(GivePointRewardModel model.GivePointRewardModel, ctx iris.Context) map[string]string {
	/*
		jwt := ctx.Values().Get("jwt").(*jwt.Token)
		createBy := utils.Decode(jwt.Raw).(string)
		now := time.Now()
		MiniposPointRewardEntity.Create_by = createBy
		MiniposMenuCategoryEntity.Create_date = &now;
	*/
	menuItems := GivePointRewardModel.MenuItems
	minipos_setting_id := GivePointRewardModel.MiniposPointReward.Minipos_Setting_Id
	reward_item_type := GivePointRewardModel.MiniposPointReward.Reward_Item_Type
	mc_id := GivePointRewardModel.MiniposPointReward.MC_Id

	err := config.DBsql().Create(&GivePointRewardModel.MiniposPointReward).Error
	for index := 0; index < len(menuItems); index++ {
		var miniposItemRewardEntity entity.MiniposItemRewardEntity

		miniposItemRewardEntity.Mpos_Menu_Id = menuItems[index].Mpos_Menu_Id
		miniposItemRewardEntity.Minipos_Setting_Id = minipos_setting_id
		miniposItemRewardEntity.Mc_Id = mc_id
		miniposItemRewardEntity.Reward_Item_Type = reward_item_type
		miniposItemRewardEntity.Ord = index + 1

		config.DBsql().Create(&miniposItemRewardEntity)

	}
	//err := config.DBsql().Create(&MiniposPointRewardEntity).Error
	return utils.ResDataAdd(err)
}
func UpdateGiveRewardSetting(GivePointRewardModel model.GivePointRewardModel, ctx iris.Context) map[string]string {
	/*
		jwt := ctx.Values().Get("jwt").(*jwt.Token)
		now := time.Now()
		MiniposPointRewardEntity.Update_by = utils.Decode(jwt.Raw).(string)
		MiniposPointRewardEntity.Update_date = &now
	*/
	menuItems := GivePointRewardModel.MenuItems
	minipos_setting_id := GivePointRewardModel.MiniposPointReward.Minipos_Setting_Id
	reward_item_type := GivePointRewardModel.MiniposPointReward.Reward_Item_Type

	mc_id := GivePointRewardModel.MiniposPointReward.MC_Id

	//reward_type := GivePointRewardModel.MiniposPointReward.Reward_Type

	miniposPointReward := GivePointRewardModel.MiniposPointReward

	err := config.DBsql().Model(&miniposPointReward).Where("minipos_point_reward_id = ? AND mc_id = ? ",
		GivePointRewardModel.MiniposPointReward.Minipos_Point_reward_Id,
		GivePointRewardModel.MiniposPointReward.MC_Id).Update(&miniposPointReward).Error

	//db.Model(&user).Updates(User{Name: "hello", Age: 18})
	err = config.DBsql().Model(&miniposPointReward).Where("minipos_point_reward_id = ? AND mc_id = ? ",
		GivePointRewardModel.MiniposPointReward.Minipos_Point_reward_Id,
		GivePointRewardModel.MiniposPointReward.MC_Id).Update(
		map[string]interface{}{"reward_amount_amt": miniposPointReward.Reward_Amount_Amt,
			"reward_amount_point": miniposPointReward.Reward_Amount_Point}).Error
	config.DBsql().Where("mc_id = ?", mc_id).Delete(&entity.MiniposItemRewardEntity{})
	//if( reward_type =="I"){
	// delete old item
	for index := 0; index < len(menuItems); index++ {
		var miniposItemRewardEntity entity.MiniposItemRewardEntity

		miniposItemRewardEntity.Mpos_Menu_Id = menuItems[index].Mpos_Menu_Id
		miniposItemRewardEntity.Minipos_Setting_Id = minipos_setting_id
		miniposItemRewardEntity.Mc_Id = mc_id
		miniposItemRewardEntity.Reward_Item_Type = reward_item_type
		miniposItemRewardEntity.Ord = index + 1

		config.DBsql().Create(&miniposItemRewardEntity)

	}
	/*
		}else{

		}
	*/

	return utils.ResDataEdit(err)
}
