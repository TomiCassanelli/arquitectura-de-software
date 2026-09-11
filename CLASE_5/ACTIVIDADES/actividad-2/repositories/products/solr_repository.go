package products

import (
	"context"
	"net/http"

	models "main/models/products"
)

type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{baseURL: baseURL, http: httpClient}
}

func (c *Client) Search(ctx context.Context, q models.Query) ([]models.Result, error) {
	// TODO:
	// - construir <baseURL>/select
	// - q.Text -> q; defType=edismax; qf=title description
	// - category/brand -> fq con {!term f=category} / {!term f=brand} (valor literal)
	// - limit -> rows
	// - fl=id,title,brand,category,price,score
	// - wt=json
	// - request con context
	// - validar status
	// - parsear response
	// - mapear documentos
	return nil, nil
}
