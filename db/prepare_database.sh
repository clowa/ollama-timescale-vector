#!/bin/bash

set -e
set -u

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS embeddings (
        id uuid PRIMARY KEY,
        md5 text UNIQUE,
        embedding vector,
        text text,
        created_at timestamptz DEFAULT now()
    );
EOSQL
