#!/usr/bin/env bash

set -euo pipefail
SOLR_URL="http://localhost:8983/solr/products/select"

# Sin fq se observan todos los candidatos. Al comparar con la consulta
# filtrada de la actividad anterior, recordar que los scores se calculan en
# el contexto de la consulta: lo importante es el orden de cada respuesta.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'fl=id,title,score' \
  --data-urlencode 'wt=json'
printf '\n'

# El boost hace explícita una regla de relevancia del producto: una aparición
# en title pesa cinco veces la señal equivalente de description.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title^5 description' \
  --data-urlencode 'fl=id,title,score' \
  --data-urlencode 'wt=json'
printf '\n'
