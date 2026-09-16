# Actividad 1 resuelta — Consultas, filtros y relevancia

Los scripts contienen las consultas completas y comentarios sobre cada parámetro. Ejecutarlos desde esta carpeta:

```bash
bash consultas/01_busquedas.sh
bash consultas/02_filtros_y_campos.sh
bash consultas/03_relevancia.sh
```

## Qué comprobar

| Concepto | Qué hace |
|---|---|
| `q` | El texto que busca la persona. Recupera candidatos y participa en el score. |
| `fq` | Un filtro obligatorio; elimina documentos no permitidos. Puede enviarse más de una vez. |
| `qf` | Los campos de texto donde Solr debe buscar `q`. |
| `rows` | Cuántos resultados devuelve como máximo; no cambia el total de coincidencias. |
| `fl` | Qué campos aparecen dentro de cada resultado. |
| `score` | La relevancia relativa que Solr calculó para esa consulta. |

## Conclusión esperada

`q` expresa la intención de búsqueda; `fq` agrega condiciones obligatorias; `score` ordena por relevancia a los documentos que siguen siendo elegibles. Un score no es una nota absoluta: se interpreta dentro de la misma consulta y configuración.
