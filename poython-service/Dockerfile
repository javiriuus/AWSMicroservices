# Usa una imagen base de Python 3
FROM python:3.10-slim

# Establece la carpeta de trabajo
WORKDIR /app

# Copia el archivo requirements.txt
COPY requirements.txt .

# Instala las dependencias
RUN pip install --no-cache-dir -r requirements.txt

# Copia el código fuente de la aplicación
COPY . .

# Expone el puerto 5000
EXPOSE 5000

# Ejecuta la aplicación
CMD ["python", "main.py"]
