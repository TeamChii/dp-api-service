package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func LoadDashboard(ctx iris.Context) {

	var ReportReq model.ReportReq
	err := ctx.ReadJSON(&ReportReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.LoadDashboard(ReportReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func LoadReportList(ctx iris.Context) {
	var a = service.LoadReportList()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func LoadReportByCategory(ctx iris.Context) {

	var ReportDateReq model.ReportDateReq
	err := ctx.ReadJSON(&ReportDateReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.LoadReportByCategory(ReportDateReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func LoadCampaignMonitor(ctx iris.Context) {

	var ReportReq model.ReportReq
	err := ctx.ReadJSON(&ReportReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.LoadCampaignMonitor(ReportReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
