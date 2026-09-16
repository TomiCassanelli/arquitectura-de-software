package products

import "context"

// Query es el pedido interno de búsqueda. No usa los nombres q, fq o rows:
// esos son detalles de Solr que se traducen recién en el repositorio.
type Query struct {
	Text     string
	Category string
	Brand    string
	Limit    int
}

// Result es el producto que nuestra API devolverá en JSON al cliente.
type Result struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Score    float64 `json:"score"`
}

// Engine es el contrato del buscador. Service pide “buscar productos” sin
// saber si la implementación concreta consulta Solr, una cache o un fake.
type Engine interface {
	Search(context.Context, Query) ([]Result, error)
}
