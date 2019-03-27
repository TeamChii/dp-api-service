package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func GetCRMReport(ctx iris.Context) {
	var CRMReportReq model.CRMReportReq
	err := ctx.ReadJSON(&CRMReportReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.GetCRMReport(CRMReportReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
