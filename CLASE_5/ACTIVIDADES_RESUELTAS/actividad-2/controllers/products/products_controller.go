package products

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	models "main/models/products"
	service "main/services/products"
)

type Handler struct{ service *service.Service }

func NewHandler(s *service.Service) *Handler { return &Handler{service: s} }

func (h *Handler) Register(r *gin.Engine) {
	r.GET("/products/search", h.search)
}

func (h *Handler) search(c *gin.Context) {
	text := strings.TrimSpace(c.Query("q"))
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "q es obligatorio"})
		return
	}

	limit, err := parseLimit(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit debe ser un entero positivo"})
		return
	}

	response, err := h.service.Search(c.Request.Context(), models.Query{
		Text:     text,
		Category: strings.TrimSpace(c.Query("category")),
		Brand:    strings.TrimSpace(c.Query("brand")),
		Limit:    limit,
	})
	if err != nil {
		// La petición del cliente es válida; la dependencia de búsqueda falló.
		c.JSON(http.StatusBadGateway, gin.H{"error": "búsqueda no disponible"})
		return
	}
	c.JSON(http.StatusOK, response)
}

func parseLimit(raw string) (int, error) {
	if raw == "" {
		return 0, nil // Service lo convierte al valor por defecto: 10.
	}
	limit, err := strconv.Atoi(raw)
	if err != nil || limit <= 0 {
		return 0, strconv.ErrSyntax
	}
	return limit, nil
}
