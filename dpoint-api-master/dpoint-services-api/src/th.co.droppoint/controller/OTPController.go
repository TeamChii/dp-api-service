package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func SendOTPCode(ctx iris.Context) {
	mobile := ctx.Params().Get("mobile")
	category := ctx.Params().Get("category")
	message := service.FileMessageForSend(category)
	//fmt.Println(category)
	code := utils.RandSeq(6)
	message = strings.Replace(message, "xxxxxx", code, -1)
	/* */
	resp, err := http.PostForm("http://www.thaibulksms.com/sms_api.php",
		url.Values{
			"username": {"0818904560"},
			"password": {"029262"},
			"msisdn":   {mobile},
			//"message":  {"SMS OTP for transaction approval on DropPoint \r<OTP: " + code + "> will be expried in 5 minutes"},
			"message": {message},
			"sender":  {"NOTICE"},
			"force":   {"standard"},
		})
	/* */
	/*
		    <SMS>
		    <QUEUE>
		        <Msisdn>0966408458</Msisdn>
		        <Status>1</Status>
		        <Transaction>23a7e0c3bef949be927f01b70852727c</Transaction>
		        <UsedCredit>1</UsedCredit>
		        <RemainCredit>1532</RemainCredit>
		    </QUEUE>
		    <!-- 0.015279054641724 -->
			</SMS>
	*/
	/* */
	if nil != err {
		fmt.Println("errorination happened getting the response", err)
		return
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return
	}
	/* */
	var a = service.AddAuthVerify(code, "SMS")

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func VerifyOTPCode(ctx iris.Context) {
	token := ctx.Params().Get("token")
	code := ctx.Params().Get("code")
	var a = service.VerifyByCode(token, code, "OTP")
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func QRGenCode(ctx iris.Context) {
	code := ctx.Params().Get("customerId")
	/* */
	var a = service.AddAuthVerify(code, "CUSTOMER_QR")

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func QRVerifyOTPCode(ctx iris.Context) {
	token := ctx.Params().Get("token")
	//code := ctx.Params().Get("code")
	s := strings.Split(token, "_")
	token = s[0]
	var code = s[1]
	var a = service.VerifyByCode(token, code, "QR_VERIFY")
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
