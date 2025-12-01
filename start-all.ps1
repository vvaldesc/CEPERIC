# Script para iniciar todo el proyecto (Frontend + Backend)
# Para Windows PowerShell

Write-Host "🚀 Iniciando proyecto completo CEPERIC..." -ForegroundColor Cyan
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
Write-Host "📡 Iniciando Backend API (Go)..." -ForegroundColor Yellow
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PWD\backend'; docker-compose up"

Start-Sleep -Seconds 3

Write-Host "🎨 Iniciando Frontend (Angular)..." -ForegroundColor Yellow
if (Test-Path "ceperic-app/docker-compose.yml") {
    Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PWD\ceperic-app'; docker-compose up dev"
} else {
    Write-Host "⚠️  Frontend no encontrado en ceperic-app/" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "✅ Servicios iniciándose en ventanas separadas..." -ForegroundColor Green
Write-Host ""
Write-Host "📝 URLs disponibles:" -ForegroundColor Cyan
Write-Host "   Backend API:  http://localhost:8080/api/v1/health" -ForegroundColor White
Write-Host "   Frontend:     http://localhost:4200" -ForegroundColor White
Write-Host ""
Write-Host "⚠️  Usa stop-all.ps1 para detener todos los servicios" -ForegroundColor Yellow
