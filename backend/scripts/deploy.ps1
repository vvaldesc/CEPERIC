# Script para hacer deploy manual a Cloud Run
# Para Windows PowerShell

Write-Host "🚀 Desplegando Backend a Cloud Run..." -ForegroundColor Cyan
Write-Host ""

$PROJECT_ID = "ceperic-68bcd"
$BACKEND_DIR = Split-Path -Parent $PSScriptRoot

# Verificar autenticación
try {
    gcloud config get-value project | Out-Null
    Write-Host "✓ Autenticado en GCP" -ForegroundColor Green
} catch {
    Write-Host "✗ No estás autenticado. Ejecuta: gcloud auth login" -ForegroundColor Red
    exit 1
}

# Navegar a la carpeta backend
Set-Location $BACKEND_DIR

Write-Host ""
Write-Host "📦 Iniciando build y deploy desde ./backend..." -ForegroundColor Yellow
Write-Host "   (Esto puede tardar 2-3 minutos)" -ForegroundColor White
Write-Host ""

gcloud builds submit --config=cloudbuild.yaml .

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "✅ Deploy exitoso!" -ForegroundColor Green
    Write-Host ""
    Write-Host "🌍 Obteniendo URL del servicio..." -ForegroundColor Cyan
    
    $SERVICE_URL = gcloud run services describe ceperic-backend `
        --region=europe-west1 `
        --format="value(status.url)"
    
    Write-Host ""
    Write-Host "📝 Tu API está disponible en:" -ForegroundColor Cyan
    Write-Host "   $SERVICE_URL" -ForegroundColor White
    Write-Host ""
    Write-Host "🔍 Endpoints disponibles:" -ForegroundColor Cyan
    Write-Host "   $SERVICE_URL/api/v1/health" -ForegroundColor White
    Write-Host "   $SERVICE_URL/api/v1/users" -ForegroundColor White
    Write-Host ""
    Write-Host "📊 Ver logs:" -ForegroundColor Yellow
    Write-Host "   gcloud run services logs read ceperic-backend --region=europe-west1" -ForegroundColor White
    Write-Host ""
} else {
    Write-Host ""
    Write-Host "✗ Deploy falló. Ver logs arriba." -ForegroundColor Red
    Write-Host ""
    Write-Host "🔍 Troubleshooting:" -ForegroundColor Yellow
    Write-Host "   1. Verifica que cloudbuild.yaml existe en ./backend" -ForegroundColor White
    Write-Host "   2. Verifica que el secret ceperic_db_password existe" -ForegroundColor White
    Write-Host "   3. Revisa los logs de build arriba" -ForegroundColor White
    exit 1
}
