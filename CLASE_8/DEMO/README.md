# Demo: observabilidad local con Prometheus

Esta demo arranca una API Go, Prometheus, MySQL y Grafana con **un solo archivo de orquestación**, `compose.yaml`. `main.go` es el único punto de entrada de la API; los paquetes `controllers`, `observability` y `logstore` separan las responsabilidades. Prometheus guarda métricas y MySQL guarda los logs en una tabla. Los archivos de configuración contienen comentarios junto a cada decisión. El dashboard es JSON, que no admite comentarios, y se explica aquí. El proyecto Compose se llama `clase8-observabilidad` para no cruzarse con las demos de otras clases.

## Iniciar

Desde esta carpeta:

```sh
docker compose up --build -d
```

La primera ejecución descarga imágenes y compila Go. Necesitás Docker Desktop o un daemon Docker activo.

| Servicio | Dirección | Para qué sirve |
| --- | --- | --- |
| API | http://localhost:8080 | Genera respuestas y logs. |
| Prometheus | http://localhost:9090/targets | Comprueba si la app entrega métricas. |
| Grafana | http://localhost:3000 | Muestra métricas y logs. Usuario y clave: `admin` / `admin`. |
| MySQL | Dentro de Compose: `mysql:3306` | Guarda una fila por cada evento de la app. |

Grafana crea automáticamente las fuentes de datos y el dashboard **Presentación / Demo de observabilidad**. La clave simple sirve exclusivamente para esta demo local.

## Probar

Abrí `http://localhost:8080/` para ver las rutas. Ejecutá algunas solicitudes:

```sh
curl http://localhost:8080/ok
curl http://localhost:8080/slow
curl http://localhost:8080/error
curl http://localhost:8080/metrics
```

`/ok` devuelve 200, `/slow` tarda unos 750 ms y `/error` devuelve 500. `/metrics` muestra el texto que consulta Prometheus cada cinco segundos. El dashboard también actualiza cada cinco segundos. Para generar más datos, repetí esos tres `curl`. En Prometheus, el target `demo-app` debe figurar como **UP** y la consulta `requests_total{job="demo-app"}` muestra el contador. En Grafana, el panel **Logs de la aplicación** lee la tabla de MySQL. También podés abrir Explore, elegir **MySQL Logs**, seleccionar formato **Table** y consultar:

```sql
SELECT created_at AS time, message AS body, level
FROM app_logs
WHERE $__timeFilter(created_at)
ORDER BY created_at DESC;
```

Para comprobar la tabla directamente sin instalar un cliente MySQL en tu equipo:

```sh
docker compose exec mysql mysql -u demo -pdemo observabilidad -e "SELECT * FROM app_logs ORDER BY id DESC LIMIT 10;"
```

## Cómo viajan los datos

```text
GET /ok, /slow, /error → API → contador e histograma → /metrics → Prometheus → Grafana
                           └→ INSERT en app_logs → MySQL → Grafana
```

- La API crea un contador y un histograma en `observability/metrics.go`. `requests.WithLabelValues(route, status).Inc()` suma una solicitud; `duration.WithLabelValues(route, status).Observe(...)` registra cuánto tardó. Las etiquetas permiten separar rutas y estados HTTP.
- `main.go` registra el middleware solo en el grupo `api`, que contiene `/`, `/ok`, `/slow` y `/error`. Gin lo ejecuta al entrar en esas rutas. En `metrics.go`, `c.Next()` ejecuta el controller y después vuelve al middleware para contar la respuesta y medir cuánto tardó. `/health` y `/metrics` quedan fuera de ese grupo, así que no contaminan el contador. `promhttp.Handler()` sirve `/metrics` en el formato que Prometheus entiende.
- `prometheus.yml` indica que Prometheus debe consultar `app:8080/metrics`. `app` es un nombre de servicio dentro de Compose, no `localhost` del contenedor de Prometheus.
- `logstore/mysql.go` conecta la app con MySQL y hace un `INSERT` por mensaje. `controllers/demo.go` registra los eventos `info`, `warn` y `error`. `mysql/001_schema.sql` crea la tabla `app_logs` y un usuario de solo lectura para Grafana. Ese SQL se ejecuta únicamente cuando se crea un volumen MySQL nuevo.
- `grafana/provisioning/datasources/datasources.yml` crea los accesos a Prometheus y MySQL. `grafana/provisioning/dashboards/dashboards.yml` carga `grafana/dashboards/demo.json`, que dibuja solicitudes por segundo, duración media y los mensajes consultados con SQL.
- Los volúmenes de Compose conservan métricas, filas de MySQL y configuración de Grafana entre reinicios.

Las credenciales `demo`, `grafana_demo` y `admin` son solo para esta demo local. Grafana puede leer la tabla, pero no escribirla. Como la app guarda mensajes directamente en MySQL, si la base de datos falla el evento no queda en la tabla; en ese caso la app escribe un aviso en `docker compose logs app`.

Si venías de la versión con Alloy y Loki, `docker compose up --build -d --remove-orphans` retira sus contenedores antiguos. Sus volúmenes anteriores no se borran automáticamente.

## Detener y limpiar

```sh
docker compose down
```

Si también querés borrar las métricas, logs y datos guardados en los volúmenes:

```sh
docker compose down -v
```
