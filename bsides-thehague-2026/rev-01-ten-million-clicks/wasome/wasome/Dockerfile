FROM docker.io/library/httpd:2.4

RUN printf '\nAddType application/wasm .wasm\n' >> /usr/local/apache2/conf/httpd.conf

COPY index.html /usr/local/apache2/htdocs/index.html
COPY challenge.wasm /usr/local/apache2/htdocs/challenge.wasm

COPY cheat.html /usr/local/apache2/htdocs/cheat.html
COPY challenge.js /usr/local/apache2/htdocs/challenge.js
