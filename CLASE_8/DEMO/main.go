package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"observabilidad-demo/controllers"
	"observabilidad-demo/logstore"
	"observabilidad-demo/observability"
)

// main es el único punto de entrada de la aplicación. Docker Compose inicia los demás servicios.
func main() {
	// Gin en modo release evita los mensajes internos de depuración durante la demo.
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// La app comprueba MySQL al arrancar porque los controllers guardan logs allí.
	logs, err := logstore.Open()
	if err != nil {
		log.Fatalf("no se pudo conectar a MySQL: %v", err)
	}
	defer logs.Close()
	demo := &controllers.Demo{Logs: logs}
	// Recovery convierte un panic en HTTP 500 para que el proceso siga vivo.
	router.Use(gin.Recovery())
	// Estas dos rutas quedan fuera del grupo medido para no contar las sondas.
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/health", controllers.Health)

	// Gin ejecuta este middleware antes y después de cada controller del grupo.
	api := router.Group("/")
	api.Use(observability.Middleware())
	api.GET("/", controllers.Home)
	api.GET("/ok", demo.OK)
	api.GET("/slow", demo.Slow)
	api.GET("/error", demo.Error)

	// La dirección 0.0.0.0 permite que otros contenedores lleguen a la API.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("app lista en :%s a las %s", port, time.Now().Format(time.RFC3339))
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
