# Actividad 3: Dead Letter Queue

## Objetivo

Separar automáticamente los mensajes corruptos para que no queden bloqueando la cola principal.

## Consigna

Completá los `TODO` del consumidor para:

1. Declarar el exchange de mensajes muertos.
2. Declarar y vincular la cola `q.pedidos.venenosos`.
3. Configurar `q.pedidos` para usar la DLQ.
4. Confirmar los mensajes válidos y rechazar los corruptos sin reencolarlos.

## Pistas

- La cola principal necesita `x-dead-letter-exchange` y `x-dead-letter-routing-key`.
- `Nack` recibe un parámetro que indica si el mensaje debe volver a la cola.
- El productor ya envía mensajes válidos y un mensaje que no es JSON.

## Ejecutar

Desde esta carpeta, abrí tres terminales:

Terminal 1:

```bash
go run ./consumidor
```

Terminal 2:

```bash
go run ./productor
```

Terminal 3, opcional:

```bash
go run ./visor-dlq
```

El tercer programa es opcional y permite observar los mensajes que llegan a la DLQ.