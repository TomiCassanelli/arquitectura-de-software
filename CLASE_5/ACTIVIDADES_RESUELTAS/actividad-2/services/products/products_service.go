package products

import (
	"context"

	models "main/models/products"
)

// Service contiene reglas de aplicación, no detalles de HTTP ni de Solr.
type Service struct{ engine models.Engine }

func NewService(engine models.Engine) *Service { return &Service{engine: engine} }

func (s *Service) Search(ctx context.Context, query models.Query) ([]models.Result, error) {
	// Si no llegó un límite, elegimos una página razonable para la API.
	if query.Limit <= 0 {
		query.Limit = 10
	}
	return s.engine.Search(ctx, query)
}
