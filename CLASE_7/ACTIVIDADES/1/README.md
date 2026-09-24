# Actividad 1 — Migración de una API a tres réplicas

## Situación

Una API de clima funcionaba en una única instancia. Un equipo intentó pasarla
a tres réplicas (`api1`, `api2` y `api3`) y colocar NGINX delante, pero el
traspaso quedó a medio terminar.

No hay una lista de errores ni marcas que indiquen dónde están. El objetivo es
investigar el escenario, encontrar los problemas y dejar una arquitectura que
funcione con tres servidores.

Los problemas pueden estar en `main.go`, en el service, en el handler, en
Docker Compose o en NGINX. La aplicación debe seguir consultando Open-Meteo y
la corrección no requiere agregar funcionalidades nuevas.

## Bloque de personalización asignado

Abrí el formulario que te asignó el docente. La primera sección muestra un
bloque de constantes propio. Copialo completo y reemplazá el bloque `const` de
`config/config.go` **antes de levantar** la actividad. No modifiques esos
valores ni agregues otros: identifican tu variante de entrega.

El bloque cambia ciudad, coordenadas y mensaje. Cambia los datos de la entrega,
no la arquitectura que debe corregirse.