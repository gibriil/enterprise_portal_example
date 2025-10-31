FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y apache2 curl libapache2-mod-auth-openidc \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN a2enmod mpm_event proxy proxy_fcgi auth_openidc headers rewrite deflate brotli mime env include

COPY httpd.conf /etc/apache2/sites-available/000-default.conf
COPY startup.sh /startup.sh

RUN chmod +x /startup.sh

EXPOSE 80

CMD ["/startup.sh"]