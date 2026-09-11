# Actividad 1 resuelta — Consultas, filtros y relevancia

Los scripts dejan las respuestas JSON completas para que se puedan observar `numFound`, `docs` y `score`. Ejecutarlos desde esta carpeta:

```bash
bash consultas/01_busquedas.sh
bash consultas/02_filtros_y_campos.sh
bash consultas/03_relevancia.sh
```

## Conclusión esperada

- `q` representa el texto de la búsqueda: recupera candidatos y participa en el cálculo de relevancia.
- `fq` es una condición obligatoria sobre los candidatos, como categoría o marca. Restringe el conjunto elegible y no es una señal para el score.
- `score` expresa la relevancia calculada por la consulta y la configuración de campos. Sirve para ordenar resultados de una misma consulta; no debe interpretarse como una medida absoluta ni compararse ciegamente entre consultas distintas.

El boost `title^5` del desafío muestra una política sencilla: una coincidencia en título aporta más señal que una coincidencia únicamente en descripción.
