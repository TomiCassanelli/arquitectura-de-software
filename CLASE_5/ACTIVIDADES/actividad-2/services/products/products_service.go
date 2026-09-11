package products

import (
	"context"

	models "main/models/products"
)

type Service struct {
	engine models.Engine
}

func NewService(engine models.Engine) *Service {
	return &Service{engine: engine}
}

func (s *Service) Search(ctx context.Context, q models.Query) ([]models.Result, error) {
	if q.Limit <= 0 {
		q.Limit = 10
	}
	return s.engine.Search(ctx, q)
}
