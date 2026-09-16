#!/usr/bin/env bash

# Esta demo permite ejecutar cada consulta de la clase sin armar URLs a mano.
# Requiere haber completado el setup de ../../README.md.
set -euo pipefail

SOLR_URL="http://localhost:8983/solr/products/select"

consulta() {
  local titulo="$1"
  shift
  printf '\n=== %s ===\n' "$titulo"
  curl -fsS -G "$SOLR_URL" "$@" --data-urlencode 'wt=json'
  printf '\n'
}

# q=*:* no es una búsqueda por texto: sirve para comprobar que los cinco
# documentos llegaron al índice.
consulta '1. Todos los documentos' \
  --data-urlencode 'q=*:*' \
  --data-urlencode 'fl=id,title,brand,category,price'

# q participa en la recuperación y el ranking. qf decide dónde buscar y el
# boost ^5 hace que una coincidencia en title pese más que description.
consulta '2. Texto libre: zapatillas' \
  --data-urlencode 'q=zapatillas' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title^5 description' \
  --data-urlencode 'fl=id,title,score'

# AND se escribe de forma explícita: con eDisMax, separar términos con un
# espacio no obliga a que ambos estén presentes.
consulta '3. Texto + filtro + respuesta limitada' \
  --data-urlencode 'q=zapatillas AND running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title^5 description' \
  --data-urlencode 'fq=category:calzado' \
  --data-urlencode 'rows=5' \
  --data-urlencode 'fl=id,title,brand,category,price,score'

# fq restringe qué documentos son elegibles. No agrega una señal al score;
# por eso es adecuado para filtros obligatorios como categoría y marca.
consulta '4. Dos filtros acumulados' \
  --data-urlencode 'q=running' \
  --data-urlencode 'defType=edismax' \
  --data-urlencode 'qf=title^5 description' \
  --data-urlencode 'fq=category:calzado' \
  --data-urlencode 'fq=brand:Adidas' \
  --data-urlencode 'rows=5' \
  --data-urlencode 'fl=id,title,brand,category,price,score'
