# Demo — Un balanceador de carga simple

Esta demo enseña una sola idea: **NGINX reparte peticiones entre tres copias de
una API**.

```text
Cliente ──> NGINX ──> Servidor A
                     Servidor B
                     Servidor C
```

Los servicios `api1`, `api2` y `api3` ejecutan exactamente el mismo código Go.
Solo cambia la variable `NOMBRE_SERVIDOR`, para que sea fácil ver quién atendió
cada solicitud.

## Archivos

- `main.go`: conecta Gin, la ruta y el handler.
- `handlers/`: recibe la petición HTTP y devuelve la respuesta.
- `dto/`: define el JSON de salida de `GET /test`.
- `Dockerfile`: empaqueta la API en una imagen Docker simple.
- `compose.yaml`: crea las tres APIs y el balanceador.
- `nginx.conf`: lista los tres backends y deja que NGINX use round robin.

## Ejecutar la demo

Desde este directorio, con Docker Desktop iniciado:

```sh
docker compose up --build -d
```

Esperar unos segundos y hacer nueve peticiones:

```sh
for i in {1..9}; do
  curl -s http://localhost:8080/test
done
```

La respuesta irá alternando entre `Servidor A`, `Servidor B` y `Servidor C`.
El punto importante no es que el orden visual sea perfecto, sino que el cliente
siempre llama a una sola dirección (`localhost:8080`) y NGINX decide a qué API
enviarla.

Para terminar:

```sh
docker compose down
```

## Probar primero la API sin Docker (opcional)

```sh
NOMBRE_SERVIDOR="Servidor local" go run .
```

En otra terminal:

```sh
curl http://localhost:3000/test
```

No hay autoescalado, DNS dinámico ni health checks en esta demo. Tampoco se
crea un service: la respuesta es estática y no consulta una dependencia externa.
Esos temas se pueden presentar más adelante, una vez que el reparto básico esté
claro.
