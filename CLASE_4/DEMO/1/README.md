# Demo 1: productor y consumidor

Esta demo muestra el flujo más simple de RabbitMQ:

```text
productor -> cola-pedidos -> consumidor
```

El productor envía un texto simple (`pedido-101`). El consumidor lo recibe, simula diez segundos de procesamiento y recién entonces envía un **ACK manual**. Es una forma concreta de ver que recibir no es lo mismo que terminar de procesar.

Esta primera demo es intencionalmente mínima: no usa JSON, exchanges ni manejo de errores. Esos temas aparecen recién después.

## Ejecutar

Desde esta carpeta, abrí dos terminales.

Terminal 1:

```bash
go run ./consumidor
```

Terminal 2:

```bash
go run ./productor
```

Primero ejecutá el consumidor. Después ejecutá el productor y observá el mensaje en la primera terminal.

## RabbitMQ

Desde la terminal, descargá de forma manual RabbitMQ

```bash
docker run -d --name clase-rabbitmq \
  -p 5672:5672 -p 15672:15672 \
  -e RABBITMQ_DEFAULT_USER=user \
  -e RABBITMQ_DEFAULT_PASS=pass \
  rabbitmq:3-management
```

o

Desde `CLASE_4`, iniciá RabbitMQ una sola vez:

```bash
docker compose up -d
```

La consola de administración queda disponible en <http://localhost:15672> con usuario `user` y contraseña `pass`.
