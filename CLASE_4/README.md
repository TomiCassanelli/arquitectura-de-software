# Clase 4: asincronismo con RabbitMQ

La clase avanza de lo más simple a lo más completo. El objetivo no es memorizar la API: es observar qué ocurre cuando un productor, una cola y uno o varios consumidores trabajan a ritmos distintos.

## Demo 1: una cola

```text
productor -> cola -> consumidor
```

Muestra cómo publicar, recibir y confirmar un mensaje con ACK manual.

## Actividades centrales

Después de la primera demo, completar:

1. `ACTIVIDADES/actividad-1`: Fair Dispatch y QoS.
2. `ACTIVIDADES/actividad-2`: Fanout Exchange.

Cada actividad agrega una sola idea y se puede observar en el panel de RabbitMQ.

## Demo 2: cierre guiado — DLQ mínima

Muestra una sola idea:

- Separar un mensaje rechazado para que no frene a los demás.

Usarla como cierre guiado: integra ACK, NACK y DLQ sin JSON, reintentos ni código compartido.

## Demo 3: resiliencia — backoff exponencial

Solo muestra reintentos de conexión con esperas de 1, 2 y 4 segundos. No mezcla colas, consumidores ni mensajes: primero se entiende el patrón de backoff y recién después se conversa sobre aplicaciones resilientes.

Cada demo y actividad tiene su propio README con los comandos de ejecución.