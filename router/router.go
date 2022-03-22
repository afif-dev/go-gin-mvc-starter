package router

import (
	"gin-mvc/controllers"
	"gin-mvc/middlewares"
	"gin-mvc/models"
	"io"
	"os"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {

	// # Uncomment this code to load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	gin.SetMode(os.Getenv("GIN_MODE"))

	// # Disable log's color
    gin.DisableConsoleColor()
	// # Force log's color
	// gin.ForceConsoleColor()

    // # Logging to a file.
    f, _ := os.Create("logs/gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

	models.ConnectDatabase(os.Getenv("DB_CONN_URL"))

	r := gin.Default()
	
	r.SetTrustedProxies(nil) //disabled Trusted Proxies
	
	r.Static("/static", "./static")
	
	// # Template Engine
	r.HTMLRender = ginview.Default()

	// # Page Router
	r.GET("/", controllers.IndexPage, middlewares.TestMw())
	r.GET("/structure", controllers.StructurePage)
	r.GET("/credits", controllers.CreditPage)
	r.GET("/articles", controllers.GetArticlesPage)
	r.GET("/article/:id", controllers.GetArticlePage)
	r.GET("/api", controllers.ApiPage)

	// # API Group Router
	api := r.Group("/api") 
	{
		api.GET("/articles", controllers.GetArticlesAPI)
		api.GET("/article/:id", controllers.GetArticle)
		api.POST("/article", controllers.AddArticle)
		api.PUT("/article", controllers.UpdateArticle)
		api.DELETE("/article", controllers.DeleteArticle)
	}
	
	return r
}
