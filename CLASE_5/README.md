# Clase 5: búsquedas con Apache Solr

La clase se compone de dos actividades consecutivas:

1. [Actividad 1: consultas, filtros y relevancia](ACTIVIDADES/actividad-1/README.md).
2. [Actividad 2: API Go + Gin](ACTIVIDADES/actividad-2/README.md).

Seguir estos pasos en orden. Los comandos de los pasos 1 a 4 se ejecutan una sola vez por clase, desde esta carpeta (`CLASE_5`).

## 1. Abrir Docker Desktop

## 2. Levantar Solr

Desde `CLASE_5` ejecutar:

```bash
docker compose up -d --wait
docker compose ps
```

El servicio `solr` debe figurar como `running` o `healthy`. El panel de Solr queda disponible en <http://localhost:8983/solr/>.

## 3. Crear el esquema

Ejecutar este comando una vez por cada Solr nuevo:

```bash
curl -fsS -X POST \
  -H 'Content-Type: application/json' \
  --data-binary @schema.json \
  'http://localhost:8983/solr/products/schema'
```

Debe responder con `"status":0`. Si informa que los campos ya existen, se puede continuar: el esquema ya fue creado.

## 4. Cargar los productos

```bash
curl -fsS -X POST \
  -H 'Content-Type: application/json' \
  --data-binary @products.json \
  'http://localhost:8983/solr/products/update?commit=true'
```

Verificar que se cargaron los quince productos:

```bash
curl -fsS -G 'http://localhost:8983/solr/products/select' \
  --data-urlencode 'q=*:*' \
  --data-urlencode 'rows=15' \
  --data-urlencode 'wt=json'
```

En la respuesta debe aparecer `"numFound":15`.

## 5. Resolver la Actividad 1

```bash
cd ACTIVIDADES/actividad-1
```

Construir las seis consultas indicadas en los cuatro archivos dentro de
`consultas` y ejecutarlos de a uno. La guía explica qué debe resolverse en
cada archivo y deja un único ejemplo incompleto de la forma del comando:

```bash
bash consultas/01_busquedas.sh
bash consultas/02_filtros_y_campos.sh
bash consultas/03_relevancia.sh
bash consultas/04_consultas_adicionales.sh
```

Cada estudiante debe usar la variante asignada en el [formulario de la clase](FORMULARIO_GOOGLE/README.md). El dataset es común; las consultas y preguntas de comprobación cambian según la variante.

## 6. Resolver la Actividad 2

```bash
cd ACTIVIDADES/actividad-2
go mod tidy
go run .
```

Completar antes los `TODO` de `controllers/products` y `repositories/products`. Con la API en ejecución, probar desde otra terminal:

```bash
curl 'http://localhost:8080/products/search?q=running&category=calzado&brand=Adidas&limit=5'
```

## 7. Detener los servicios

Para detener Solr y conservar los datos:

```bash
cd ../..
docker compose stop
```

Para borrar el contenedor e índice y comenzar desde cero:

```bash
docker compose down
```

Después de `down`, repetir los pasos 2, 3 y 4.
