package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetMiniPOSSettingById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.GetMiniPOSSettingById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func AddMiniPOSSetting(ctx iris.Context) {
	var MiniposSettingEntity entity.MiniposSettingEntity
	err := ctx.ReadJSON(&MiniposSettingEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMiniPOSSetting(MiniposSettingEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateMiniPOSSetting(ctx iris.Context) {
	var MiniposSettingEntity entity.MiniposSettingEntity
	err := ctx.ReadJSON(&MiniposSettingEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.UpdateMiniPOSSetting(MiniposSettingEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
