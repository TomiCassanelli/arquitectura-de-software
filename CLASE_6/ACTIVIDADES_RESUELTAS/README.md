# Ejercicios — soluciones

Mismo proyecto que [`ejercicios-base/`](../ejercicios-base/), con cada `TODO` resuelto. El código de infraestructura (simulación de latencias, generación de datos) es idéntico a la base a propósito: lo único que cambia es la parte de concurrencia.

Cada función/archivo modificado tiene un comentario arriba explicando **qué cambió respecto a la base y por qué**, pensado como guía para revisar después de haber intentado resolverlo (no para copiar y pegar antes).

## Cómo correr cada solución

Parado en esta carpeta (`soluciones/`):

```bash
go run ./handson1
go run -race ./handson1        # ver el reporte de race condition

go run ./handson2
go run -race ./handson2        # el reporte ya no aparece

go run ./ejercicio1
go run ./ejercicio2             # tarda ~1 minuto en total (ver nota abajo)
go run ./ejercicio3

go run ./ejercicio4/inseguro
go run -race ./ejercicio4/inseguro   # ver el reporte de race condition

go run ./ejercicio4/mutex
go run -race ./ejercicio4/mutex      # el reporte ya no aparece

go run ./ejercicio4/channel
go run -race ./ejercicio4/channel    # el reporte ya no aparece
```

> `-race` necesita cgo + un compilador de C instalado. Ver los prerrequisitos en el [README general](../README.md#prerrequisitos) si falla con `-race requires cgo`.

## Cómo comparar contra la base

La forma más simple es un diff archivo por archivo. `git diff --no-index` funciona aunque los dos proyectos sean módulos Go separados (no hace falta que estén en el mismo repo git para usarlo así):

```bash
git diff --no-index ../ejercicios-base/ejercicio1/main.go ejercicio1/main.go
git diff --no-index ../ejercicios-base/ejercicio2/main.go ejercicio2/main.go
git diff --no-index ../ejercicios-base/ejercicio3/main.go ejercicio3/main.go
git diff --no-index ../ejercicios-base/ejercicio4/mutex/main.go   ejercicio4/mutex/main.go
git diff --no-index ../ejercicios-base/ejercicio4/channel/main.go ejercicio4/channel/main.go
```

(En Windows con PowerShell, el mismo comando `git diff --no-index ...` funciona igual si tenés Git instalado.)

También se pueden abrir los dos archivos lado a lado en el editor — cada ejercicio tiene el mismo nombre de archivo en ambas carpetas para facilitar justo esta comparación.

## Qué mirar en cada solución

| Ejercicio | Qué cambia respecto a la base |
|---|---|
| Hands On 1 | Nada — es el mismo código buggy, se deja para observar el reporte de `-race` antes de pasar al Hands On 2. Comentario explica qué información da el reporte. |
| Hands On 2 | Se agrega `sync.Mutex` alrededor de `stock--`. |
| Ejercicio 1 | `obtenerFichaConcurrente` pasa de llamar tres funciones en secuencia a lanzarlas como goroutines y recolectar por channels (fan-out/fan-in). |
| Ejercicio 2 | `procesarConWorkerPool` agrega un channel de `jobs`, N goroutines fijas consumiéndolo, un channel de `resultados`, y una goroutine auxiliar que cierra `resultados` cuando el `WaitGroup` confirma que todos los workers terminaron. |
| Ejercicio 3 | `obtenerFichaConTimeout` agrega `context.WithTimeout(300ms)` y un `select` entre el channel de reviews y `ctx.Done()`. |
| Ejercicio 4 (a) | El contador inseguro se protege con `sync.Mutex` alrededor de `stock--`. |
| Ejercicio 4 (b) | El contador deja de ser una variable compartida: pasa a vivir dentro de una goroutine "dueña" que recibe pedidos de decremento por un channel `chan struct{}`. |

> ⏱ El Ejercicio 2 corre la versión secuencial (1000 productos × 50ms ≈ 50s) y después el worker pool con N=5 (~10s) y N=20 (~2.5s). Es normal que tarde alrededor de un minuto en total.
