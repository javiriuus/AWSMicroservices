#!/bin/bash

# Variables
REPO_PREFIX="tu-registro-ecr"  # Reemplaza con el prefijo de tu registro ECR

# Inicia sesión en el registro ECR
echo "Iniciando sesión en Amazon ECR..."
aws ecr get-login-password --region tu-región | docker login --username AWS --password-stdin "$REPO_PREFIX"

# Enviar la imagen Docker para el microservicio en Golang
echo "Enviando la imagen Docker para el microservicio en Golang..."
docker push "$REPO_PREFIX/go-service:latest"

# Enviar la imagen Docker para el microservicio en Python
echo "Enviando la imagen Docker para el microservicio en Python..."
docker push "$REPO_PREFIX/python-service:latest"

# Enviar la imagen Docker para la aplicación front-end
echo "Enviando la imagen Docker para la aplicación front-end..."
docker push "$REPO_PREFIX/frontend:latest"

echo "Envío de imágenes Docker completado."
