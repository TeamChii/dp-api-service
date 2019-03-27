package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/kataras/iris"
)

func PointCheck(ctx iris.Context) {
	phone := ctx.Params().Get("phone")
	var a = service.PointCheck(phone)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
