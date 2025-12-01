# Script para iniciar el backend Go en modo desarrollo
# Para Windows PowerShell

Write-Host "🚀 Iniciando CEPERIC Backend API (Go)..." -ForegroundColor Green
Write-Host ""

# Verificar si Docker está en ejecución
try {
    docker info | Out-Null
    Write-Host "✓ Docker está en ejecución" -ForegroundColor Green
} catch {
    Write-Host "✗ Docker no está en ejecución. Por favor, inicia Docker Desktop." -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "📦 Iniciando Backend API..." -ForegroundColor Yellow
Set-Location backend
docker-compose up

# Si el comando anterior falla
if ($LASTEXITCODE -ne 0) {
    Write-Host ""
    Write-Host "✗ Error al iniciar el backend" -ForegroundColor Red
    Set-Location ..
    exit 1
}

Set-Location ..
