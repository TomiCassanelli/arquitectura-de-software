package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// -----------------------------------------------------------------------------
// MODELOS:
// -----------------------------------------------------------------------------
// Query es lo que el controlador le entrega al servicio.
type Query struct {
	Text, Category, Brand string
	Limit                 int
}

// Result es un producto que la API devolverá al cliente.
type Result struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Score    float64 `json:"score"`
}

// SearchResponse es el contrato público de la API. Results es la página que
// devolvemos; Total es la cantidad total de coincidencias antes de paginar.
type SearchResponse struct {
	Results []Result `json:"results"`
	Total   int      `json:"total"`
}

// -----------------------------------------------------------------------------
// INTERFAZ ABSTRACTA DEL REPOSITORY:
// -----------------------------------------------------------------------------
// Engine dice: “necesito alguien que pueda buscar”. No dice Solr, HTTP ni URL.
// Por eso Service puede depender de esta interfaz y no de una tecnología.
// Un test podría usar un buscador falso; en producción usaremos SolrClient.
type Engine interface {
	Search(context.Context, Query) (SearchResponse, error)
}

// -----------------------------------------------------------------------------
// CLIENTE SOLR: adaptador que habl con Apache Solr
// -----------------------------------------------------------------------------
type SolrClient struct {
	baseURL string
	http    *http.Client
}

func (c SolrClient) Search(ctx context.Context, query Query) (SearchResponse, error) {
	// Acá se traduce el lenguaje de nuestra aplicación al lenguaje de Solr.
	// query.Text se convierte en q; query.Limit se convierte en rows.
	params := url.Values{
		"q":       {query.Text},
		"defType": {"edismax"},
		"qf":      {"title^5 description"},
		"rows":    {strconv.Itoa(query.Limit)},
		"fl":      {"id,title,brand,category,price,score"},
		"wt":      {"json"},
	}
	// Si llegaron filtros opcionales, se agregan como fq. {!term} hace que
	// category/brand sean valores literales y no sintaxis Solr escrita por quien
	// usa nuestra API.
	if query.Category != "" {
		params.Add("fq", "{!term f=category}"+query.Category)
	}
	if query.Brand != "" {
		params.Add("fq", "{!term f=brand}"+query.Brand)
	}

	// Se construye una URL como:
	// http://localhost:8983/solr/products/select?q=zapatillas&rows=10...
	// El context permite cancelar la llamada si la petición original se corta.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/select?"+params.Encode(), nil)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("crear consulta a Solr: %w", err)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("consultar Solr: %w", err)
	}
	// Siempre cerramos el cuerpo HTTP cuando terminamos de leer la respuesta.
	defer resp.Body.Close()
	// Solr respondió, pero un status 4xx/5xx todavía es un error para la API.
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return SearchResponse{}, fmt.Errorf("Solr respondió HTTP %d", resp.StatusCode)
	}

	// Solr responde JSON con response.numFound y response.docs. Esta estructura
	// temporal representa esa forma externa; no la devolvemos tal cual al cliente.
	var payload struct {
		Response struct {
			NumFound int      `json:"numFound"`
			Docs     []Result `json:"docs"`
		} `json:"response"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return SearchResponse{}, fmt.Errorf("decodificar respuesta de Solr: %w", err)
	}
	// Traducimos la respuesta de Solr al contrato de NUESTRA API: results/total.
	return SearchResponse{Results: payload.Response.Docs, Total: payload.Response.NumFound}, nil
}

// -----------------------------------------------------------------------------
// SERVICIO:
// -----------------------------------------------------------------------------
// Service recibe un Engine. No sabe si ese Engine es Solr, Elasticsearch, una
// cache o un fake. Su trabajo es aplicar reglas antes de pedir la búsqueda.
type Service struct{ engine Engine }

func (s Service) Search(ctx context.Context, query Query) (SearchResponse, error) {
	// Si el cliente no envió limit, la regla de nuestra API es devolver 10.
	// Esta regla no pertenece al Controller ni al cliente de Solr.
	if query.Limit == 0 {
		query.Limit = 10
	}
	// El servicio delega la búsqueda. No conoce cómo se arma una URL de Solr.
	return s.engine.Search(ctx, query)
}

// -----------------------------------------------------------------------------
// CONTROLADOR:
// -----------------------------------------------------------------------------
// searchController recibe GET /products/search. extrae q y
// limit de la URL, valida datos y llama al servicio. No conoce qf, rows ni la
// URL de Solr.
func searchController(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Leemos ?q=zapatillas y quitamos espacios accidentales.
		text := strings.TrimSpace(r.URL.Query().Get("q"))
		if text == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "q es obligatorio"})
			return
		}

		// Leemos ?limit=5 solo si fue enviado. Cero significa “usar default”.
		limit := 0
		if rawLimit := r.URL.Query().Get("limit"); rawLimit != "" {
			var err error
			limit, err = strconv.Atoi(rawLimit)
			if err != nil || limit <= 0 {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "limit debe ser un entero positivo"})
				return
			}
		}

		// Convertimos parámetros HTTP en Query y delegamos al servicio.
		result, err := service.Search(r.Context(), Query{
			Text: text, Category: r.URL.Query().Get("category"),
			Brand: r.URL.Query().Get("brand"), Limit: limit,
		})
		if err != nil {
			// La petición es válida; el problema ocurrió al buscar aguas abajo.
			// 502 comunica que falló una dependencia, por ejemplo Solr.
			writeJSON(w, http.StatusBadGateway, map[string]string{"error": "búsqueda no disponible"})
			return
		}
		writeJSON(w, http.StatusOK, result)
	}
}

// writeJSON es un detalle HTTP repetido: configura el tipo de contenido y
// convierte cualquier valor Go a JSON para enviarlo en la respuesta.
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

// -----------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------
func main() {
	// Esta es la única parte que conoce la dirección de Solr.
	client := SolrClient{
		baseURL: "http://localhost:8983/solr/products",
		http:    &http.Client{Timeout: 2 * time.Second},
	}

	// mux es el router: decide qué función atenderá cada URL HTTP.
	mux := http.NewServeMux()

	// Conectamos las capas: Controller -> service -> SolrClient.
	mux.Handle("GET /products/search", searchController(Service{engine: client}))

	log.Println("Search API disponible en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
