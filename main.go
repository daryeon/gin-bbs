package main

import (
	"github.com/daryeon/gin-bbs/controllers"
	"github.com/daryeon/gin-bbs/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initRouter(router *gin.Engine) {
	/* 绑定模板，静态文件 */
	router.LoadHTMLGlob("static/template/*.html")
	router.StaticFile("script.js", "./static/js/script.js")
	router.StaticFile("style.css", "./static/css/style.css")

	/* 定义路由 */
	router.GET("/", controllers.Get)
	router.GET("/add", controllers.Add)
	router.GET("/posts", controllers.GetPosts)
	router.POST("/posts", controllers.AddPost)
}

// go mod edit -require github.com/gin-gonic/gin@latest
func main() {
	defer func() {
		var c *gin.Context
		if err := recover(); err != nil {
			c.JSON(500, gin.H{"err": err})
		}
	}()

	db, err := models.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	initRouter(router)
	router.Run(":8080")
}
