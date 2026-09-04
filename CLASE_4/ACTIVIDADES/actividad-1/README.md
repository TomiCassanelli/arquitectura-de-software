# Actividad 1: Fair Dispatch y QoS

## Objetivo

Repartir los mensajes entre dos consumidores sin sobrecargar al consumidor más lento.

## Consigna

Completá el `TODO` de `consumidor/main.go` para configurar QoS antes de `Consume`.

Ejecutá un consumidor lento y otro rápido, y luego el productor. El productor enviará 10 pedidos y terminará solo. Probá primero sin tu cambio y luego con tu cambio: sin QoS, el consumidor lento puede quedar con varios mensajes sin confirmar; con QoS=1, el rápido recibe los que el lento todavía no puede confirmar.

## Pistas

- El método que necesitás es `Qos`.
- El primer parámetro indica cuántos mensajes puede recibir cada consumidor sin confirmar.
- El mensaje debe confirmarse con `ACK` después de ser procesado.

## Ejecutar

Desde esta carpeta, abrí tres terminales:

Terminal 1, consumidor lento:

```bash
CONSUMIDOR=lento DELAY_SEGUNDOS=5 go run ./consumidor
```

Terminal 2, consumidor rápido:

```bash
CONSUMIDOR=rapido DELAY_SEGUNDOS=1 go run ./consumidor
```

Terminal 3:

```bash
go run ./productor
```

Observá en los logs cómo se distribuyen los 10 pedidos.
