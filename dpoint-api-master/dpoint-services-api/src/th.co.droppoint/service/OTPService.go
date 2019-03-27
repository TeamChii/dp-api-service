package service

import (
	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
		"th.co.droppoint/service/authentication"
	*/

	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func AddAuthVerify(code string, authen_type string) map[string]interface{} {
	token := utils.RandString(40)
	now := time.Now()

	authVerify := &entity.AuthVerifyEntity{
		Auth_token:   token,
		Auth_code:    code,
		Auth_type:    authen_type,
		Expire_time:  5,
		Created_time: &now,
	}

	messageCode := "S"
	statusCode := 1
	messageAbbr := "Send " + authen_type + " Success"
	messageDesc := ""
	db := config.DBsql()
	err := db.Create(&authVerify).Error
	if err != nil {
		messageCode = "E"
		statusCode = 0
		messageAbbr = "Send " + authen_type + " Error"
		//return "", ""
	} else {
		messageDesc = token
		if authen_type == "CUSTOMER_QR" {
			messageDesc = messageDesc + "_" + code
		}
		//return token,code;
	}
	return map[string]interface{}{
		"statusCode":  statusCode,
		"messageCode": messageCode,
		"messageAbbr": messageAbbr,
		"messageDesc": messageDesc}
}

func VerifyByCode(token string, code string, authen_type string) map[string]interface{} {
	messageCode := "S"
	statusCode := 1
	//messageAbbr := "valid OTP Code"
	messageAbbr := "valid " + authen_type + " Code"
	messageDesc := ""

	var data entity.AuthVerifyEntity
	check := config.DBsql().Where("auth_token = ? AND auth_code = ? ", token, code).Find(&data).RecordNotFound()

	if check == false {
		expire_time := data.Expire_time
		create_time := data.Created_time

		var pastDate = time.Date(create_time.Year(), create_time.Month(), create_time.Day(), create_time.Hour(),
			create_time.Minute(), create_time.Second(), 0, time.UTC)

		//fmt.Println(pastDate)
		now := time.Now() //.In(loc)

		var nowDate = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(),
			now.Minute(), now.Second(), 0, time.UTC)
		//fmt.Println(nowDate)
		duration := nowDate.Sub(pastDate)
		expire_hr := int(duration.Hours())
		expire_minute := int(duration.Minutes())
		//fmt.Println(expire_hr)
		//fmt.Printf("Diffrence in Hours : %d Hours\n", expire_hr)
		//fmt.Println(expire_minute)
		if !(expire_hr == 0 && int(expire_minute)-expire_time <= 0) {
			messageCode = "E"
			statusCode = 0
			messageAbbr = "Expired"
		} else {
			if authen_type == "QR_VERIFY" {
				var data model.CustomerResp
				check := config.DBsql().Where("cust_id = ? ", code).Find(&data).RecordNotFound()
				if check == true {
					messageCode = "E"
					statusCode = 0
					messageAbbr = "Invalid " + authen_type + " Code"
				} else if data.Cust_mobile == "" || data.Country_code == "" {
					messageCode = "E"
					statusCode = 0
					messageDesc = "Invalid Mobile Number"
				} else {
					messageDesc = data.Country_code + "_" + data.Cust_mobile
				}
				//check := config.DBsql().Where("auth_token = ? AND auth_code = ? ", token, code).Find(&data).RecordNotFound()
			}
		}
	} else {
		messageCode = "E"
		statusCode = 0
		messageAbbr = "Invalid " + authen_type + " Code"
	}
	return map[string]interface{}{
		"statusCode":  statusCode,
		"messageCode": messageCode,
		"messageAbbr": messageAbbr,
		"messageDesc": messageDesc}
}

func FileMessageForSend(category string) string {
	message := ""

	var data entity.SystemParamEntity
	check := config.DBsql().Where("category_name = ? AND key_code = ? ", "OTP_MSG", category).Find(&data).RecordNotFound()

	if check == false {
		message = data.Key_value
	} else {
		message = "SMS OTP for transaction approval on DropPoint \r<OTP: xxxxxx> will be expried in 5 minutes"
	}
	return message
}
