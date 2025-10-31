#!/bin/bash
# wait-for-keycloak.sh
until curl -s http://keycloak:8080/realms/myportal/.well-known/openid-configuration; do
    echo "Waiting for Keycloak..."
    sleep 5
done
echo "Keycloak ready! Starting Apache..."
apache2ctl -D FOREGROUND
