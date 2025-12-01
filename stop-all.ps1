# Script para detener todos los servicios Docker del proyecto
# Para Windows PowerShell

Write-Host "🛑 Deteniendo servicios CEPERIC..." -ForegroundColor Yellow
Write-Host ""

# Detener backend
Write-Host "Deteniendo Backend API..." -ForegroundColor Cyan
Set-Location backend
docker-compose down
Set-Location ..

# Detener frontend (si existe)
if (Test-Path "ceperic-app/docker-compose.yml") {
    Write-Host "Deteniendo Frontend..." -ForegroundColor Cyan
    Set-Location ceperic-app
    docker-compose down
    Set-Location ..
}

Write-Host ""
Write-Host "✓ Todos los servicios detenidos" -ForegroundColor Green
