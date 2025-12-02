# Script para deploy manual del frontend
# Para Windows PowerShell

Write-Host "🚀 Desplegando Frontend a Firebase Hosting..." -ForegroundColor Cyan
Write-Host ""

$FRONTEND_DIR = "C:\Users\Usuario\Documents\DAW\VSC_Workspace\CEPERIC\frontend"

# Verificar que estamos en el directorio correcto
if (-Not (Test-Path $FRONTEND_DIR)) {
    Write-Host "✗ Directorio frontend no encontrado" -ForegroundColor Red
    exit 1
}

Set-Location $FRONTEND_DIR

Write-Host "📦 Instalando dependencias..." -ForegroundColor Yellow
npm install

if ($LASTEXITCODE -ne 0) {
    Write-Host "✗ Error instalando dependencias" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "🔨 Building para producción..." -ForegroundColor Yellow
npm run build -- --configuration=production

if ($LASTEXITCODE -ne 0) {
    Write-Host "✗ Error en build" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "🚀 Desplegando a Firebase..." -ForegroundColor Yellow
firebase deploy --only hosting

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "✅ Deploy exitoso!" -ForegroundColor Green
    Write-Host ""
    Write-Host "🌍 Tu aplicación está disponible en:" -ForegroundColor Cyan
    Write-Host "   https://ceperic-68bcd.web.app" -ForegroundColor White
    Write-Host "   https://ceperic-68bcd.firebaseapp.com" -ForegroundColor White
    Write-Host ""
} else {
    Write-Host ""
    Write-Host "✗ Deploy falló" -ForegroundColor Red
    exit 1
}
