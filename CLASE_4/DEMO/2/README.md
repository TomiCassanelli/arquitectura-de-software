# Demo 2: una DLQ mínima

Esta demo muestra una sola decisión importante: qué hacer con un mensaje que el consumidor no puede procesar.

## Estructura

```text
2/
├── go.mod
├── productor/
│   └── main.go         Publica dos pedidos y un mensaje ERROR
└── consumidor/
    └── main.go         Acepta pedidos y envía ERROR a la DLQ
```

## Flujo

```text
productor
    |
    | publica un mensaje
    v
RabbitMQ: demo.pedidos
    |
    | entrega el pedido
    v
consumidor
    |
    +--> pedido: consumidor lo confirma (ACK)
    |
    +--> ERROR: consumidor envía NACK sin reencolar
                                      |
                                      v
                              demo.pedidos.dlq
```

## Qué hace cada parte

### `productor/main.go`

El productor publica `pedido-1`, `ERROR` y `pedido-2`. No declara colas: por eso se inicia después del consumidor.

### `consumidor/main.go`

El consumidor declara la cola principal y la DLQ. Si recibe `ERROR`, lo rechaza sin reencolarlo con `Nack(false, false)`. RabbitMQ lo mueve automáticamente a la DLQ y sigue con el siguiente pedido.

## Ejecutar

RabbitMQ debe estar ejecutándose primero.

Abrí una terminal para el consumidor:

```bash
cd "CLASE_4/DEMO/2"
go run ./consumidor
```

Abrí otra terminal para el productor:

```bash
cd "CLASE_4/DEMO/2"
go run ./productor
```

Primero ejecutá el consumidor y después el productor. El consumidor queda esperando y el productor publica tres mensajes.

En los logs se ve que `pedido-1` y `pedido-2` se confirman con ACK. El mensaje `ERROR` no vuelve a la cola: llega a `demo.pedidos.dlq`.

## Verificar

Desde la carpeta `DEMO/2`:

```bash
go test ./...
```

Este comando compila el productor y el consumidor.