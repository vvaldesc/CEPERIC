# CEPERIC App

Aplicación Angular 18 con Material Design, Firebase y Tailwind CSS, completamente dockerizada.

## 🚀 Stack Tecnológico

- **Angular 18** - Framework principal
- **Angular Material** - Componentes UI con tema Azure/Blue
- **Tailwind CSS** - Utilidades CSS
- **Firebase** - Backend as a Service (AngularFire)
- **Docker** - Contenedorización
- **Nginx** - Servidor web para producción

## 📋 Prerequisitos

- Docker Desktop instalado
- Docker Compose instalado
- (Opcional) Node.js 20+ y npm para desarrollo local

## 🐳 Desarrollo con Docker

### Iniciar entorno de desarrollo
```bash
docker-compose up dev
```

La aplicación estará disponible en: `http://localhost:4200`

**Características del entorno de desarrollo:**
- Hot-reload activado
- Los cambios en el código se reflejan automáticamente
- Volúmenes montados para persistencia

### Detener el servicio
```bash
docker-compose down
```

## 🏭 Producción con Docker

### Construir y ejecutar imagen de producción
```bash
docker-compose up prod
```

La aplicación estará disponible en: `http://localhost:8080`

**Características del entorno de producción:**
- Build optimizado de Angular
- Servidor Nginx configurado
- Gzip habilitado
- Headers de seguridad
- Caché para recursos estáticos

### Build manual
```bash
# Construir imagen
docker build -t ceperic-app:prod .

# Ejecutar contenedor
docker run -p 8080:80 ceperic-app:prod
```

## 💻 Desarrollo Local (sin Docker)

### Instalar dependencias
```bash
npm install
```

### Servidor de desarrollo
```bash
npm start
```

Navega a `http://localhost:4200`

### Build de producción
```bash
npm run build
```

Los artefactos se generarán en `dist/`

## 🔧 Configuración de Firebase

1. Edita `src/app/app.config.ts` con tus credenciales de Firebase
2. Los servicios de Firebase ya están configurados en el proyecto

## 📦 Estructura del Proyecto

```
ceperic-app/
├── src/                    # Código fuente
│   ├── app/               # Módulos y componentes
│   ├── assets/            # Recursos estáticos
│   └── styles.scss        # Estilos globales (con Tailwind)
├── Dockerfile             # Imagen de producción (multi-stage)
├── Dockerfile.dev         # Imagen de desarrollo
├── docker-compose.yml     # Orquestación de servicios
├── nginx.conf            # Configuración de Nginx
├── tailwind.config.js    # Configuración de Tailwind
└── angular.json          # Configuración de Angular
```

## 🎨 Uso de Tailwind con Angular Material

Puedes combinar clases de Tailwind con componentes de Material:

```html
<mat-toolbar class="bg-blue-600 text-white">
  <span class="font-bold text-xl">CEPERIC</span>
</mat-toolbar>

<div class="container mx-auto p-4">
  <mat-card class="shadow-lg">
    <mat-card-content class="space-y-4">
      <!-- Tu contenido -->
    </mat-card-content>
  </mat-card>
</div>
```

## 📝 Scripts Disponibles

- `npm start` - Inicia servidor de desarrollo
- `npm run build` - Build de producción
- `npm test` - Ejecuta tests
- `npm run watch` - Build en modo watch
- `docker-compose up dev` - Desarrollo en Docker
- `docker-compose up prod` - Producción en Docker

## 🔐 Seguridad

El servidor Nginx incluye headers de seguridad:
- X-Frame-Options
- X-Content-Type-Options
- X-XSS-Protection

## 📚 Recursos

- [Angular Documentation](https://angular.dev)
- [Angular Material](https://material.angular.io)
- [Tailwind CSS](https://tailwindcss.com)
- [Firebase](https://firebase.google.com)
- [Docker](https://www.docker.com)

## 🤝 Contribuciones

Este proyecto fue generado con Angular CLI 18.

## 📄 Licencia

MIT
