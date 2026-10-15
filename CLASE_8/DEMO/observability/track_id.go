package observability

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddTrackID asigna un ID distinto a cada petición del grupo /compra.
func AddTrackID(c *gin.Context) {
	// 16 bytes aleatorios se muestran como 32 caracteres hexadecimales.
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "no se pudo crear el track_id"})
		return
	}
	id := hex.EncodeToString(bytes[:])
	// El controller lo lee con GetString; el cliente lo recibe en el header.
	c.Set("track_id", id)
	c.Header("X-Track-ID", id)
	c.Next()
}
