package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/blocks"
	"github.com/pengdst/golang-file-upload/config"
	"github.com/pengdst/golang-file-upload/controller"
	webController "github.com/pengdst/golang-file-upload/controller/web"
	"github.com/pengdst/golang-file-upload/exception"
	"github.com/pengdst/golang-file-upload/repository"
	"github.com/pengdst/golang-file-upload/service"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"path"
)

func main() {
	env, errConf := config.LoadEnv()
	if errConf != nil {
		log.Fatal(errConf)
	}

	db, errDb := config.NewDatabase(env)
	if errDb != nil {
		log.Fatal(errDb)
	}

	router := gin.Default()

	userRepo := repository.NewUserRepository(db)

	fileService := service.NewFilesService(env)
	authService := service.NewAuthService(userRepo)

	filesController := controller.NewFilesController(fileService)
	authController := controller.NewAuthController(authService)

	router.StaticFS("public", http.Dir("public"))

	api := router.Group("api")
	api.Use(gin.CustomRecovery(exception.ErrorHandler))

	api.POST("file", filesController.Upload)

	auth := api.Group("auth")
	auth.POST("login", authController.Login)
	auth.POST("register", authController.Register)

	rootDir, errDir := os.Getwd()
	if errDir != nil {
		panic(errDir)
	}
	views := blocks.New(path.Join(rootDir, "web/view")).
		Extension(".gohtml").
		Reload(true).
		Funcs(
			template.FuncMap{
				"add": func(result int, numbs ...int) int {
					for _, numb := range numbs {
						result += numb
					}

					return result
				},
			},
		)

	homeController := webController.NewHomeController(views, userRepo)
	webAuthController := webController.NewAuthController(views, userRepo)

	router.GET("/", homeController.Index)

	router.GET("/login", webAuthController.Login)
	router.GET("/register", webAuthController.Register)

	router.POST("/register", webAuthController.RegisterProcess)
	router.POST("/login", webAuthController.LoginProcess)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
