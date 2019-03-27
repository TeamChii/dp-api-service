package main

import (
	"os"

	//"./src/th.co.droppoint/config"
	//"./src/th.co.droppoint/routes"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/routes"
	"github.com/kataras/iris"
)

func main() {
	config.InitPostgres()
	app := routes.GetRoute()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}
	// test
	/*
		app.WrapRouter(cors.WrapNext(cors.Options{
			AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "X-Requested-With", "Application", "X-Z-Header", "X-Header"},
		}))
	*/
	app.Run(iris.Addr(":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}
