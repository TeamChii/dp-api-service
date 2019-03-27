package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func MerchantCustomerMapByIdMcCust(ctx iris.Context) {
	var MerchantCustomerMapReq model.MerchantCustomerMapReq
	err := ctx.ReadJSON(&MerchantCustomerMapReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.MerchantCustomerMapByIdMcCust(MerchantCustomerMapReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func AddMerchantCustomerMap(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var ContainerMaster entity.MerchantCustomerMapEntity
	err := ctx.ReadJSON(&ContainerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMerchantCustomerMap(ContainerMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateMerchantCustomerMap(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var ContainerMaster entity.MerchantCustomerMapEntity
	err := ctx.ReadJSON(&ContainerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//ContainerMaster.UpdatedBy = username
		var a = service.UpdateMerchantCustomerMap(ContainerMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
