# Actividad 1 — Construcción de consultas a Solr

En esta actividad vas a escribir y ejecutar consultas HTTP directamente contra Solr. No hay código Go ni una API intermedia: el objetivo es decidir qué parámetros necesita cada consulta, construir el comando `curl`, observar su respuesta y entregarla en tu formulario individual.

La actividad tiene **seis consultas** y se resuelve en los cuatro archivos de la carpeta `consultas`. Todos los archivos deben quedar completados.

## Antes de comenzar

1. Seguí los pasos 1 a 4 del [README de la Clase 5](../../README.md).
2. Comprobá que Solr esté disponible en <http://localhost:8983/solr/>.
3. Verificá que el core `products` tenga 15 documentos. Podés hacer una consulta con `q=*:*` y confirmar que `numFound` sea `15`.
4. Abrí el formulario que te asignó el docente. Allí figuran los textos, filtros y rangos que corresponden a tu variante.

Los comandos se ejecutan desde esta carpeta. Por ejemplo:

```bash
bash consultas/01_busquedas.sh
```

Cada script usa la variable `SOLR_URL`; escribí debajo de cada consigna tus propios comandos `curl`. Cuando termines un archivo, ejecutalo y pegá en el formulario tanto el comando construido como la salida JSON completa.

## Único ejemplo de forma

Este ejemplo muestra solamente la estructura de una petición GET. Está incompleto a propósito: no contiene todos los parámetros que requiere una consulta de la actividad ni valores de una variante.

```bash
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=TU_TEXTO' \
  # agregá aquí los demás parámetros necesarios
```

Usá `--data-urlencode` para cada parámetro: evita problemas cuando el valor tiene espacios, operadores o caracteres especiales. Agregá `wt=json` para recibir una respuesta fácil de leer y, si lo necesitás, `printf '\n'` después de cada comando para separar respuestas en la terminal.

## Qué resolver en cada archivo

| Archivo | Consultas que tenés que construir |
|---|---|
| `01_busquedas.sh` | Una búsqueda de texto de tu variante. |
| `02_filtros_y_campos.sh` | La búsqueda de tu variante con el filtro asignado, límite de 5 y los campos solicitados. |
| `03_relevancia.sh` | Dos veces la misma búsqueda: una sin boost y otra priorizando `title`. |
| `04_consultas_adicionales.sh` | Una consulta que exija dos términos con `AND` y otra con texto más un rango de precios. |

No reutilices el mismo comando sin revisarlo: cada consigna cambia los parámetros necesarios. Conservá en los scripts las consultas finales que usaste.

## Guía de decisiones

| Parámetro | Para qué sirve | Cuándo usarlo |
|---|---|---|
| `q` | Expresa el texto que se quiere buscar. Sus coincidencias participan en el orden por relevancia. | En todas las consultas de la actividad. |
| `defType=edismax` | Indica que Solr interprete la búsqueda con el parser eDisMax. | En las búsquedas por texto. |
| `qf` | Define en qué campos de texto se busca el contenido de `q`. | Usá `title description`; en la comparación de relevancia también usarás una versión que da más peso a `title`. |
| `fq` | Impone una condición obligatoria y descarta resultados que no la cumplan. Puede repetirse si hay varios filtros. | Para marca, categoría o rango de precio. |
| `rows` | Limita cuántos documentos aparecen en la respuesta. No cambia `numFound`. | En la consulta con filtros, con valor `5`. |
| `fl` | Elige qué campos muestra cada documento devuelto. | Cuando la consigna pide campos específicos o cuando necesitás comparar resultados. |
| `wt=json` | Pide la respuesta en JSON. | En todas las consultas. |
| `score` | Es la relevancia relativa calculada para esa consulta. | Incluilo en `fl` cuando la consigna solicita observar el ranking. |

Para un rango inclusivo de precios, la sintaxis de Solr es `price:[MIN TO MAX]`; esa condición corresponde a `fq`, no a `q`.
