FROM httpd:2.4-alpine

RUN rm -rf \
    conf/extra/ \
    conf/magic \
    conf/original/

WORKDIR /tmp

RUN echo "https://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
    && apk add --no-cache curl \
    && apk add --no-cache apache2-utils \
    && apk add --no-cache apache2 \
    && apk add --no-cache apache-mod-auth-openidc \
    && cp /usr/lib/apache2/mod_auth_openidc.so /usr/local/apache2/modules/ \
    && rm -rf /var/cache/apk/* /tmp/*


COPY /etc /usr/local/apache2/conf/
COPY wait-for-keycloak.sh /wait-for-keycloak.sh

RUN chmod +x /wait-for-keycloak.sh

EXPOSE 80

CMD ["/wait-for-keycloak.sh"]