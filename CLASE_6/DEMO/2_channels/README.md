# Channels

Un **channel** sirve para que una goroutine le pase un dato a otra. Por ejemplo,
el servicio de delivery sabe que un pedido está listo y necesita avisarle a la
pantalla.

El channel no hace el trabajo ni crea una goroutine: solamente transporta el
mensaje entre quien lo envía y quien lo recibe.

```go
orderChannel := make(chan string)
```

Esto crea un channel que transporta textos. Las flechas se leen así:

```go
orderChannel <- "Order is ready" // poner un mensaje en el channel
order := <-orderChannel           // sacar un mensaje del channel
```

## Sin buffer y con buffer

Un channel sin buffer es como pasar una caja de mano a mano: quien envía espera
hasta que alguien esté listo para recibir. No deja el mensaje guardado.

Un channel con buffer tiene una pequeña fila. Por ejemplo,
`make(chan string, 2)` puede guardar hasta dos mensajes. Mientras quede lugar,
el emisor puede continuar aunque el receptor todavía no haya escuchado.

## Recorrido

- `main_1.go`: una goroutine envía un pedido y `main` lo recibe.
- `main_2.go`: `main` tarda cinco segundos en recibir. El envío queda detenido
  hasta que aparece alguien que escuche.
- `main_3.go`: es el mismo caso que el paso 2, pero solo cambia a
  `make(chan string, 2)`. Ahora el emisor puede continuar aunque `main` todavía
  no haya recibido el pedido.
- `main_4.go`: aplica la misma idea a dos pedidos, con funciones separadas para
  enviar y recibir.
- `main_5.go`: es igual al paso 4, pero el receptor tarda un segundo en procesar
  cada pedido. Como no hay buffer, el emisor debe esperar antes de continuar.
- `main_6.go`: es igual al paso 5; la única diferencia es el buffer para dos
  pedidos, que funciona como una pequeña fila.

Ejecutar: `go run main_1.go`.

## Flujo de `main_2.go`

```mermaid
sequenceDiagram
    participant Main as main
    participant Sender as sendOrderTo
    participant Channel as orderChannel

    Main->>Sender: go sendOrderTo(orderChannel)
    Sender->>Channel: intenta enviar "Order is ready"
    Note over Sender,Channel: espera: main todavía no recibe
    Main->>Main: espera 5 segundos
    Main->>Channel: recibe con <-orderChannel
    Channel-->>Main: entrega "Order is ready"
    Channel-->>Sender: el envío termina
```
