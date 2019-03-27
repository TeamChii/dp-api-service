package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func UserById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.UserById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func UserCheck(ctx iris.Context) {
	phone := ctx.Params().Get("phone")
	var a = service.UserCheck(phone)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func UpdateUser(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var UserEntity entity.UserEntity
	err := ctx.ReadJSON(&UserEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		UserEntity.User_name = UserEntity.User_first_name
		var a = service.UpdateUser(UserEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UserByRole(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var UserEntity model.UserLoadRoleReq
	err := ctx.ReadJSON(&UserEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.UserByRole(UserEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UserLoadCategory(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var UserEntity model.UserLoadRoleReq
	err := ctx.ReadJSON(&UserEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.UserLoadCategory(UserEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func AddUser(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var UserEntity model.UserAddReq
	err := ctx.ReadJSON(&UserEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.AddUser(UserEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func AddUserStaff(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var UserEntity model.UserAddReq
	err := ctx.ReadJSON(&UserEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.AddUserStaff(UserEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateUserPin(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var PinReq model.PinReq
	err := ctx.ReadJSON(&PinReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.UpdateUserPin(PinReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func DeleteUser(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var UserEntity []entity.UserEntity
	err := ctx.ReadJSON(&UserEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.DeleteUser(UserEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
