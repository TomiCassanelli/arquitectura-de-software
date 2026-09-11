# Actividad 2 — API Go + Gin

Completar primero la [Actividad 1](../actividad-1/README.md) y los pasos 1 a 4 del [README de la Clase 5](../../README.md). Esta actividad integra Solr detrás de una API, manteniendo la separación controlador → servicio → interfaz → repositorio de Solr.

## Consigna

Implementar los `TODO` de estos archivos:

- `controllers/products/products_controller.go`: endpoint HTTP y validación de parámetros.
- `repositories/products/solr_repository.go`: consulta HTTP a Solr y conversión de la respuesta.

No hace falta modificar `main.go`, `models`, `services` ni `repositories/products/response.go`: son el código base que conecta las capas y representa la respuesta de Solr.

El endpoint esperado es:

```text
GET /products/search?q=running&category=calzado&brand=Adidas&limit=5
```

Una vez implementado el endpoint, cada estudiante recibe una variante (`V01` a `V10`) en el formulario. Debe transformar la indicación asignada en parámetros HTTP, ejecutar el endpoint y responder el resultado observado. Las variantes están generadas en [`../../FORMULARIO_GOOGLE`](../../FORMULARIO_GOOGLE/README.md).

- `q` es obligatorio.
- `limit` debe ser entero positivo; por defecto vale `10`.
- `category` y `brand` son opcionales y se envían a Solr como filtros exactos `fq`.
- La búsqueda debe usar `defType=edismax` y `qf=title description`.
- Responder `400` ante parámetros inválidos, `502` ante fallas de Solr y `200` con `results` ante una búsqueda exitosa.

## Ejecutar

Desde esta carpeta:

```bash
go mod tidy
go test ./...
go run .
```

La API usa `http://localhost:8983/solr/products`, por lo que el Solr iniciado desde `CLASE_5` debe seguir activo. El test de repositorio falla hasta que se implemente el cliente: es parte de la consigna.

También hay pruebas en `controllers/products/products_controller_test.go`. Al inicio fallan porque el endpoint todavía no está registrado; deben pasar al completar el controlador.
