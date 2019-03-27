package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func GenJWT(ctx iris.Context) {

	/*var Json map[string]interface{}
	ctx.ReadJSON(&Json)
	var user Model.User
	var count float64
	config.DBsql().
		Model(&Model.User{}).
		Where("user_name = ? AND password = ?", Json["username"], HashPassword(Json["password"].(string))).
		Count(&count).
		Find(&user)
	if count > 0 {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": Json["username"],
			"password": Json["password"],
		})

		// Sign and get the complete encoded token as a string using the secret
		mySigningKey := []byte("mc-learning")

		tokenString, _ := token.SignedString(mySigningKey)

		user.Password = ""

		var dataToJson map[string]interface{}
		b, _ := json.Marshal(user)
		json.Unmarshal(b, &dataToJson)

		ctx.JSON(map[string]interface{}{"statusCode": "S", "messageCode": "S003", "messageAbbr": "Success", "token": tokenString, "data": dataToJson})
	} else {
		ctx.JSON(map[string]string{"statusCode": "E", "messageCode": "E003", "messageAbbr": "Error", "messageDesc": "username or password not found"})
	}*/

}
func Decode(tokenString string) interface{} {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("mc-learning"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_phone"].(string)

	} else {
		return err
	}

}
