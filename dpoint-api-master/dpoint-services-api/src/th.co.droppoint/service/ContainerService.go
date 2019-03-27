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
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	/* */ /* */)

func ContainerByIdMc(containerMaster model.ContainerReq) map[string]interface{} {

	var data []model.ContainerEntityReq

	var count int
	//count := 2
	var err error
	if containerMaster.Active_Status != "" {
		if containerMaster.Contaienr_type != "" {
			err = config.DBsql().Joins("INNER JOIN dp_ms_container_type ON dp_ms_container_type.container_type_id = dp_ms_container.container_type_id").
				Where("( dp_ms_container.container_type_id = ? OR dp_ms_container_type.ref_id = ? ) AND dp_ms_container.active_status = ? AND dp_ms_container.mc_id = ?",
					containerMaster.Contaienr_type, containerMaster.Contaienr_type, containerMaster.Active_Status, containerMaster.Mc_id).
				Order("dp_ms_container." + containerMaster.Paging.OrderBy + " " + containerMaster.Paging.SortBy).
				Find(&data).
				Offset(((containerMaster.Paging.PageNo * int(containerMaster.Paging.PageSize)) - int(containerMaster.Paging.PageSize))).
				Limit(containerMaster.Paging.PageSize).
				Count(&count).Error
		} else {
			err = config.DBsql().Where("mc_id = ? AND active_status = ? ", containerMaster.Mc_id, containerMaster.Active_Status).
				Order("dp_ms_container." + containerMaster.Paging.OrderBy + " " + containerMaster.Paging.SortBy).
				Find(&data).
				Offset(((containerMaster.Paging.PageNo * int(containerMaster.Paging.PageSize)) - int(containerMaster.Paging.PageSize))).
				Limit(containerMaster.Paging.PageSize).
				Count(&count).Error
		}
	} else {
		if containerMaster.Contaienr_type != "" {
			err = config.DBsql().Joins("INNER JOIN dp_ms_container_type ON dp_ms_container_type.container_type_id = dp_ms_container.container_type_id").
				Where("( dp_ms_container.container_type_id = ? OR dp_ms_container_type.ref_id = ? ) AND dp_ms_container.mc_id = ?",
					containerMaster.Contaienr_type, containerMaster.Contaienr_type, containerMaster.Mc_id).
				Order("dp_ms_container." + containerMaster.Paging.OrderBy + " " + containerMaster.Paging.SortBy).
				Find(&data).
				Offset(((containerMaster.Paging.PageNo * int(containerMaster.Paging.PageSize)) - int(containerMaster.Paging.PageSize))).
				Limit(containerMaster.Paging.PageSize).
				Count(&count).Error
		} else {
			err = config.DBsql().Where("mc_id = ? ", containerMaster.Mc_id).
				Order("dp_ms_container." + containerMaster.Paging.OrderBy + " " + containerMaster.Paging.SortBy).
				Find(&data).
				Offset(((containerMaster.Paging.PageNo * int(containerMaster.Paging.PageSize)) - int(containerMaster.Paging.PageSize))).
				Limit(containerMaster.Paging.PageSize).
				Count(&count).Error
		}
	}

	for index := 0; index < len(data); index++ {
		var containerType entity.ContainerTypeEntity

		config.DBsql().Where("container_type_id = ?", data[index].Container_type_id).
			Find(&containerType)
		data[index].ContainerTypeEntity = &containerType
		data[index].Image_ref = utils.CONTENT_URL + "/" + data[index].Image_ref + ""
		expire_flag := "N"
		if data[index].Expire_mode == "FX" {
			timeFormat := "02/01/2006"
			expireDate, err := time.Parse(timeFormat, data[index].Expire_value)
			if err == nil {
				if expireDate.UnixNano() < time.Now().UnixNano() {
					expire_flag = "Y"
				}
			}
		}
		data[index].Expire_flag = expire_flag
		if data[index].Create_date != nil {
			data[index].Create_date_str = data[index].Create_date.Format("02/01/2006 15:04:05")
		}
		if data[index].Update_date != nil {
			data[index].Update_date_str = data[index].Update_date.Format("02/01/2006 15:04:05")
		}
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, containerMaster.Paging, count)
}

/*func ContainerByIdMcCust(containerMaster model.ContainerReq2) map[string]interface{} {

	var data []model.ContainerEntityReq

	var count int
	err := config.DBsql().Where("mc_id = ?", containerMaster.Mc_id).Find(&data).
		Preload("MerchantEntity").
		Count(&count).Error

	var containerType entity.ContainerTypeEntity
	var sumPoint entity.SumCustomerPointEntity
	for index := 0; index < len(data); index++ {
		config.DBsql().Where("container_type_id = ?", data[index].Container_type_id).Find(&containerType)
		data[index].ContainerTypeEntity = containerType

		config.DBsql().Where("cust_id = ? AND mc_id = ? AND container_id = ?", containerMaster.Cust_id, containerMaster.Mc_id, data[index].Container_id).
			Find(&sumPoint)

		data[index].Total_point = sumPoint.Current_point_amt
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, containerMaster.Paging, count)
}*/
func ContainerId(id string) map[string]interface{} {
	//fmt.Println(id)
	var data model.ContainerEntityReq
	check := config.DBsql().Where("container_id = ?", id).Find(&data).RecordNotFound()

	var ContainerTypeEntity entity.ContainerTypeEntity
	config.DBsql().Where("container_type_id = ?", data.Container_type_id).
		Find(&ContainerTypeEntity)

	data.ContainerTypeEntity = &ContainerTypeEntity

	var containerRewardEntity []entity.ContainerRewardEntity
	config.DBsql().Where("container_id = ?", id).Order("point_amt asc").
		Find(&containerRewardEntity)

	data.ContainerRewardEntity = containerRewardEntity
	if data.Create_date != nil {
		data.Create_date_str = data.Create_date.Format("02/01/2006 15:04:05")
	}
	if data.Update_date != nil {
		data.Update_date_str = data.Update_date.Format("02/01/2006 15:04:05")
	}
	if len(data.Image_ref) > 0 {
		data.Image_ref = utils.CONTENT_URL + "/" + data.Image_ref
	}

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func ContainerById(ContainerByIdReq model.ContainerByIdReq) map[string]interface{} {

	var data model.ContainerEntityReq
	check := config.DBsql().Where("container_id = ?", ContainerByIdReq.Container_id).Find(&data).RecordNotFound()

	var ContainerTypeEntity entity.ContainerTypeEntity
	config.DBsql().Where("container_type_id = ?", data.Container_type_id).
		Find(&ContainerTypeEntity)

	data.ContainerTypeEntity = &ContainerTypeEntity

	var mc_group entity.MerchantEntity

	config.DBsql().
		Where("mc_id = ? ", ContainerByIdReq.Mc_id).Find(&mc_group)
	var sumPoint entity.SumCustomerPointEntity

	config.DBsql().Where("cust_id = ? AND mc_id = ? AND container_id = ?", ContainerByIdReq.Cust_id, mc_group.Mc_group_id, ContainerByIdReq.Container_id).
		Find(&sumPoint)
	var totalPoint = 0
	if sumPoint.Current_point_amt != nil {
		totalPoint = *sumPoint.Current_point_amt
	}
	data.Total_point = totalPoint // *sumPoint.Current_point_amt

	data.Image_ref = utils.CONTENT_URL + "/" + data.Image_ref + ""

	if data.Create_date != nil {
		data.Create_date_str = data.Create_date.Format("02/01/2006 15:04:05")
	}
	if data.Update_date != nil {
		data.Update_date_str = data.Update_date.Format("02/01/2006 15:04:05")
	}

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func AddContainer(containerMaster model.ContainerEntityReq, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//	var resModel model.ResponseModel
	//id, _ := strconv.Atoi(utils.RandSeq(32))
	//ContainerMaster.Mc_id = id
	now := time.Now()
	containerMaster.Create_date = &now
	containerMaster.Create_by = utils.Decode(jwt.Raw).(string)
	err := config.DBsql().Create(&containerMaster).Error
	if err == nil {

		var containerReward entity.ContainerRewardEntity
		for index := 0; index < len(containerMaster.ContainerRewardEntity); index++ {
			containerReward.Container_id = containerMaster.Container_id
			containerReward.Point_amt = containerMaster.ContainerRewardEntity[index].Point_amt
			containerReward.Reward_detail = containerMaster.ContainerRewardEntity[index].Reward_detail
			config.DBsql().Create(&containerReward)
		}

		//date := utils.GenIssueExpireDate(containerMaster.Expire_value, containerMaster.Expire_mode)
		//issue_date := date["issue_date"].(time.Time)
		//expire_date := date["expire_date"].(*time.Time)

		var merchantCustMap entity.MerchantCustomerMapEntity
		merchantCustMap.Mc_id = containerMaster.Mc_id
		merchantCustMap.Cust_id = containerMaster.Cust_id
		merchantCustMap.Container_id = containerMaster.Container_id
		merchantCustMap.Expire_date = nil
		merchantCustMap.Issue_date = nil
		config.DBsql().Create(&merchantCustMap)
	}

	return utils.ResDataAdd(err)

}
func AddContainerMcOnly(containerMaster model.ContainerEntityReq, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//	var resModel model.ResponseModel
	//id, _ := strconv.Atoi(utils.RandSeq(32))
	//ContainerMaster.Mc_id = id
	//ContainerMaster.Create_date = time.Now()
	now := time.Now()
	containerMaster.Create_date = &now
	containerMaster.Create_by = utils.Decode(jwt.Raw).(string)
	err := config.DBsql().Create(&containerMaster).Error

	if err == nil {
		var containerReward entity.ContainerRewardEntity
		for index := 0; index < len(containerMaster.ContainerRewardEntity); index++ {
			containerReward.Container_id = containerMaster.Container_id
			containerReward.Point_amt = containerMaster.ContainerRewardEntity[index].Point_amt
			containerReward.Reward_detail = containerMaster.ContainerRewardEntity[index].Reward_detail
			config.DBsql().Create(&containerReward)
		}
	}

	return utils.ResDataAdd(err)

}
func UpdateContainer(containerMaster entity.ContainerEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	now := time.Now()
	containerMaster.Update_date = &now
	containerMaster.Update_by = utils.Decode(jwt.Raw).(string)
	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	var count int
	config.DBsql().Where("container_id = ? ", containerMaster.Container_id).
		Table("dp_ms_container_reward").Count(&count).Delete(entity.ContainerRewardEntity{})

	var containerReward entity.ContainerRewardEntity
	for index := 0; index < len(containerMaster.ContainerRewardEntity); index++ {
		containerReward.Container_id = containerMaster.Container_id
		containerReward.Point_amt = containerMaster.ContainerRewardEntity[index].Point_amt
		containerReward.Reward_detail = containerMaster.ContainerRewardEntity[index].Reward_detail
		config.DBsql().Create(&containerReward)
	}
	err := config.DBsql().Model(&containerMaster).Where("container_id = ?", containerMaster.Container_id).Update(&containerMaster).Error

	return utils.ResDataEdit(err)
}
