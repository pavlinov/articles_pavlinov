package routes

import (
	"articles_pavlinov/controllers"
	"articles_pavlinov/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/logout", middleware.AuthMiddleware(), controllers.Logout)
	}

	article := r.Group("/articles")
	{
		article.GET("/", controllers.GetArticles)
		article.GET("/:id", controllers.GetArticle)
		article.POST("/", middleware.AuthMiddleware(), controllers.AddArticle)
		article.DELETE("/:id", middleware.AuthMiddleware(), controllers.RemoveArticle)
	}

	preference := r.Group("/preferences")
	{
		preference.GET("/", middleware.AuthMiddleware(), controllers.GetPreferences)
		preference.POST("/", middleware.AuthMiddleware(), controllers.SetPreference)
	}

	return r
}
