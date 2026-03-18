
echo "Esperando que PostgreSQL esté listo..."
until pg_isready -h db -p 5432 -U postgres; do
  sleep 1
done

echo "Ejecutando ddl.sql..."
psql -h db -p 5432 -U postgres -d istla -f /app/ddl.sql

echo "Ejecutando dml.sql..."
psql -h db -p 5432 -U postgres -d istla -f /app/dml.sql

# Ahora ejecutamos la aplicación Go
echo "Iniciando la aplicación Go..."
./istla