
set -e

echo "--- Iniciando proceso de arranque (Between Bytes Software) ---"

echo "Esperando a que el servicio de base de datos responda..."
until pg_isready -h db -p 5432 -U postgres; do
  echo "La base de datos todavía se está inicializando... reintentando en 2s"
  sleep 2
done

echo "¡Base de datos detectada y lista!"

echo "Lanzando aplicación istla..."
exec ./istla