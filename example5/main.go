package main

import (
	"log"
	"net/http"

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
	// XXX try to change the order of these two lines,
	// and see how many handles are registered to the public endpoints.
	registerPublic(r)
	registerPrivate(r)
}

func registerPublic(r *gin.Engine) {
	corsMiddleware := func() func(*gin.Context) {
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AddAllowHeaders("Authorization")
		return cors.New(config)
	}()

	publicGroup := r.Group("/public")

	// register to grouped router
	publicGroup.Use(
		func(ctx *gin.Context) {
			log.Println("register to grouped router")
		},
		corsMiddleware)

	publicGroup.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "POST",
		})
	})

	publicGroup.OPTIONS("/:path", func(c *gin.Context) {
		log.Println(c.FullPath())

		// by pass gin's default OPTIONS handler
		// c.AbortWithStatus(http.StatusNoContent)
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}

func registerPrivate(r *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5000"}

	corsMiddleware := cors.New(corsConfig)

	// register to root router
	r.Use(func(ctx *gin.Context) {
		log.Println("register to root router")
	},
		corsMiddleware)

	r.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "POST",
		})
	})
}
