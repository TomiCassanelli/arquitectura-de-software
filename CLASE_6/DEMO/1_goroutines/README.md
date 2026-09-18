# Goroutines

Una **goroutine** es una tarea que Go puede ejecutar sin detener a quien la
creó. Se inicia agregando `go` antes de llamar una función:

```go
go processUser(userID)
```

Ejemplo: una app puede subir una foto mientras la persona sigue leyendo sus
mensajes. Subir la foto es una tarea; la pantalla es otra.

## Recorrido

- `main_1.go`: `main` procesa los tres usuarios de a uno. Cada `End` aparece
  antes del `Start` siguiente.
- `main_2.go`: se agrega solo `go processUsers()`. `main` termina enseguida y
  el programa se cierra sin esperar a la goroutine.
- `main_3.go`: se agrega una espera temporal para que la goroutine pueda
  terminar. Los usuarios siguen procesándose de a uno dentro de ella.
- `main_4.go`: se agrega `go` al procesar cada usuario. Ahora los tres usuarios
  están activos a la vez: pueden aparecer varios `Start` antes de cualquier
  `End`.

Una goroutine no garantiza el orden de los mensajes. Go decide cuándo avanza
cada tarea. Por eso el orden de los `Start` y `End` puede variar, pero en el
paso 4 las tres tareas tienen oportunidad de avanzar al mismo tiempo.

Ejecutar: `go run main_1.go`.
