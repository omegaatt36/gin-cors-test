package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFS("/", http.Dir("."))
	r.Run(":5000")
	log.Println("a")
}
