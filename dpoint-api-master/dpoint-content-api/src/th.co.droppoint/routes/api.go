package routes

import (
	"fmt"
	"os"

	"th.co.droppoint/controller"
	"github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	log "github.com/sirupsen/logrus"
)

func GetRoute() *iris.Application {
	app := iris.New()
	//app.Use(recover.New())
	//app.Use(logger.New())
	log_level := os.Getenv("LOG_LEVEL")
	fmt.Println("LOG_LEVEL:", log_level)
	/*
		debug , info , warn , error , fatal , panic
	*/
	if log_level != "" {
		log.ParseLevel(log_level)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	/*
		if log_level != "error" {
			log.SetLevel(log.InfoLevel)

		} else {
			log.SetLevel(log.ErrorLevel)
		}
	*/
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("mc-learning"), nil
		},
		//Debug: true,

		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	//crs := cors.Default()
	//crs.Log = iris.Logger().Logger
	/* */
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // allows everything, use that to change the hosts.
		//AllowCredentials: true,
		AllowedMethods: []string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization", "X-Requested-With", "Application", "X-Z-Header", "X-Header"},
		//Debug:            true,
	})
	//cors.AllowAll()
	//app.AllowMethods(iris.MethodOptions)
	/* */
	//app.Use(crs) // not used here

	/*api1 := app.Party("/api//authen")
	{
		api1.Post("/login", utils.GenJWT)
		//api1.Post("/login-learner", utils.GenJWTLeanrner)
	}*/


	api := app.Party("/api/content", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api.Use(jwtHandler.Serve)
	{
		//api.Get("/get/{content_id:int}", controller.ContentById) // ok
		api.Get("/remove/{content_id:int}", controller.RemoveContent) // ok
		api.Post("/upload", controller.SaveContent) //  ok remain manage remove file
		api.Post("/update/{content_id:int}", controller.UpdateContent) //  ok remain manage remove file

		//api.Get("/inline/{content_id:int}",controller.RedirectToContent)
	}


	return app
}
