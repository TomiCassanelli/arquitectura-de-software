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

type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{baseURL: baseURL, http: httpClient}
}

func (c *Client) Search(ctx context.Context, query models.Query) (models.Response, error) {
	parameters := url.Values{
		"q":       {query.Text},
		"defType": {"edismax"},
		"qf":      {"title description"},
		"rows":    {strconv.Itoa(query.Limit)},
		"fl":      {"id,title,brand,category,price,score"},
		"wt":      {"json"},
	}

	// {!term} trata category/brand como valores exactos. Esto evita que un
	// parámetro externo sea interpretado como sintaxis libre de Solr.
	if query.Category != "" {
		parameters.Add("fq", "{!term f=category}"+query.Category)
	}
	if query.Brand != "" {
		parameters.Add("fq", "{!term f=brand}"+query.Brand)
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/select?"+parameters.Encode(),
		nil,
	)
	if err != nil {
		return models.Response{}, fmt.Errorf("crear request a Solr: %w", err)
	}

	response, err := c.http.Do(request)
	if err != nil {
		return models.Response{}, fmt.Errorf("ejecutar búsqueda en Solr: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return models.Response{}, fmt.Errorf("Solr respondió HTTP %d", response.StatusCode)
	}

	var decoded solrResponse
	if err := json.NewDecoder(response.Body).Decode(&decoded); err != nil {
		return models.Response{}, fmt.Errorf("decodificar respuesta de Solr: %w", err)
	}

	results := make([]models.Result, 0, len(decoded.Response.Docs))
	for _, document := range decoded.Response.Docs {
		results = append(results, models.Result{
			ID: document.ID, Title: document.Title, Brand: document.Brand,
			Category: document.Category, Price: document.Price, Score: document.Score,
		})
	}
	return models.Response{Results: results, Total: decoded.Response.NumFound}, nil
}
