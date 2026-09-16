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
	// Esta es la composición de la aplicación: el Client de Solr cumple el
	// contrato Engine; el Service lo usa; el Handler expone el endpoint HTTP.
	solrClient := repository.NewClient(
		"http://localhost:8983/solr/products",
		&http.Client{Timeout: 2 * time.Second},
	)
	productsService := service.NewService(solrClient)

	router := gin.Default()
	controllers.NewHandler(productsService).Register(router)
	_ = router.Run(":8080")
}
