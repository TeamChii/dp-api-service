package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetCustomerByMobile(ctx iris.Context) {
	var customerObj model.CustomerReq
	err := ctx.ReadJSON(&customerObj)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.GetCustomerByMobileMcId(customerObj)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func GetCustomerById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.GetCustomerById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func UpdateCustomer(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var customerMaster entity.CustomerEntity
	err := ctx.ReadJSON(&customerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.UpdateCustomer(customerMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

/*
func CustomerSearch(ctx iris.Context) {
	var searchObject model.SearchObjectModel
	err := ctx.ReadJSON(&searchObject)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.CustomerSearch(searchObject)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func GetCustomerById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.GetCustomerById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}

func AddCustomer(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var customerMaster entity.CustomerEntity
	err := ctx.ReadJSON(&customerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddCustomer(customerMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateCustomer(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var customerMaster entity.CustomerEntity
	err := ctx.ReadJSON(&customerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.UpdateCustomer(customerMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func DeleteCustomer(ctx iris.Context) {
	var customerMaster []entity.CustomerEntity
	err := ctx.ReadJSON(&customerMaster)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.DeleteCustomer(customerMaster)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
*/
