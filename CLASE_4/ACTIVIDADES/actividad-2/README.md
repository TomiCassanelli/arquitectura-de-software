# Actividad 2: Fanout Exchange

## Objetivo

Enviar un mismo pedido a dos sistemas: facturación y logística.

## Consigna

Completá los `TODO` para:

1. Declarar un exchange de tipo `fanout`.
2. Crear una cola para facturación y otra para logística.
3. Vincular ambas colas al exchange.
4. Publicar los pedidos en el exchange.

## Pistas

- Cada consumidor debe tener su propia cola.
- Usá `ExchangeDeclare`, `QueueDeclare` y `QueueBind`.
- En un exchange `fanout`, la routing key no define el destino.

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

Ejecutá cada comando en una terminal diferente y verificá que el mismo mensaje aparezca en facturación y logística.