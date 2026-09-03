# Demo 1: productor y consumidor

Esta demo muestra el flujo más simple de RabbitMQ:

```text
productor -> cola-pedidos -> consumidor
```

El productor envía un pedido en formato JSON. El consumidor escucha la cola y muestra el mensaje recibido.

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

RabbitMQ debe estar ejecutándose en `localhost:5672` con las credenciales configuradas en los archivos.