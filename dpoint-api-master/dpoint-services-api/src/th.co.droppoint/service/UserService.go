package service

import (
	"encoding/json"
	"math/rand"
	"strconv"

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

func UserById(id string) map[string]interface{} {

	var data entity.UserEntity
	check := config.DBsql().Where("user_id = ?", id).Find(&data).RecordNotFound()
	data.Pin = ""
	data.Image_ref = utils.CONTENT_URL + "/" + data.Image_ref + ""
	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func UserByRole(UserLoadRoleReq model.UserLoadRoleReq) map[string]interface{} {

	var data []model.UserMerchantMapEntityResp2
	err := config.DBsql().Raw(`SELECT *
	FROM "dp_mp_user_merchant"
	JOIN dp_ms_user
	 ON dp_mp_user_merchant.user_id = dp_ms_user.user_id
	WHERE dp_mp_user_merchant.role_id = ? AND dp_mp_user_merchant.mc_id = ?`, UserLoadRoleReq.Role_id, UserLoadRoleReq.Mc_id).
		Scan(&data).Error
	for index := 0; index < len(data); index++ {
		data[index].Image_ref = utils.CONTENT_URL + "/" + data[index].Image_ref + ""
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func UserLoadCategory(UserLoadRoleReq model.UserLoadRoleReq) map[string]interface{} {

	type CategoryRole struct {
		Key_code  string `json:"key_code"`
		Key_value string `json:"key_value"`
		Count     int    `json:"count"`
	}

	var data []CategoryRole

	err := config.DBsql().Raw(`select 
	sysnp.key_code, sysnp.key_value,
	(select count(*) from dp_mp_user_merchant as usermap
			 where usermap.role_id = sysnp.key_code and usermap.mc_id = ?
	) as count
	from dp_ms_system_param as sysnp
	where category_name = 'ROLE' order by sysnp.ord`, UserLoadRoleReq.Mc_id).
		Scan(&data).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}

func AddUser(UserAddReq model.UserAddReq) map[string]string {
	var err2 error
	var UserEntity entity.UserEntity
	UserEntity.User_name = "Unnamed User"
	UserEntity.User_first_name = "Unnamed User"
	UserEntity.User_last_name = ""
	UserEntity.User_phone = UserAddReq.User_Phone
	err := config.DBsql().Create(&UserEntity).Error
	if err == nil {
		var UserMerchantMapEntity entity.UserMerchantMapEntity
		UserMerchantMapEntity.User_id = UserEntity.User_id
		UserMerchantMapEntity.Mc_id = UserAddReq.Mc_id
		UserMerchantMapEntity.Role_id = UserAddReq.Role_id

		err2 = config.DBsql().Create(&UserMerchantMapEntity).Error

	}
	return utils.ResDataAdd(err2)

}
func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
func AddUserStaff(UserAddReq model.UserAddReq) map[string]string {
	var err2 error

	pin := strconv.Itoa(rangeIn(1000, 9999))

	var UserEntity entity.UserEntity
	UserEntity.User_name = UserAddReq.User_name
	UserEntity.User_first_name = UserAddReq.User_name // "Unnamed User"
	UserEntity.User_last_name = ""
	UserEntity.Staff_pin = pin
	err := config.DBsql().Create(&UserEntity).Error
	if err == nil {
		var UserMerchantMapEntity entity.UserMerchantMapEntity
		UserMerchantMapEntity.User_id = UserEntity.User_id
		UserMerchantMapEntity.Mc_id = UserAddReq.Mc_id
		UserMerchantMapEntity.Role_id = UserAddReq.Role_id

		err2 = config.DBsql().Create(&UserMerchantMapEntity).Error

	}
	return utils.ResDataAdd(err2)

}
func UpdateUser(UserMaster entity.UserEntity) map[string]string {

	err := config.DBsql().Model(&UserMaster).Where("user_id = ?", UserMaster.User_id).Update(&UserMaster).Error

	return utils.ResDataEdit(err)
}
func UpdateUserPin(PinReq model.PinReq) map[string]string {

	var UserMaster entity.UserEntity
	UserMaster.Pin = utils.HashPin(PinReq.Pin)
	err := config.DBsql().Model(&UserMaster).Where("user_id = ?", PinReq.User_id).Update("pin", utils.HashPin(PinReq.Pin)).Error

	return utils.ResDataEdit(err)
}
func DeleteUser(data []entity.UserEntity) map[string]string {

	var count int
	var deleted = 0
	conn := config.DBsql()
	for i := 0; i < len(data); i++ {
		conn.Where("user_id = ?", data[i].User_id).
			Table("dp_ms_user").Count(&count).Delete(entity.UserEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
	}

	return utils.ResDataDel(len(data), deleted)

}
func UserCheck(phone string) map[string]interface{} {

	var data entity.UserEntity
	checkPhone := config.DBsql().Where("user_phone = ?", phone).Find(&data).RecordNotFound()
	checkPin := config.DBsql().Where("user_phone = ? AND pin = '' ", phone).Find(&data).RecordNotFound()
	if checkPhone == false {
		return map[string]interface{}{
			"statusCode":   "S",
			"messageCode":  "400",
			"status_phone": !checkPhone,
			"status_pin":   checkPin,
			"user_id":      data.User_id,
			"messageDesc":  "Found User",
		}
	} else {
		return map[string]interface{}{
			"statusCode":   "E",
			"messageCode":  "400",
			"status_phone": false,
			"status_pin":   false,
			"user_id":      data.User_id,
			"messageDesc":  "Not Found User",
		}
	}

}
