# 🚀 Deploy Frontend con Nginx (Ubuntu)

## 1. Instalar dependencias

```bash
npm install

2. Generar build
npm run build

Esto generará una carpeta como:

dist/ (Vue, Vite, Angular)
build/ (React)

3. Copiar archivos a Nginx

Crear carpeta destino:
sudo mkdir -p /var/www/frontend

Copiar archivos (ejemplo con dist):
sudo cp -r dist/* /var/www/frontend/

Si usas React:
sudo cp -r build/* /var/www/frontend/

4. Configurar Nginx

Crear archivo:
sudo nano /etc/nginx/sites-available/frontend

Contenido:
server {
    listen 80;
    server_name tu-dominio.com;

    root /var/www/frontend;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }
}

5. Activar configuración

sudo ln -s /etc/nginx/sites-available/frontend /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx

6. Actualizar el frontend

Cada vez que hagas cambios:
npm run build
sudo rm -rf /var/www/frontend/*
sudo cp -r dist/* /var/www/frontend/
sudo systemctl reload nginx


✅ Resultado

Tu aplicación estará disponible en:
http://tu-dominio.com en este caso seria como localhost en fururas actualizaciones podriamos configurar usando un dominio + dns o deploy automático (CI/CD)


O Tambien podremos solo usar el comando npm run dev para correr la aplicacion del frontend
