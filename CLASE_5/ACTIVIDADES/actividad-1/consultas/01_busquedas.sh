#!/usr/bin/env bash

SOLR_URL="http://localhost:8983/solr/products/select"

# TODO 1
# Reemplazar TODO_BUSQUEDA_1 por el texto indicado en la consigna.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=TODO_BUSQUEDA_1' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'wt=json'

# TODO 2
# Reemplazar TODO_BUSQUEDA_2 por el segundo texto indicado en la consigna.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=TODO_BUSQUEDA_2' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'wt=json'
