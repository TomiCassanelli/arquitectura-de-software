# Trabajo práctico integrador: e-commerce y logística

## Idea

El punto de partida es un monolito simple de e-commerce. Clientes, productos y pedidos viven dentro de la misma aplicación, en el puerto `8080`.

El objetivo es transformarlo gradualmente en un pequeño monorepo de microservicios y publicar un evento cuando un pedido queda confirmado.

```text
Antes:
cliente -> monolito :8080

Después:
cliente -> clientes :8081
        -> pedidos  :8082

pedidos -- evento pedido.confirmado --> RabbitMQ --> logística (conceptual)
```

Logística no debe implementarse. Solo representa el sistema que, en un caso real, tomaría el evento para preparar un envío.

## Qué tienen que hacer

1. Ejecutar y entender el monolito de la carpeta `monolito`.
2. Separar las responsabilidades en dos microservicios:
   - `clientes`: alta y consulta de clientes.
   - `pedidos`: consulta de productos y confirmación de pedidos.
3. Levantar cada microservicio en un puerto diferente.
4. Mantener una organización por capas: controlador, servicio y repositorio.
5. Agregar una caché simple para las consultas de productos.
6. Publicar el evento `pedido.confirmado` en RabbitMQ cuando un pedido se confirme.

No es necesario implementar autenticación, frontend, pagos, stock real, una base de datos real ni el microservicio de logística. Se puede trabajar con datos en memoria.

## Estructura inicial

```text
TRABAJO_PRACTICO_1_4/
├── README.md
├── ARQUITECTURA.md
├── monolito/
│   ├── go.mod
│   ├── main.go
│   └── controllers.go
└── microservicios/
    ├── clientes/
    ├── pedidos/
    └── eventos/
```

Las carpetas de `microservicios` son el destino del trabajo. Los alumnos deben crear allí los archivos y paquetes necesarios.

## Endpoints del monolito

| Método | Endpoint | Función |
| --- | --- | --- |
| `POST` | `/clientes` | Crear un cliente. |
| `GET` | `/clientes/:id` | Consultar un cliente. |
| `GET` | `/productos` | Listar productos. |
| `POST` | `/pedidos` | Confirmar un pedido. |

## Requisitos mínimos

### Microservicio `clientes`

- Escuchar en el puerto `8081`.
- Implementar `POST /clientes`.
- Implementar `GET /clientes/:id`.
- Usar al menos controlador, servicio y repositorio.

### Microservicio `pedidos`

- Escuchar en el puerto `8082`.
- Implementar `GET /productos`.
- Implementar `POST /pedidos`.
- Usar al menos controlador, servicio y repositorio.
- Agregar una caché en memoria para `GET /productos`.
- Publicar un evento cuando un pedido se confirme.

## El punto de RabbitMQ

La cola se implementa **dentro del microservicio `pedidos`**, después de que el servicio confirme la compra:

```text
POST /pedidos
    -> servicio confirma el pedido
    -> publica pedido.confirmado
    -> cola pedidos-confirmados
    -> logística podría preparar el envío
```

El evento puede tener este formato:

```json
{
  "tipo": "pedido.confirmado",
  "pedido_id": "PED-1",
  "cliente_id": "C-1",
  "producto_id": "P-1"
}
```

Para la primera versión alcanza con publicar en una cola llamada `pedidos-confirmados`. No hace falta crear consumidor, DLQ, Fanout, Backoff ni un sistema complejo de reintentos en este TP.

Una organización posible dentro de `pedidos` es:

```text
pedidos/
├── controllers/
├── services/
│   └── pedido_service.go
├── repositories/
├── messaging/
│   └── rabbitmq_publisher.go
└── main.go
```

`rabbitmq_publisher.go` es donde se usan `amqp.Dial`, `QueueDeclare("pedidos-confirmados", ...)` y `Publish(...)`.

## Ejecución inicial

Desde la carpeta `monolito`:

```bash
go mod tidy
go run .
```

El monolito queda disponible en `http://localhost:8080`.

## Ejecución esperada al finalizar

Cada microservicio debe ejecutarse en una terminal distinta:

```bash
cd microservicios/clientes
go run .
```

```bash
cd microservicios/pedidos
go run .
```

RabbitMQ debe estar activo para probar el evento de pedido. Desde `CLASE_4` se puede iniciar con:

```bash
docker compose up -d
```

## Entrega

Entregar el código dentro de `microservicios`.

## Criterio de finalización

El trabajo está terminado cuando el monolito original puede reemplazarse por los dos microservicios ejecutándose en puertos distintos, y un pedido confirmado genera el evento correspondiente en RabbitMQ.
