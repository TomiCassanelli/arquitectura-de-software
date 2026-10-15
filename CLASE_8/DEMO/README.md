# Demo: observabilidad local con Prometheus

Esta demo arranca una API Go y cuatro herramientas con **un solo archivo de orquestación**, `compose.yaml`. `main.go` es el único punto de entrada de la API; los paquetes `controllers` y `observability` separan las responsabilidades del código. La librería oficial de Prometheus para Go cuenta las solicitudes y mide su duración; Prometheus consulta esos valores en `/metrics`. Los archivos de configuración contienen comentarios junto a cada decisión. El dashboard es JSON, que no admite comentarios, y se explica aquí. El proyecto Compose se llama `clase8-observabilidad` para no cruzarse con las demos de otras clases.

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
| Loki | http://localhost:3100/ready | Confirma que el depósito de logs está listo. |
| Alloy | http://localhost:12345 | Muestra el estado de la recolección de logs. |

Grafana crea automáticamente las fuentes de datos y el dashboard **Presentación / Demo de observabilidad**. La clave simple sirve exclusivamente para esta demo local.

## Probar

Abrí `http://localhost:8080/` para ver las rutas. Ejecutá algunas solicitudes:

```sh
curl http://localhost:8080/ok
curl http://localhost:8080/slow
curl http://localhost:8080/error
curl http://localhost:8080/metrics
```

`/ok` devuelve 200, `/slow` tarda unos 750 ms y `/error` devuelve 500. `/metrics` muestra el texto que consulta Prometheus cada cinco segundos. El dashboard también actualiza cada cinco segundos. Para generar más datos, repetí esos tres `curl`. En Prometheus, el target `demo-app` debe figurar como **UP** y la consulta `demo_http_requests_total{job="demo-app"}` muestra el contador. En Grafana Explore, elegí **Loki** y consultá `{service="app"}` para ver los mensajes. También podés verlos directamente con `docker compose logs -f app`.

## Cómo viajan los datos

```text
GET /ok, /slow, /error → API → contador e histograma → /metrics → Prometheus → Grafana
                           └→ stdout → Alloy → Loki → Grafana
```

- La API crea un contador y un histograma en `observability/metrics.go`. `requests.WithLabelValues(route, status).Inc()` suma una solicitud; `duration.WithLabelValues(route, status).Observe(...)` registra cuánto tardó. Las etiquetas permiten separar rutas y estados HTTP.
- `main.go` registra el middleware solo en el grupo `api`, que contiene `/`, `/ok`, `/slow` y `/error`. Gin lo ejecuta al entrar en esas rutas. En `metrics.go`, `c.Next()` ejecuta el controller y después vuelve al middleware para contar la respuesta y medir cuánto tardó. `/health` y `/metrics` quedan fuera de ese grupo, así que no contaminan el contador. `promhttp.Handler()` sirve `/metrics` en el formato que Prometheus entiende.
- `prometheus.yml` indica que Prometheus debe consultar `app:8080/metrics`. `app` es un nombre de servicio dentro de Compose, no `localhost` del contenedor de Prometheus.
- `alloy.config` descubre contenedores por el socket de Docker, conserva solo el servicio `app` y envía sus logs a Loki con la etiqueta `service="app"`.
- `grafana/provisioning/datasources/datasources.yml` crea los accesos a Prometheus y Loki. `grafana/provisioning/dashboards/dashboards.yml` carga `grafana/dashboards/demo.json`, que dibuja solicitudes por segundo, duración media y logs.
- Los volúmenes de Compose conservan métricas, logs y configuración de Grafana entre reinicios.

La app continúa usando logs en stdout para que Alloy los envíe a Loki. Prometheus almacena las métricas y Grafana las muestra. Esta configuración es una demostración local.

El socket de Docker concede acceso amplio al daemon incluso montado como solo lectura. Esta configuración es para una máquina de aprendizaje; antes de usarla en un entorno compartido hay que revisar ese acceso. El colector de logs solo muestra los logs de la aplicación, no los demás contenedores.

## Detener y limpiar

```sh
docker compose down
```

Si también querés borrar las métricas, logs y datos guardados en los volúmenes:

```sh
docker compose down -v
```
