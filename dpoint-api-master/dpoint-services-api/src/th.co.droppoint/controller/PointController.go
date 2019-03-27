package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func PointByIdMc(ctx iris.Context) {
	var PointReq model.PointReq
	err := ctx.ReadJSON(&PointReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.PointByIdMc(PointReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func AddPoint(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var PointMaster model.PointEntityReq
	err := ctx.ReadJSON(&PointMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddPoint(PointMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdatePoint(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var PointMaster entity.PointEntity
	err := ctx.ReadJSON(&PointMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//PointMaster.UpdatedBy = username
		var a = service.UpdatePoint(PointMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func CheckMobileCustomer(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var PointCheckMobileReq model.PointCheckMobileReq
	err := ctx.ReadJSON(&PointCheckMobileReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//PointMaster.UpdatedBy = username
		var a = service.CheckMobileCustomer(PointCheckMobileReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
