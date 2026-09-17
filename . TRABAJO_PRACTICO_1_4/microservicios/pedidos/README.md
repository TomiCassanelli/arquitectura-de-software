# Microservicio de pedidos

Crear aquí el servicio responsable de productos y pedidos.

Debe:

- Escuchar en `:8082`.
- Implementar `GET /productos`.
- Implementar `POST /pedidos`.
- Separar controlador, servicio y repositorio.
- Usar datos en memoria.
- Incorporar una caché simple para el listado de productos.
- Publicar `pedido.confirmado` en RabbitMQ después de confirmar un pedido.

## Punto de RabbitMQ

La publicación se hace en el servicio, después de confirmar el pedido:

```text
confirmar pedido -> publicar pedido.confirmado -> cola pedidos-confirmados
```

Una carpeta sugerida es `messaging/rabbitmq_publisher.go`. Allí se declara la cola y se publica el evento. No hace falta implementar el consumidor de logística.
