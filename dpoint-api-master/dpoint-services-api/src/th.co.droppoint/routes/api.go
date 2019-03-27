package routes

import (
	"os"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/controller"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/controller/authentication"
	"github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	log "github.com/sirupsen/logrus"
	//"../ControllerMysql"
	//"../Util"
)

func GetRoute() *iris.Application {
	app := iris.New()
	//app.Use(recover.New())
	//app.Use(logger.New())
	log_level := os.Getenv("LOG_LEVEL")
	///fmt.Println("LOG_LEVEL:", log_level)
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

	/*	api1 := app.Party("/api/authen")
		{
			api1.Post("/login", authentication.FirstLogin)
			//api1.Post("/login-learner", utils.GenJWTLeanrner)
		}
	*/
	api2 := app.Party("/api/merchant", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)

	api2.Use(jwtHandler.Serve)
	{

		api2.Get("/id/{id:int}", controller.MerchantById)
		api2.Get("/userid/{userid:string}", controller.MerchantByPhone)
		api2.Post("/add", controller.AddMerchant)
		api2.Post("/update", controller.UpdateMerchant)
		api2.Post("/set-head-office", controller.SetHeadOfficeMerchant)

	}

	api3 := app.Party("/api/container", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api3.Use(jwtHandler.Serve)
	{

		api3.Post("/load-by-mc-id", controller.ContainerByIdMc)
		//api3.Post("/load-by-mc-cust", controller.ContainerByIdMcCust)
		api3.Post("/load-id", controller.ContainerById)
		api3.Get("/id/{id:string}", controller.ContainerId)
		api3.Post("/add", controller.AddContainer)
		api3.Post("/add-mc", controller.AddContainerMcOnly)
		api3.Post("/update", controller.UpdateContainer)

	}

	api4 := app.Party("/api/container-type", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api4.Use(jwtHandler.Serve)
	{
		api4.Post("/list", controller.ContainerTypeSearch)
		api4.Get("/id/{id:int}", controller.ContainerTypeById)
		api4.Post("/add", controller.AddContainerType)
		api4.Post("/update", controller.UpdateContainerType)

	}

	api5 := app.Party("/api/give-point", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api5.Use(jwtHandler.Serve)
	{
		api5.Post("/load-by-mc-id", controller.PointByIdMc)
		api5.Post("/add", controller.AddPoint)
		api5.Post("/update", controller.UpdatePoint)
		api5.Post("/check-customer", controller.CheckMobileCustomer)

	}

	api6 := app.Party("/api/customer", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api6.Use(jwtHandler.Serve)
	{
		api6.Post("/load-by-mobile", controller.GetCustomerByMobile)
		api6.Get("/id/{id:int}", controller.GetCustomerById)
		api6.Post("/update", controller.UpdateCustomer)

	}

	api7 := app.Party("/api/merchant-map-customer", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api7.Use(jwtHandler.Serve)
	{
		api7.Post("/load", controller.MerchantCustomerMapByIdMcCust)
	}

	api8 := app.Party("/api/customer-map-point", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api8.Use(jwtHandler.Serve)
	{
		api8.Post("/load", controller.CustomerPointMapByIdMcCust)
	}

	api9 := app.Party("/api/redeem", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api9.Use(jwtHandler.Serve)
	{
		api9.Post("/add", controller.AddRedeem)
	}

	api10 := app.Party("/api/special-request", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api10.Use(jwtHandler.Serve)
	{
		api10.Post("/list", controller.LoadRequestPoint)
		api10.Post("/send-card-all", controller.SendCardAll)
		api10.Post("/send-card", controller.SendCard)
	}
	api11 := app.Party("/api/authen", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	{
		api11.Post("/login", authentication.FirstLogin)
		api11.Post("/login-staff", authentication.SecondLogin)
		api11.Post("/pin", authentication.CheckPin)
	}
	/*
		// for forward inline content
		proxy_inline := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: utils.CONTENT_SCHEMA,
			//Host:   "stackoverflow.com",
			Host: utils.CONTENT_HOST,
		})
		proxy_inline.Director = controller.InlineContent
	*/
	api12 := app.Party("/api/content", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api12.Use(jwtHandler.Serve)
	{
		api12.Get("/get/{content_id:int}", controller.ContentById)       // ok
		api12.Get("/remove/{content_id:int}", controller.RemoveContent)  // ok
		api12.Post("/upload", controller.SaveContent)                    //  ok remain manage remove file
		api12.Post("/update/{content_id:int}", controller.UpdateContent) //  ok remain manage remove file

		api12.Get("/inline/{content_id:int}", controller.RedirectToContent)
	}

	api13 := app.Party("/api/otp", crs).AllowMethods(iris.MethodOptions)

	//api13.Use(jwtHandler.Serve)
	{
		api13.Get("/send/{mobile:string}/{category:string}", controller.SendOTPCode) // ok
		api13.Get("/verify/{token:string}/{code:string}", controller.VerifyOTPCode)  // ok

		api13.Get("/qr/gen/{customerId:string}", controller.QRGenCode)     // ok
		api13.Get("/qr/verify/{token:string}", controller.QRVerifyOTPCode) // ok
		//api13.Get("/verify/{code:string}", controller.VerifyOTPCode)
	}

	api14 := app.Party("/api/system-param", crs).AllowMethods(iris.MethodOptions)

	api14.Use(jwtHandler.Serve)
	{
		api14.Post("/load", controller.SystemParamSearch)
		api14.Post("/add", controller.SystemParamAdd)
		api14.Post("/update", controller.SystemParamUpdate)
		api14.Post("/delete", controller.SystemParamDelete)

	}

	api15 := app.Party("/api/promotion", crs).AllowMethods(iris.MethodOptions)

	api15.Use(jwtHandler.Serve)
	{
		api15.Post("/list", controller.LoadMerchantCampaignMap)
		api15.Post("/send-card", controller.SendCardPromotion)
		api15.Post("/load-by-id", controller.LoadMerchantCampaignMapById)
		api15.Post("/set-favourite", controller.SetFavourite)
		api15.Post("/unset-favourite", controller.UnSetFavourite)
		api15.Post("/list-customer", controller.LoadCustoemrInCampaignMap)

	}

	api16 := app.Party("/api/user", crs).AllowMethods(iris.MethodOptions)

	api16.Use(jwtHandler.Serve)
	{
		api16.Get("/id/{id:int}", controller.UserById)
		api16.Get("/check/{phone:string}", controller.UserCheck)
		api16.Post("/update", controller.UpdateUser)
		api16.Post("/update-pin", controller.UpdateUserPin)
		api16.Post("/add", controller.AddUser)
		api16.Post("/add-staff", controller.AddUserStaff)
		api16.Post("/load-by-role", controller.UserByRole)
		api16.Post("/delete", controller.DeleteUser)
		api16.Post("/load-cat", controller.UserLoadCategory)

	}

	api17 := app.Party("/api/package", crs).AllowMethods(iris.MethodOptions)

	api17.Use(jwtHandler.Serve)
	{
		api17.Post("/load", controller.PackageLoad)
		api17.Post("/add", controller.PackageAdd)
		api17.Post("/update", controller.PackageUpdate)
		api17.Post("/delete", controller.PackageDelete)

	}

	api18 := app.Party("/api/merchant-non", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	//api2.Use(jwtHandler.Serve)
	{

		api18.Post("/add", controller.AddMerchant2)

	}

	api19 := app.Party("/api/merchant-image", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api19.Use(jwtHandler.Serve)
	{

		api19.Post("/add", controller.AddMerchantImage)
		api19.Post("/delete", controller.DeleteMerchantImage)

	}

	api20 := app.Party("/api/reports", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api20.Use(jwtHandler.Serve)
	{
		api20.Post("/dashboard", controller.LoadDashboard)
		api20.Get("/report-list", controller.LoadReportList)
		api20.Post("/report-by-cat", controller.LoadReportByCategory)
		api20.Post("/campaign-monitor", controller.LoadCampaignMonitor)
		api20.Post("/crm", controller.GetCRMReport)
	}

	api21 := app.Party("/api/devices", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api21.Use(jwtHandler.Serve)
	{
		api21.Get("/id/{id:int}", controller.LoadDevice)
		api21.Post("/add", controller.AddDevice)
		api21.Post("/update", controller.UpdateDevice)
		api21.Post("/delete", controller.DeleteDevice)
		api21.Post("/check", controller.CheckUidDevice)

	}

	api22 := app.Party("/api/notification", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api22.Use(jwtHandler.Serve)
	{
		api22.Post("/load", controller.NotiLoad)
		api22.Get("/load/category/{id:int}", controller.NotiCategory)
		api22.Get("/id/{id:int}", controller.NotiById)
		api22.Post("/read", controller.NotiRead)
		api22.Post("/delete", controller.NotiDelete)
	}

	api23 := app.Party("/api/minipos", crs).AllowMethods(iris.MethodOptions)
	//api.Use(crs)
	//handler = crs.Handler(handler)
	api23.Use(jwtHandler.Serve)
	{
		api23.Post("/load/category-menu", controller.LoadCategoryMenu)
		api23.Post("/load/menu", controller.LoadMenuByCategory)
		api23.Post("/load/active/menu", controller.LoadActiveMenuByCategory)
		api23.Post("/load/all-menu", controller.LoadAllMenu)
		api23.Post("/add/transaction-receive", controller.AddTransactionAndReceive)

		api23.Get("/id/category-menu/{id:int}", controller.GetMiniPOSCategoryById)
		api23.Post("/add/category-menu", controller.AddMiniPOSCategory)
		api23.Post("/update/category-menu", controller.UpdateMiniPOSCategory)
		api23.Post("/delete/category-menu", controller.DeleteMiniPOSCategory)

		api23.Get("/id/menu/{id:int}", controller.GetMiniPOSMenuById)
		api23.Post("/add/menu", controller.AddMiniPOSMenu)
		api23.Post("/update/menu", controller.UpdateMiniPOSMenu)
		api23.Post("/delete/menu", controller.DeleteMiniPOSMenu)

		api23.Post("/add/mapping-menu", controller.AddMiniposMappingMenu)
		api23.Post("/delete/mapping-menu", controller.DeleteMiniposMappingMenu)

		api23.Get("/id/setting/{id:int}", controller.GetMiniPOSSettingById)
		api23.Post("/add/setting", controller.AddMiniPOSSetting)
		api23.Post("/update/setting", controller.UpdateMiniPOSSetting)

		api23.Get("/id/thai-prompt-pay/{id:int}", controller.GetMiniposThaiPromptPayById)
		api23.Post("/add/thai-prompt-pay", controller.AddMiniposThaiPromptPay)
		api23.Post("/update/thai-prompt-pay", controller.UpdateMiniposThaiPromptPay)

		api23.Get("/id/give-reward-setting/{id:int}", controller.GetMiniposGiveRewardSettingById)
		api23.Post("/add/give-reward-setting", controller.AddMiniposGiveRewardSetting)
		api23.Post("/update/give-reward-setting", controller.UpdateMiniposGiveRewardSetting)

		api23.Post("/search/receipt-history", controller.SearchReceiptHistory)
		api23.Post("/update/receipt", controller.UpdateReceiptStatus)
		api23.Post("/report", controller.GetMiniPOSReport)

	}

	api24 := app.Party("/api/user-non", crs).AllowMethods(iris.MethodOptions)

	{
		api24.Get("/check/{phone:string}", controller.UserCheck)
		api24.Post("/update-pin", controller.UpdateUserPin)

	}

	api25 := app.Party("/api/devices-non", crs).AllowMethods(iris.MethodOptions)

	{
		api25.Post("/check", controller.CheckUidDevice)

	}

	api26 := app.Party("/api/system-param-non", crs).AllowMethods(iris.MethodOptions)

	{
		api26.Post("/load", controller.SystemParamSearch)

	}

	api27 := app.Party("/api/customer-non", crs).AllowMethods(iris.MethodOptions)

	{
		api27.Get("/check/{phone:string}", controller.PointCheck)

	}
	api28 := app.Party("api/increas-point", crs).AllowMethods(iris.MethodOptions)

	// api28.Use(jwtHandler.Serve)
	{

		api28.Post("/update", controller.IncreasPoint)
	}

	return app
}
