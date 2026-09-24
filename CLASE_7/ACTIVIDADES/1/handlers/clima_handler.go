package handlers

import (
	"net/http"

	"actividad-balanceador-roto/config"
	"actividad-balanceador-roto/dto"
	"actividad-balanceador-roto/services"
	"github.com/gin-gonic/gin"
)

// El handler usa este service para consultar la API externa.
var climaService = services.OpenMeteoService{}

// GetClima maneja GET /clima.
func GetClima(c *gin.Context) {
	temperaturaActual, err := climaService.ObtenerTemperatura()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{
			"error": "no se pudo consultar la temperatura",
		})
		return
	}

	nombreServidor := "Clima A"

	c.IndentedJSON(http.StatusOK, dto.RespuestaClima{
		Ciudad:       config.Ciudad,
		TemperaturaC: temperaturaActual.TemperaturaC,
		Servidor:     nombreServidor,
		Mensaje:      config.MensajeRespuesta,
	})
}
