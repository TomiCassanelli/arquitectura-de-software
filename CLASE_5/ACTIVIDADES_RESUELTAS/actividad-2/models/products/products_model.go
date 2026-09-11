package products

import "context"

// Query es el lenguaje de la aplicación. No expone q, fq ni rows: esos son
// detalles del adaptador Solr.
type Query struct {
	Text     string
	Category string
	Brand    string
	Limit    int
}

type Result struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Score    float64 `json:"score"`
}

// Response conserva el total antes de paginar. results puede tener menos
// elementos que total porque Solr aplica rows.
type Response struct {
	Results []Result `json:"results"`
	Total   int      `json:"total"`
}

// Engine permite que Service no dependa de Solr ni de HTTP.
type Engine interface {
	Search(context.Context, Query) (Response, error)
}
