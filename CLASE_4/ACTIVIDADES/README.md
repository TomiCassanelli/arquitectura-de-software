# Actividades de RabbitMQ

Hay dos actividades independientes. En cada carpeta encontrarán el código base y una consigna.

Antes de comenzar:

- Tener RabbitMQ ejecutándose en `localhost:5672`. Desde `CLASE_4`, se puede iniciar con `docker compose up -d`.
- Las demos usan `amqp://user:pass@localhost:5672`; el panel queda en <http://localhost:15672> con `user` / `pass`.
- Tener Go instalado.

Cada programa se ejecuta en una terminal distinta. Para detenerlo, usar `CTRL+C`.

Si se cambia la configuración de una cola durable durante una prueba, RabbitMQ no permite redeclararla con argumentos distintos. Detené los programas y reiniciá el contenedor.