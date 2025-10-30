#!/bin/sh
set -e

MELLON_DIR=/etc/httpd/mellon
KEYCLOAK_URL=http://keycloak:8080/realms/myrealm/protocol/saml/descriptor

# Wait for Keycloak to be ready
echo "==> Waiting for Keycloak..."
while ! curl -s "$KEYCLOAK_URL" > /dev/null; do
    echo "Waiting for Keycloak SAML metadata..."
    sleep 5
done
echo "Keycloak is ready."

# Download IdP metadata
echo "==> Downloading IdP metadata..."
curl -sSL "$KEYCLOAK_URL" -o "$MELLON_DIR/idp-metadata.xml"
echo "IdP metadata saved to $MELLON_DIR/idp-metadata.xml"

# Generate SP metadata if missing
if [ ! -f "$MELLON_DIR/sp-metadata.xml" ]; then
    echo "==> Generating SP metadata..."
    mellon_create_metadata.sh "http://localhost:80/mellon" "http://localhost:80/mellon" \
        > "$MELLON_DIR/sp-metadata.xml"
    echo "SP metadata generated."
fi

# Start Apache
echo "==> Starting Apache..."
httpd-foreground
