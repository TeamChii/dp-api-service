package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetMiniposGiveRewardSettingById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.GetGiveRewardSettingById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func AddMiniposGiveRewardSetting(ctx iris.Context) {
	var GivePointRewardModel model.GivePointRewardModel
	err := ctx.ReadJSON(&GivePointRewardModel)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddGiveRewardSetting(GivePointRewardModel, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateMiniposGiveRewardSetting(ctx iris.Context) {
	var GivePointRewardModel model.GivePointRewardModel
	err := ctx.ReadJSON(&GivePointRewardModel)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.UpdateGiveRewardSetting(GivePointRewardModel, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
