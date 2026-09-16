package products

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	models "main/models/products"
	service "main/services/products"
)

// Handler es la frontera HTTP: entiende URLs, parámetros y códigos de estado.
// No conoce la URL de Solr ni arma parámetros q/fq/rows.
type Handler struct{ service *service.Service }

func NewHandler(s *service.Service) *Handler { return &Handler{service: s} }

func (h *Handler) Register(router *gin.Engine) {
	router.GET("/products/search", h.search)
}

func (h *Handler) search(c *gin.Context) {
	// q representa la intención de búsqueda y no puede faltar.
	text := strings.TrimSpace(c.Query("q"))
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "q es obligatorio"})
		return
	}

	// limit es opcional. Cero significa “no fue enviado”: el Service aplicará
	// el valor por defecto. Si fue enviado, debe ser un entero positivo.
	limit, err := readLimit(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit debe ser un entero positivo"})
		return
	}

	// Convertimos parámetros HTTP al modelo interno y delegamos la búsqueda.
	results, err := h.service.Search(c.Request.Context(), models.Query{
		Text:     text,
		Category: strings.TrimSpace(c.Query("category")),
		Brand:    strings.TrimSpace(c.Query("brand")),
		Limit:    limit,
	})
	if err != nil {
		// El request era válido; falló una dependencia aguas abajo (Solr).
		c.JSON(http.StatusBadGateway, gin.H{"error": "búsqueda no disponible"})
		return
	}

	// La API expone su propio contrato; no el JSON crudo de Solr.
	c.JSON(http.StatusOK, gin.H{"results": results})
}

func readLimit(raw string) (int, error) {
	if raw == "" {
		return 0, nil
	}
	limit, err := strconv.Atoi(raw)
	if err != nil || limit <= 0 {
		return 0, strconv.ErrSyntax
	}
	return limit, nil
}
