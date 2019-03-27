package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func LoadCategoryMenu(ctx iris.Context) {
	var MenuReq model.MenuReq
	err := ctx.ReadJSON(&MenuReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadCategoryMenu(MenuReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func LoadMenuByCategory(ctx iris.Context) {
	var MenuReq model.MenuReq
	err := ctx.ReadJSON(&MenuReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadMenuByCategory(MenuReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func LoadActiveMenuByCategory(ctx iris.Context) {
	var MenuReq model.MenuReq
	err := ctx.ReadJSON(&MenuReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadActiveMenuByCategory(MenuReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func LoadAllMenu(ctx iris.Context) {
	var MenuReq model.MenuReq
	err := ctx.ReadJSON(&MenuReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadAllMenu(MenuReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func AddTransactionAndReceive(ctx iris.Context) {
	var TransactionReceiveReq model.TransactionReceiveReq
	err := ctx.ReadJSON(&TransactionReceiveReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddTransactionAndReceive(TransactionReceiveReq, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func GetMiniPOSCategoryById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.GetMiniPOSCategoryById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func AddMiniPOSCategory(ctx iris.Context) {
	var MiniposMenuCategoryEntity entity.MiniposMenuCategoryEntity
	err := ctx.ReadJSON(&MiniposMenuCategoryEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMiniPOSCategory(MiniposMenuCategoryEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateMiniPOSCategory(ctx iris.Context) {
	var MiniposMenuCategoryEntity entity.MiniposMenuCategoryEntity
	err := ctx.ReadJSON(&MiniposMenuCategoryEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.UpdateMiniPOSCategory(MiniposMenuCategoryEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func DeleteMiniPOSCategory(ctx iris.Context) {
	//decoded := context.Get(ctx.Request(), "username")
	//var username = decoded.(string)
	var MiniposMenuCategoryEntity []entity.MiniposMenuCategoryEntity
	err := ctx.ReadJSON(&MiniposMenuCategoryEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//customerMaster.UpdatedBy = username
		var a = service.DeleteMiniPOSCategory(MiniposMenuCategoryEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func GetMiniPOSMenuById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var a = service.GetMiniposMenuById(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func AddMiniPOSMenu(ctx iris.Context) {
	var MiniposMenuEntity entity.MiniposMenuEntity
	err := ctx.ReadJSON(&MiniposMenuEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMiniposMenu(MiniposMenuEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UpdateMiniPOSMenu(ctx iris.Context) {
	var MiniposMenuEntity entity.MiniposMenuEntity
	err := ctx.ReadJSON(&MiniposMenuEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.UpdateMiniposMenu(MiniposMenuEntity, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func DeleteMiniPOSMenu(ctx iris.Context) {
	var MiniposMenuEntity []entity.MiniposMenuEntity
	err := ctx.ReadJSON(&MiniposMenuEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {

		var a = service.DeleteMiniposMenu(MiniposMenuEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func AddMiniposMappingMenu(ctx iris.Context) {
	var MiniposMenuEntity []entity.MiniposMenuEntity
	err := ctx.ReadJSON(&MiniposMenuEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.AddMiniposMappingMenu(MiniposMenuEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func DeleteMiniposMappingMenu(ctx iris.Context) {
	var MiniposMenuEntity []entity.MiniposMenuEntity
	err := ctx.ReadJSON(&MiniposMenuEntity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.DeleteMiniposMappingMenu(MiniposMenuEntity)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func SearchReceiptHistory(ctx iris.Context) {
	var MiniposReceiptHistoryReq model.MiniposReceiptHistoryReq
	err := ctx.ReadJSON(&MiniposReceiptHistoryReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SearchReceiptHistory(MiniposReceiptHistoryReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}

func UpdateReceiptStatus(ctx iris.Context) {
	var miniposReceiptHistoryReq model.MiniposReceiptHistoryReq
	err := ctx.ReadJSON(&miniposReceiptHistoryReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.UpdateReceiptStatus(miniposReceiptHistoryReq, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func GetMiniPOSReport(ctx iris.Context) {
	var MiniposReportReq model.MiniposReportReq
	err := ctx.ReadJSON(&MiniposReportReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.GetMiniPOSReport(MiniposReportReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
