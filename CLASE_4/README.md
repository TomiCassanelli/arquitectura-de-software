# Clase 4: RabbitMQ

La clase avanza de lo más simple a lo más completo.

## Demo 1: una cola

```text
productor -> cola -> consumidor
```

Muestra cómo publicar y recibir un mensaje.

## Demo 2: resiliencia y DLQ

Agrega tres ideas, sin mezclar demasiadas responsabilidades:

- Reintentos de conexión.
- Espera creciente entre intentos: backoff.
- DLQ para mensajes que no se pueden procesar.

## Actividades

Después de entender las demos, completar:

- `ACTIVIDADES/actividad-1`: Fair Dispatch y QoS.
- `ACTIVIDADES/actividad-2`: Fanout Exchange.
- `ACTIVIDADES/actividad-3`: Dead Letter Queue.

Cada demo y actividad tiene su propio README con los comandos de ejecución.