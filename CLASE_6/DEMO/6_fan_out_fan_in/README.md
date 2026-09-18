# Fan-out / Fan-in

Imaginá una app que compara precios. Para buscar un producto, consulta varios
comercios a la vez y muestra cada respuesta cuando llega.

**Fan-out** significa repartir una misma tarea en varias goroutines. En esta
demo, la tarea “buscar el precio” sale hacia varios comercios:

```go
go searchStore("Frávega", 1200*time.Millisecond, pricesChannel)
go searchStore("Garbarino", 500*time.Millisecond, pricesChannel)
```

**Fan-in** significa juntar resultados que vienen de varias goroutines. En esta
demo, los comercios responden al mismo `pricesChannel` y `main` los recibe:

```go
fmt.Println(<-pricesChannel)
```

El channel se crea siempre en `main`. Así se ve claramente dónde se juntarán
las respuestas:

```go
pricesChannel := make(chan string)
```

## Recorrido

- `main_1.go`: una sola consulta y un solo channel. Es el caso base: un
  comercio responde y `main` muestra su respuesta.
- `main_2.go`: dos consultas empiezan al mismo tiempo, pero cada una tiene su
  propio channel. `main` espera primero a Frávega, aunque Garbarino responda
  antes. Aquí aparece el fan-out, pero todavía no el fan-in.
- `main_3.go`: las dos consultas siguen empezando a la vez, pero ahora ambas
  envían al mismo `pricesChannel`. `main` recibe la primera respuesta que llega
  y luego la segunda. Esto une fan-out y fan-in.
- `main_4.go`: se suma un tercer comercio. Solo cambia la cantidad de tareas y
  respuestas; la idea es la misma.
- `main_5.go`: se usa el mismo patrón con cinco comercios, como una app de
  comparación de precios más real.

El orden de las respuestas depende de cuál comercio tarda menos. No es un
error: justamente la app muestra cada una apenas llega, sin esperar a las demás.

Ejecutar: `go run main_1.go`.
