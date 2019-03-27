package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func MerchantById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.MerchantById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func MerchantByPhone(ctx iris.Context) {
	userid := ctx.Params().Get("userid")
	var a = service.MerchantByPhone(userid)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func AddMerchant(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var merchantMaster model.MerchantEntityReq
	err := ctx.ReadJSON(&merchantMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMerchant(merchantMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func AddMerchant2(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var merchantMaster model.MerchantEntityReq
	err := ctx.ReadJSON(&merchantMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMerchant2(merchantMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateMerchant(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var merchantMaster entity.MerchantEntity
	err := ctx.ReadJSON(&merchantMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//merchantMaster.UpdatedBy = username
		var a = service.UpdateMerchant(merchantMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func SetHeadOfficeMerchant(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var merchantMaster entity.MerchantEntity
	err := ctx.ReadJSON(&merchantMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//merchantMaster.UpdatedBy = username
		var a = service.SetHeadOfficeMerchant(merchantMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
