# Usa una imagen base de Golang
FROM golang:1.20 AS builder

# Establece la carpeta de trabajo
WORKDIR /app

# Copia el código fuente de la aplicación
COPY . .

# Compila la aplicación
RUN go build -o myapp

# Usa una imagen base más pequeña de Alpine Linux para la fase de ejecución
FROM alpine:latest

# Instala certificados de CA raíz
RUN apk --no-cache add ca-certificates

# Establece la carpeta de trabajo
WORKDIR /app

# Copia el ejecutable de la aplicación desde la fase de compilación
COPY --from=builder /app/myapp /app/

# Expone el puerto 8080
EXPOSE 8080

# Ejecuta la aplicación
CMD ["./myapp"]
