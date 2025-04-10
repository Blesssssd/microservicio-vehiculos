# Obtener fecha actual en formato YYYY-MM-DD
$fecha = Get-Date -Format "yyyy-MM-dd"

# Nombre del directorio destino
$directorio = "respaldo-mongo-$fecha"

Write-Host "Generando respaldo de la base de datos MongoDB 'vehiculosdb'..."

# Ejecutar mongodump dentro del contenedor
docker exec mongo-db mongodump --db=vehiculosdb --out=/backup

# Copiar respaldo a máquina local con fecha
docker cp mongo-db:/backup "./$directorio"

Write-Host "Respaldo completado exitosamente. Carpeta creada: $directorio"
