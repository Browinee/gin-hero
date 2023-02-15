package router

import (
	"master-gin/controllers"
	"master-gin/logger"
	"master-gin/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
		v1.POST("/posts", controllers.CreatePostHandler)
		v1.GET("/posts/", controllers.GetPostListHandler)
		v1.GET("/posts/:id", controllers.GetPostDetailHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "404"})
	})
	return r
}
