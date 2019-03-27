package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func AddMerchantImage(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var MerchantImageMapEntity entity.MerchantImageMapEntity
	err := ctx.ReadJSON(&MerchantImageMapEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMerchantImage(MerchantImageMapEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func DeleteMerchantImage(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var MerchantImageMapEntity entity.MerchantImageMapEntity
	err := ctx.ReadJSON(&MerchantImageMapEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.DeleteMerchantImage(MerchantImageMapEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
