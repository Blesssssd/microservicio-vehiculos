# Script para automatizar construcción y ejecución del microservicio

Write-Host "`n🔁 Deteniendo contenedores antiguos..."
docker-compose down -v

Write-Host "`n🔨 Reconstruyendo imagen del microservicio..."
docker-compose build

Write-Host "`n🐳 Levantando contenedores..."
docker-compose up -d

Start-Sleep -Seconds 3

Write-Host "`n✅ Microservicio en ejecución. Accede a: http://localhost:8080"
