#!/usr/bin/env bash

set -euo pipefail

SOLR_URL="http://localhost:8983/solr/products/select"

# CONSULTA 1
# q es el texto que escribió la persona. qf indica que Solr debe buscarlo
# tanto en title como en description. Todavía no hay filtros: queremos ver
# todos los productos relacionados con “zapatillas”.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=zapatillas' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'wt=json'
printf '\n'

# CONSULTA 2
# Cambiamos únicamente q. Así podemos comparar los resultados y scores de
# “running” sin mezclar cambios de filtros ni de campos de búsqueda.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'wt=json'
printf '\n'
