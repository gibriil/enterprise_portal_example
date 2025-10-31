#!/bin/sh
set -e

KEYCLOAK_URL="http://keycloak:8080/realms/myportal/.well-known/openid-configuration"

echo "Waiting for Keycloak at $KEYCLOAK_URL ..."
until curl -fsS "$KEYCLOAK_URL" > /dev/null 2>&1; do
    echo "Keycloak not ready yet..."
    sleep 5
done

echo "✅ Keycloak is ready!"
exec httpd-foreground