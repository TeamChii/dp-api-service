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
		"th.co.droppoint/service/authentication"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func LoadRequestPoint(requestpointMaster model.RequestPointReq) map[string]interface{} {

	var data []model.RequestPointEntityReq

	var count int
	var err error
	if requestpointMaster.Category_type == "" {
		err = config.DBsql().Where("mc_id = ? AND reqeust_type = ? AND request_status IS NULL", requestpointMaster.Mc_id, requestpointMaster.Reqeust_type).Find(&data).
			Count(&count).Error
	} else {
		fmt.Println(requestpointMaster.Reqeust_type)
		err = config.DBsql().
			Where("mc_id = ? AND category_type = ?  AND reqeust_type = ? AND request_status IS NULL", requestpointMaster.Mc_id, requestpointMaster.Category_type, requestpointMaster.Reqeust_type).
			Find(&data).
			Count(&count).Error
	}

	for i := 0; i < len(data); i++ {
		var CustomerEntity entity.CustomerEntity

		config.DBsql().Where("cust_id = ?", data[i].Cust_id).Find(&CustomerEntity)
		CustomerEntity.Image_ref = utils.CONTENT_URL + "/" + CustomerEntity.Image_ref + ""
		data[i].CustomerEntity = &CustomerEntity
		if data[i].Request_date != nil && data[i].Request_date.Year() != 1 {
			data[i].Request_date_str = data[i].Request_date.Format("02/01/2006 15:04:05")

		}

	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, requestpointMaster.Paging, count)
}
func SendCardAll(RequestPointAddReq model.RequestPointAddReq, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//err := config.DBsql().Create(&RequestPointAddReq).Error
	now := time.Now()
	var err error
	/*for index := 0; index < len(RequestPointAddReq.ContainerDetail); index++ {
		MerchantCustomerMapMaster.Mc_id = RequestPointAddReq.Mc_id
		MerchantCustomerMapMaster.Cust_id = RequestPointAddReq.Cust_id
		MerchantCustomerMapMaster.Container_id = RequestPointAddReq.ContainerDetail[index].Container_id
		MerchantCustomerMapMaster.Issue_date = RequestPointAddReq.ContainerDetail[index].Issue_date
		MerchantCustomerMapMaster.Expire_date = RequestPointAddReq.ContainerDetail[index].Expire_date
		MerchantCustomerMapMaster.Cust_tag = RequestPointAddReq.ContainerDetail[index].Cust_tag
		MerchantCustomerMapMaster.Cust_frg = RequestPointAddReq.ContainerDetail[index].Cust_frg
		MerchantCustomerMapMaster.Cust_status = RequestPointAddReq.ContainerDetail[index].Cust_status
		MerchantCustomerMapMaster.Cust_first_visit_date = RequestPointAddReq.ContainerDetail[index].Cust_first_visit_date
		MerchantCustomerMapMaster.Cust_last_visit_date = RequestPointAddReq.ContainerDetail[index].Cust_last_visit_date
		MerchantCustomerMapMaster.Create_by = utils.Decode(jwt.Raw).(string)
		MerchantCustomerMapMaster.Create_date = &now

		config.DBsql().Create(&MerchantCustomerMapMaster)

	}*/
	var data []model.RequestPointEntityReq
	config.DBsql().Where("mc_id = ? AND request_status != ?", RequestPointAddReq.Mc_id, "Y").Find(&data)
	for index := 0; index < len(data); index++ {

		for indexj := 0; indexj < len(RequestPointAddReq.ContainerDetail); indexj++ {
			var MerchantCustomerMapMaster entity.MerchantCustomerMapEntity
			MerchantCustomerMapMaster.Mc_id = RequestPointAddReq.Mc_id
			MerchantCustomerMapMaster.Cust_id = data[index].Cust_id
			MerchantCustomerMapMaster.Container_id = RequestPointAddReq.ContainerDetail[indexj].Container_id

			if RequestPointAddReq.ContainerDetail[indexj].Container_type == "CC" {
				date := utils.GenIssueExpireDate(RequestPointAddReq.ContainerDetail[indexj].Expire_value, RequestPointAddReq.ContainerDetail[indexj].Expire_mode)
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
			MerchantCustomerMapMaster.Create_by = utils.Decode(jwt.Raw).(string)
			MerchantCustomerMapMaster.Create_date = &now

			err = config.DBsql().Create(&MerchantCustomerMapMaster).Error
		}
		if err == nil {
			var updateStatus model.RequestPointEntityReq
			config.DBsql().Model(&updateStatus).Where("request_point_id = ?", data[index].Request_point_id).
				Update(model.RequestPointEntityReq{Request_status: "Y", Request_message: RequestPointAddReq.Request_message})

		}

	}

	return utils.ResDataAdd(err)

}
func SendCard(RequestPointAddReq model.RequestPointAddReq2, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//err := config.DBsql().Create(&RequestPointAddReq).Error
	now := time.Now()
	var err error
	/*for index := 0; index < len(RequestPointAddReq.ContainerDetail); index++ {
		MerchantCustomerMapMaster.Mc_id = RequestPointAddReq.Mc_id
		MerchantCustomerMapMaster.Cust_id = RequestPointAddReq.Cust_id
		MerchantCustomerMapMaster.Container_id = RequestPointAddReq.ContainerDetail[index].Container_id
		MerchantCustomerMapMaster.Issue_date = RequestPointAddReq.ContainerDetail[index].Issue_date
		MerchantCustomerMapMaster.Expire_date = RequestPointAddReq.ContainerDetail[index].Expire_date
		MerchantCustomerMapMaster.Cust_tag = RequestPointAddReq.ContainerDetail[index].Cust_tag
		MerchantCustomerMapMaster.Cust_frg = RequestPointAddReq.ContainerDetail[index].Cust_frg
		MerchantCustomerMapMaster.Cust_status = RequestPointAddReq.ContainerDetail[index].Cust_status
		MerchantCustomerMapMaster.Cust_first_visit_date = RequestPointAddReq.ContainerDetail[index].Cust_first_visit_date
		MerchantCustomerMapMaster.Cust_last_visit_date = RequestPointAddReq.ContainerDetail[index].Cust_last_visit_date
		MerchantCustomerMapMaster.Create_by = utils.Decode(jwt.Raw).(string)
		MerchantCustomerMapMaster.Create_date = &now

		config.DBsql().Create(&MerchantCustomerMapMaster)

	}*/
	//var data []model.RequestPointEntityReq
	for index := 0; index < len(RequestPointAddReq.CustAr); index++ {

		for indexj := 0; indexj < len(RequestPointAddReq.ContainerDetail); indexj++ {
			var MerchantCustomerMapMaster entity.MerchantCustomerMapEntity
			MerchantCustomerMapMaster.Mc_id = RequestPointAddReq.Mc_id
			MerchantCustomerMapMaster.Cust_id = RequestPointAddReq.CustAr[index].Cust_id
			MerchantCustomerMapMaster.Container_id = RequestPointAddReq.ContainerDetail[indexj].Container_id

			if RequestPointAddReq.ContainerDetail[indexj].Container_type == "CC" {
				date := utils.GenIssueExpireDate(RequestPointAddReq.ContainerDetail[indexj].Expire_value, RequestPointAddReq.ContainerDetail[indexj].Expire_mode)
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
			MerchantCustomerMapMaster.Create_by = utils.Decode(jwt.Raw).(string)
			MerchantCustomerMapMaster.Create_date = &now

			err = config.DBsql().Create(&MerchantCustomerMapMaster).Error
		}
		if err == nil {
			var updateStatus model.RequestPointEntityReq
			config.DBsql().Model(&updateStatus).Where("request_point_id = ?", RequestPointAddReq.CustAr[index].Request_point_id).
				Update(model.RequestPointEntityReq{Request_status: "Y", Request_message: RequestPointAddReq.Request_message})
		}

	}

	return utils.ResDataAdd(err)

}
