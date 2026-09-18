# Ejercicios — base

Proyecto Go ya armado para trabajar los ejercicios de la Clase Práctica 4 (Concurrencia y Paralelismo en Go). Cada carpeta es un `main` independiente con la parte "de infraestructura" (simulación de latencias, generación de datos, etc.) ya resuelta, y la parte de concurrencia marcada con `// TODO` para completar.

## Cómo correr cada parte

Parado en esta carpeta (`ejercicios-base/`):

```bash
go run ./handson1
go run -race ./handson1        # ver el reporte de race condition

go run -race ./handson2        # ya viene resuelto con Mutex, para comparar

go run ./ejercicio1
go run ./ejercicio2
go run ./ejercicio3
go run ./ejercicio4/inseguro
go run -race ./ejercicio4/inseguro   # ver el reporte de race condition

go run ./ejercicio4/mutex      # completar el TODO
go run ./ejercicio4/channel    # completar el TODO
```

> `-race` necesita cgo + un compilador de C instalado. Ver los prerrequisitos en el [README general](../README.md#prerrequisitos) si falla con `-race requires cgo`.

Los ejercicios 1, 2 y 3 corren con `go run` normal (sin `-race`): no involucran memoria compartida sin proteger, así que no hay nada que el detector de races pueda marcar ahí.

## Hands On 1 — Reproducir una Race Condition

**Código deliberadamente incorrecto** (ver [handson1/race.go](handson1/race.go)):

```go
var stock int = 100
for i := 0; i < 1000; i++ {
  go func() {
    stock-- // ← data race
  }()
}
fmt.Println(stock) // ¿cuánto vale?
```

Las `totalSales` goroutines modifican `stock` sin sincronización. El resultado es impredecible.

**Pasos para reproducir:**
1. Guardar el código en `race.go` (ya está en `handson1/race.go`).
2. Ejecutar con el detector: `go run -race race.go` (o `go run -race ./handson1` desde esta carpeta).
3. Leer el reporte: Go muestra las dos goroutines en conflicto y la línea exacta.

## Hands On 2 — Corregir con sync.Mutex

**Solución con Mutex** (ver [handson2/race_mutex.go](handson2/race_mutex.go)):

```go
var (
  stock int = 100
  mu    sync.Mutex
)
for i := 0; i < 1000; i++ {
  go func() {
    mu.Lock()
    stock--
    mu.Unlock()
  }()
}
```

El Mutex garantiza acceso excluyente a la variable compartida.

**Verificación:** correr de nuevo con `go run -race`: el reporte desaparece y el resultado es predecible (siempre -900).

> 🔁 Alternativa: un channel como semáforo con capacidad 1. Comparar ambas soluciones en el Ejercicio 4.

## Ejercicio 1 — Fan-Out/Fan-In: ficha de producto

Simulá tres funciones que consultan precio, stock y reviews de un producto (cada una con `time.Sleep` de distinta duración, imitando latencia real). Lanzá las tres como goroutines independientes (fan-out) y usá un channel por cada una (o un channel compartido con `sync.WaitGroup`) para recolectar los resultados (fan-in) antes de armar y devolver la ficha completa del producto. Medí el tiempo total y compará contra la versión secuencial (llamando a las tres funciones una tras otra).

**Objetivo:** entender fan-out/fan-in y verificar que el tiempo total se acerca al de la fuente más lenta, no a la suma de las tres.

**Dónde completar:** [ejercicio1/main.go](ejercicio1/main.go), función `obtenerFichaConcurrente`.

## Ejercicio 2 — Worker Pool: actualización masiva de precios

Usá la lista `productIDs`. Cada producto tarda `processDelay` y el pool debe tener `workerCount` workers fijos. Los valores concretos los asigna el formulario mediante el bloque de variante.

**Objetivo:** entender por qué no conviene lanzar una goroutine por tarea sin control, y cómo el número de workers impacta el throughput.

**Dónde completar:** [ejercicio2/main.go](ejercicio2/main.go), función `procesarConWorkerPool`.

> ⏱ Compará el tiempo secuencial con el del worker pool usando los valores de `productIDs`, `processDelay` y `workerCount` de tu variante.

## Ejercicio 3 — Context: cortar una fuente lenta

Usá `reviewsDelay`, `timeoutLimit` y `defaultReviews` de tu variante. Si Reviews no responde antes de `timeoutLimit`, devolvé `defaultReviews` en vez de seguir esperando.

**Objetivo:** entender cómo Context evita que una fuente lenta bloquee todo el sistema, y practicar `select` con `ctx.Done()`.

**Dónde completar:** [ejercicio3/main.go](ejercicio3/main.go), función `obtenerFichaConTimeout`. La función `consultarReviews` de este archivo ya simula la fuente lenta (2s).

## Ejercicio 4 — Contador seguro: Mutex vs Channel (bonus)

Usá `initialStock`, `totalSales` y `saleDelay` de tu variante. Primero observá la versión sin protección y después resolvela con Mutex o, como bonus, con una goroutine dueña del estado.

**Objetivo:** ver una race condition real, entender dos estrategias válidas para resolverla (memoria compartida protegida vs. estado que solo cambia por mensajes) y cuándo conviene cada una.

**Dónde completar:**
- [ejercicio4/inseguro/main.go](ejercicio4/inseguro/main.go) — ya está completo, es el punto de partida con el bug.
- [ejercicio4/mutex/main.go](ejercicio4/mutex/main.go) — completar con `sync.Mutex`.
- [ejercicio4/channel/main.go](ejercicio4/channel/main.go) — completar con una goroutine dueña del estado + channel.

> El resultado correcto es `initialStock - totalSales`.
