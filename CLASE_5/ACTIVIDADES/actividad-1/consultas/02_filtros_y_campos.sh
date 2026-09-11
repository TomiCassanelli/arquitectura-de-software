#!/usr/bin/env bash

SOLR_URL="http://localhost:8983/solr/products/select"

# TODO 3
# Completar el filtro de categoría. Mantener la búsqueda textual dada.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'fq=TODO_CATEGORIA' \
  --data-urlencode 'wt=json'

# TODO 4, 5 y 6
# Completar el segundo filtro, el límite de resultados y los campos de salida.
# Cada filtro debe tener su propio parámetro fq.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'fq=TODO_CATEGORIA' \
  --data-urlencode 'fq=TODO_MARCA' \
  --data-urlencode 'rows=TODO_LIMITE' \
  --data-urlencode 'fl=TODO_CAMPOS' \
  --data-urlencode 'wt=json'
