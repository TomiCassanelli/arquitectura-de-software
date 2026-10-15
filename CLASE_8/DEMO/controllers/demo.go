package controllers

import (
	"net/http"
	"time"

	"observabilidad-demo/logs"

	"github.com/gin-gonic/gin"
)

// DemoController agrupa las operaciones de la demo que escriben logs en MySQL.
type DemoController struct{ Logs *logs.Store }

// Health permite comprobar que el proceso responde; también sirve al healthcheck de Compose.
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Home enumera las rutas para que la demo se pueda explorar desde el navegador.
func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Demo de observabilidad", "rutas": []string{"/ok", "/slow", "/error", "/compra/ok", "/compra/error", "/compra/error-pedidos", "/metrics", "/health"}})
}

// OperacionExitosa simula una operación exitosa y produce un log de aplicación.
func (controller *DemoController) OperacionExitosa(ctx *gin.Context) {
	controller.Logs.Write("", "api", "info", "operacion_exitosa")
	ctx.JSON(http.StatusOK, gin.H{"resultado": "ok"})
}

// OperacionLenta simula latencia para que la duración de las peticiones cambie en Grafana.
func (controller *DemoController) OperacionLenta(ctx *gin.Context) {
	controller.Logs.Write("", "api", "warn", "operacion_lenta")
	time.Sleep(750 * time.Millisecond)
	ctx.JSON(http.StatusOK, gin.H{"resultado": "lento", "duracion_ms": 750})
}

// OperacionFallida simula un fallo controlado. Su HTTP 500 y su log se pueden correlacionar.
func (controller *DemoController) OperacionFallida(ctx *gin.Context) {
	controller.Logs.Write("", "api", "error", "operacion_fallida: fallo_simulado")
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "fallo simulado para la demostración"})
}
