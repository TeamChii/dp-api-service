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
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service/authentication"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func MerchantById(id string) map[string]interface{} {

	var data model.MerchantEntityResp
	check := config.DBsql().Where("mc_id = ?", id).Find(&data).RecordNotFound()
	var dataPackage entity.PackageEntity
	config.DBsql().Where("package_id = ?", data.Package_id).Find(&dataPackage)
	var dataContainer entity.ContainerEntity
	var countCon int
	config.DBsql().Where("mc_id = ?", id).Find(&dataContainer).Count(&countCon)
	data.PackageEntity = &dataPackage
	data.Count_container = countCon
	var MerchantImageMapResp []model.MerchantImageMapResp

	config.DBsql().
		Table("dp_mp_merchant_image").Select("dp_tb_content.content_id, dp_tb_content.content_path").
		Joins("JOIN dp_tb_content ON dp_mp_merchant_image.content_id = dp_tb_content.content_id").
		Where("dp_mp_merchant_image.mc_id = ? ", data.Mc_id).
		Find(&MerchantImageMapResp)

	data.MerchantImageMapResp = MerchantImageMapResp

	//data.Image_ref = utils.CONTENT_URL+"/" + data.Image_ref + ""
	//data.Logo_ref = utils.CONTENT_URL+"/" + data.Logo_ref + ""
	data.Logo_ref = utils.CONTENT_URL + "/" + data.Logo_ref + ""
	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func MerchantByPhone(userid string) map[string]interface{} {
	var data []model.UserMerchantMapEntityResp

	err := config.DBsql().
		Where("user_id = ? ", userid).
		Order("mc_id desc").
		Find(&data).Error

	for index := 0; index < len(data); index++ {
		var MerchantEntity model.MerchantEntityResp
		config.DBsql().
			Where("mc_id = ? ", data[index].Mc_id).
			Find(&MerchantEntity)

		var MerchantImageMapResp []model.MerchantImageMapResp

		config.DBsql().
			Table("dp_mp_merchant_image").Select("dp_tb_content.content_id, dp_tb_content.content_path").
			Joins("JOIN dp_tb_content ON dp_mp_merchant_image.content_id = dp_tb_content.content_id").
			Where("dp_mp_merchant_image.mc_id = ? ", MerchantEntity.Mc_id).
			Find(&MerchantImageMapResp)

		MerchantEntity.MerchantImageMapResp = MerchantImageMapResp

		//MerchantEntity.Image_ref = utils.CONTENT_URL+"/" + MerchantEntity.Image_ref + ""
		MerchantEntity.Logo_ref = utils.CONTENT_URL + "/" + MerchantEntity.Logo_ref + ""

		dataMer := MerchantEntity
		data[index].MerchantEntityResp = &dataMer

		//data[index].MerchantEntityResp = &dataMer

	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func AddMerchant(merchantMaster model.MerchantEntityReq, ctx iris.Context) map[string]interface{} {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	merchantMaster.Create_by = utils.Decode(jwt.Raw).(string)
	now := time.Now()
	merchantMaster.Create_date = &now
	var UserMerchantMapEntity2 []model.UserMerchantMapEntityResp

	err := config.DBsql().Create(&merchantMaster).Error
	if err == nil {
		var UserEntity2 entity.UserEntity
		check := config.DBsql().Where("user_phone = ?", merchantMaster.Mc_phone).Find(&UserEntity2).RecordNotFound()
		if check == true {
			var UserEntity entity.UserEntity
			UserEntity.User_name = "Unnamed User"
			UserEntity.User_first_name = "Unnamed User"
			UserEntity.User_last_name = ""
			UserEntity.Pin = utils.HashPin(merchantMaster.Pin)
			UserEntity.User_phone = merchantMaster.Mc_phone
			UserEntity.User_email = merchantMaster.Mc_email
			err = config.DBsql().Create(&UserEntity).Error

			var UserMerchantMapEntity entity.UserMerchantMapEntity
			UserMerchantMapEntity.User_id = UserEntity.User_id
			UserMerchantMapEntity.Mc_id = merchantMaster.Mc_id
			UserMerchantMapEntity.Role_id = merchantMaster.Role_id

			err = config.DBsql().Create(&UserMerchantMapEntity).Error

			config.DBsql().
				Where("user_id = ? ", UserEntity.User_id).
				Find(&UserMerchantMapEntity2)
		} else {

			var UserMerchantMapEntity entity.UserMerchantMapEntity
			UserMerchantMapEntity.User_id = UserEntity2.User_id
			UserMerchantMapEntity.Mc_id = merchantMaster.Mc_id
			UserMerchantMapEntity.Role_id = merchantMaster.Role_id

			err = config.DBsql().Create(&UserMerchantMapEntity).Error

			config.DBsql().
				Where("user_id = ? ", UserEntity2.User_id).
				Order("mc_id desc").
				Find(&UserMerchantMapEntity2)

		}

		for index := 0; index < len(UserMerchantMapEntity2); index++ {
			var MerchantEntity model.MerchantEntityResp
			config.DBsql().
				Where("mc_id = ? ", UserMerchantMapEntity2[index].Mc_id).Find(&MerchantEntity)

			//MerchantEntity.Image_ref = utils.CONTENT_URL+"/" + MerchantEntity.Image_ref + ""
			MerchantEntity.Logo_ref = utils.CONTENT_URL + "/" + MerchantEntity.Logo_ref + ""
			var MerchantImageMapResp []model.MerchantImageMapResp

			config.DBsql().
				Table("dp_mp_merchant_image").Select("dp_tb_content.content_id, dp_tb_content.content_path").
				Joins("JOIN dp_tb_content ON dp_mp_merchant_image.content_id = dp_tb_content.content_id").
				Where("dp_mp_merchant_image.mc_id = ? ", MerchantEntity.Mc_id).
				Find(&MerchantImageMapResp)

			MerchantEntity.MerchantImageMapResp = MerchantImageMapResp

			dataMer := MerchantEntity
			UserMerchantMapEntity2[index].MerchantEntityResp = &dataMer

		}

	}
	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(UserMerchantMapEntity2)
	json.Unmarshal(b, &dataToJson)

	tokenString := authentication.GenToken(merchantMaster.Mc_phone, merchantMaster.Pin)

	return utils.ResDataAddMerchant(err, dataToJson, tokenString)

}
func AddMerchant2(merchantMaster model.MerchantEntityReq) map[string]interface{} {
	now := time.Now()
	merchantMaster.Create_date = &now
	var UserMerchantMapEntity2 []model.UserMerchantMapEntityResp

	err := config.DBsql().Create(&merchantMaster).Error
	if err == nil {
		var UserEntity2 entity.UserEntity
		check := config.DBsql().Where("user_phone = ?", merchantMaster.Mc_phone).Find(&UserEntity2).RecordNotFound()
		if check == true {
			var UserEntity entity.UserEntity
			UserEntity.User_name = "Unnamed User"
			UserEntity.User_first_name = "Unnamed User"
			UserEntity.User_last_name = ""
			UserEntity.Pin = utils.HashPin(merchantMaster.Pin)
			UserEntity.User_phone = merchantMaster.Mc_phone
			UserEntity.User_email = merchantMaster.Mc_email
			err = config.DBsql().Create(&UserEntity).Error

			var UserMerchantMapEntity entity.UserMerchantMapEntity
			UserMerchantMapEntity.User_id = UserEntity.User_id
			UserMerchantMapEntity.Mc_id = merchantMaster.Mc_id
			UserMerchantMapEntity.Role_id = merchantMaster.Role_id

			err = config.DBsql().Create(&UserMerchantMapEntity).Error

			config.DBsql().
				Where("user_id = ? ", UserEntity.User_id).
				Find(&UserMerchantMapEntity2)
		} else {

			var UserMerchantMapEntity entity.UserMerchantMapEntity
			UserMerchantMapEntity.User_id = UserEntity2.User_id
			UserMerchantMapEntity.Mc_id = merchantMaster.Mc_id
			UserMerchantMapEntity.Role_id = merchantMaster.Role_id

			err = config.DBsql().Create(&UserMerchantMapEntity).Error

			config.DBsql().
				Where("user_id = ? ", UserEntity2.User_id).
				Order("mc_id desc").
				Find(&UserMerchantMapEntity2)

		}

		err = config.DBsql().Model(&model.MerchantEntityReq{}).Where("mc_id = ? ",
			merchantMaster.Mc_id).Update(
			map[string]interface{}{"mc_group_id": merchantMaster.Mc_id}).Error

		for index := 0; index < len(UserMerchantMapEntity2); index++ {
			var MerchantEntity model.MerchantEntityResp
			config.DBsql().
				Where("mc_id = ? ", UserMerchantMapEntity2[index].Mc_id).Find(&MerchantEntity)

			//MerchantEntity.Image_ref = utils.CONTENT_URL+"/" + MerchantEntity.Image_ref + ""
			MerchantEntity.Logo_ref = utils.CONTENT_URL + "/" + MerchantEntity.Logo_ref + ""
			var MerchantImageMapResp []model.MerchantImageMapResp

			config.DBsql().
				Table("dp_mp_merchant_image").Select("dp_tb_content.content_id, dp_tb_content.content_path").
				Joins("JOIN dp_tb_content ON dp_mp_merchant_image.content_id = dp_tb_content.content_id").
				Where("dp_mp_merchant_image.mc_id = ? ", MerchantEntity.Mc_id).
				Find(&MerchantImageMapResp)

			MerchantEntity.MerchantImageMapResp = MerchantImageMapResp

			dataMer := MerchantEntity
			UserMerchantMapEntity2[index].MerchantEntityResp = &dataMer

		}

	}
	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(UserMerchantMapEntity2)
	json.Unmarshal(b, &dataToJson)

	tokenString := authentication.GenToken(merchantMaster.Mc_phone, merchantMaster.Pin)

	return utils.ResDataAddMerchant(err, dataToJson, tokenString)

}

func UpdateMerchant(merchantMaster entity.MerchantEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	merchantMaster.Update_by = utils.Decode(jwt.Raw).(string)
	now := time.Now()
	merchantMaster.Update_date = &now
	err := config.DBsql().Model(&merchantMaster).Where("mc_id = ?", merchantMaster.Mc_id).Update(&merchantMaster).Error

	return utils.ResDataEdit(err)
}

func SetHeadOfficeMerchant(data entity.MerchantEntity) map[string]string {

	conn := config.DBsql()
	mc_group_id := data.Mc_group_id
	mc_id := data.Mc_id

	conn.Where("mc_id = ?", mc_group_id).
		Table("dp_ms_merchant").Update(entity.MerchantEntity{Mc_is_head_office: "N"})

	conn.Where("mc_group_id = ?", mc_group_id).
		Table("dp_ms_merchant").Update(entity.MerchantEntity{Mc_group_id: &mc_id, Mc_is_head_office: "N"})

	err := conn.Where("mc_id = ?", mc_id).
		Table("dp_ms_merchant").Update(entity.MerchantEntity{Mc_is_head_office: "Y"}).Error

	return utils.ResDataEdit(err)

}
