# Actividad 1 — Consultas, filtros y relevancia

Esta actividad es un laboratorio de consultas a Solr. No hay API ni código Go: el trabajo consiste en completar los `TODO` de los scripts dentro de `consultas`.

## Estructura de trabajo

```text
consultas/
├── 01_busquedas.sh          # TODO 1 y 2: búsqueda por texto
├── 02_filtros_y_campos.sh  # TODO 3 a 6: filtros y forma de respuesta
└── 03_relevancia.sh         # TODO 7 y desafío: score y peso de campos
```

## Antes de empezar

Completar los pasos 1 a 4 del [README de la Clase 5](../../README.md). El core `products` debe tener los cinco productos cargados antes de ejecutar los scripts.

Resolver los archivos en orden. Después de reemplazar los `TODO` de un archivo, ejecutarlo desde esta carpeta:

```bash
bash consultas/01_busquedas.sh
```

## Consigna

1. En `01_busquedas.sh`, realizar las búsquedas por `zapatillas` y `running`. Observar la cantidad, los títulos y el `score` de los resultados.
2. En `02_filtros_y_campos.sh`, partir de `running`, filtrar primero por `category:calzado` y luego sumar `brand:Adidas`. Limitar la respuesta a cinco documentos y pedir solamente `id`, `title`, `brand`, `category`, `price` y `score`.
3. En `03_relevancia.sh`, comparar el `score` antes y después de filtrar. Luego, en el desafío, dar mayor peso a una coincidencia en `title` que a una coincidencia en `description`.

## Conceptos a comprobar

- `q` contiene el texto que se busca.
- `qf` define los campos donde Solr busca ese texto.
- `fq` filtra resultados; se puede enviar más de una vez.
- `rows` limita la cantidad de documentos devueltos.
- `fl` limita los campos incluidos en la respuesta.

Como entrega, conservar las consultas completadas y escribir una conclusión breve sobre la diferencia entre `q`, `fq` y `score`.
