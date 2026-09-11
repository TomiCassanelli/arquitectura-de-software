package products

// solrResponse representa solamente la parte de JSON que esta API necesita.
// Mantenerlo privado evita filtrar la forma del proveedor al resto de capas.
type solrResponse struct {
	Response struct {
		NumFound int `json:"numFound"`
		Docs     []struct {
			ID       string  `json:"id"`
			Title    string  `json:"title"`
			Brand    string  `json:"brand"`
			Category string  `json:"category"`
			Price    float64 `json:"price"`
			Score    float64 `json:"score"`
		} `json:"docs"`
	} `json:"response"`
}
