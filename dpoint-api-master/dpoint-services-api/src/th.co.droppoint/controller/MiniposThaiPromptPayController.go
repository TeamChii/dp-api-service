package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetMiniposThaiPromptPayById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.GetMiniposThaiPromptPayById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func AddMiniposThaiPromptPay(ctx iris.Context) {
	var MiniposThaiPromptPayEntity entity.MiniposThaiPromptPayEntity
	err := ctx.ReadJSON(&MiniposThaiPromptPayEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMiniposThaiPromptPay(MiniposThaiPromptPayEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateMiniposThaiPromptPay(ctx iris.Context) {
	var MiniposThaiPromptPayEntity entity.MiniposThaiPromptPayEntity
	err := ctx.ReadJSON(&MiniposThaiPromptPayEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.UpdateMiniposThaiPromptPay(MiniposThaiPromptPayEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
