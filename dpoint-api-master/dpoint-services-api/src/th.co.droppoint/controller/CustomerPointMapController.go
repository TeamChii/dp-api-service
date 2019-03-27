package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func CustomerPointMapByIdMcCust(ctx iris.Context) {
	var CustomerPointMapReq model.CustomerPointMapReq
	err := ctx.ReadJSON(&CustomerPointMapReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.CustomerPointMapByIdMcCust(CustomerPointMapReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func AddCustomerPointMap(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var ContainerMaster entity.CustomerPointMapEntity
	err := ctx.ReadJSON(&ContainerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddCustomerPointMap(ContainerMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateCustomerPointMap(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var ContainerMaster entity.CustomerPointMapEntity
	err := ctx.ReadJSON(&ContainerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//ContainerMaster.UpdatedBy = username
		var a = service.UpdateCustomerPointMap(ContainerMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
