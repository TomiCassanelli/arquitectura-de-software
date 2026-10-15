# Demo: observabilidad local con Prometheus

Esta demo arranca una API Go, Prometheus, MySQL y Grafana con **un solo archivo de orquestación**, `compose.yaml`. `main.go` es el único punto de entrada de la API; los paquetes `controllers`, `observability` y `logs` separan las responsabilidades. Prometheus guarda métricas y MySQL guarda los logs en una tabla. Los archivos de configuración contienen comentarios junto a cada decisión. El dashboard es JSON, que no admite comentarios, y se explica aquí. El proyecto Compose se llama `clase8-observabilidad` para no cruzarse con las demos de otras clases.

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

`/ok` devuelve 200, `/slow` tarda unos 750 ms y `/error` devuelve 500. `/metrics` muestra el texto que consulta Prometheus cada cinco segundos. El dashboard también actualiza cada cinco segundos. Para generar más datos, repetí esos tres `curl`. En Prometheus, el target `demo-app` debe figurar como **UP** y la consulta `requests_total{job="demo-app"}` muestra el contador.

El panel **Porcentaje de fallos HTTP 5xx** compara las respuestas 5xx con el resto durante el rango de tiempo seleccionado en Grafana.

### Seguir una compra con track_id

Estas rutas están separadas de las tres rutas básicas. **No hay microservicios adicionales**: la misma API simula las etapas `gateway`, `pedidos` y `pagos`. Cada llamada guarda sus pasos con el mismo `track_id`:

```sh
curl http://localhost:8080/compra/ok
curl http://localhost:8080/compra/error
curl http://localhost:8080/compra/error-pedidos
```

La respuesta incluye, por ejemplo, `"track_id":"a1b2..."`; también aparece en el header `X-Track-ID`. Copiá el ID completo y buscá sus pasos en Grafana: **Explore → MySQL Logs → formato Table**. La consulta puede ser:

```sql
SELECT created_at, stage, track_id, level, message
FROM app_logs
WHERE track_id = 'PEGAR_EL_ID_AQUI'
ORDER BY id;
```

`/compra/ok` recorre las tres etapas sin errores. `/compra/error` recorre las tres y registra `pago rechazado` con `stage = pagos`. `/compra/error-pedidos` registra `fallo al verificar stock` con `stage = pedidos` y termina ahí: no habrá un paso de Pagos para ese ID. Ambos errores devuelven HTTP 500, así que también aparecen en la torta de fallos.

El panel **Fallos por etapa en el período** compara las respuestas 500 de cada ruta de error dentro del rango de tiempo seleccionado. Lee el contador HTTP que ya existe: la barra **Pedidos** usa `/compra/error-pedidos` y la barra **Pagos** usa `/compra/error`. No se agregó otra métrica al código Go. `increase()` estima el cambio entre muestras y Grafana muestra el valor sin decimales. Si acabás de hacer las primeras llamadas, esperá dos consultas de Prometheus para que pueda calcular ese cambio.

También podés pegar el ID en el campo **Track ID (filtra logs)** arriba del dashboard. **Logs del recorrido simulado** mostrará los pasos y **Fallo por etapa simulada** señalará `pagos` o `pedidos` según la ruta que llamaste. El campo vacío muestra todos los logs. El filtro afecta solo a los paneles de logs; las métricas siguen agregando todas las peticiones. Los logs de las rutas básicas tienen `track_id` vacío y `stage = api`. En un sistema real, cada microservicio escribiría su propio nombre y recibiría el mismo ID en un header.

### Ver todos los logs

En Grafana, los paneles de logs leen la tabla de MySQL. También podés abrir Explore, elegir **MySQL Logs**, seleccionar formato **Table** y consultar:

```sql
SELECT created_at AS time, stage, track_id, message AS body, level
FROM app_logs
WHERE $__timeFilter(created_at)
ORDER BY created_at DESC, id DESC;
```

Para comprobar la tabla directamente sin instalar un cliente MySQL en tu equipo:

```sh
docker compose exec mysql mysql -u demo -pdemo observabilidad -e "SELECT * FROM app_logs ORDER BY id DESC LIMIT 10;"
```

## Cómo viajan los datos

```text
GET /ok, /slow, /error, /compra/... → API → contador e histograma → /metrics → Prometheus → Grafana
                                      └→ INSERT en app_logs → MySQL → Grafana

En /compra, stage indica qué etapa se está simulando dentro de la misma API.
```

- La API crea un contador y un histograma en `observability/metrics.go`. `httpRequests.WithLabelValues(route, statusCode).Inc()` suma una solicitud; `requestDuration.WithLabelValues(route, statusCode).Observe(...)` registra cuánto tardó. Las etiquetas permiten separar rutas y estados HTTP.
- `main.go` registra el middleware de métricas en el grupo `apiRoutes`, que contiene las rutas básicas y `/compra/...`. Gin lo ejecuta al entrar en esas rutas. En `metrics.go`, `ctx.Next()` ejecuta el controller y después vuelve al middleware para contar la respuesta y medir cuánto tardó. `/health` y `/metrics` quedan fuera de ese grupo, así que no contaminan el contador. `promhttp.Handler()` sirve `/metrics` en el formato que Prometheus entiende.
- `observability/track_id.go` contiene `AddTrackID`, el middleware que agrega `track_id` solo a las peticiones de `/compra/...`. Crea el ID antes de entrar al controller y lo devuelve en `X-Track-ID`. `controllers/compra.go` escribe cada paso usando ese ID; por eso se pueden filtrar juntos en MySQL y Grafana. Tanto este middleware como `RecordMetrics` reciben directamente `*gin.Context`; `main.go` los pasa a `Use` sin paréntesis porque Gin los ejecuta al llegar cada petición.
- `prometheus.yml` indica que Prometheus debe consultar `app:8080/metrics`. `app` es un nombre de servicio dentro de Compose, no `localhost` del contenedor de Prometheus.
- `logs/mysql.go` conecta la app con MySQL mediante GORM. `Log` contiene `track_id`, `stage`, nivel y mensaje; `Table("app_logs").Create(&Log{...})` guarda cada evento en esa tabla. GORM genera el `INSERT` internamente. `AutoMigrate` agrega la columna `stage` si ya existía un volumen MySQL de la demo anterior. `controllers/demo.go` registra los eventos básicos y `controllers/compra.go` las etapas simuladas. `mysql/001_schema.sql` crea la tabla y un usuario de solo lectura para Grafana cuando el volumen es nuevo.
- `grafana/provisioning/datasources/datasources.yml` crea los accesos a Prometheus y MySQL. `grafana/provisioning/dashboards/dashboards.yml` carga `grafana/dashboards/demo.json`, que dibuja solicitudes por segundo, duración media, porcentaje de fallos, fallos por etapa y los logs. La variable de texto `track_id` filtra los dos paneles SQL; `${track_id:sqlstring}` cita el valor antes de usarlo en MySQL.
- Los volúmenes de Compose conservan métricas, filas de MySQL y configuración de Grafana entre reinicios.

Las credenciales `demo`, `grafana_demo` y `admin` son solo para esta demo local. Grafana puede leer la tabla, pero no escribirla. Como la app guarda mensajes directamente en MySQL, si la base de datos falla el evento no queda en la tabla; en ese caso la app escribe un aviso en `docker compose logs app`.

## Detener y limpiar

```sh
docker compose down
```

Si también querés borrar las métricas, logs y datos guardados en los volúmenes:

```sh
docker compose down -v
```
