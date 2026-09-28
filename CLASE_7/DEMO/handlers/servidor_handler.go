package handlers

import (
	"net/http"
	"os"

	"balanceador-simple/dto"
	"github.com/gin-gonic/gin"
)

// GetTest maneja GET /test.
func GetTest(c *gin.Context) {
	nombreServidor := os.Getenv("NOMBRE_SERVIDOR")
	if nombreServidor == "" {
		nombreServidor = "Servidor local"
	}

	c.IndentedJSON(http.StatusOK, dto.RespuestaTest{
		Servidor: nombreServidor,
		Mensaje:  "¡Hola Mundo!",
	})
}
