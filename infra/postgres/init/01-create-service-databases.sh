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
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-SQL
    CREATE DATABASE "$database" OWNER "$POSTGRES_USER";
SQL
done
