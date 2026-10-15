package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Health permite comprobar que el proceso responde; también sirve al healthcheck de Compose.
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Home enumera las rutas para que la demo se pueda explorar desde el navegador.
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mensaje": "Demo de observabilidad", "rutas": []string{"/ok", "/slow", "/error", "/metrics", "/health"}})
}

// OK simula una operación exitosa y produce un log de aplicación.
func OK(c *gin.Context) {
	log.Print("nivel=info evento=operacion_exitosa")
	c.JSON(http.StatusOK, gin.H{"resultado": "ok"})
}

// Slow simula latencia para que la duración de las peticiones cambie en Grafana.
func Slow(c *gin.Context) {
	log.Print("nivel=warn evento=operacion_lenta")
	time.Sleep(750 * time.Millisecond)
	c.JSON(http.StatusOK, gin.H{"resultado": "lento", "duracion_ms": 750})
}

// Error simula un fallo controlado. Su HTTP 500 y su log se pueden correlacionar.
func Error(c *gin.Context) {
	log.Print("nivel=error evento=operacion_fallida motivo=fallo_simulado")
	c.JSON(http.StatusInternalServerError, gin.H{"error": "fallo simulado para la demostración"})
}
