# Worker Pool

Un **worker pool** es un grupo fijo de goroutines que se reparte muchos
trabajos. En lugar de abrir una goroutine por cada producto, se decide cuántos
trabajadores habrá y esos mismos trabajadores toman los productos disponibles.

Ejemplo: una tienda debe actualizar el precio de cientos de productos. Tener
tres personas encargadas de hacerlo en paralelo es más controlable que contratar
cientos de personas, una para cada producto.

## Recorrido

- `main_1.go`: los productos se actualizan uno por uno. Es el punto de partida:
  no hay goroutines.
- `main_2.go`: se abre una goroutine por producto. Es rápido para pocos
  productos, pero con miles de productos también se abrirían miles de
  goroutines.
- `main_3.go`: hay cinco productos, pero solo tres workers. Los productos se
  envían por `productChannel` y cada worker toma el próximo que esté disponible.

## La idea de `main_3.go`

```go
productChannel := make(chan int)
```

El channel se crea en `main`: es la fila de productos pendientes. Después se
abren tres workers. Cada uno espera un producto, lo actualiza y vuelve a la fila
para tomar otro.

```text
Productos pendientes → productChannel → Worker 1
                                      → Worker 2
                                      → Worker 3
```

No importa qué worker actualice cada producto. Un worker puede actualizar dos
productos y otro solo uno: el siguiente producto lo toma quien vuelve a estar
libre primero. Lo importante es que nunca hay más de tres actualizaciones
activas a la vez.

`main` conoce toda la lista de productos. Cuando termina de enviarla, hace:

```go
close(productChannel)
```

Eso no significa que los workers ya terminaron; solo significa “no llegarán más
productos”. Cada worker sigue procesando los que ya recibió. Cuando el channel
está cerrado y vacío, su `for productID := range productChannel` termina.

El `WaitGroup` no reparte productos ni los protege. Solo hace que `main` espere
a que terminen los tres workers antes de imprimir `All products updated`.
