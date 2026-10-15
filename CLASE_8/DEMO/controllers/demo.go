package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"observabilidad-demo/logstore"
)

// Demo agrupa los controllers que escriben logs en MySQL.
type Demo struct{ Logs *logstore.Store }

// Health permite comprobar que el proceso responde; también sirve al healthcheck de Compose.
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Home enumera las rutas para que la demo se pueda explorar desde el navegador.
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mensaje": "Demo de observabilidad", "rutas": []string{"/ok", "/slow", "/error", "/metrics", "/health"}})
}

// OK simula una operación exitosa y produce un log de aplicación.
func (d *Demo) OK(c *gin.Context) {
	d.Logs.Write(c.Request.Context(), "info", "operacion_exitosa")
	c.JSON(http.StatusOK, gin.H{"resultado": "ok"})
}

// Slow simula latencia para que la duración de las peticiones cambie en Grafana.
func (d *Demo) Slow(c *gin.Context) {
	d.Logs.Write(c.Request.Context(), "warn", "operacion_lenta")
	time.Sleep(750 * time.Millisecond)
	c.JSON(http.StatusOK, gin.H{"resultado": "lento", "duracion_ms": 750})
}

// Error simula un fallo controlado. Su HTTP 500 y su log se pueden correlacionar.
func (d *Demo) Error(c *gin.Context) {
	d.Logs.Write(c.Request.Context(), "error", "operacion_fallida: fallo_simulado")
	c.JSON(http.StatusInternalServerError, gin.H{"error": "fallo simulado para la demostración"})
}
