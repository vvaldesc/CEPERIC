# Script para iniciar el proyecto en modo producción con Docker
# Para Windows PowerShell

Write-Host "🏭 Iniciando CEPERIC App en modo producción..." -ForegroundColor Cyan
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
Write-Host "Construyendo imagen de producción..." -ForegroundColor Yellow
docker-compose -f ceperic-app/docker-compose.yml up --build prod

# Si el comando anterior falla
if ($LASTEXITCODE -ne 0) {
    Write-Host ""
    Write-Host "✗ Error al iniciar el proyecto" -ForegroundColor Red
    exit 1
}
