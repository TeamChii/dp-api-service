package service

import (
	"encoding/json"
	"fmt"
	"time"

	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func LoadMerchantCampaignMap(MerchantCampaignMapReq model.MerchantCampaignMapReq) map[string]interface{} {

	var data []model.MerchantCampaignMapEntityResp
	var count int
	var err error
	if MerchantCampaignMapReq.Category_type == "" {
		err = config.DBsql().Where("campaign_status IS NULL AND mc_id = ?", MerchantCampaignMapReq.Mc_id).Find(&data).
			Offset(((MerchantCampaignMapReq.Paging.PageNo * int(MerchantCampaignMapReq.Paging.PageSize)) - int(MerchantCampaignMapReq.Paging.PageSize))).
			Limit(MerchantCampaignMapReq.Paging.PageSize).
			Order(MerchantCampaignMapReq.Paging.OrderBy + " " + MerchantCampaignMapReq.Paging.SortBy).
			Count(&count).Error
	} else {
		err = config.DBsql().Where("category_type = ? AND campaign_status IS NULL AND mc_id = ?", MerchantCampaignMapReq.Category_type, MerchantCampaignMapReq.Mc_id).Find(&data).
			Offset(((MerchantCampaignMapReq.Paging.PageNo * int(MerchantCampaignMapReq.Paging.PageSize)) - int(MerchantCampaignMapReq.Paging.PageSize))).
			Limit(MerchantCampaignMapReq.Paging.PageSize).
			Order(MerchantCampaignMapReq.Paging.OrderBy + " " + MerchantCampaignMapReq.Paging.SortBy).
			Count(&count).Error
	}

	var data2 []model.MerchantCampaignMapEntityResp
	for index := 0; index < len(data); index++ {
		var CampaignEntity entity.CampaignEntity
		config.DBsql().Where("campaign_id = ?", data[index].Campaign_id).Find(&CampaignEntity)
		camData := CampaignEntity
		data[index].CampaignEntity = &camData

		var customer_count int
		var MerchantCustomerCampaignMapEntity entity.MerchantCustomerCampaignMapEntity
		if MerchantCampaignMapReq.Category_type == "" {
			config.DBsql().Where("mc_id = ? AND campaign_id = ? AND send_status IS NULL",
				MerchantCampaignMapReq.Mc_id, data[index].Campaign_id).Find(&MerchantCustomerCampaignMapEntity).
				Count(&customer_count)
		} else {
			config.DBsql().Where("category_type = ? AND mc_id = ? AND campaign_id = ? AND send_status IS NULL",
				data[index].Category_type, MerchantCampaignMapReq.Mc_id, data[index].Campaign_id).Find(&MerchantCustomerCampaignMapEntity).
				Count(&customer_count)
		}

		data[index].Customer_count = customer_count
		if data[index].Customer_count > 0 {
			data2 = append(data2, data[index])
		}

	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data2)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, MerchantCampaignMapReq.Paging, count)
}
func SendCardPromotion(MerchantCustomerCampaignMapAddReq model.MerchantCustomerCampaignMapAddReq, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//err := config.DBsql().Create(&RequestPointAddReq).Error
	now := time.Now()
	createdBy := utils.Decode(jwt.Raw).(string)
	var err error

	var data []entity.MerchantCustomerCampaignMapEntity
	if MerchantCustomerCampaignMapAddReq.Category_type == "" {
		config.DBsql().Where("mc_id = ? AND send_status IS NULL AND campaign_id = ?",
			MerchantCustomerCampaignMapAddReq.Mc_id, MerchantCustomerCampaignMapAddReq.Campaign_id).Find(&data)
	} else {
		config.DBsql().Where("mc_id = ? AND send_status IS NULL AND category_type = ? AND campaign_id = ?",
			MerchantCustomerCampaignMapAddReq.Mc_id, MerchantCustomerCampaignMapAddReq.Category_type, MerchantCustomerCampaignMapAddReq.Campaign_id).Find(&data)
	}
	for index := 0; index < len(data); index++ {

		for indexj := 0; indexj < len(MerchantCustomerCampaignMapAddReq.ContainerDetail); indexj++ {
			var MerchantCustomerMapMaster entity.MerchantCustomerMapEntity
			MerchantCustomerMapMaster.Mc_id = MerchantCustomerCampaignMapAddReq.Mc_id
			MerchantCustomerMapMaster.Cust_id = data[index].Cust_id
			MerchantCustomerMapMaster.Container_id = MerchantCustomerCampaignMapAddReq.ContainerDetail[indexj].Container_id

			if MerchantCustomerCampaignMapAddReq.ContainerDetail[indexj].Container_type == "CC" {
				date := utils.GenIssueExpireDate(MerchantCustomerCampaignMapAddReq.ContainerDetail[indexj].Expire_value, MerchantCustomerCampaignMapAddReq.ContainerDetail[indexj].Expire_mode)
				issue_date := date["issue_date"].(*time.Time)
				expire_date := date["expire_date"].(*time.Time)

				MerchantCustomerMapMaster.Issue_date = issue_date
				MerchantCustomerMapMaster.Expire_date = expire_date
			} else {
				MerchantCustomerMapMaster.Issue_date = nil
				MerchantCustomerMapMaster.Expire_date = nil
			}

			//MerchantCustomerMapMaster.Cust_tag = RequestPointAddReq.ContainerDetail[indexj].Cust_tag
			//MerchantCustomerMapMaster.Cust_frg = RequestPointAddReq.ContainerDetail[indexj].Cust_frg
			//MerchantCustomerMapMaster.Cust_status = RequestPointAddReq.ContainerDetail[indexj].Cust_status
			//MerchantCustomerMapMaster.Cust_first_visit_date = RequestPointAddReq.ContainerDetail[indexj].Cust_first_visit_date
			//MerchantCustomerMapMaster.Cust_last_visit_date = RequestPointAddReq.ContainerDetail[indexj].Cust_last_visit_date
			MerchantCustomerMapMaster.Create_by = createdBy
			MerchantCustomerMapMaster.Create_date = &now

			err = config.DBsql().Create(&MerchantCustomerMapMaster).Error
		}

		if err == nil {
			var updateStatus entity.MerchantCustomerCampaignMapEntity
			config.DBsql().Model(&updateStatus).Where("mc_id = ? AND campaign_id = ? AND cust_id = ? AND category_type = ?",
				data[index].Mc_id, data[index].Campaign_id, data[index].Cust_id, data[index].Category_type).
				Update(entity.MerchantCustomerCampaignMapEntity{Send_status: "Y", Message_detail: MerchantCustomerCampaignMapAddReq.Message_detail, Send_date: &now})

		}

	}
	if err == nil {
		var MerchantCampaignMapEntity entity.MerchantCampaignMapEntity
		config.DBsql().Model(&MerchantCampaignMapEntity).Where("mc_id = ? AND campaign_id = ?",
			MerchantCustomerCampaignMapAddReq.Mc_id, MerchantCustomerCampaignMapAddReq.Campaign_id).Update("campaign_status", "Y")
	}

	return utils.ResDataAdd(err)

}
func LoadMerchantCampaignMapById(MerchantCampaignMapReq2 model.MerchantCampaignMapReq2) map[string]interface{} {

	var data []model.MerchantCampaignMapEntityResp

	err := config.DBsql().Where("mc_id = ? AND campaign_id = ?", MerchantCampaignMapReq2.Mc_id, MerchantCampaignMapReq2.Campaign_id).Find(&data).Error

	fmt.Println(data)

	var CampaignEntity entity.CampaignEntity
	config.DBsql().Where("campaign_id = ?", MerchantCampaignMapReq2.Campaign_id).Find(&CampaignEntity)
	camData := CampaignEntity
	data[0].CampaignEntity = &camData

	var customer_count int
	var MerchantCustomerCampaignMapEntity entity.MerchantCustomerCampaignMapEntity

	if MerchantCampaignMapReq2.Category_type == "" {
		config.DBsql().Where("mc_id = ? AND campaign_id = ? AND send_status != ?",
			MerchantCampaignMapReq2.Mc_id, MerchantCampaignMapReq2.Campaign_id, "Y").Find(&MerchantCustomerCampaignMapEntity).
			Count(&customer_count)
	} else {
		config.DBsql().Where("category_type = ? AND mc_id = ? AND campaign_id = ? AND send_status != ?",
			MerchantCampaignMapReq2.Category_type, MerchantCampaignMapReq2.Mc_id, MerchantCampaignMapReq2.Campaign_id, "Y").Find(&MerchantCustomerCampaignMapEntity).
			Count(&customer_count)
	}
	data[0].Customer_count = customer_count

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}

func SetFavourite(MerchantCampaignMapUpadteReq model.MerchantCampaignMapUpadteReq) map[string]string {

	var err error

	var MerchantCampaignMapEntity entity.MerchantCampaignMapEntity
	err = config.DBsql().Model(&MerchantCampaignMapEntity).Where("mc_id = ? AND campaign_id = ?", MerchantCampaignMapUpadteReq.Mc_id, MerchantCampaignMapUpadteReq.Campaign_id).
		Update("category_type", "FV").Error

	var MerchantCustomerCampaignMapEntity entity.MerchantCustomerCampaignMapEntity
	err = config.DBsql().Model(&MerchantCustomerCampaignMapEntity).Where("mc_id = ? AND campaign_id = ?", MerchantCampaignMapUpadteReq.Mc_id, MerchantCampaignMapUpadteReq.Campaign_id).
		Update("category_type", "FV").Error

	return utils.ResDataEdit(err)
}
func UnSetFavourite(MerchantCampaignMapUpadteReq model.MerchantCampaignMapUpadteReq) map[string]string {

	var err error

	var MerchantCampaignMapEntity entity.MerchantCampaignMapEntity
	err = config.DBsql().Model(&MerchantCampaignMapEntity).Where("mc_id = ? AND campaign_id = ?", MerchantCampaignMapUpadteReq.Mc_id, MerchantCampaignMapUpadteReq.Campaign_id).
		Update("category_type", nil).Error

	var MerchantCustomerCampaignMapEntity entity.MerchantCustomerCampaignMapEntity
	err = config.DBsql().Model(&MerchantCustomerCampaignMapEntity).Where("mc_id = ? AND campaign_id = ?", MerchantCampaignMapUpadteReq.Mc_id, MerchantCampaignMapUpadteReq.Campaign_id).
		Update("category_type", nil).Error

	return utils.ResDataEdit(err)
}
func LoadCustoemrInCampaignMap(MerchantCampaignMapReq model.MerchantCustomerCampaignMapLoadCustReq) map[string]interface{} {

	var data []entity.MerchantCustomerCampaignMapEntity
	var count int

	err := config.DBsql().Where("mc_id = ? AND campaign_id = ?", MerchantCampaignMapReq.Mc_id, MerchantCampaignMapReq.Campaign_id).Find(&data).
		Offset(((MerchantCampaignMapReq.Paging.PageNo * int(MerchantCampaignMapReq.Paging.PageSize)) - int(MerchantCampaignMapReq.Paging.PageSize))).
		Limit(MerchantCampaignMapReq.Paging.PageSize).
		Order(MerchantCampaignMapReq.Paging.OrderBy + " " + MerchantCampaignMapReq.Paging.SortBy).
		Count(&count).Error

	for index := 0; index < len(data); index++ {
		var CustomerEntity entity.CustomerEntity
		config.DBsql().Where("cust_id = ?", data[index].Cust_id).Find(&CustomerEntity)
		CustomerEntity.Image_ref = utils.CONTENT_URL + "/" + CustomerEntity.Image_ref + ""

		data[index].CustomerEntity = &CustomerEntity
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, MerchantCampaignMapReq.Paging, count)
}
