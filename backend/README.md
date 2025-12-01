# CEPERIC Backend API (Go + Fiber)

Backend API en Go para el proyecto CEPERIC, completamente dockerizado.

## 🚀 Características

- **Framework:** Fiber (Express-like para Go)
- **ORM:** GORM con PostgreSQL
- **Arquitectura:** Clean Architecture (Repository, Service, Handler)
- **Hot-reload:** Air para desarrollo
- **Docker:** Completamente dockerizado (no necesitas Go instalado)

## 📋 Requisitos

- Docker
- Docker Compose
- (Opcional) Go 1.21+ si quieres desarrollo local sin Docker

## 🏗️ Estructura del Proyecto

```
backend/
├── cmd/
│   └── api/
│       └── main.go              # Entry point
├── internal/
│   ├── config/                  # Configuración
│   ├── domain/                  # Modelos y DTOs
│   ├── repository/              # Capa de datos
│   ├── service/                 # Lógica de negocio
│   ├── handler/                 # HTTP handlers
│   └── router/                  # Rutas
├── pkg/
│   └── response/                # Utilidades compartidas
├── Dockerfile                   # Producción
├── Dockerfile.dev               # Desarrollo
└── docker-compose.yml           # Orquestación
```

## 🐳 Desarrollo con Docker

### Configurar conexión a Cloud SQL:

1. Copia `.env.example` a `.env`
2. Configura la IP pública de tu Cloud SQL o usa Cloud SQL Proxy
3. Levanta el backend:

```powershell
docker-compose up
```

Esto levanta solo el **Backend Go** en http://localhost:8080 (sin PostgreSQL local)

### Ver logs:

```powershell
docker-compose logs -f backend
```

### Rebuild:

```powershell
docker-compose up --build backend
```

## 📡 Endpoints API

### Health Check
```
GET /api/v1/health
```

### Users
```
GET    /api/v1/users       - Listar usuarios
GET    /api/v1/users/:id   - Obtener usuario
POST   /api/v1/users       - Crear usuario
PUT    /api/v1/users/:id   - Actualizar usuario
DELETE /api/v1/users/:id   - Eliminar usuario
```

### Ejemplo de uso:

```powershell
# Crear usuario
curl -X POST http://localhost:8080/api/v1/users `
  -H "Content-Type: application/json" `
  -d '{\"email\":\"test@ceperic.com\",\"name\":\"Test User\"}'

# Listar usuarios
curl http://localhost:8080/api/v1/users
```

## 🔧 Variables de Entorno

Copia `.env.example` a `.env` y configura:

```bash
DB_HOST=postgres
DB_PORT=5432
DB_USER=ceperic_user
DB_PASSWORD=ceperic_pass
DB_NAME=ceperic_db

FIREBASE_PROJECT_ID=ceperic-68bcd
ENVIRONMENT=development
PORT=8080
ALLOWED_ORIGINS=http://localhost:4200
```

## 🧪 Testing

```powershell
# Ejecutar tests
docker-compose exec backend go test ./...

# Con coverage
docker-compose exec backend go test -cover ./...
```

## 📦 Build para Producción

```powershell
# Build imagen
docker build -t ceperic-backend .

# Run producción
docker run -p 8080:8080 --env-file .env ceperic-backend
```

## 🚀 Deploy a Cloud Run

```powershell
# Build y push
gcloud builds submit --tag gcr.io/ceperic-68bcd/backend

# Deploy
gcloud run deploy ceperic-backend `
  --image gcr.io/ceperic-68bcd/backend `
  --region europe-west1 `
  --platform managed `
  --allow-unauthenticated
```

## 📝 Notas

- **Hot-reload:** El código se recarga automáticamente en desarrollo
- **Migraciones:** GORM AutoMigrate crea/actualiza tablas automáticamente
- **Sin Go local:** Todo funciona dentro de Docker, no necesitas instalar Go
