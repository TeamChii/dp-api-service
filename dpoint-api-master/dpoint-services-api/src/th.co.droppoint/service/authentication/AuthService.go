package authentication

import (
	"encoding/json"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
		"th.co.droppoint/service/authentication"
	*/)

/*
func AuthMobileAndRegister(user model.User) dto.AuthDTO  {

	var authDto dto.AuthDTO
	var successStatus = false
	var userF model.User

	if user.UserIdMobile != ""{  // auth mobile to return token
		nf := config.GetDB().Where("user_id = ?", user.UserIdMobile).Find(&userF).RecordNotFound()
		if nf == false {
			successStatus = true
		}else{
			successStatus = false
		}
	}else{ // register with new mobile

		userF.UserId = utils.RandSeq(20)
		userF.UserType = "mobile"
		userF.MobileUUID = user.MobileUUID
		userF.Username = userF.UserId
		userF.Email = userF.UserId
		userF.CreatedBy = userF.UserId
		userF.CreatedDate = time.Now()
		activeFlag := utils.ACTIVE_FLAG
		userF.Active = &activeFlag

		err := config.GetDB().Create(&userF).Error

		if err != nil {
			successStatus = false
			log.Error("Error register mobile : ", err.Error())
		}else{
			successStatus = true
		}
	}

	if successStatus {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": userF.UserId,
			"username": userF.UserId,
			"type": "mobile",
		})
		tokenString, error := token.SignedString([]byte(utils.TOKEN_SECRET_KEY))
		if error != nil {
			log.Error(error)
		}

		authDto.StatusCode = "S"
		var authModel model.Auth
		var authInfoModel model.AuthInfo
		authInfoModel.ID = string(userF.UserId)
		authInfoModel.Username = userF.UserId
		authInfoModel.SignInType = "mobile"
		authModel.AuthInfo = authInfoModel
		authModel.Authorization = tokenString

		var dataList []model.Auth
		dataList= append(dataList, authModel)
		authDto.Data = dataList
	}else{
		authDto.StatusCode = "E"

		var alertMsg dto.AlertMessage
		alertMsg.MessageCode = "403"
		alertMsg.MessageDesc = "Can not register this mobile"
		var errList []dto.AlertMessage
		errList= append(errList, alertMsg)
		authDto.ErrorList = errList
	}

	return authDto

}

func checkForceSignIn(user model.User, mode string, signUserId string) bool {

	var status = true

	if user.ForceSignIn == true || user.UserIdMobile == "" {
		status = true
	}else{

		var userMappingList []model.UserMapping
		config.GetDB().Where("user_id_mobile = ?", user.UserIdMobile).Find(&userMappingList).RecordNotFound()

		if mode == "new" {
			if len(userMappingList) == 0 {
				status = true
			}else{
				status = false
			}
		}else{  // old user condition
			if len(userMappingList) == 0 {
				status = true
			}else{
				for i:=0 ; i<len(userMappingList) ; i++{
					if userMappingList[i].UserIdUser == signUserId{
						status = true
						break
					}else{
						status = false
					}
				}
			}
		}
	}

	return status
}

func AuthUserSocialAndRegister(user model.User) dto.AuthDTO  {

	var authDto dto.AuthDTO
	var successStatus = false
	var errMsg = ""
	var userF model.User
	nf := config.GetDB().Where("email = ?", user.Email).Find(&userF).RecordNotFound()
	if nf == true { // found condition

		var forceSignInStatus = checkForceSignIn(user, "new", "") // new = user
		if forceSignInStatus {
			// register if not found
			userF.UserId = utils.RandSeq(20)
			userF.UserType = "user"
			userF.Username = user.Email
			userF.Email = user.Email
			userF.FirstName = user.FirstName
			userF.LastName = user.LastName
			userF.Image = user.Image
			userF.Connected = user.Connected
			userF.ConnectedId = user.ConnectedId
			userF.CreatedBy = userF.UserId
			userF.CreatedDate = time.Now()
			activeFlag := utils.ACTIVE_FLAG
			userF.Active = &activeFlag

			tx := config.GetDB().Begin()
			err := tx.Create(&userF).Error

			if err == nil {
				if user.UserIdMobile != "" {
					// copy card in current mobile to this user account

					var budUserList []model.BudUser

					notFound := config.GetDB().Where("user_id = ?", user.UserIdMobile).Find(&budUserList).RecordNotFound()

					if notFound == false {

						var mapStatus = true

						for i := 0; i < len(budUserList); i++ {

							var budUser model.BudUser
							budUser = budUserList[i]
							budUser.UserId = userF.UserId

							errMapping := tx.Where(model.BudUser{UserId: budUser.UserId, BudId: budUser.BudId}).FirstOrCreate(&budUser).Error  // init if not found
							if errMapping != nil {
								mapStatus = false
								tx.Rollback()

								log.Error("Error copy bud to user : ", err.Error())
								break
							} else {
								mapStatus = true
							}
						}

						if mapStatus {
							successStatus = true
						}
					}

					// Map user mobile to user user
					var userMapping model.UserMapping
					userMapping.UserIdUser = userF.UserId
					userMapping.UserIdMobile = user.UserIdMobile
					errMapping := tx.FirstOrCreate(&userMapping).Error
					if errMapping != nil {
						tx.Rollback()

						successStatus = false
						errMsg = "Error mapping account "+ errMapping.Error()
					}else{
						successStatus = true
						tx.Commit()

					}

				}else{
					successStatus = true
					tx.Commit()

				}
			}else {
				tx.Rollback()

				successStatus = false
				errMsg = "Can not sign in with this account "+ err.Error()
			}
		}else{
			successStatus = false
			errMsg = "This mobile is connect to other account"
		}

	}else{ // found condition

		var forceSignInStatus = checkForceSignIn(user, "old", userF.UserId)
		if forceSignInStatus {
			successStatus = true
			tx := config.GetDB().Begin()

			if user.UserIdMobile != "" {
				// Merge card to each other (mobile and user)

				var budUserList []model.BudUser

				notFound := config.GetDB().Where("user_id in (?)", []string{user.UserIdMobile, userF.UserId}).Find(&budUserList).RecordNotFound()

				if notFound == false{
					var mapStatus = true
					for i := 0; i < len(budUserList); i++ {
						var budUser model.BudUser
						budUser = budUserList[i]

						// copy mobile'bud to user
						budUser.UserId = userF.UserId
						errMapping := tx.Where(model.BudUser{UserId: budUser.UserId, BudId: budUser.BudId}).FirstOrCreate(&budUser).Error  // init if not found
						if errMapping != nil {
							mapStatus = false
							tx.Rollback()

							errMsg = "Error copy mobile card to user: "+ errMapping.Error()
							log.Error("Error copy mobile card to user : ", errMapping.Error())
							break
						} else {
							mapStatus = true
						}

						// copy user'bud to mobile
						budUser.UserId = user.UserIdMobile
						errMapping1 := tx.Where(model.BudUser{UserId: budUser.UserId, BudId: budUser.BudId}).FirstOrCreate(&budUser).Error  // init if not found
						if errMapping1 != nil {
							mapStatus = false
							tx.Rollback()

							errMsg = "Error copy user card to mobile : "+ errMapping1.Error()
							log.Error("Error copy user card to mobile : ", errMapping1.Error())
							break
						} else {
							mapStatus = true
						}
					}

					if mapStatus {
						successStatus = true
					}
				}

				// Map user mobile to user user
				var userMapping model.UserMapping
				userMapping.UserIdUser = userF.UserId
				userMapping.UserIdMobile = user.UserIdMobile
				errMapping := tx.FirstOrCreate(&userMapping).Error
				if errMapping != nil {
					tx.Rollback()

					successStatus = false
					errMsg = "Error mapping account "+ errMapping.Error()
				}else{
					successStatus = true
					tx.Commit()

				}

			}else{
				tx.Commit()

			}
		}else{
			successStatus = false
			errMsg = "This mobile is connect to other account"
		}

	}

	if successStatus {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": userF.UserId,
			"username": userF.Username,
			"type": "user",
		})
		tokenString, error := token.SignedString([]byte(utils.TOKEN_SECRET_KEY))
		if error != nil {
			log.Error(error)
		}

		authDto.StatusCode = "S"
		var authModel model.Auth
		var authInfoModel model.AuthInfo
		authInfoModel.ID = string(userF.UserId)
		authInfoModel.Username = userF.Username
		authInfoModel.Name = userF.FirstName + " " + userF.LastName
		authInfoModel.Image = userF.Image
		authInfoModel.Connected = userF.Connected
		authInfoModel.SignInType = "user"
		authModel.AuthInfo = authInfoModel
		authModel.Authorization = tokenString

		var dataList []model.Auth
		dataList= append(dataList, authModel)
		authDto.Data = dataList
	}else{
		authDto.StatusCode = "E"

		var alertMsg dto.AlertMessage
		alertMsg.MessageCode = "403"
		alertMsg.MessageDesc = errMsg
		var errList []dto.AlertMessage
		errList= append(errList, alertMsg)
		authDto.ErrorList = errList
	}

	return authDto

}

func AuthAccount(account model.Account) dto.AuthDTO  {

	var authDto dto.AuthDTO

	var accountF model.Account


	if config.GetDB().Where("account_id_code = ? OR account_username = ?", account.AccountUsername, account.AccountUsername).Find(&accountF).RecordNotFound() {
		authDto.StatusCode = "E"

		var alertMsg dto.AlertMessage
		alertMsg.MessageCode = "404"
		alertMsg.MessageDesc = "Account not found"
		var errList []dto.AlertMessage
		errList= append(errList, alertMsg)
		authDto.ErrorList = errList

	}else{

		if CheckPasswordHash(account.AccountPassword , accountF.AccountPassword){
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id" : accountF.AccountId,
				"username": accountF.AccountUsername,
				"type": "account",
			})
			tokenString, error := token.SignedString([]byte(utils.TOKEN_SECRET_KEY))
			if error != nil {
				fmt.Println(error)
			}

			authDto.StatusCode = "S"
			var authModel model.Auth
			var authInfoModel model.AuthInfo
			authInfoModel.ID = accountF.AccountId
			authInfoModel.Username = accountF.AccountUsername
			authInfoModel.Name = accountF.AccountName
			authInfoModel.Image = accountF.AccountImage
			authInfoModel.SignInType = "account"
			authModel.AuthInfo = authInfoModel
			authModel.Authorization = tokenString

			var dataList []model.Auth
			dataList= append(dataList, authModel)
			authDto.Data = dataList
		}else{
			authDto.StatusCode = "E"

			var alertMsg dto.AlertMessage
			alertMsg.MessageCode = "403"
			alertMsg.MessageDesc = "Incorrect Password"
			var errList []dto.AlertMessage
			errList= append(errList, alertMsg)
			authDto.ErrorList = errList
		}
	}

	return authDto

}

func SignUpAccount(account model.Account) dto.AuthDTO  {
	var authDto dto.AuthDTO

	account.AccountId = utils.RandSeq(20)
	account.AccountQRCode = account.AccountId
	account.AccountIDCode = account.AccountId
	account.AccountPassword, _ = HashPassword(account.AccountPassword)
	account.CreatedBy = account.AccountId
	account.CreatedDate = time.Now()
	activeFlag := utils.ACTIVE_FLAG
	account.Active = &activeFlag
	err := config.GetDB().Create(&account).Error

	if err != nil{
		authDto.StatusCode = "E"

		var alertMsg dto.AlertMessage
		alertMsg.MessageCode = "500"
		alertMsg.MessageDesc = "Can not sign up account"
		var errList []dto.AlertMessage
		errList= append(errList, alertMsg)
		authDto.ErrorList = errList
	}else{
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id" : account.AccountId,
			"username": account.AccountUsername,
			"type": "account",
		})
		tokenString, error := token.SignedString([]byte(utils.TOKEN_SECRET_KEY))
		if error != nil {
			fmt.Println(error)
		}

		authDto.StatusCode = "S"
		var authModel model.Auth
		var authInfoModel model.AuthInfo
		authInfoModel.ID = account.AccountId
		authInfoModel.Username = account.AccountUsername
		authInfoModel.Name = account.AccountName
		authInfoModel.Image = account.AccountImage
		authInfoModel.SignInType = "account"
		authModel.AuthInfo = authInfoModel
		authModel.Authorization = tokenString

		var dataList []model.Auth
		dataList= append(dataList, authModel)
		authDto.Data = dataList
	}

	return authDto
}
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GenToken(user_phone string, pin string) string {
	pinEncrypt := utils.HashPin(pin)
	var data entity.UserEntity
	check := config.DBsql().
		Where("user_phone = ? AND pin = ?", user_phone, pinEncrypt).
		Find(&data).RecordNotFound()

	if check == false {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_phone": user_phone,
			"pin":        pinEncrypt,
		})

		// Sign and get the complete encoded token as a string using the secret
		mySigningKey := []byte("mc-learning")

		tokenString, _ := token.SignedString(mySigningKey)

		return tokenString
	} else {
		return ""
	}
}
func Auth(PinAuthenReq model.PinAuthenReq) map[string]interface{} {
	pinEncrypt := utils.HashPin(PinAuthenReq.Pin)
	var data entity.UserEntity
	check := config.DBsql().
		Where("user_phone = ? AND pin = ?", PinAuthenReq.User_Phone, pinEncrypt).
		Find(&data).RecordNotFound()

	if check == false {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_phone": PinAuthenReq.User_Phone,
			"pin":        pinEncrypt,
		})

		// Sign and get the complete encoded token as a string using the secret
		mySigningKey := []byte("mc-learning")

		tokenString, _ := token.SignedString(mySigningKey)

		var UserMerchantMapEntity []model.UserMerchantMapEntityResp
		config.DBsql().
			Where("user_id = ? ", data.User_id).
			Order("mc_id desc").
			Find(&UserMerchantMapEntity)

		for index := 0; index < len(UserMerchantMapEntity); index++ {
			var MerchantEntity model.MerchantEntityResp
			config.DBsql().
				Where("mc_id = ? ", UserMerchantMapEntity[index].Mc_id).
				Find(&MerchantEntity)

			var MerchantImageMapResp []model.MerchantImageMapResp

			config.DBsql().
				Table("dp_mp_merchant_image").Select("dp_tb_content.content_id, dp_tb_content.content_path").
				Joins("JOIN dp_tb_content ON dp_mp_merchant_image.content_id = dp_tb_content.content_id").
				Where("dp_mp_merchant_image.mc_id = ? ", UserMerchantMapEntity[index].Mc_id).
				Find(&MerchantImageMapResp)

			MerchantEntity.MerchantImageMapResp = MerchantImageMapResp

			//MerchantEntity.Image_ref = utils.CONTENT_URL+"/" + MerchantEntity.Image_ref + ""
			MerchantEntity.Logo_ref = utils.CONTENT_URL + "/" + MerchantEntity.Logo_ref + ""

			dataMer := MerchantEntity
			UserMerchantMapEntity[index].MerchantEntityResp = &dataMer

		}

		var dataToJson []map[string]interface{}
		b, _ := json.Marshal(UserMerchantMapEntity)
		json.Unmarshal(b, &dataToJson)

		return map[string]interface{}{
			"statusCode":  "S",
			"messageCode": "S003",
			"messageAbbr": "Success",
			"token":       tokenString,
			"data":        dataToJson,
		}
	} else {
		return map[string]interface{}{
			"statusCode":  "E",
			"messageCode": "E003",
			"messageAbbr": "Error",
			"messageDesc": "username or password not found",
		}
	}
}
func CheckPin(PinAuthenReq model.PinAuthenReq) map[string]interface{} {

	var data entity.UserEntity
	check := config.DBsql().
		Where("user_phone = ? AND pin = ?", PinAuthenReq.User_Phone, utils.HashPin(PinAuthenReq.Pin)).
		Find(&data).RecordNotFound()
	if check == false {
		return map[string]interface{}{
			"statusCode":  "S",
			"messageCode": "S003",
			"messageAbbr": "Success",
			"messageDesc": "Authen Success",
		}
	} else {
		return map[string]interface{}{
			"statusCode":  "E",
			"messageCode": "E003",
			"messageAbbr": "Error",
			"messageDesc": "username or password not found",
		}
	}
}
func AuthSecond(AuthenSecReq model.AuthenSecReq) map[string]interface{} {
	pinEncrypt := utils.HashPin(AuthenSecReq.Pin)
	var data model.UserAuthenResp
	check := config.DBsql().
		Raw(`select * 
	from dp_mp_user_merchant as usermapp  
	 join dp_ms_user as u 
			on usermapp.user_id=u.user_id
	 join dp_mp_device_merchant as device 
			on (usermapp.mc_id=device.mc_id)
	where u.staff_pin= ?
	and device.device_uid= ?;`, AuthenSecReq.Pin, AuthenSecReq.Device_uid).
		Scan(&data).RecordNotFound()

	data.Image_ref = utils.CONTENT_URL + "/" + data.Image_ref + ""
	//var dataDevice DeviceMerchantMapEntity
	var dataToJson map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	if check == false {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_phone": data.User_phone,
			"pin":        pinEncrypt,
		})

		// Sign and get the complete encoded token as a string using the secret
		mySigningKey := []byte("mc-learning")

		tokenString, _ := token.SignedString(mySigningKey)

		return map[string]interface{}{
			"statusCode":  "S",
			"messageCode": "S003",
			"messageAbbr": "Success",
			"token":       tokenString,
			"data":        dataToJson,
		}
	} else {
		return map[string]interface{}{
			"statusCode":  "E",
			"messageCode": "E003",
			"messageAbbr": "Error",
			"messageDesc": "username or password not found",
		}
	}
}
