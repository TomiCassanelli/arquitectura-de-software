#!/usr/bin/env bash

set -euo pipefail
SOLR_URL="http://localhost:8983/solr/products/select"

# q participa en recuperación y ranking. qf selecciona los campos de texto;
# todavía no aplicamos filtros para poder comparar los resultados puros.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=zapatillas' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'wt=json'
printf '\n'

# La segunda consulta usa el mismo parser y campos: así la diferencia que se
# observa proviene del texto buscado (`running`) y no de otra configuración.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'wt=json'
printf '\n'
