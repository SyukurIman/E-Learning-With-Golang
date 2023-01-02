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
	UserAPIHandler api.UserAPI
	TaskAPIHandler api.TaskAPI
}

type ClientHandler struct {
	AuthWeb       web.AuthWeb
	HomeWeb       web.HomeWeb
	DashboardUser web.DashboardUser
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

	// Service Set
	userService := service.NewUserService(userRepo, taskRepo)

	// API Set
	userAPIHandler := api.NewUserAPI(userService)
	apiHandler := APIHandler{
		UserAPIHandler: userAPIHandler,
	}

	// User Set
	userClient := client.NewUserClient()

	authWeb := web.NewAuthWeb(userClient)
	homeWeb := web.NewHomeWeb()
	dashboardUser := web.NewDashboardUser(userRepo)

	client := ClientHandler{
		AuthWeb:       authWeb,
		HomeWeb:       homeWeb,
		DashboardUser: dashboardUser,
	}

	auth := r.Group("/")
	{
		auth.GET("/", client.HomeWeb.Index)
		auth.GET("/login", client.AuthWeb.Login)
		auth.POST("/login/proses", client.AuthWeb.LoginProses)

		auth.GET("/register", client.AuthWeb.Register)
		auth.POST("/register/proses", client.AuthWeb.RegisterProses)

		auth.GET("/user/dashboard", utils.Authentication(), client.DashboardUser.Dashboard)
		auth.GET("/logout", client.AuthWeb.Logout)

		// Bagian API
		api := auth.Group("/api")
		{
			// Set Routing User API
			userAPI := api.Group("/user")
			{
				userAPI.POST("/login", apiHandler.UserAPIHandler.Login)
				userAPI.POST("/register", apiHandler.UserAPIHandler.Register)
			}

		}
	}

	return r
}

// func RunClient(r *gin.Engine) *gin.Engine {

// 	return r
// }
