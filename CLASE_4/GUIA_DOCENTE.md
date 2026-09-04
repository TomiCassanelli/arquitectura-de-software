# Guía docente — revisión del PDF v3 y recorrido de clase

## Orden definitivo de las demos

Mantener **dos demos separadas** después de la Demo 1:

1. **Demo 2 — DLQ:** el broker funciona; el problema es un mensaje que el consumidor rechaza.
2. **Demo 3 — Backoff:** el problema ocurre antes; RabbitMQ todavía no responde y la aplicación debe esperar para volver a conectar.

El orden correcto es **DLQ primero y backoff después**. DLQ continúa el mismo flujo que ACK/NACK y las colas. Backoff cambia de problema: ya no hablamos de un mensaje, sino de la disponibilidad del broker. Mezclarlos en una sola demo haría más difícil saber por qué falla cada cosa.

## Diapositiva 1 — Portada

**Estado:** bien.

**Cambio sugerido:** cambiar `PRÁCTICO` por `Asincronismo con RabbitMQ`.

**Decir:** “Vamos a ver un mensaje desde que se publica hasta que se confirma, se rechaza o no se puede ni siquiera conectar.”

## Diapositiva 2 — Diccionario

**Estado:** bien y mucho más claro que las versiones anteriores.

**No agregar términos nuevos.** Dejar solo Productor, Broker, Cola, Consumidor, ACK y DLQ.

**Decir:** “Estas son las únicas seis palabras que necesitamos al principio. NACK, QoS y Exchange aparecerán cuando hagan falta.”

## Diapositiva 3 — Sincronismo vs. asincronismo

**Estado:** bien.

**Ajuste mínimo:** corregir la puntuación visual del recuadro final para que quede en una sola frase:

> En asincronismo, “recibido” no significa “terminado”. Significa: “guardé el trabajo y lo procesaré después”.

**Decir:** “No es necesariamente más rápido; permite que el emisor no espere.”

## Diapositiva 4 — Tres momentos de un trabajo asíncrono

**Estado:** todavía tiene una contradicción. El texto inferior dice “primero guardamos”, pero los bloques visuales comienzan con Aceptación.

**Cambiar el orden visual, de izquierda a derecha, por:**

1. **Registro durable:** “Guardamos el trabajo en una cola o base de datos.”
2. **Aceptación:** “Respondemos: recibido. Podemos entregar un ID.”
3. **Finalización:** “Un worker procesa el trabajo más tarde.”

**Mantener el párrafo inferior actual**, porque ya explica correctamente ese orden.

**Decir:** “La Demo 1 muestra la finalización: el consumidor termina y recién entonces envía ACK.”

## Diapositiva 5 — Los tres desacoples

**Estado:** bien.

**No agregar contenido.** Es la cantidad correcta de teoría para esta clase.

**Decir:** “En unos minutos vamos a ver el desacople de ritmo: qué pasa cuando un consumidor es lento y otro rápido.”

## Diapositiva 6 — Docker

**Estado:** pendiente. El comando sigue rompiéndose visualmente por las variables de entorno partidas en varias líneas.

**Reemplazar todo el bloque de comando por:**

```bash
cd CLASE_4
docker compose up -d
```

**Mantener debajo:**

> Panel: http://localhost:15672 — Usuario: `user` — Contraseña: `pass`

**Decir:** “Este comando inicia el broker. No vamos a dedicar tiempo a instalar ni configurar RabbitMQ.”

## Diapositiva 7 — Panel de RabbitMQ

**Estado:** muy bien simplificada.

**No cambiar contenido.**

**Decir:** “Miramos solo Queues. Ready es espera; Unacked es trabajo entregado que todavía no terminó.”

## Diapositiva 8 — Productor de la Demo 1

**Estado:** bien. Reemplazar la captura de GitHub por cuatro pasos fue una mejora importante.

**Cambio mínimo:** debajo del diagrama agregar:

> Ejecutá primero el productor sin consumidor y verificá `Ready = 1` en el panel.

**Decir:** “El productor no procesa. Solo deja `pedido-101` en la cola.”

## Diapositiva 9 — Consumidor y ACK manual

**Estado:** bien.

**Cambio mínimo:** agregar al final:

> Ejecutá ahora el consumidor y observá: `Ready → Unacked → 0`.

**Decir:** “Durante dos segundos el mensaje ya no está esperando, pero tampoco terminó. ACK cierra ese trabajo.”

## Diapositiva 10 — Ejercicios centrales

**Estado:** bien. QoS y Fanout son los dos ejercicios obligatorios correctos.

**Cambio mínimo:** debajo del título agregar:

> Recorrido obligatorio: Actividad 1 y Actividad 2. La DLQ se verá primero en una demo guiada.

**Reemplazar la consigna de Fanout por:**

> Una tienda confirma un pedido. Facturación debe generar la factura y logística debe preparar el envío. Los programas ya crean el exchange y las colas. Completá primero el vínculo de facturación y publicá un pedido. Después completá el vínculo de logística y publicá otro: ahora ambos deben recibir una copia.

**Decir:** “En QoS miramos reparto de trabajo; en Fanout miramos una copia del mismo evento para dos destinos.”

**Cómo guiarla:** primero hacer el vínculo de facturación con el curso y publicar un pedido. Luego pedir que completen el vínculo de logística solos y vuelvan a publicar. El resultado esperado pasa de un destinatario a dos, sin introducir más APIs nuevas.

## Diapositiva 11 — DLQ en acción / Demo 2

**Estado:** el concepto es correcto, pero está repetido: aparecen los cinco bloques, el flujo de tres líneas y dos frases finales que vuelven a explicar lo mismo.

**Dejar solo este contenido:**

```text
pedido-1 → ACK
ERROR    → NACK sin reencolar → DLQ
pedido-2 → ACK
```

Debajo:

> `ERROR` es una marca de la demo para representar un mensaje que no se puede procesar. La DLQ lo aparta; no lo corrige.

**Quitar:** los bloques “Fallo de procesamiento”, “Operación continua” y “Revisión posterior”, porque el diagrama ya comunica esas ideas.

**Decir:** “Esto es Demo 2. El broker está funcionando; la decisión es del consumidor: ACK para continuar o NACK sin reencolar para apartar.”

## Diapositiva 12 — Backoff exponencial / Demo 3

**Estado:** el orden es correcto: va después de DLQ. Hay dos frases que deben ajustarse para que coincidan con el código.

**Cambiar el inicio:**

> Imaginá que la aplicación se inicia mientras RabbitMQ todavía está apagado o arrancando.

**Reemplazar “la conexión se restablece” por:**

> En el siguiente intento disponible, la aplicación se conecta cuando RabbitMQ ya está listo.

**Mantener el bloque 1 s → 2 s → 4 s.**

**Agregar al pie:**

> Demo 3: reintento de conexión inicial. No es una reconexión completa de una aplicación que ya estaba procesando mensajes.

**Decir:** “DLQ resuelve un mal mensaje. Backoff resuelve una dependencia que todavía no está disponible. Son problemas distintos; por eso son dos demos.”

## Guion práctico sugerido

1. Diapositivas 1 a 7: 8 minutos.
2. Diapositivas 8 y 9 + Demo 1: 12 minutos.
3. Diapositiva 10 + Actividades 1 y 2: 35 minutos.
4. Diapositiva 11 + Demo 2: 12 minutos.
5. Diapositiva 12 + Demo 3: 10 minutos.

Con este orden, cada demo introduce exactamente un tipo de fallo: procesamiento normal, mensaje rechazado y broker no disponible.
