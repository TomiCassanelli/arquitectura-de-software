package products

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	models "main/models/products"
)

// Client implementa models.Engine. Es el único lugar de esta solución que
// conoce la URL, parámetros y respuesta HTTP de Solr.
type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{baseURL: baseURL, http: httpClient}
}

func (c *Client) Search(ctx context.Context, query models.Query) ([]models.Result, error) {
	// Traducimos Query (nuestro modelo) a parámetros Solr.
	params := url.Values{
		"q":       {query.Text},
		"defType": {"edismax"},
		"qf":      {"title description"},
		"rows":    {strconv.Itoa(query.Limit)},
		"fl":      {"id,title,brand,category,price,score"},
		"wt":      {"json"},
	}

	// fq es un filtro obligatorio. {!term} trata el valor como literal: quien
	// llama a la API envía “Adidas”, no una expresión libre del lenguaje Solr.
	if query.Category != "" {
		params.Add("fq", "{!term f=category}"+query.Category)
	}
	if query.Brand != "" {
		params.Add("fq", "{!term f=brand}"+query.Brand)
	}

	// NewRequestWithContext hace que una cancelación del request original pueda
	// interrumpir también la llamada HTTP que está esperando a Solr.
	request, err := http.NewRequestWithContext(
		ctx, http.MethodGet, c.baseURL+"/select?"+params.Encode(), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("crear request a Solr: %w", err)
	}

	responseHTTP, err := c.http.Do(request)
	if err != nil {
		return nil, fmt.Errorf("consultar Solr: %w", err)
	}
	defer responseHTTP.Body.Close()

	if responseHTTP.StatusCode < http.StatusOK || responseHTTP.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("Solr respondió HTTP %d", responseHTTP.StatusCode)
	}

	var payload response
	if err := json.NewDecoder(responseHTTP.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("decodificar respuesta de Solr: %w", err)
	}

	// Traducimos el JSON externo a []models.Result, el tipo que entiende la API.
	results := make([]models.Result, 0, len(payload.Response.Docs))
	for _, document := range payload.Response.Docs {
		results = append(results, models.Result{
			ID: document.ID, Title: document.Title, Brand: document.Brand,
			Category: document.Category, Price: document.Price, Score: document.Score,
		})
	}
	return results, nil
}
