# Usa una imagen base de Node.js
FROM node:18 AS builder

# Establece la carpeta de trabajo
WORKDIR /app

# Copia los archivos necesarios
COPY package.json yarn.lock ./

# Instala dependencias
RUN yarn install

# Copia el código fuente de la aplicación
COPY . .

# Compila la aplicación Next.js
RUN yarn build

# Usa una imagen base de Nginx
FROM nginx:latest

# Copia los archivos de compilación al servidor Nginx
COPY --from=builder /app/out /usr/share/nginx/html

# Expone el puerto 80
EXPOSE 80

# Copia el archivo de configuración personalizado de Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Inicia el servidor Nginx
CMD ["nginx", "-g", "daemon off;"]
