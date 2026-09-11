#!/usr/bin/env bash

set -euo pipefail
SOLR_URL="http://localhost:8983/solr/products/select"

# fq no modifica el texto buscado: descarta los candidatos que no pertenecen
# a la categoría solicitada. Es el lugar apropiado para una restricción dura.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'fq=category:calzado' \
  --data-urlencode 'wt=json'
printf '\n'

# Se envía cada fq por separado: Solr los acumula. rows solo pagina la
# respuesta; fl reduce el contrato al conjunto de campos que necesita la UI.
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
