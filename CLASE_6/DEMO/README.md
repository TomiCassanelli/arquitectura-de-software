# Demos de Clase 6

Cada archivo es un programa independiente. Ejecutalo desde su carpeta, por
ejemplo: `go run main_1.go`. Los archivos de cada directorio están ordenados:
cada paso agrega una idea al anterior.

| Tema del PDF | Demo | Progresión |
| --- | --- | --- |
| Goroutines | `1_goroutines` | Secuencial, `main` que termina antes y usuarios que avanzan a la vez. |
| Channels | `2_channels` | Envío/recepción, espera sin buffer, varios pedidos y buffer como fila. |
| `select` | `3_select` | Espera entre mensajes/fin y luego agrega `default` no bloqueante. |
| Timeout | `4_timeout` | `time.After` y espera limitada de una respuesta. |
| Race Conditions, Mutex y `WaitGroup` | `5_race_conditions_mutex_y_waitgroup` | Race deliberada, corrección con `sync.Mutex` y espera con `WaitGroup`. |
| Fan-out / Fan-in | `6_fan_out_fan_in` | Consultas en paralelo y respuestas que vuelven a un channel común. |
| Worker Pool | `7_worker_pool` | Secuencial, una goroutine por producto y pool fijo de workers. |
| Context: timeouts y cancelación | `8_context_timeouts_y_cancelacion` | Timeout de una fuente y propagación de cancelación a varias fuentes. |

`EXTRA_melisearch` queda fuera del recorrido principal. Es un ejercicio más
avanzado para retomar después de la clase.

Para verificar la race y su corrección:

```sh
go run -race main_1.go
go run -race main_2.go
```

desde `5_race_conditions_mutex_y_waitgroup`.

La distinción entre concurrencia y paralelismo se explica en la presentación:
no tiene una demo determinista, porque el paralelismo efectivo depende del
hardware y del runtime.
