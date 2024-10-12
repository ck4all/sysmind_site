package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers/masterdata"
	"goravel/app/http/controllers/utility"
	"goravel/app/http/controllers/webmpp"
	"goravel/app/http/middleware"
	"goravel/app/services"

	"goravel/app/http/controllers"
)

func Api() {
	//declare package service variable
	authService := services.NewAuthService()
	appInfoService := services.NewAppInfoService()
	fileMediaService := services.NewFileManagementService()
	userService := services.NewUserService()
	fileManagementService := services.NewFileManagementService()
	menuService := services.NewMenuService()
	categoryService := services.NewCategoryService()

	//declare package controller variable
	authController := controllers.NewAuthController(authService)
	appInfoController := masterdata.NewAppInfoController(appInfoService)
	mediaController := utility.NewMediaController(fileMediaService)
	userController := utility.NewUserController(userService)
	fileManagementController := utility.NewFileManagementController(fileManagementService)
	menuController := controllers.NewMenuController(menuService)
	categoryController := webmpp.NewCategoryController(categoryService)

	//create route here
	facades.Route().Prefix("api/v1/").Group(func(router route.Router) {
		//landingpage route
		router.Get("app-info", appInfoController.Index)

		//auth router page
		router.Prefix("auth").Group(func(router route.Router) {
			router.Post("register", authController.Register)
			router.Post("login", authController.Login)
			router.Post("reset-password", authController.Reset)
			router.Middleware(middleware.Auth()).Post("logout", authController.Logout)
			router.Middleware(middleware.Auth()).Get("user-info", authController.UserInfo)
			router.Middleware(middleware.Auth()).Post("change-password", authController.ChangePassword)
			router.Middleware(middleware.Auth()).Get("menus", menuController.Index)
		})

		//media route
		router.Prefix("media").Group(func(router route.Router) {
			router.Get("show-file/{doctype}/{filename}", mediaController.ShowFile)
			router.Middleware(middleware.Auth()).Post("upload", mediaController.Upload)
		})

		//router masterdata
		router.Middleware(middleware.Auth()).Prefix("master-data").Group(func(router route.Router) {
			//route application information
			router.Get("app-info", appInfoController.Index)
			router.Put("app-info/{id}", appInfoController.Update)
		})

		//Router WebMpp
		router.Middleware(middleware.Auth()).Prefix("webmpp").Group(func(router route.Router) {
			//route category
			router.Get("category", categoryController.Index)
			router.Post("category", categoryController.Store)
			router.Get("category/{id}", categoryController.Show)
			router.Put("category/{id}", categoryController.Update)
			router.Delete("category/{id}", categoryController.Destroy)
		})

		//Route Utility
		router.Middleware(middleware.Auth()).Prefix("utility").Group(func(router route.Router) {
			//Route User Management
			router.Get("users", userController.Index)
			router.Post("users", userController.Store)
			router.Get("users/{id}", userController.Show)
			router.Put("users/{id}", userController.Update)
			router.Delete("users/{id}", userController.Destroy)
			router.Post("user-update-profile", userController.UpdateProfile)

			//Route FileManagement
			router.Get("file-managements", fileManagementController.Index)

		})

	})

}
