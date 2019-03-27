package controller

import (
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"

	"github.com/kataras/iris"
)

func LoadMerchantCampaignMap(ctx iris.Context) {
	var MerchantCampaignMapReq model.MerchantCampaignMapReq
	err := ctx.ReadJSON(&MerchantCampaignMapReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadMerchantCampaignMap(MerchantCampaignMapReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func SendCardPromotion(ctx iris.Context) {
	var MerchantCustomerCampaignMapAddReq model.MerchantCustomerCampaignMapAddReq
	err := ctx.ReadJSON(&MerchantCustomerCampaignMapAddReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SendCardPromotion(MerchantCustomerCampaignMapAddReq, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func LoadMerchantCampaignMapById(ctx iris.Context) {
	var MerchantCampaignMapReq2 model.MerchantCampaignMapReq2
	err := ctx.ReadJSON(&MerchantCampaignMapReq2)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadMerchantCampaignMapById(MerchantCampaignMapReq2)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func SetFavourite(ctx iris.Context) {
	var MerchantCampaignMapUpadteReq model.MerchantCampaignMapUpadteReq
	err := ctx.ReadJSON(&MerchantCampaignMapUpadteReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SetFavourite(MerchantCampaignMapUpadteReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func UnSetFavourite(ctx iris.Context) {
	var MerchantCampaignMapUpadteReq model.MerchantCampaignMapUpadteReq
	err := ctx.ReadJSON(&MerchantCampaignMapUpadteReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.UnSetFavourite(MerchantCampaignMapUpadteReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func LoadCustoemrInCampaignMap(ctx iris.Context) {
	var MerchantCustomerCampaignMapLoadCustReq model.MerchantCustomerCampaignMapLoadCustReq
	err := ctx.ReadJSON(&MerchantCustomerCampaignMapLoadCustReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.LoadCustoemrInCampaignMap(MerchantCustomerCampaignMapLoadCustReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
