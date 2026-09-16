#!/usr/bin/env bash

set -euo pipefail

SOLR_URL="http://localhost:8983/solr/products/select"

# CONSULTA 7
# Sin filtros: muestra todos los candidatos de “running”. Comparar estos scores
# con los de la actividad anterior, recordando que score es relativo al contexto
# de cada consulta; lo que interesa es entender cómo se ordena cada respuesta.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title description' \
  --data-urlencode 'fl=id,title,score' \
  --data-urlencode 'wt=json'
printf '\n'

# DESAFÍO
# title^5 aplica un boost: una coincidencia en title aporta cinco veces más
# señal que la equivalente en description. No es un filtro: puede cambiar el
# orden y el score de los documentos que ya eran candidatos.
curl -fsS -G "$SOLR_URL" \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title^5 description' \
  --data-urlencode 'fl=id,title,score' \
  --data-urlencode 'wt=json'
printf '\n'
