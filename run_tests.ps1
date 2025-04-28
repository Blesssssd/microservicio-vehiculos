Write-Host "==> Cargando variables de entorno desde .env.test..."

# Leer .env.test y cargar como variables de entorno
Get-Content ".env.test" | ForEach-Object {
    if ($_ -and ($_ -notmatch "^\s*#") -and ($_ -match "=")) {
        $pair = $_ -split "=", 2
        $key = $pair[0].Trim()
        $value = $pair[1].Trim()
        [System.Environment]::SetEnvironmentVariable($key, $value, "Process")
    }
}

Write-Host "==> Variables cargadas correctamente ✅"
Write-Host "==> Ejecutando pruebas de integración y generando cobertura..."

# Ejecutar tests y guardar el resultado
go test ./controllers_test -coverprofile=coverage.out -v
$testStatus = $LASTEXITCODE

if ($testStatus -eq 0) {
    Write-Host "==> Generando reporte de cobertura en HTML..."
    go tool cover -html=coverage.out -o coverage.html
    Write-Host "✅ Reporte generado en: coverage.html"
} else {
    Write-Host "❌ Las pruebas fallaron. No se generó el reporte de cobertura."
}
