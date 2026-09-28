package main

import (
	"github.com/gin-gonic/gin"

	"balanceador-simple/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/test", handlers.GetTest)

	router.Run(":3000")
}
