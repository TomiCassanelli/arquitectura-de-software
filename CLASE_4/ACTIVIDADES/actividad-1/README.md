# Actividad 1: Fair Dispatch y QoS

## Objetivo

Repartir los mensajes entre dos consumidores sin sobrecargar al consumidor más lento.

## Consigna

Completá el `TODO` de `consumidor/main.go` para configurar QoS antes de `Consume`.

Ejecutá dos consumidores y un productor. El productor enviará 10 pedidos y terminará solo. Probá primero sin tu cambio y luego con tu cambio.

## Pistas

- El método que necesitás es `Qos`.
- El primer parámetro indica cuántos mensajes puede recibir cada consumidor sin confirmar.
- El mensaje debe confirmarse con `ACK` después de ser procesado.

## Ejecutar

Desde esta carpeta, abrí tres terminales:

Terminales 1 y 2:

```bash
go run ./consumidor
```

Terminal 3:

```bash
go run ./productor
```

Observá en los logs cómo se distribuyen los 10 pedidos.