package observability

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Promauto registra estos instrumentos en el registro por defecto de Prometheus.
// CounterVec permite guardar un contador distinto por ruta y estado HTTP.
var requests = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "requests_total",
	Help: "Cantidad de respuestas HTTP.",
}, []string{"route", "status"})

// HistogramVec guarda las duraciones; Prometheus genera _bucket, _sum y _count.
var duration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "request_duration_seconds",
	Help: "Duración de las respuestas HTTP en segundos.",
}, []string{"route", "status"})

// RecordMetrics mide una vez cada petición atendida por un controller.
func RecordMetrics(c *gin.Context) {
	// Empezamos a medir antes de que responda el controller.
	start := time.Now()
	c.Next() // Gin ejecuta el controller y luego vuelve a la línea siguiente.

	// Al volver, ya conocemos la ruta y el estado de la respuesta.
	route := c.FullPath()
	status := strconv.Itoa(c.Writer.Status())

	// Inc suma una solicitud; Observe registra cuánto tardó.
	requests.WithLabelValues(route, status).Inc()
	duration.WithLabelValues(route, status).Observe(time.Since(start).Seconds())
}
