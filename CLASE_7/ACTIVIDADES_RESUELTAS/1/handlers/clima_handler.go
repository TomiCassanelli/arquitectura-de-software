package handlers

import (
	"net/http"
	"os"

	"actividad-balanceador-roto-resuelta/config"
	"actividad-balanceador-roto-resuelta/dto"
	"actividad-balanceador-roto-resuelta/services"
	"github.com/gin-gonic/gin"
)

var climaService = services.OpenMeteoService{}

func GetClima(c *gin.Context) {
	temperaturaActual, err := climaService.ObtenerTemperatura()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{
			"error": "no se pudo consultar la temperatura",
		})
		return
	}

	nombreServidor := os.Getenv("NOMBRE_SERVIDOR")
	if nombreServidor == "" {
		nombreServidor = "Servidor local"
	}

	c.IndentedJSON(http.StatusOK, dto.RespuestaClima{
		Ciudad:       config.Ciudad,
		TemperaturaC: temperaturaActual.TemperaturaC,
		Servidor:     nombreServidor,
		Mensaje:      config.MensajeRespuesta,
	})
}
