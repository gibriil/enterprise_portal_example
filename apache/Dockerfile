FROM httpd:2.4-alpine

RUN apk add --no-cache curl xmlsec-dev libxml2-dev openssl mod_auth_mellon bash

COPY httpd.conf /usr/local/apache2/conf/httpd.conf
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh \
    && mkdir -p /etc/httpd/mellon \
    && chmod 700 /etc/httpd/mellon

ENTRYPOINT ["/entrypoint.sh"]