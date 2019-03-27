package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func SystemParamSearch(ctx iris.Context) {
	var SearchObjectSystemParamModel model.SearchObjectSystemParamModel
	err := ctx.ReadJSON(&SearchObjectSystemParamModel)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SystemParamSearch(SearchObjectSystemParamModel)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func SystemParamAdd(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var SystemParamEntity entity.SystemParamEntity
	err := ctx.ReadJSON(&SystemParamEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SystemParamAdd(SystemParamEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func SystemParamUpdate(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var SystemParamEntity entity.SystemParamEntity
	err := ctx.ReadJSON(&SystemParamEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//RedeemMaster.UpdatedBy = username
		var a = service.SystemParamUpdate(SystemParamEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func SystemParamDelete(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var SystemParamEntity []entity.SystemParamEntity
	err := ctx.ReadJSON(&SystemParamEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//RedeemMaster.UpdatedBy = username
		var a = service.SystemParamDelete(SystemParamEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
