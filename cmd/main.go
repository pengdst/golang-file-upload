package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/blocks"
	"github.com/pengdst/golang-file-upload/config"
	"github.com/pengdst/golang-file-upload/controller"
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

	router.GET("/", func(context *gin.Context) {
		data := map[string]interface{}{
			"Peoples": []map[string]string{
				{
					"Name":      "One Piece",
					"Position":  "Troublemaker",
					"Office":    "New World",
					"Age":       "33",
					"StartDate": "2008-11-28",
					"Salary":    "162,700",
				},
			},
		}
		err := views.ExecuteTemplate(context.Writer, "index", "admin", data)
		if err != nil {
			panic(err)
		}
	})

	router.GET("/login", func(context *gin.Context) {
		data := map[string]interface{}{
			"Title":    "Login",
			"ImageUrl": "http://www.w3.org/2000/svg",
		}
		err := views.ExecuteTemplate(context.Writer, "auth/login", "guest", data)
		if err != nil {
			panic(err)
		}
	})

	router.GET("/register", func(context *gin.Context) {
		data := map[string]interface{}{
			"Title":    "Login",
			"ImageUrl": "http://www.w3.org/2000/svg",
		}
		err := views.ExecuteTemplate(context.Writer, "auth/register", "guest", data)
		if err != nil {
			panic(err)
		}
	})

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
