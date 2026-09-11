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
	// El timeout limita la espera por el proveedor. Sin él, una dependencia
	// lenta podría retener recursos de la API por demasiado tiempo.
	client := repository.NewClient(
		"http://localhost:8983/solr/products",
		&http.Client{Timeout: 2 * time.Second},
	)

	router := gin.Default()
	controllers.NewHandler(service.NewService(client)).Register(router)
	_ = router.Run(":8080")
}
