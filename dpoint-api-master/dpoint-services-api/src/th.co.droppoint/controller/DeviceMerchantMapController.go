package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func LoadDevice(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.LoadDevice(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func CheckUidDevice(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var DeviceMerchantMapEntity entity.DeviceMerchantMapEntity
	err := ctx.ReadJSON(&DeviceMerchantMapEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.CheckUidDevice(DeviceMerchantMapEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func AddDevice(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var DeviceMerchantMapEntity entity.DeviceMerchantMapEntity
	err := ctx.ReadJSON(&DeviceMerchantMapEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.AddDevice(DeviceMerchantMapEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateDevice(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var DeviceMerchantMapEntity entity.DeviceMerchantMapEntity
	err := ctx.ReadJSON(&DeviceMerchantMapEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.UpdateDevice(DeviceMerchantMapEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func DeleteDevice(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var DeviceMerchantMapEntity []entity.DeviceMerchantMapEntity
	err := ctx.ReadJSON(&DeviceMerchantMapEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.DeleteDevice(DeviceMerchantMapEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
