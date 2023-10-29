package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// with two CORS policies

	r := gin.New()
	r.Use(gin.Logger())

	registerEndpoint(r)

	r.Run(":8080")
}

func registerEndpoint(r *gin.Engine) {
	registerPublic(r)
	registerPrivate(r)
}

func registerPublic(r *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsMiddleware := cors.New(corsConfig)

	publicGroup := r.Group("/public")

	// register to grouped router
	publicGroup.Use(corsMiddleware)

	publicGroup.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "POST",
		})
	})
}

func registerPrivate(r *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5000"}

	corsMiddleware := cors.New(corsConfig)

	// register to root router
	r.Use(corsMiddleware)

	r.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "POST",
		})
	})
}
