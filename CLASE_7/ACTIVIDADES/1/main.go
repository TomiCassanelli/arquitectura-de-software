package main

import (
	"actividad-balanceador-roto/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/clima", handlers.GetClima)
	router.Run(":3000")
}
