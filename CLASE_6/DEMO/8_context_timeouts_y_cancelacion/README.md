# Context: cancelación y timeout

Un `context` es una forma de avisar a una tarea que ya no debe seguir
trabajando. El aviso puede ser manual, por ejemplo porque una persona cerró una
pantalla, o automático porque se agotó un tiempo máximo.

Ejemplo: una ficha de producto busca precio, stock y reviews. Si la persona
cierra la ficha, no tiene sentido seguir buscando. Si reviews tarda demasiado,
tampoco queremos dejar esa consulta abierta para siempre.

## Recorrido

- `main_1.go`: cancelación manual. La persona cierra la pantalla y `main` llama
  a `cancelRequest()`. La consulta de reviews recibe el aviso y termina.
- `main_2.go`: tiene prácticamente el mismo código que el paso 1, pero ahora
  el context se cancela solo después de 300 milisegundos.
- `main_3.go`: precio, stock y reviews comparten un mismo context con timeout.
  Precio y stock llegan a tiempo; reviews recibe la cancelación.

## `main_1.go`: cancelación manual

```go
requestContext, cancelRequest := context.WithCancel(context.Background())
```

`requestContext` es el aviso que recibirá la tarea. `cancelRequest` es la
función que permite enviar ese aviso. Cuando la persona cierra la pantalla:

```go
cancelRequest()
```

la función que está esperando puede detectarlo con:

```go
case <-requestContext.Done():
    // la tarea fue cancelada
```

`finishedChannel` no es parte del context. Solo le permite a `main` esperar y
confirmar: “la tarea recibió el aviso y ya terminó”.

## `main_2.go`: timeout

```go
context.WithTimeout(context.Background(), 300*time.Millisecond)
```

crea un context que hace el mismo aviso automáticamente cuando se cumple el
tiempo. El context no es “un timeout”: el timeout es solo una forma de cancelar
un context.

Ejecutar: `go run main_1.go`.
