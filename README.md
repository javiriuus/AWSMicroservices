# AWSMicroservices

Estructura de carpetas y archivos:

project-root/
│
├── README.md
├── docker-compose.yml      # Opcional, para desarrollo local
│
├── go-service/              # Carpeta para el microservicio en Golang
│   ├── Dockerfile
│   ├── main.go
│   ├── go.mod               # Archivo de dependencias de Go
│   ├── go.sum               # Archivo de sumas de verificación de dependencias
│   └── ...
│
├── python-service/          # Carpeta para el microservicio en Python
│   ├── Dockerfile
│   ├── main.py
│   ├── requirements.txt     # Lista de dependencias de Python
│   └── ...
│
├── frontend/                # Carpeta para la aplicación front-end (Next.js)
│   ├── Dockerfile
│   ├── package.json
│   ├── yarn.lock
│   ├── pages/               # Páginas de Next.js
│   │   └── index.js
│   ├── components/          # Componentes de React
│   │   └── ...
│   └── ...
│
├── deployment/              # Carpeta para los manifiestos de Kubernetes
│   ├── go-service.yaml      # Manifiesto de Kubernetes para el microservicio en Golang
│   ├── python-service.yaml  # Manifiesto de Kubernetes para el microservicio en Python
│   ├── frontend.yaml        # Manifiesto de Kubernetes para la aplicación front-end
│   └── ...
│
└── scripts/                 # Scripts útiles para automatizar tareas
    ├── build-images.sh      # Script para construir imágenes Docker
    ├── push-images.sh       # Script para enviar imágenes Docker a un registro
    └── deploy.sh            # Script para desplegar aplicaciones en Kubernetes



Kubernetes--
Usa kubectl apply -f para aplicar los manifiestos YAML y desplegar las aplicaciones en el clúster de Amazon EKS.