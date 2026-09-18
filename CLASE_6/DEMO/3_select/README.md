# Select

`select` sirve cuando una goroutine debe escuchar más de un channel. Es como
tener dos orejas: una escucha mensajes y la otra una señal para salir. Usa el
primer caso que esté listo.

Ejemplo: una pantalla de chat espera un mensaje nuevo, pero también debe poder
cerrarse cuando la persona toca “Salir”. No sabemos cuál de las dos cosas
ocurrirá primero.

```go
select {
case message := <-messages:
    // llegó un mensaje
case <-exit:
    // llegó la señal de salida
}
```

`select` no da prioridad fija a `messages` ni a `exit`. Si uno está listo, usa
ese. Si ambos estuvieran listos exactamente al mismo tiempo, Go elige uno; no
conviene depender de cuál será.

## Recorrido

- `main_1.go`: el mensaje tarda un segundo y `exit` dos; llega primero el
  mensaje.
- `main_2.go`: se invierten las demoras; llega primero `exit`.
- `main_3.go`: el `for` vuelve a escuchar una y otra vez. Llegan varios
  mensajes hasta que llega `exit`; `return` termina el programa.
- `main_4.go`: nadie envía un mensaje y no hay `default`. El programa queda
  bloqueado y Go muestra el error `all goroutines are asleep - deadlock!`.
- `main_5.go`: es el mismo código que el paso 4, pero agrega `default`. Así
  `select` continúa de inmediato aunque no haya mensajes.

Sin `default`, `select` espera hasta que un caso esté listo. Con `default`,
puede continuar aunque aún no haya mensajes.

En el paso 3, el `for` no es un tipo especial de select: simplemente vuelve a
ejecutar el mismo `select` después de cada mensaje. La señal `exit` hace que se
ejecute `return`, por eso el ciclo termina.

Ejecutar: `go run main_1.go`.
