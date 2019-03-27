package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func ContainerTypeSearch(ctx iris.Context) {
	var searchObject model.SearchObjectModel
	err := ctx.ReadJSON(&searchObject)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.ContainerTypeSearch(searchObject)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func ContainerTypeById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.ContainerTypeById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func AddContainerType(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var containerTypeMaster entity.ContainerTypeEntity
	err := ctx.ReadJSON(&containerTypeMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddContainerType(containerTypeMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateContainerType(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var containerTypeMaster entity.ContainerTypeEntity
	err := ctx.ReadJSON(&containerTypeMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//ContainerTypeMaster.UpdatedBy = username
		var a = service.UpdateContainerType(containerTypeMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
