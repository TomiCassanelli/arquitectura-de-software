# Actividad 2 resuelta — API Go + Gin + Solr

La solución mantiene la separación propuesta por la consigna:

```text
HTTP/Gin Handler/Controller -> Service -> Engine (interfaz) -> Solr Client -> Solr
```

## Qué resuelve cada capa

- El **controller** valida `q` y `limit`, convierte parámetros HTTP a `models.Query` y traduce errores a `400`, `502` o `200`.
- El **service** aplica el límite por defecto (`10`) y depende de la interfaz `Engine`, no de Solr.
- El **cliente Solr** construye y codifica la consulta, aplica `fq` literales para filtros, verifica el status y traduce el JSON externo al contrato de la API.
- `models.Response` conserva `total` desde `numFound`, incluso si `rows` limita la cantidad de `results`.

## Ejecutar

Con Solr preparado desde `CLASE_5`:

```bash
go mod tidy
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

## Puntos para explicar al corregir

1. `q` es obligatorio porque sin texto no hay intención de búsqueda; `category` y `brand` son opcionales porque restringen una intención existente.
2. `fq` no es otro `q`: elimina documentos no elegibles y no representa una señal de relevancia.
3. `url.Values` se ocupa del encoding de URL y `{!term}` trata los filtros como valores literales.
4. Un fallo o timeout de Solr produce `502`: el request del cliente era correcto, pero falló una dependencia aguas abajo.
