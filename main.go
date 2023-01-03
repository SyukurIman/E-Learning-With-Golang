package main

import (
	"e-learning/client"
	"e-learning/handler/api"
	"e-learning/handler/web"
	"e-learning/repository"
	"e-learning/service"
	"e-learning/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler  api.UserAPI
	TaskAPIHandler  api.TaskAPI
	AdminAPIHandler api.AdminAPI
}

type ClientHandler struct {
	AuthWeb        web.AuthWeb
	HomeWeb        web.HomeWeb
	DashboardUser  web.DashboardUser
	DashboardAdmin web.DashboardAdmin
}

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:123@localhost:5432/syukur")

	r := gin.New()
	err := utils.ConnectDB()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db := utils.GetDBConnection()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("views/**/*.tmpl")

	//Set Static Data
	r.Static("/assets", "./assets")
	r = RunServer(db, r)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func RunServer(db *gorm.DB, r *gin.Engine) *gin.Engine {
	// Repository Set
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	// Service Set
	userService := service.NewUserService(userRepo, taskRepo)
	adminService := service.NewAdminService(adminRepo)

	// API Set
	userAPIHandler := api.NewUserAPI(userService)
	adminAPIHandler := api.NewAdminAPI(adminService)
	apiHandler := APIHandler{
		UserAPIHandler:  userAPIHandler,
		AdminAPIHandler: adminAPIHandler,
	}

	// Client Set
	userClient := client.NewUserClient()
	adminClient := client.NewAdminClient()
	authWeb := web.NewAuthWeb(userClient, adminClient)
	homeWeb := web.NewHomeWeb()

	//User Set
	dashboardUser := web.NewDashboardUser(userRepo)

	// Admin Set
	dashboardAdmin := web.NewDashboardAdmin()

	client := ClientHandler{
		AuthWeb:        authWeb,
		HomeWeb:        homeWeb,
		DashboardUser:  dashboardUser,
		DashboardAdmin: dashboardAdmin,
	}

	// Route
	auth := r.Group("/")
	{
		auth.GET("/", client.HomeWeb.Index)
		auth.GET("/logout", client.AuthWeb.Logout)

		// User
		user := auth.Group("/user")
		{
			user.GET("/dashboard", utils.Authentication("/user/login"), client.DashboardUser.Dashboard)
			user.GET("/login", client.AuthWeb.Login)
			user.POST("/login/proses", client.AuthWeb.LoginProses)

			user.GET("/register", client.AuthWeb.Register)
			user.POST("/register/proses", client.AuthWeb.RegisterProses)
		}

		// Admin
		admin := auth.Group("/admin")
		{
			admin.GET("/login", client.AuthWeb.LoginAdmin)
			admin.POST("/login/proses", client.AuthWeb.LoginAdminProses)
			admin.GET("/dashboard", utils.Authentication("/admin/login"), client.DashboardAdmin.Dashboard)
		}

		// Bagian API
		api := auth.Group("/api")
		{
			// Set Routing User API
			userAPI := api.Group("/user")
			{
				userAPI.POST("/login", apiHandler.UserAPIHandler.Login)
				userAPI.POST("/register", apiHandler.UserAPIHandler.Register)
			}

			// Set Routing Admin API
			adminAPI := api.Group("/admin/v1")
			{
				adminAPI.POST("/login", apiHandler.AdminAPIHandler.Login)
				adminAPI.POST("/new-admin", apiHandler.AdminAPIHandler.AddNewAdmin)
			}
		}
	}

	return r
}
