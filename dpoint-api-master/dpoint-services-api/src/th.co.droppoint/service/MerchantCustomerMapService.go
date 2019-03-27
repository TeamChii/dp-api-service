package service

import (
	"encoding/json"
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

func MerchantCustomerMapByIdMcCust(MerchantCustomerMapMaster model.MerchantCustomerMapReq) map[string]interface{} {

	var data []model.MerchantCustomerMapEntity

	var count int
	var err error

	var mc_group entity.MerchantEntity

	config.DBsql().
		Where("mc_id = ? ", MerchantCustomerMapMaster.Mc_id).Find(&mc_group)
	if MerchantCustomerMapMaster.Contaienr_type != "" {

		err = config.DBsql().
			Select("dp_mp_merchant_customer.container_id").
			Joins("INNER JOIN dp_ms_container ON dp_mp_merchant_customer.container_id = dp_ms_container.container_id INNER JOIN dp_ms_container_type ON dp_ms_container.container_type_id = dp_ms_container_type.container_type_id").
			Where(" "+
				" (dp_mp_merchant_customer.mc_id in (select mc_id from dp_ms_merchant where mc_group_id = ? ) "+
				" AND dp_ms_container.container_type_id = ? ) "+
				" or (dp_mp_merchant_customer.mc_id in (select mc_id from dp_ms_merchant where mc_group_id = ? ) "+
				" AND dp_ms_container_type.ref_id = ? ) "+
				//" OR dp_ms_container_type.ref_id = ? " +
				" AND dp_mp_merchant_customer.cust_id = ?",
				mc_group.Mc_group_id, MerchantCustomerMapMaster.Contaienr_type,
				mc_group.Mc_group_id, MerchantCustomerMapMaster.Contaienr_type,
				//MerchantCustomerMapMaster.Contaienr_type,
				MerchantCustomerMapMaster.Cust_id).
			Offset(((MerchantCustomerMapMaster.Paging.PageNo * int(MerchantCustomerMapMaster.Paging.PageSize)) - int(MerchantCustomerMapMaster.Paging.PageSize))).
			Limit(MerchantCustomerMapMaster.Paging.PageSize).
			Group("dp_mp_merchant_customer.container_id").
			Find(&data).Count(&count).Error
		/* oh comment
		for index := 0; index < len(data); index++ {
			var data3 model.MerchantCustomerMapEntity
			config.DBsql().
				Where("mc_id in (select mc_id from dp_ms_merchant where mc_group_id = ? ) AND cust_id = ? AND container_id = ?",
				mc_group.Mc_group_id, MerchantCustomerMapMaster.Cust_id, data[index].Container_id).
				Order(MerchantCustomerMapMaster.Paging.OrderBy + " " + MerchantCustomerMapMaster.Paging.SortBy).
				Find(&data3)
			data[index] = data3

		}
		*/
		//fmt.Println(data)
		/*err = config.DBsql().
		Joins("INNER JOIN dp_ms_container ON dp_mp_merchant_customer.container_id = dp_ms_container.container_id INNER JOIN dp_ms_container_type ON dp_ms_container.container_type_id = dp_ms_container_type.container_type_id").
		Where("(dp_mp_merchant_customer.mc_id = ? AND dp_mp_merchant_customer.cust_id = ? AND dp_ms_container.container_type_id = ?) OR dp_ms_container_type.ref_id = ?",
			MerchantCustomerMapMaster.Mc_id, MerchantCustomerMapMaster.Cust_id, MerchantCustomerMapMaster.Contaienr_type, MerchantCustomerMapMaster.Contaienr_type).
		Offset(((MerchantCustomerMapMaster.Paging.PageNo * int(MerchantCustomerMapMaster.Paging.PageSize)) - int(MerchantCustomerMapMaster.Paging.PageSize))).
		Limit(MerchantCustomerMapMaster.Paging.PageSize).
		Order(MerchantCustomerMapMaster.Paging.OrderBy + " " + MerchantCustomerMapMaster.Paging.SortBy).
		Find(&data).Count(&count).Error*/
	} else {
		err = config.DBsql().
			Select("container_id").
			Where("mc_id in (select mc_id from dp_ms_merchant where mc_group_id = ? )  AND cust_id = ?", mc_group.Mc_group_id, MerchantCustomerMapMaster.Cust_id).
			Offset(((MerchantCustomerMapMaster.Paging.PageNo * int(MerchantCustomerMapMaster.Paging.PageSize)) - int(MerchantCustomerMapMaster.Paging.PageSize))).
			Limit(MerchantCustomerMapMaster.Paging.PageSize).
			Group("container_id").
			Find(&data).Count(&count).Error
		/* oh comment
		for index := 0; index < len(data); index++ {
			var data3 model.MerchantCustomerMapEntity
			config.DBsql().
				Where("mc_id in (select mc_id from dp_ms_merchant where mc_group_id = ? ) AND cust_id = ? AND container_id = ?",
				mc_group.Mc_group_id, MerchantCustomerMapMaster.Cust_id, data[index].Container_id).
				Order(MerchantCustomerMapMaster.Paging.OrderBy + " " + MerchantCustomerMapMaster.Paging.SortBy).
				Find(&data3)
			data[index] = data3

		}
		*/
	}
	var dataRejectPoint []model.MerchantCustomerMapEntity
	for index := 0; index < len(data); index++ {
		var container model.ContainerEntityReq
		var sumPoint entity.SumCustomerPointEntity
		config.DBsql().Select("sum(current_point_amt) as current_point_amt ").Where("cust_id = ? "+
			" AND mc_id in (select mc_id from dp_ms_merchant where mc_group_id = ? ) AND container_id = ?",
			MerchantCustomerMapMaster.Cust_id, mc_group.Mc_group_id, data[index].Container_id).
			Find(&sumPoint)

		config.DBsql().Where("container_id = ?", data[index].Container_id).Find(&container)
		data[index].ContainerEntityReq = &container

		expire_flag := "N"
		if container.Expire_mode == "FX" {
			timeFormat := "02/01/2006"
			expireDate, err := time.Parse(timeFormat, container.Expire_value)
			if err == nil {
				if expireDate.UnixNano() < time.Now().UnixNano() {
					expire_flag = "Y"
				}
			}
		}

		data[index].ContainerEntityReq.Expire_flag = expire_flag

		var containerReward []entity.ContainerRewardEntity
		config.DBsql().Where("container_id = ?", data[index].Container_id).Find(&containerReward)

		var ContainerTypeEntity entity.ContainerTypeEntity
		config.DBsql().Where("container_type_id = ?", data[index].ContainerEntityReq.Container_type_id).
			Find(&ContainerTypeEntity)

		data[index].ContainerEntityReq.ContainerTypeEntity = &ContainerTypeEntity

		if sumPoint.Current_point_amt != nil {
			data[index].ContainerEntityReq.Total_point = *sumPoint.Current_point_amt
		}
		//data[index].ContainerEntityReq.Total_point = *sumPoint.Current_point_amt;
		data[index].ContainerEntityReq.Image_ref = utils.CONTENT_URL + "/" + data[index].ContainerEntityReq.Image_ref + ""
		data[index].ContainerEntityReq.ContainerRewardEntity = containerReward

		if sumPoint.Current_point_amt != nil {
			if *sumPoint.Current_point_amt > 0 && data[index].ContainerEntityReq.Container_type_id == 2 {
				dataRejectPoint = append(dataRejectPoint, data[index])
			}
		}
		/*
			if *sumPoint.Current_point_amt > 0 && data[index].ContainerEntityReq.Container_type_id == 2 {
				dataRejectPoint = append(dataRejectPoint, data[index])
			}
		*/
	}

	if MerchantCustomerMapMaster.Contaienr_type == "2" {
		MerchantCustomerMapMaster.Paging.TotalRecord = len(dataRejectPoint)
		MerchantCustomerMapMaster.Paging.TotalPage = len(dataRejectPoint) / MerchantCustomerMapMaster.Paging.PageSize
		var dataToJson []map[string]interface{}
		b, _ := json.Marshal(dataRejectPoint)
		json.Unmarshal(b, &dataToJson)
		return utils.ResDataLoad(err, dataToJson, MerchantCustomerMapMaster.Paging, count)
	} else {
		var dataToJson []map[string]interface{}
		b, _ := json.Marshal(data)
		json.Unmarshal(b, &dataToJson)
		return utils.ResDataLoad(err, dataToJson, MerchantCustomerMapMaster.Paging, count)
	}
}
func AddMerchantCustomerMap(MerchantCustomerMapMaster entity.MerchantCustomerMapEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//	var resModel model.ResponseModel
	//id, _ := strconv.Atoi(utils.RandSeq(32))
	//MerchantCustomerMapMaster.Mc_id = id
	MerchantCustomerMapMaster.Create_by = utils.Decode(jwt.Raw).(string)
	now := time.Now()
	MerchantCustomerMapMaster.Create_date = &now

	err := config.DBsql().Create(&MerchantCustomerMapMaster).Error
	return utils.ResDataAdd(err)

}

func UpdateMerchantCustomerMap(MerchantCustomerMapMaster entity.MerchantCustomerMapEntity) map[string]string {

	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	err := config.DBsql().Model(&MerchantCustomerMapMaster).Where("mc_id = ? AND cust_id = ? AND container_id = ?",
		MerchantCustomerMapMaster.Mc_id, MerchantCustomerMapMaster.Cust_id, MerchantCustomerMapMaster.Container_id).
		Update(&MerchantCustomerMapMaster).Error

	return utils.ResDataEdit(err)
}
