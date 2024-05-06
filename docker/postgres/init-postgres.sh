#!/bin/bash
# Copy any SQL init files to /docker-entrypoint-initdb.d
# Adjust log file permissions
chmod -R 644 /var/lib/postgresql/data/pg_logs/*.log

# Call original entry point script
docker-entrypoint.sh postgres
