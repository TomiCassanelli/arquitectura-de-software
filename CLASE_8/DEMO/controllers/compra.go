package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CompraExitosa simula tres etapas de una compra exitosa dentro de esta API.
func (controller *DemoController) CompraExitosa(ctx *gin.Context) {
	trackID := ctx.GetString("track_id")
	controller.Logs.Write(trackID, "gateway", "info", "compra recibida")
	controller.Logs.Write(trackID, "pedidos", "info", "stock verificado")
	controller.Logs.Write(trackID, "pagos", "info", "pago aprobado")
	ctx.JSON(http.StatusOK, gin.H{"resultado": "compra confirmada", "track_id": trackID})
}

// CompraErrorPagos simula el mismo recorrido y marca Pagos como etapa del fallo.
func (controller *DemoController) CompraErrorPagos(ctx *gin.Context) {
	trackID := ctx.GetString("track_id")
	controller.Logs.Write(trackID, "gateway", "info", "compra recibida")
	controller.Logs.Write(trackID, "pedidos", "info", "stock verificado")
	controller.Logs.Write(trackID, "pagos", "error", "pago rechazado")
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "pago rechazado", "track_id": trackID})
}

// CompraErrorPedidos detiene la compra en Pedidos; Pagos no llega a ejecutarse.
func (controller *DemoController) CompraErrorPedidos(ctx *gin.Context) {
	trackID := ctx.GetString("track_id")
	controller.Logs.Write(trackID, "gateway", "info", "compra recibida")
	controller.Logs.Write(trackID, "pedidos", "error", "fallo al verificar stock")
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "fallo al verificar stock", "track_id": trackID})
}
