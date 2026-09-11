#!/usr/bin/env bash

SOLR_URL="http://localhost:8983/solr/products/select"

# CONSULTA 2 — FILTRO, LÍMITE Y CAMPOS
#
# Usá el texto y el filtro asignados en el formulario. Construí aquí una única
# consulta que:
#   - busque el texto en title y description con eDisMax;
#   - aplique la condición de marca o categoría como filtro obligatorio;
#   - limite la respuesta a 5 documentos;
#   - devuelva id, title, brand, category, price y score en JSON.
#
# Identificá qué parte es q, cuál es fq, y qué parámetros controlan el límite
# y los campos de salida. Ejecutala y entregá el comando y la respuesta.
