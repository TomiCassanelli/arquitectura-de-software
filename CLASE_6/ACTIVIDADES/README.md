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

1000 goroutines modifican `stock` sin sincronización. El resultado es impredecible.

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

Tenés una lista de 1000 IDs de productos a los que hay que aplicarles un aumento de precio (simulado con `time.Sleep(50 * time.Millisecond)` por producto). Implementá un worker pool con N workers fijos (probá con 5 y con 20) que toman IDs de un channel de trabajos y devuelven el resultado por un channel de resultados. Usá `sync.WaitGroup` para saber cuándo terminaron todos los workers y cerrar los channels correctamente. Medí el tiempo total con distintos valores de N.

**Objetivo:** entender por qué no conviene lanzar una goroutine por tarea sin control, y cómo el número de workers impacta el throughput.

**Dónde completar:** [ejercicio2/main.go](ejercicio2/main.go), función `procesarConWorkerPool`.

> ⏱ La versión secuencial de este ejercicio tarda ~50 segundos (1000 productos × 50ms). Es intencional: hace evidente la diferencia contra el worker pool.

## Ejercicio 3 — Context: cortar una fuente lenta

Sobre el Ejercicio 1, hacé que la consulta de "reviews" tarde deliberadamente más de lo razonable (por ejemplo 2 segundos). Agregá un `context.WithTimeout` de 300ms a esa consulta específica: si no responde a tiempo, la ficha del producto se devuelve igual, pero sin reviews (o con un valor por defecto), en vez de bloquear toda la respuesta. Usá `select` para elegir entre el resultado del channel y `ctx.Done()`.

**Objetivo:** entender cómo Context evita que una fuente lenta bloquee todo el sistema, y practicar `select` con `ctx.Done()`.

**Dónde completar:** [ejercicio3/main.go](ejercicio3/main.go), función `obtenerFichaConTimeout`. La función `consultarReviews` de este archivo ya simula la fuente lenta (2s).

## Ejercicio 4 — Contador seguro: Mutex vs Channel (bonus)

Escribí un contador de stock compartido que 50 goroutines decrementan al mismo tiempo (simulando 50 compras simultáneas del mismo producto). Primero corré la versión sin protección con `go run -race` y observá el reporte de race condition. Después resolvelo de dos formas distintas: (a) protegiendo el contador con `sync.Mutex`, y (b) reemplazando el contador por una goroutine dueña del estado que recibe pedidos de decremento por un channel. Compará ambas soluciones.

**Objetivo:** ver una race condition real, entender dos estrategias válidas para resolverla (memoria compartida protegida vs. estado que solo cambia por mensajes) y cuándo conviene cada una.

**Dónde completar:**
- [ejercicio4/inseguro/main.go](ejercicio4/inseguro/main.go) — ya está completo, es el punto de partida con el bug.
- [ejercicio4/mutex/main.go](ejercicio4/mutex/main.go) — completar con `sync.Mutex`.
- [ejercicio4/channel/main.go](ejercicio4/channel/main.go) — completar con una goroutine dueña del estado + channel.

> El stock arranca en 100 y se decrementa 50 veces, así que el resultado correcto es **50**, no 0.
