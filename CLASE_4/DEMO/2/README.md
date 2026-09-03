# Demo 2: JSON, resiliencia y DLQ

Esta demo muestra cómo enviar pedidos JSON y cómo manejar dos errores básicos: una conexión que puede fallar y un mensaje que no se puede leer.

## Estructura

```text
2/
├── go.mod
├── go.sum
├── shared/
│   ├── config.go       Configuración común
│   ├── model.go        Estructura Pedido
│   └── rabbitmq.go     Conexión con reintentos y colas
├── productor/
│   └── main.go         Programa que publica pedidos
└── consumidor/
    └── main.go         Programa que procesa pedidos
```

## Flujo

```text
productor
    |
    | publica un pedido JSON
    v
RabbitMQ: cola-pedidos-resiliente
    |
    | entrega el pedido
    v
consumidor
    |
    +--> JSON válido: consumidor convierte el mensaje y lo confirma
    |
    +--> JSON inválido: consumidor envía NACK sin reencolar
                                      |
                                      v
                              cola-pedidos-dlq
```

## Qué hace cada parte

### `shared/config.go`

Guarda la dirección, nombres de las colas y límites de reintentos. La conexión actual es:

```text
amqp://user:pass@localhost:5672
```

El usuario y la contraseña deben existir en tu RabbitMQ. Si tu instalación usa otras credenciales, cambiá `AMQPURI`.

### `shared/model.go`

Define el tipo `Pedido`, con un identificador, un monto y un cliente.

### `shared/rabbitmq.go`

Contiene la conexión con reintentos y backoff lineal. Si RabbitMQ no está disponible, espera un poco más entre cada intento. También declara la cola principal con su DLQ.

### `productor/main.go`

El productor crea un pedido válido y también publica un texto que no es JSON. Usa un timeout para no quedar esperando indefinidamente.

### `consumidor/main.go`

El consumidor escucha la cola, recibe el JSON y lo convierte nuevamente a un `Pedido`. Si el JSON es inválido, lo rechaza sin reencolarlo para que RabbitMQ lo envíe a la DLQ.

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

Primero ejecutá el consumidor y después el productor. El consumidor queda esperando y el productor envía un único pedido.

La resiliencia se observa al iniciar la demo sin RabbitMQ: el programa intenta conectarse tres veces y espera cada vez un poco más. La DLQ se utiliza cuando el consumidor rechaza el texto inválido con `Nack(false, false)`.

## Verificar

Desde la carpeta `DEMO/2`:

```bash
go test ./...
```

Este comando compila el productor, el consumidor y el paquete compartido.