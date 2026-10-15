# Primera etapa: compila el main.go y sus paquetes con la misma versión usada en clase.
FROM golang:1.22-alpine AS build

# El directorio de trabajo queda dentro de la imagen, no modifica el host.
WORKDIR /app

# Copiar dependencias primero permite reutilizar la caché de Docker.
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código y compilar un binario Linux estático.
COPY . .
RUN CGO_ENABLED=0 go build -o /demo .

# Segunda etapa: contiene solo el binario, sin compilador ni código fuente.
FROM alpine:3.21
COPY --from=build /demo /demo
EXPOSE 8080
CMD ["/demo"]
