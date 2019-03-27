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

func NotiLoad(NotiCategoryReq model.NotiCategoryReq) map[string]interface{} {

	var data []entity.NotificationEntity

	err := config.DBsql().Where("mc_id = ? AND noti_category = ? AND noti_flag != 'D' ", NotiCategoryReq.Mc_id, NotiCategoryReq.Noti_category).Find(&data).Error

	for index := 0; index < len(data); index++ {
		if data[index].Noti_send_time.Year() != 1 {
			notidate := data[index].Noti_send_time.Format("02/01/2006")
			data[index].Noti_send_time_str = &notidate
		}
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func NotiCategory(id string) map[string]interface{} {

	type DataResp struct {
		Key_code  string `json:"key_code"`
		Key_value string `json:"key_value"`
		Count     int    `json:"count"`
	}
	var dataresp []DataResp
	var err error

	err = config.DBsql().Raw(`select 
	sysnp.key_code, sysnp.key_value,
	(select count(*) from dp_tb_notification as noti
			 where noti.noti_category = sysnp.key_code and noti.mc_id = ? AND noti.noti_flag = 'S'
	) as count
	from dp_ms_system_param as sysnp
	where category_name = 'NOTIFICATION' order by sysnp.ord`, id).
		Scan(&dataresp).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(dataresp)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func NotiById(id string) map[string]interface{} {

	var data entity.NotificationEntity
	check := config.DBsql().Where("noti_id = ?", id).Find(&data).RecordNotFound()

	if data.Noti_send_time.Year() != 1 {
		notidate := data.Noti_send_time.Format("02/01/2006")
		data.Noti_send_time_str = &notidate
	}

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func NotiRead(NotificationEntity entity.NotificationEntity) map[string]string {

	var data entity.NotificationEntity

	data.Noti_flag = "R"

	err := config.DBsql().Model(&data).Where("noti_id = ?", NotificationEntity.Noti_id).Update(&data).Error

	return utils.ResDataEdit(err)
}
func NotiDelete(NotificationEntity entity.NotificationEntity) map[string]string {

	var data entity.NotificationEntity

	data.Noti_flag = "D"

	err := config.DBsql().Model(&data).Where("noti_id = ?", NotificationEntity.Noti_id).Update(&data).Error

	return utils.ResDataEdit(err)
}
