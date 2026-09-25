FROM docker.io/library/httpd:2.4

# enable Rewrite module
RUN sed -i 's/#\(LoadModule .*mod_rewrite.so\)/\1/' /usr/local/apache2/conf/httpd.conf
RUN sed -i 's/AllowOverride None/AllowOverride All/' /usr/local/apache2/conf/httpd.conf

# enable Rewrite logging
RUN echo 'LogLevel alert rewrite:trace2' >> /usr/local/apache2/conf/httpd.conf

COPY .htstaccess /usr/local/apache2/htdocs/.htaccess
