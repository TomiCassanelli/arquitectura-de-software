## Consigna

### Desarmando el monolito del gimnasio

#### Contexto

El sistema del gimnasio creció mucho más rápido que su arquitectura. Lo que comenzó como una aplicación simple en Go terminó concentrando en un único proceso la lógica de usuarios, actividades, operaciones comerciales, beneficios y administración general. El resultado es un monolito difícil de escalar, de mantener y de desplegar.

Para este trabajo práctico se parte de una base inicial en la que todo convive en un mismo `main.go` y en un mismo `controllers.go`, con un único servidor escuchando en un solo puerto. A partir de ese punto, el objetivo es reorganizar la solución para que refleje una arquitectura más modular, con responsabilidades separadas y *ejecución independientes*.

#### Objetivo

Diseñar y construir una versión del sistema que permita separar las responsabilidades por dominio y ejecutar cada parte de manera autónoma. La solución debe demostrar una división clara del sistema, de forma tal que cada módulo pueda levantarse y evolucionar sin depender del resto en tiempo de ejecución.

#### Actividad

Reescribir la aplicación para que el sistema quede dividido en varios servicios o módulos ejecutables de forma independiente. Cada uno debe representar una capacidad de negocio distinta y exponerse a través de su propia API.

La solución debe contemplar, como mínimo:

- separación por responsabilidades claramente identificables;
- un proyecto o módulo independiente por cada dominio definido;
- endpoints coherentes con el contexto funcional de cada servicio;
- arranque independiente de cada componente;
- validación de que el sistema funciona con más de un puerto encendido al mismo tiempo.


#### Pista

Si la solución solo necesita un puerto para correr, todavía no está resuelta como se espera.