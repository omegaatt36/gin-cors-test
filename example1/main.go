package main

import "github.com/gin-gonic/gin"

func main() {
	// without CORS middleware

	r := gin.New()
	r.Use(gin.Logger())
	r.POST("/public/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.POST("/private/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.Run(":8080")
}
