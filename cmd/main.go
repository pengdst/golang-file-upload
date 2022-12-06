package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/blocks"
	"github.com/pengdst/golang-file-upload/config"
	"github.com/pengdst/golang-file-upload/controller"
	"github.com/pengdst/golang-file-upload/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	env, errConf := config.LoadEnv()
	if errConf != nil {
		log.Fatal(errConf)
	}

	router := gin.Default()

	fileService := service.NewFilesService(env)
	filesController := controller.NewFilesController(fileService)

	router.StaticFS("public", http.Dir("public"))

	api := router.Group("api")
	api.POST("file", filesController.Upload)

	rootDir, errDir := os.Getwd()
	if errDir != nil {
		panic(errDir)
	}
	var views = blocks.New(path.Join(rootDir, "web/view")).
		Reload(true).
		Funcs(map[string]interface{}{
			"year": func() int {
				return time.Now().Year()
			},
		})

	router.GET("/", func(context *gin.Context) {
		data := map[string]interface{}{
			"Title": "Page Title",
		}
		err := views.ExecuteTemplate(context.Writer, "index", "main", data)
		if err != nil {
			panic(err)
		}
	})

	router.GET("/login", func(context *gin.Context) {
		data := map[string]interface{}{
			"Title":    "Login",
			"ImageUrl": "http://www.w3.org/2000/svg",
		}
		err := views.ExecuteTemplate(context.Writer, "auth/login", "main", data)
		if err != nil {
			panic(err)
		}
	})

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
