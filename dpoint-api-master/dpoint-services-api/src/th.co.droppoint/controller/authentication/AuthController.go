package authentication

import (
	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
		"th.co.droppoint/service/authentication"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service/authentication"
	"github.com/kataras/iris"
)

/*
func AuthMobileAndRegister(c *gin.Context)  {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		var a = authentication.AuthMobileAndRegister(user)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	}

}

func AuthUserSocialAndRegister(c *gin.Context)  {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		var a = authentication.AuthUserSocialAndRegister(user)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	}

}

func AuthAccount(c *gin.Context)  {
	var account model.Account
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		var a = authentication.AuthAccount(account)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	}

}

func SignUpAccount(c *gin.Context)  {
	var account model.Account
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		var a = authentication.SignUpAccount(account)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	}

}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorizationHeader := c.GetHeader("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}
					return []byte(utils.TOKEN_SECRET_KEY), nil
				})
				if error != nil {
					c.JSON(200, gin.H{"errcode": 400, "data":utils.SetErrReturn("105", error.Error())})
					c.AbortWithStatus(401)
					return
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					context.Set(c.Request, "id", claims["id"].(string))
					context.Set(c.Request, "username", claims["username"].(string))
					context.Set(c.Request, "type", claims["type"].(string))
					c.Next()
				} else {
					c.JSON(200, gin.H{"errcode": 400, "data": utils.SetErrReturn("103", "Invalid authorization token")})
					c.AbortWithStatus(401)
				}
			} else {
				c.JSON(200, gin.H{"errcode": 400, "data": utils.SetErrReturn("103", "Invalid authorization token")})
				c.AbortWithStatus(401)
			}
		} else {
			c.JSON(200, gin.H{"errcode": 400, "data":utils.SetErrReturn("102", "An authorization header is required")})
			c.AbortWithStatus(401)
		}
	}
}

func RefreshToken(user model.User) dto.AuthDTO  {

	var authDto dto.AuthDTO

	var userF model.User


	if config.GetDB().Where("username = ?", user.Username).Find(&userF).RecordNotFound() {
		authDto.StatusCode = "1"

		var alertMsg dto.AlertMessage
		alertMsg.MessageCode = "404"
		alertMsg.MessageDesc = "User not found"
		var errList []dto.AlertMessage
		errList= append(errList, alertMsg)
		authDto.ErrorList = errList

	}else{

		if user.Password == userF.Password{
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": userF.Username,
				"password": userF.Password,
			})
			tokenString, error := token.SignedString([]byte(utils.TOKEN_SECRET_KEY))
			if error != nil {
				fmt.Println(error)
			}

			authDto.StatusCode = "0"
			var authModel model.Auth
			var authInfoModel model.AuthInfo
			authInfoModel.Username = userF.Username
			authModel.AuthInfo = authInfoModel
			authModel.Authorization = tokenString

			var dataList []model.Auth
			dataList= append(dataList, authModel)
			authDto.Data = dataList

		}else{
			authDto.StatusCode = "1"

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

*/

func FirstLogin(ctx iris.Context) {
	var PinAuthenReq model.PinAuthenReq
	err := ctx.ReadJSON(&PinAuthenReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = authentication.Auth(PinAuthenReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func SecondLogin(ctx iris.Context) {
	var AuthenSecReq model.AuthenSecReq
	err := ctx.ReadJSON(&AuthenSecReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = authentication.AuthSecond(AuthenSecReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
func CheckPin(ctx iris.Context) {
	var PinAuthenReq model.PinAuthenReq
	err := ctx.ReadJSON(&PinAuthenReq)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = authentication.CheckPin(PinAuthenReq)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
}
