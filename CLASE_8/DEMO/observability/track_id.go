package observability

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddTrackID asigna un ID distinto a cada petición del grupo /compra.
func AddTrackID(ctx *gin.Context) {
	// 16 bytes aleatorios se muestran como 32 caracteres hexadecimales.
	var randomBytes [16]byte
	if _, err := rand.Read(randomBytes[:]); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "no se pudo crear el track_id"})
		return
	}
	trackID := hex.EncodeToString(randomBytes[:])
	// El controller lo lee con GetString; el cliente lo recibe en el header.
	ctx.Set("track_id", trackID)
	ctx.Header("X-Track-ID", trackID)
	ctx.Next()
}
