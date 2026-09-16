# Demo 1 — Consultas Solr y Search API en Go

La demo sigue el recorrido de las diapositivas 19, 20 y 25:

```text
cliente HTTP -> handler -> Service -> Engine -> SolrClient -> Solr
```

`consultas.sh` permite mostrar primero los parámetros de Solr de forma directa. `main.go` muestra después cómo una API oculta esos parámetros detrás de `GET /products/search`.

## Preparación

Desde `CLASE_5`, levantar Solr, crear el schema y cargar los productos según el [setup común](../../README.md). Deben quedar cinco documentos indexados.

## Parte A: consultas directas

Desde esta carpeta:

```bash
bash consultas.sh
```

El script muestra, en este orden: verificación de documentos, texto libre, búsqueda con `AND` y categoría, y filtros acumulados. Leer cada comentario antes de ejecutar la consulta siguiente.

## Parte B: API en Go

En una terminal, desde esta carpeta:

```bash
go run .
```

En otra terminal:

```bash
curl 'http://localhost:8080/products/search?q=running'
curl 'http://localhost:8080/products/search?q=running&category=calzado'
curl 'http://localhost:8080/products/search?q=zapatillas&brand=Adidas&limit=5'
```

Para hacer visibles las validaciones:

```bash
curl -i 'http://localhost:8080/products/search'
curl -i 'http://localhost:8080/products/search?q=running&limit=0'
```

La respuesta usa el contrato de la API (`results` y `total`), no la forma interna de Solr (`response.docs` y `numFound`). El handler conoce HTTP; `SolrClient` conoce URLs y parámetros de Solr.
