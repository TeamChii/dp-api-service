package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func NotiLoad(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var NotificationEntity model.NotiCategoryReq
	err := ctx.ReadJSON(&NotificationEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.NotiLoad(NotificationEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func NotiCategory(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.NotiCategory(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}

func NotiById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.NotiById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func NotiRead(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var NotificationEntity entity.NotificationEntity
	err := ctx.ReadJSON(&NotificationEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.NotiRead(NotificationEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func NotiDelete(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var NotificationEntity entity.NotificationEntity
	err := ctx.ReadJSON(&NotificationEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.NotiDelete(NotificationEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
