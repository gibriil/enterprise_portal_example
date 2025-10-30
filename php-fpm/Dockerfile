FROM php:8.3.20-fpm-alpine

ADD --chmod=0755 https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions /usr/local/bin/

RUN install-php-extensions gd xdebug

COPY etc/xdebug/xdebug.ini /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini
COPY etc/xdebug/error_reporting.ini /usr/local/etc/php/conf.d/error_reporting.ini
COPY etc/www.conf /usr/local/etc/php-fpm.d/www.conf
COPY src/ /var/www/html