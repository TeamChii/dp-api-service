package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func PackageLoad(ctx iris.Context) {
	var PackageReq model.PackageReq
	err := ctx.ReadJSON(&PackageReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.PackageLoad(PackageReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func PackageAdd(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var PackageEntity entity.PackageEntity
	err := ctx.ReadJSON(&PackageEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.PackageAdd(PackageEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func PackageUpdate(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var PackageEntity entity.PackageEntity
	err := ctx.ReadJSON(&PackageEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//PointMaster.UpdatedBy = username
		var a = service.PackageUpdate(PackageEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func PackageDelete(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var PackageEntity []entity.PackageEntity
	err := ctx.ReadJSON(&PackageEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//PointMaster.UpdatedBy = username
		var a = service.PackageDelete(PackageEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
