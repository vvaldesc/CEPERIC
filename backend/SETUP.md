# CEPERIC Backend - Setup Guide

## 🎯 Arquitectura

Backend completamente independiente y dockerizado que se conecta a:
- **Cloud SQL PostgreSQL** (IP pública o Cloud SQL Proxy)
- **Firebase** para autenticación y storage
- Sin bases de datos locales

## 🚀 Inicio Rápido

### 1. Configurar Variables de Entorno

Copia `.env.example` a `.env` y configura tu conexión a Cloud SQL:

```bash
cp .env.example .env
```

Edita `.env` con tus credenciales:

```env
# Opción 1: Conexión directa a IP pública de Cloud SQL
DB_HOST=34.175.xxx.xxx  # Tu IP pública de Cloud SQL
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=ceperic_db

# Opción 2: Cloud SQL Proxy (recomendado para desarrollo)
# DB_HOST=host.docker.internal
# DB_PORT=5432
```

### 2. Levantar el Backend

```powershell
# En la carpeta backend/
docker-compose up
```

El backend estará disponible en http://localhost:8080

### 3. Probar la API

```powershell
# Health check
curl http://localhost:8080/api/v1/health

# Crear usuario
curl -X POST http://localhost:8080/api/v1/users `
  -H "Content-Type: application/json" `
  -d '{\"email\":\"test@ceperic.com\",\"name\":\"Test User\"}'

# Listar usuarios
curl http://localhost:8080/api/v1/users
```

## 🔌 Opciones de Conexión a Cloud SQL

### Opción 1: IP Pública (Más Simple)

1. Obtén la IP pública de tu instancia Cloud SQL:
   ```powershell
   gcloud sql instances describe ceperic-db --format="value(ipAddresses[0].ipAddress)"
   ```

2. Añade tu IP a las redes autorizadas:
   ```powershell
   gcloud sql instances patch ceperic-db --authorized-networks=TU_IP
   ```

3. Usa la IP directamente en `.env`:
   ```env
   DB_HOST=34.175.xxx.xxx
   ```

### Opción 2: Cloud SQL Proxy (Más Seguro)

1. Instala Cloud SQL Proxy localmente:
   ```powershell
   curl -o cloud-sql-proxy.exe https://storage.googleapis.com/cloud-sql-connectors/cloud-sql-proxy/v2.8.0/cloud-sql-proxy.x64.exe
   ```

2. Ejecuta el proxy:
   ```powershell
   .\cloud-sql-proxy.exe ceperic-68bcd:europe-west1:ceperic-db
   ```

3. Configura `.env` para usar el proxy:
   ```env
   DB_HOST=host.docker.internal
   DB_PORT=5432
   ```

## 📦 Comandos Útiles

```powershell
# Ver logs
docker-compose logs -f

# Rebuild
docker-compose up --build

# Parar
docker-compose down

# Limpiar todo
docker-compose down -v
```

## 🌍 Estructura de Carpetas

```
backend/
├── cmd/api/              # Entry point
├── internal/
│   ├── config/          # Configuración
│   ├── domain/          # Modelos
│   ├── repository/      # Base de datos
│   ├── service/         # Lógica
│   ├── handler/         # HTTP handlers
│   └── router/          # Rutas
├── pkg/                 # Utilidades
├── docker-compose.yml   # Solo backend
├── Dockerfile.dev       # Desarrollo
├── Dockerfile           # Producción
└── .env                 # Variables (no commitear)
```

## 🚢 Deploy a Cloud Run

```powershell
# Build
gcloud builds submit --tag gcr.io/ceperic-68bcd/backend

# Deploy con Cloud SQL
gcloud run deploy ceperic-backend `
  --image gcr.io/ceperic-68bcd/backend `
  --region europe-west1 `
  --add-cloudsql-instances ceperic-68bcd:europe-west1:ceperic-db `
  --set-env-vars DB_HOST=/cloudsql/ceperic-68bcd:europe-west1:ceperic-db `
  --allow-unauthenticated
```

## 📝 Notas

- **No hay PostgreSQL local**: Todo conecta a Cloud SQL
- **Hot-reload**: El código se recarga automáticamente
- **Separación total**: Backend independiente del frontend
- **Listo para producción**: Multi-stage build optimizado
