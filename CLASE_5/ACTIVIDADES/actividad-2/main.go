package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	controllers "main/controllers/products"
	repository "main/repositories/products"
	service "main/services/products"
)

func main() {
	client := repository.NewClient(
		"http://localhost:8983/solr/products",
		&http.Client{Timeout: 2 * time.Second},
	)
	productsService := service.NewService(client)

	r := gin.Default()
	controllers.NewHandler(productsService).Register(r)
	_ = r.Run(":8080")
}
