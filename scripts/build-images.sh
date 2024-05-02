#!/bin/bash

# Variables
REPO_PREFIX="tu-registro-ecr"  # Reemplaza con el prefijo de tu registro ECR

# Construcción de la imagen Docker para el microservicio en Golang
echo "Construyendo la imagen Docker para el microservicio en Golang..."
docker build -t "$REPO_PREFIX/go-service:latest" go-service/

# Construcción de la imagen Docker para el microservicio en Python
echo "Construyendo la imagen Docker para el microservicio en Python..."
docker build -t "$REPO_PREFIX/python-service:latest" python-service/

# Construcción de la imagen Docker para la aplicación front-end
echo "Construyendo la imagen Docker para la aplicación front-end..."
docker build -t "$REPO_PREFIX/frontend:latest" frontend/

echo "Construcción de imágenes Docker completada."
