#!/usr/bin/env bash

set -euo pipefail

SOLR_URL="http://localhost:8983/solr/products/select"

# CONSULTA 3
# q=running busca por texto. fq=category:calzado no cambia ese texto: descarta
# los candidatos que no son calzado. Por eso fq es una condición obligatoria.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'fq=category:calzado' \
  --data-urlencode 'wt=json'
printf '\n'

# CONSULTA 4, 5 y 6
# Los dos fq se acumulan: debe ser calzado Y Adidas. rows=5 limita la página;
# fl decide qué campos verá el cliente, incluido score para observar ranking.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'fq=category:calzado' \
  --data-urlencode 'fq=brand:Adidas' \
  --data-urlencode 'rows=5' \
  --data-urlencode 'fl=id,title,brand,category,price,score' \
  --data-urlencode 'wt=json'
printf '\n'
