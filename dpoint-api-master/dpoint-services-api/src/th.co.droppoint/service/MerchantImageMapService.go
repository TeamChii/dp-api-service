package service

import (
	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/utils"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func AddMerchantImage(UserMaster entity.MerchantImageMapEntity) map[string]string {

	err := config.DBsql().Create(&UserMaster).Error

	return utils.ResDataAdd(err)

}
func UpdateMerchantImage(UserMaster entity.MerchantImageMapEntity) map[string]string {

	err := config.DBsql().Model(&UserMaster).Where("mc_id = ? AND content_id = ?", UserMaster.Mc_id, UserMaster.Content_id).Update(&UserMaster).Error

	return utils.ResDataEdit(err)
}
func DeleteMerchantImage(data entity.MerchantImageMapEntity) map[string]string {

	conn := config.DBsql()
	err := conn.Where("mc_id = ? AND content_id = ?", data.Mc_id, data.Content_id).
		Table("dp_mp_merchant_image").Delete(entity.PackageEntity{}).Error

	if err != nil {
		return map[string]string{
			"statusCode":  "E",
			"messageCode": "E001",
			"messageAbbr": "Error",
			"messageDesc": "Delete Image Fail",
		}
	} else {
		return map[string]string{
			"statusCode":  "S",
			"messageCode": "S001",
			"messageAbbr": "Success",
			"messageDesc": "Delete Image Success",
		}
	}

}
