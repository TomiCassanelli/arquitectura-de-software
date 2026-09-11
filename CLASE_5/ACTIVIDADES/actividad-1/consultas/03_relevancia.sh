#!/usr/bin/env bash

SOLR_URL="http://localhost:8983/solr/products/select"

# TODO 7
# Ejecutar esta consulta y comparar el score con los obtenidos en 02_filtros_y_campos.sh.
# Escribir la conclusión en la entrega.

# DESAFÍO
# Reemplazar TODO_PESO_TITLE por un peso que priorice title sobre description.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title^TODO_PESO_TITLE description' \
  --data-urlencode 'fl=id,title,score' \
  --data-urlencode 'wt=json'
