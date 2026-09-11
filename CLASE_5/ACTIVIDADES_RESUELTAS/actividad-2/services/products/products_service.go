package products

import (
	"context"

	models "main/models/products"
)

type Service struct{ engine models.Engine }

func NewService(engine models.Engine) *Service { return &Service{engine: engine} }

func (s *Service) Search(ctx context.Context, query models.Query) (models.Response, error) {
	// El default pertenece a la política de la aplicación. El handler valida
	// los valores explícitos; el repositorio solo traduce una Query válida.
	if query.Limit == 0 {
		query.Limit = 10
	}
	return s.engine.Search(ctx, query)
}
