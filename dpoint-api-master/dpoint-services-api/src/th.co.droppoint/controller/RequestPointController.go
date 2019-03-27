package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func LoadRequestPoint(ctx iris.Context) {
	var RequestPointReq model.RequestPointReq
	err := ctx.ReadJSON(&RequestPointReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadRequestPoint(RequestPointReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func SendCardAll(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var RequestPointAddReq model.RequestPointAddReq
	err := ctx.ReadJSON(&RequestPointAddReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SendCardAll(RequestPointAddReq, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func SendCard(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var RequestPointAddReq model.RequestPointAddReq2
	err := ctx.ReadJSON(&RequestPointAddReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SendCard(RequestPointAddReq, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
