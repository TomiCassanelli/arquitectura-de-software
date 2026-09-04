# Actividad 2: pedido confirmado para dos áreas

## Objetivo

Cuando una tienda confirma un pedido, deben enterarse al mismo tiempo dos áreas independientes:

- **Facturación:** genera la factura.
- **Logística:** prepara el envío.

El objetivo es que ambas reciban una copia del mismo evento, sin compartir una sola cola.

```text
tienda -> exchange fanout -> q.facturas  -> facturación
                         -> q.logistica  -> logística
```

## Consigna

Los programas ya crean el exchange y las dos colas. Solo falta el vínculo que permite recibir el evento.

1. En `facturacion/main.go`, completá el `TODO` con `QueueBind`.
2. Ejecutá el productor: solo facturación debe recibir el pedido.
3. En `logistica/main.go`, repetí el vínculo para su cola.
4. Ejecutá nuevamente el productor: ahora ambos sistemas deben recibir una copia.

## Pistas

- `QueueBind` conecta una cola con un exchange.
- En un exchange `fanout`, la routing key puede ser `""`.
- No copies `q.facturas` en logística: cada área usa su propia cola.

## Ejecutar

Desde esta carpeta, abrí tres terminales:

Terminal 1:

```bash
go run ./facturacion
```

Terminal 2:

```bash
go run ./logistica
```

Terminal 3:

```bash
go run ./productor
```

Primero iniciá ambos consumidores. Sin completar los TODO, el productor no genera mensajes en sus colas: no hay vínculos todavía.

Cada vez que completes un TODO, detené y volvé a iniciar ese consumidor para ejecutar el código nuevo. Después ejecutá otra vez el productor y observá qué área recibe el evento.
