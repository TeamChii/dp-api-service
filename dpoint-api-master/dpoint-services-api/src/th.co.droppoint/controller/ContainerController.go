package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func ContainerByIdMc(ctx iris.Context) {
	var containerReq model.ContainerReq
	err := ctx.ReadJSON(&containerReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.ContainerByIdMc(containerReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

/*func ContainerByIdMcCust(ctx iris.Context) {
	var containerReq model.ContainerReq2
	err := ctx.ReadJSON(&containerReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.ContainerByIdMcCust(containerReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}*/
func ContainerId(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.ContainerId(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func ContainerById(ctx iris.Context) {
	var ContainerByIdReq model.ContainerByIdReq
	err := ctx.ReadJSON(&ContainerByIdReq)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.ContainerById(ContainerByIdReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func AddContainer(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var ContainerMaster model.ContainerEntityReq
	err := ctx.ReadJSON(&ContainerMaster)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddContainer(ContainerMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func AddContainerMcOnly(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var ContainerMaster model.ContainerEntityReq
	err := ctx.ReadJSON(&ContainerMaster)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddContainerMcOnly(ContainerMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateContainer(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var ContainerMaster entity.ContainerEntity
	err := ctx.ReadJSON(&ContainerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//ContainerMaster.UpdatedBy = username
		var a = service.UpdateContainer(ContainerMaster, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
