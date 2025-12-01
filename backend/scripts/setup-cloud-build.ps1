# Script para configurar Cloud Build y Cloud Run
# Para Windows PowerShell

Write-Host "🔧 Configurando CI/CD para CEPERIC Backend..." -ForegroundColor Cyan
Write-Host ""

$PROJECT_ID = "ceperic-68bcd"
$REGION = "europe-west1"

# Verificar autenticación
Write-Host "Verificando autenticación con GCP..." -ForegroundColor Yellow
try {
    $current_project = gcloud config get-value project 2>$null
    if ($current_project -ne $PROJECT_ID) {
        Write-Host "⚠️  Proyecto actual: $current_project" -ForegroundColor Yellow
        Write-Host "   Cambiando a: $PROJECT_ID" -ForegroundColor Yellow
        gcloud config set project $PROJECT_ID
    }
    Write-Host "✓ Autenticado en proyecto: $PROJECT_ID" -ForegroundColor Green
} catch {
    Write-Host "✗ No estás autenticado. Ejecuta: gcloud auth login" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "1️⃣  Habilitando APIs necesarias..." -ForegroundColor Cyan
gcloud services enable cloudbuild.googleapis.com
gcloud services enable run.googleapis.com
gcloud services enable secretmanager.googleapis.com
gcloud services enable artifactregistry.googleapis.com
Write-Host "✓ APIs habilitadas" -ForegroundColor Green

Write-Host ""
Write-Host "2️⃣  Configurando Secret para DB Password..." -ForegroundColor Cyan
$db_password = Read-Host "Ingresa el password de Cloud SQL" -AsSecureString
$db_password_plain = [Runtime.InteropServices.Marshal]::PtrToStringAuto(
    [Runtime.InteropServices.Marshal]::SecureStringToBSTR($db_password)
)

# Crear secret
echo $db_password_plain | gcloud secrets create ceperic_db_password --data-file=- 2>$null
if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Secret creado: ceperic_db_password" -ForegroundColor Green
} else {
    Write-Host "⚠️  Secret ya existe, actualizando..." -ForegroundColor Yellow
    echo $db_password_plain | gcloud secrets versions add ceperic_db_password --data-file=-
    Write-Host "✓ Secret actualizado" -ForegroundColor Green
}

Write-Host ""
Write-Host "3️⃣  Configurando permisos IAM..." -ForegroundColor Cyan
$PROJECT_NUMBER = gcloud projects describe $PROJECT_ID --format="value(projectNumber)"

# Permisos para leer secrets
gcloud secrets add-iam-policy-binding ceperic_db_password `
  --member="serviceAccount:$PROJECT_NUMBER-compute@developer.gserviceaccount.com" `
  --role="roles/secretmanager.secretAccessor" `
  --quiet

# Permisos para Cloud Build
gcloud projects add-iam-policy-binding $PROJECT_ID `
  --member="serviceAccount:$PROJECT_NUMBER@cloudbuild.gserviceaccount.com" `
  --role="roles/run.admin" `
  --quiet

gcloud iam service-accounts add-iam-policy-binding `
  "$PROJECT_NUMBER-compute@developer.gserviceaccount.com" `
  --member="serviceAccount:$PROJECT_NUMBER@cloudbuild.gserviceaccount.com" `
  --role="roles/iam.serviceAccountUser" `
  --quiet

Write-Host "✓ Permisos configurados" -ForegroundColor Green

Write-Host ""
Write-Host "4️⃣  Configurando Cloud Build Trigger..." -ForegroundColor Cyan
Write-Host ""
Write-Host "Abriendo Cloud Console para conectar GitHub..." -ForegroundColor Yellow
Write-Host "Configura el trigger con estos valores:" -ForegroundColor White
Write-Host ""
Write-Host "  📁 Repositorio: vvaldesc/CEPERIC" -ForegroundColor Cyan
Write-Host "  🌿 Rama: ^go-starter$" -ForegroundColor Cyan
Write-Host "  📄 Tipo: Cloud Build configuration file" -ForegroundColor Cyan
Write-Host "  📝 Ubicación: /backend/cloudbuild.yaml" -ForegroundColor Cyan
Write-Host ""
Start-Sleep -Seconds 2
Start-Process "https://console.cloud.google.com/cloud-build/triggers?project=$PROJECT_ID"

Write-Host ""
Write-Host "✅ Setup completado!" -ForegroundColor Green
Write-Host ""
Write-Host "📝 Próximos pasos:" -ForegroundColor Cyan
Write-Host "   1. Configura el trigger en la consola (se abrió automáticamente)" -ForegroundColor White
Write-Host "   2. Push tu código a GitHub" -ForegroundColor White
Write-Host "   3. Cloud Build desplegará automáticamente a Cloud Run" -ForegroundColor White
Write-Host ""
Write-Host "🚀 O haz deploy manual ahora con:" -ForegroundColor Yellow
Write-Host "   cd backend" -ForegroundColor White
Write-Host "   ..\scripts\deploy.ps1" -ForegroundColor White
Write-Host ""
