# Esta versión prioriza claridad para la clase: Go compila y ejecuta la API
# dentro de la misma imagen. Una imagen multi-stage se puede ver más adelante.
FROM golang:1.22-alpine

WORKDIR /app

COPY . .
RUN go build -o api-go .

EXPOSE 3000
CMD ["./api-go"]
