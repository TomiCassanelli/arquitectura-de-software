FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go build -o api-clima .

EXPOSE 3000
CMD ["./api-clima"]
