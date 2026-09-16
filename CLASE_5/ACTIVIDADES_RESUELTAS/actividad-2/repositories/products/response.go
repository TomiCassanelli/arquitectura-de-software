package products

// response representa solamente el fragmento del JSON de Solr que necesitamos.
// Es privado porque la forma externa de Solr no debe filtrarse a otras capas.
type response struct {
	Response struct {
		Docs []struct {
			ID       string  `json:"id"`
			Title    string  `json:"title"`
			Brand    string  `json:"brand"`
			Category string  `json:"category"`
			Price    float64 `json:"price"`
			Score    float64 `json:"score"`
		} `json:"docs"`
	} `json:"response"`
}
