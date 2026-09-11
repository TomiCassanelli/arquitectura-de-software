package products

import (
	"github.com/gin-gonic/gin"
	service "main/services/products"
)

type Handler struct{ service *service.Service }

func NewHandler(s *service.Service) *Handler { return &Handler{service: s} }

func (h *Handler) Register(r *gin.Engine) {
	// TODO GET /products/search
	// q obligatorio; category, brand y limit opcionales.
	// 400 ante parámetros inválidos.
	// 502 si falla el SearchEngine.
	// 200 con results.
}
