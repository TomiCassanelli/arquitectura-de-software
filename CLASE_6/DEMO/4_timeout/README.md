# Timeout

Un **timeout** pone un límite a la espera. Ejemplo: una app de mapas pide el
tiempo de viaje a un servicio. Si no responde rápido, es mejor avisar que no
hay datos que dejar la pantalla cargando para siempre.

`time.After(time.Second)` crea un channel especial: después de un segundo,
recibe un valor. `select` escucha tanto la respuesta como ese aviso de tiempo.

```go
select {
case response := <-serviceChannel:
    // el servicio respondió a tiempo
case <-programTimeout:
    // se acabó el tiempo permitido
}
```

## Recorrido

- `main_1.go`: la respuesta tarda dos segundos y el timeout uno. Gana el
  timeout.
- `main_2.go`: solo cambia la demora: la respuesta llega en medio segundo y
  gana antes que el timeout.
- `main_3.go`: llega una respuesta rápida; la segunda tarda demasiado. Como el
  límite de un segundo empezó al inicio, se imprime `Timeout` antes de que
  llegue esa segunda respuesta.
- `main_4.go`: el servicio responde cada 700 milisegundos, pero el programa
  completo tiene un máximo de cinco segundos. Recibe varias respuestas y luego
  finaliza.

El timeout no cancela mágicamente una función lenta: esta demo deja de esperar
y termina `main`. La cancelación explícita se verá más adelante con `context`.
Las demoras son fijas para observar siempre el mismo caso.

Ejecutar: `go run main_1.go`.
