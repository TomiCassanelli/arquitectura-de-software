# Actividad 2 resuelta — API Go + Gin + Solr

La solución conserva el recorrido que se explicó en la diapositiva 20:

```text
Cliente HTTP → Handler → Service → Engine → SolrClient → Solr
```

## Cómo leer el código

1. Empezar en `main.go`: allí se conectan las capas.
2. Seguir al controlador: transforma HTTP (`q`, `limit`, `category`, `brand`) en una `Query` y valida errores del cliente.
3. Seguir al servicio: aplica el límite por defecto y delega en la interfaz `Engine`.
4. Seguir al repositorio Solr: transforma `Query` en parámetros Solr (`q`, `fq`, `rows`, `fl`) y traduce la respuesta JSON.

## Responsabilidades

| Archivo/capa | Qué sabe hacer | Qué no debe conocer |
|---|---|---|
| Controller | HTTP, parámetros y status codes | URL o JSON de Solr |
| Service | reglas de aplicación, como el limit por defecto | detalles de HTTP/Solr |
| Engine | contrato de búsqueda | implementación concreta |
| SolrClient | URL, parámetros y JSON de Solr | status HTTP de entrada |

## Ejecutar

Solr debe estar preparado desde `CLASE_5`. Luego, desde esta carpeta:

```bash
go test ./...
go run .
```

En otra terminal:

```bash
curl 'http://localhost:8080/products/search?q=running'
curl 'http://localhost:8080/products/search?q=running&category=calzado'
curl 'http://localhost:8080/products/search?q=zapatillas&brand=Adidas&limit=5'
curl -i 'http://localhost:8080/products/search'
curl -i 'http://localhost:8080/products/search?q=running&limit=0'
```
