FROM httpd:2.4

# Enable required Apache modules
RUN sed -i \
    -e 's/^#LoadModule include_module/LoadModule include_module/' \
    -e 's/^#LoadModule proxy_module/LoadModule proxy_module/' \
    -e 's/^#LoadModule proxy_fcgi_module/LoadModule proxy_fcgi_module/' \
    -e 's/^#LoadModule proxy_http_module/LoadModule proxy_http_module/' \
    -e 's/^#LoadModule rewrite_module/LoadModule rewrite_module/' \
    -e 's/^#LoadModule headers_module/LoadModule headers_module/' \
    -e 's/^#LoadModule mime_module/LoadModule mime_module/' \
    -e 's/^#LoadModule setenvif_module/LoadModule setenvif_module/' \
    /usr/local/apache2/conf/httpd.conf

COPY httpd.conf /usr/local/apache2/conf/extra/000-default.conf

RUN echo "Listen 8500" >> /usr/local/apache2/conf/httpd.conf
RUN echo "Include conf/extra/000-default.conf" >> /usr/local/apache2/conf/httpd.conf

EXPOSE 80