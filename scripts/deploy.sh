#!/bin/bash

# Despliegue del microservicio en Golang
echo "Desplegando el microservicio en Golang..."
kubectl apply -f deployment/go-service.yaml

# Despliegue del microservicio en Python
echo "Desplegando el microservicio en Python..."
kubectl apply -f deployment/python-service.yaml

# Despliegue de la aplicación front-end
echo "Desplegando la aplicación front-end..."
kubectl apply -f deployment/frontend.yaml

echo "Despliegue completado."
