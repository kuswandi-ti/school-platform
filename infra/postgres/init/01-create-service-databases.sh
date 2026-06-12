#!/usr/bin/env bash
set -euo pipefail

service_databases=(
  identity_db
  school_core_db
  admission_db
  academic_db
  finance_db
  communication_db
  reporting_db
)

for database in "${service_databases[@]}"; do
  if psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -tAc "SELECT 1 FROM pg_database WHERE datname = '$database'" | grep -q 1; then
    echo "Database $database already exists"
  else
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -c "CREATE DATABASE \"$database\" OWNER \"$POSTGRES_USER\";"
  fi
done
