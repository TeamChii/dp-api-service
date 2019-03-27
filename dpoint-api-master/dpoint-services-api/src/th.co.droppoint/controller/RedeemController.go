package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func RedeemByIdMc(ctx iris.Context) {
	var RedeemReq model.RedeemReq
	err := ctx.ReadJSON(&RedeemReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.RedeemByIdMc(RedeemReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func AddRedeem(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var RedeemMaster model.RedeemReqList
	err := ctx.ReadJSON(&RedeemMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddRedeem(RedeemMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateRedeem(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var RedeemMaster entity.RedeemEntity
	err := ctx.ReadJSON(&RedeemMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//RedeemMaster.UpdatedBy = username
		var a = service.UpdateRedeem(RedeemMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
