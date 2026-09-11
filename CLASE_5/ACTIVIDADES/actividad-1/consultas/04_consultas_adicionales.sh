#!/usr/bin/env bash

SOLR_URL="http://localhost:8983/solr/products/select"

# CONSULTA 5 — OPERADOR LÓGICO ENTRE DOS TÉRMINOS
#
# Construí la consulta indicada en el formulario. Puede pedir AND u OR: usá
# exactamente el operador asignado dentro de q, buscá con eDisMax en title y
# description y devolvé al menos id, title y score en JSON.


# CONSULTA 6 — TEXTO Y RANGO DE PRECIO
#
# Construí la búsqueda de texto asignada y restringila al rango de precios
# indicado en el formulario. El rango es un filtro obligatorio con formato
# price:[MIN TO MAX]. Devolvé id, title, price y score en JSON.
#
# Ejecutá las dos consultas y entregá ambos comandos con sus salidas completas.
