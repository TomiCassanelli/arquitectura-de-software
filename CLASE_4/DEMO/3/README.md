# Demo 3: resiliencia con backoff exponencial

Esta demo tiene un solo objetivo: volver a intentar una conexión que falló, esperando cada vez más tiempo.

```text
intento 1 falla -> espera 1 s
intento 2 falla -> espera 2 s
intento 3 falla -> espera 4 s
intento 4            éxito o error final
```

No publica mensajes ni crea colas. Así se puede leer el código como una única pregunta: “si RabbitMQ todavía no está listo, ¿cuándo vuelvo a probar?”. Los mensajes de consola indican cada intento y cada espera.

## Ejecutar

1. Desde `CLASE_4`, detener RabbitMQ:

   ```bash
   docker compose down
   ```

2. En una terminal, desde esta carpeta, iniciar la demo:

   ```bash
   go run .
   ```

3. Mientras la demo muestra sus reintentos, en otra terminal iniciar RabbitMQ:

   ```bash
   cd CLASE_4
   docker compose up -d
   ```

La demo debería conectarse en el siguiente intento disponible. Si RabbitMQ demora más de los cuatro intentos, ejecutar la demo otra vez.

## Qué no demuestra

No es una reconexión completa de una aplicación que ya estaba trabajando: solo muestra el reintento de conexión inicial. Jitter, reintentos de mensajes e idempotencia son mejoras para una clase posterior.
