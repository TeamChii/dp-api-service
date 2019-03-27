package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func PointByIdMc(PointMaster model.PointReq) map[string]interface{} {

	var data []entity.PointEntity

	var count int
	err := config.DBsql().Where("mc_id = ?", PointMaster.Mc_id).Find(&data).Preload("MerchantEntity").Count(&count).Error

	fmt.Println(data)
	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, PointMaster.Paging, count)
}
func AddPoint(PointModel model.PointEntityReq, ctx iris.Context) map[string]interface{} {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)

	var merchantMapcustomer model.MerchantCustomerMapEntity
	//fmt.Println(PointModel)
	if PointModel.Status == "1" {
		var CustomerEntity entity.CustomerEntity
		CustomerEntity.Cust_mobile = PointModel.Cust_mobile
		CustomerEntity.Cust_name = "Unnamed Customer"
		config.DBsql().Create(&CustomerEntity)

		PointModel.Cust_id = CustomerEntity.Cust_id

		var MerchantCustomerMapEntity entity.MerchantCustomerMapEntity
		MerchantCustomerMapEntity.Mc_id = PointModel.Mc_id
		MerchantCustomerMapEntity.Cust_id = CustomerEntity.Cust_id
		MerchantCustomerMapEntity.Container_id = PointModel.Container_id
		MerchantCustomerMapEntity.Create_by = utils.Decode(jwt.Raw).(string)
		now := time.Now()
		MerchantCustomerMapEntity.Create_date = &now
		config.DBsql().Create(&MerchantCustomerMapEntity)

		config.DBsql().Where("mc_id = ? AND container_id = ? AND cust_id = ?", PointModel.Mc_id, PointModel.Container_id, PointModel.Cust_id).
			Find(&merchantMapcustomer)
	} else /*if PointModel.Status == "2"*/ {
		var MerchantCustomerMapMaster entity.MerchantCustomerMapEntity
		MerchantCustomerMapMaster.Mc_id = PointModel.Mc_id
		MerchantCustomerMapMaster.Cust_id = PointModel.Cust_id
		MerchantCustomerMapMaster.Container_id = PointModel.Container_id
		MerchantCustomerMapMaster.Create_by = utils.Decode(jwt.Raw).(string)
		now := time.Now()
		MerchantCustomerMapMaster.Create_date = &now
		config.DBsql().Create(&MerchantCustomerMapMaster)

		config.DBsql().Where("mc_id = ? AND container_id = ? AND cust_id = ?", PointModel.Mc_id, PointModel.Container_id, PointModel.Cust_id).
			Find(&merchantMapcustomer)

	} /*else {

		var MerchantCustomerMapMaster entity.MerchantCustomerMapEntity
		MerchantCustomerMapMaster.Mc_id = PointModel.Mc_id
		MerchantCustomerMapMaster.Cust_id = PointModel.Cust_id
		MerchantCustomerMapMaster.Container_id = PointModel.Container_id
		MerchantCustomerMapMaster.Create_by = utils.Decode(jwt.Raw).(string)
		now := time.Now()
		MerchantCustomerMapMaster.Create_date = &now
		config.DBsql().Create(&MerchantCustomerMapMaster)

		config.DBsql().Where("mc_id = ? AND container_id = ? AND cust_id = ?", PointModel.Mc_id, PointModel.Container_id, PointModel.Cust_id).
			Find(&merchantMapcustomer)
	}*/

	var container model.ContainerEntityReq
	config.DBsql().Where("mc_id = ? AND container_id = ?", PointModel.Mc_id, PointModel.Container_id).
		Find(&container)

	var expire_date *time.Time
	//if container.Expire_mode != "NO" && merchantMapcustomer.Expire_date != nil && merchantMapcustomer.Issue_date != nil {
	if container.Expire_mode != "NO" {
		date := utils.GenIssueExpireDate(container.Expire_value, container.Expire_mode)
		issue_date := date["issue_date"].(*time.Time)
		//expire_date := date["expire_date"].(time.Time)
		expire_date = date["expire_date"].(*time.Time)

		var MerchantCustomerMapMaster entity.MerchantCustomerMapEntity
		MerchantCustomerMapMaster.Issue_date = issue_date
		MerchantCustomerMapMaster.Expire_date = expire_date
		config.DBsql().Model(&MerchantCustomerMapMaster).Where("mc_id = ? AND cust_id = ? AND container_id = ?",
			PointModel.Mc_id, PointModel.Cust_id, PointModel.Container_id).
			Update(&MerchantCustomerMapMaster)
	}
	var PointMaster entity.PointEntity
	now := time.Now()
	//date := utils.StringToDate(PointModel.Expire_date)
	PointMaster.Mc_id = PointModel.Mc_id
	PointMaster.Cust_id = PointModel.Cust_id
	PointMaster.Container_id = PointModel.Container_id
	PointMaster.Menu_id = PointModel.Menu_id
	PointMaster.Transfer_to_cust_id = PointModel.Transfer_to_cust_id
	PointMaster.Point_amt = PointModel.Point_amt
	PointMaster.Create_by = utils.Decode(jwt.Raw).(string)
	PointMaster.Create_date = &now
	//PointMaster.Expire_date = &date
	PointMaster.Expire_date = expire_date
	//PointMaster.Expire_flag = PointModel.Expire_flag
	PointMaster.Expire_flag = "N"
	PointMaster.Redeem_flag = "N"

	err := config.DBsql().Create(&PointMaster).Error
	config.DBsql().Table("dp_tt_point").Create(&PointMaster)
	if err == nil {
		var mc_group entity.MerchantEntity

		config.DBsql().
			Where("mc_id = ? ", PointModel.Mc_id).Find(&mc_group)
		var sumcustomerPoint entity.SumCustomerPointEntity
		var check int
		config.DBsql().Where("mc_id = ? AND cust_id = ? AND container_id = ?", mc_group.Mc_group_id, PointModel.Cust_id, PointModel.Container_id).
			Find(&sumcustomerPoint).Count(&check)

		pointCurrent := sumcustomerPoint.Current_point_amt
		sumcustomerPoint.Cust_id = PointMaster.Cust_id
		sumcustomerPoint.Mc_id = PointMaster.Mc_id
		sumcustomerPoint.Container_id = PointMaster.Container_id
		point := PointMaster.Point_amt
		sumcustomerPoint.Current_point_amt = &point
		sumcustomerPoint.Created_by = utils.Decode(jwt.Raw).(string)
		sumcustomerPoint.Created_date = &now

		if check > 0 {
			sumpoint := *pointCurrent + PointMaster.Point_amt
			sumcustomerPoint.Current_point_amt = &sumpoint

			config.DBsql().Model(&sumcustomerPoint).
				Where("mc_id = ? AND cust_id = ? AND container_id = ?", PointMaster.Mc_id, PointMaster.Cust_id, PointMaster.Container_id).
				Update(&sumcustomerPoint)
		} else {
			config.DBsql().Create(&sumcustomerPoint)
		}

	}
	return utils.ResDataAdd2(err, PointModel.Cust_id)
}

func UpdatePoint(PointMaster entity.PointEntity) map[string]string {
	//account.UpdatedDate = time.Now()
	err := config.DBsql().Model(&PointMaster).Where("point_id = ?", PointMaster.Point_id).Update(&PointMaster).Error
	return utils.ResDataEdit(err)
}
func CheckMobileCustomer(PointCheckMobileReq model.PointCheckMobileReq) map[string]interface{} {
	var CustomerEntity entity.CustomerEntity
	var MerchantCustomerMapEntity []entity.MerchantCustomerMapEntity
	var status string
	var desc string
	//var otp_obj map[string]interface{}
	check1 := config.DBsql().
		Where("cust_mobile = ?", PointCheckMobileReq.Cust_mobile).
		Find(&CustomerEntity).RecordNotFound()

	if check1 == false {
		var mc_group entity.MerchantEntity

		config.DBsql().
			Where("mc_id = ? ", PointCheckMobileReq.Mc_id).Find(&mc_group)
		check2 := config.DBsql().
			Where("cust_id = ? AND mc_id in (select mc_id from dp_ms_merchant where mc_group_id = ? ) ",
				CustomerEntity.Cust_id, mc_group.Mc_group_id).
			Find(&MerchantCustomerMapEntity).RecordNotFound()
		//fmt.Println(MerchantCustomerMapEntity)
		if check2 == false {
			status = "3"
			desc = "found 2 table"
		} else {
			status = "2"
			desc = "found in customer only"
			//code := utils.CommonSendOTPCode(PointCheckMobileReq.Cust_mobile)
			//otp_obj = AddAuthVerify(code)
		}
	} else {
		status = "1"
		desc = "not found 2 table"
		//code := utils.CommonSendOTPCode(PointCheckMobileReq.Cust_mobile)
		//otp_obj = AddAuthVerify(code)
	}
	return map[string]interface{}{
		"status": status,
		"desc":   desc,
	}

}
