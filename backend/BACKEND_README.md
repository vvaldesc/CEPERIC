# CEPERIC Backend - Go API

API REST desarrollada en Go con Clean Architecture para el proyecto CEPERIC.

## 🏗️ Arquitectura

```
backend/
├── cmd/
│   └── api/
│       └── main.go              # Entry point
├── internal/
│   ├── config/                  # Configuración
│   ├── domain/                  # Modelos de datos
│   ├── repository/              # Acceso a datos
│   ├── service/                 # Lógica de negocio
│   ├── handler/                 # HTTP handlers
│   └── router/                  # Rutas
├── pkg/
│   └── response/                # Utilidades
├── scripts/                     # Scripts de deployment
│   ├── setup-cloud-build.ps1   # Setup inicial CI/CD
│   └── deploy.ps1              # Deploy manual
├── cloudbuild.yaml             # CI/CD config
├── Dockerfile                  # Production build
├── docker-compose.yml          # Desarrollo local
└── .env                        # Variables de entorno
```

## 🚀 Quick Start

### Desarrollo Local

```powershell
# Desde la raíz del proyecto
.\start-backend.ps1

# O desde backend/
docker-compose up

# La API estará en http://localhost:8080
```

### Endpoints

```
GET  /api/v1/health              # Health check
GET  /api/v1/users               # Listar usuarios
POST /api/v1/users               # Crear usuario
GET  /api/v1/users/:id           # Obtener usuario
PUT  /api/v1/users/:id           # Actualizar usuario
DEL  /api/v1/users/:id           # Eliminar usuario
```

## ☁️ Deploy a Cloud Run

### Setup Inicial (una sola vez)

```powershell
cd backend
.\scripts\setup-cloud-build.ps1
```

Este script:
- ✅ Habilita APIs necesarias
- ✅ Crea secret para DB password
- ✅ Configura permisos IAM
- ✅ Abre consola para conectar GitHub

### Configurar Trigger en Cloud Console

1. En la consola que se abre automáticamente
2. Click "Create Trigger"
3. Conecta GitHub → `vvaldesc/CEPERIC`
4. Configura:
   ```
   Rama: ^go-starter$
   Tipo: Cloud Build configuration file
   Ubicación: /backend/cloudbuild.yaml
   ```

### Deploy Automático

```powershell
# Simplemente haz push a go-starter
git push origin go-starter

# Cloud Build lo desplegará automáticamente
```

### Deploy Manual

```powershell
cd backend
.\scripts\deploy.ps1
```

## 🔧 Variables de Entorno

### Desarrollo Local (`.env`)

```env
DB_HOST=                    # Vacío = sin DB
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=
DB_NAME=ceperic_db
ENVIRONMENT=development
PORT=8080
```

### Producción (Cloud Run)

Configuradas automáticamente en `cloudbuild.yaml`:
- `DB_HOST` → Unix socket de Cloud SQL
- `DB_PASSWORD` → Desde Secret Manager
- `ENVIRONMENT` → production

## 🗄️ Base de Datos

### Cloud SQL

```
Instancia: ceperic-db
Región: europe-west1
Motor: PostgreSQL 15
IP pública: 34.175.213.10
```

### Conexión Local (opcional)

```powershell
# Instalar Cloud SQL Proxy
gcloud components install cloud-sql-proxy

# Conectar
cloud-sql-proxy ceperic-68bcd:europe-west1:ceperic-db
```

## 📊 Monitoreo

```powershell
# Ver logs de Cloud Run
gcloud run services logs read ceperic-backend --region=europe-west1

# Ver logs de Cloud Build
gcloud builds list --limit=10
```

## 🔐 Secrets

```powershell
# Ver secret
gcloud secrets versions access latest --secret=ceperic_db_password

# Actualizar
echo "nuevo_password" | gcloud secrets versions add ceperic_db_password --data-file=-
```

## 🛠️ Comandos Útiles

```powershell
# Estado del servicio
gcloud run services describe ceperic-backend --region=europe-west1

# Ver URL del servicio
gcloud run services describe ceperic-backend --region=europe-west1 --format="value(status.url)"

# Actualizar configuración
gcloud run services update ceperic-backend `
  --region=europe-west1 `
  --memory=1Gi

# Ver triggers de Cloud Build
gcloud builds triggers list
```

## 🧪 Testing

```powershell
# Ejecutar tests
go test ./...

# Con coverage
go test -cover ./...

# Tests específicos
go test ./internal/service/...
```

## 📦 Dependencias

- **Fiber v2.52.0** - Web framework
- **GORM v1.25.5** - ORM
- **godotenv** - Variables de entorno
- **uuid** - Generación de UUIDs

## 💰 Costos Estimados

Con configuración actual:
- **Free tier**: 2M requests/mes
- **Después**: ~$0.00002/request
- **Configuración**: 512Mi RAM, 1 CPU
- **Scaling**: 0-10 instancias

## 📚 Recursos

- [Cloud Run Docs](https://cloud.google.com/run/docs)
- [Cloud Build Docs](https://cloud.google.com/build/docs)
- [Fiber Framework](https://docs.gofiber.io/)
- [GORM ORM](https://gorm.io/)
