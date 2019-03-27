package service

import (
	"encoding/json"
	/*
		"th.co.droppoint/config"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func LoadCampaign(CampaignReq model.CampaignReq) map[string]interface{} {

	var data []model.RequestPointEntityReq

	var count int
	err := config.DBsql().Find(&data).
		Count(&count).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, CampaignReq.Paging, count)
}
