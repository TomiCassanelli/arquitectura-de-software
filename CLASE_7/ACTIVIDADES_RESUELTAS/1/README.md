# Actividad 1 resuelta — Balanceador roto

Los dos errores eran:

1. El `upstream` solo tenía `api1`, por eso NGINX no podía repartir solicitudes.
2. Las tres APIs usaban `ports`, que publica puertos hacia la notebook. Los
   backends deben permanecer internos; solo NGINX publica `8080:80`.

La corrección usa `expose: "3000"` como documentación del puerto interno y
declara las tres APIs en `upstream`.

La API está separada en `handlers`, `services`, `dto` y `config`: el handler
conoce Gin/HTTP, el service conoce Open-Meteo y los DTO representan los JSON.

```sh
docker compose up --build -d
for i in {1..9}; do curl -s http://localhost:8080/clima; done
curl -i http://localhost:3002/clima
docker compose down
```

El primer comando alterna entre `Clima A`, `Clima B` y `Clima C`; el segundo no
puede conectarse, porque `api2` no tiene un puerto publicado en el host.
