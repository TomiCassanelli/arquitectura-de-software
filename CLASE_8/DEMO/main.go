package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"observabilidad-demo/controllers"
	"observabilidad-demo/logs"
	"observabilidad-demo/observability"
)

// main es el único punto de entrada de la aplicación. Docker Compose inicia los demás servicios.
func main() {
	// Gin en modo release evita los mensajes internos de depuración durante la demo.
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// La app comprueba MySQL al arrancar porque los controllers guardan logs allí.
	logStore, err := logs.Connect()
	if err != nil {
		log.Fatalf("no se pudo conectar a MySQL: %v", err)
	}
	controller := &controllers.DemoController{Logs: logStore}
	// Recovery convierte un panic en HTTP 500 para que el proceso siga vivo.
	router.Use(gin.Recovery())
	// Estas dos rutas quedan fuera del grupo medido para no contar las sondas.
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/health", controllers.Health)

	// Gin ejecuta este middleware antes y después de cada controller del grupo.
	apiRoutes := router.Group("/")
	apiRoutes.Use(observability.RecordMetrics)
	apiRoutes.GET("/", controllers.Home)
	apiRoutes.GET("/ok", controller.OperacionExitosa)
	apiRoutes.GET("/slow", controller.OperacionLenta)
	apiRoutes.GET("/error", controller.OperacionFallida)
	// Solo las rutas de compra reciben un track_id; también registran las métricas del grupo apiRoutes.
	compraRoutes := apiRoutes.Group("/compra")
	compraRoutes.Use(observability.AddTrackID)
	compraRoutes.GET("/ok", controller.CompraExitosa)
	compraRoutes.GET("/error", controller.CompraErrorPagos)
	compraRoutes.GET("/error-pedidos", controller.CompraErrorPedidos)

	// La dirección 0.0.0.0 permite que otros contenedores lleguen a la API.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("app lista en :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
