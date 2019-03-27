package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func IncreasPoint(ctx iris.Context) {
	var increasPoint model.LoginToOpenReq
	err := ctx.ReadJSON(&increasPoint)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.CheckIncreasPoint(increasPoint)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
