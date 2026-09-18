# Race Conditions, Mutex y WaitGroup

Imaginá una tienda online con stock inicial de 10. Después de dos ventas, el
stock correcto debe ser 8.

Una **race condition** aparece cuando dos goroutines usan el mismo dato y sus
pasos se mezclan de una forma que deja un resultado incorrecto. El problema no
es tener dos goroutines: es leer y modificar el mismo `stock` sin coordinación.

## Qué sale mal sin Mutex

Cada venta hace tres pasos:

```text
leer stock → restar 1 → guardar stock
```

Sin protección puede ocurrir esto:

```text
Venta A lee 10
Venta B lee 10
Venta A guarda 9
Venta B guarda 9
```

Las dos ventas hicieron bien la cuenta, pero usaron una foto vieja del stock:
`10`. Se vendieron dos unidades, pero el resultado queda en `9` en vez de `8`.

## Qué soluciona Mutex

Un `Mutex` es una llave que protege una parte sensible del código:

```go
stockMutex.Lock()

// leer y modificar stock

stockMutex.Unlock()
```

Si las dos goroutines llegan a `Lock()` al mismo tiempo, solo una toma la
llave. La otra espera en su `Lock()`. Cuando la primera hace `Unlock()`, la
otra puede entrar. No se define si tiene prioridad A o B; Go lo decide. Lo
importante es que nunca modifican `stock` juntas.

```text
Venta A toma la llave: stock 10 → 9
Venta A devuelve la llave
Venta B toma la llave: stock 9 → 8
```

## WaitGroup: qué es

Un `WaitGroup` sirve para que `main` espere a que terminen otras goroutines.
No protege el stock y no ordena las ventas: solo evita que el programa termine
demasiado pronto.

## Recorrido

- `main_1.go`: ventas una después de otra, sin goroutines. Resultado: `8`.
- `main_2.go`: ventas en goroutines y sin mutex. Ambas pueden leer `10` y
  terminar guardando `9`: es la race condition.
- `main_3.go`: agrega `stockMutex`. Una venta actualiza y sale; la otra espera.
  El resultado vuelve a ser `8`.
- `main_4.go`: agrega `WaitGroup` para que `main` espere exactamente a las dos
  ventas antes de imprimir el resultado final.

`Mutex` protege el stock. `WaitGroup` no protege nada: solo espera a que las
goroutines terminen. En los pasos 2 y 3 se usa `Sleep` solo para que `main` no
termine antes de mostrar la demo; el paso 4 es la forma correcta de esperar.

## Cómo leer `main_4.go`

```go
sales.Add(2)
```

significa: “hay dos ventas que deben terminar”. El `&sales` que se pasa a cada
venta significa que ambas reciben el mismo `WaitGroup`, no una copia.

```go
defer sales.Done()
```

avisa “esta venta terminó”. Finalmente, `sales.Wait()` detiene a `main` hasta
que las dos ventas hayan hecho `Done()`. Es distinto del mutex: el mutex cuida
el stock; el WaitGroup solo espera las ventas.
