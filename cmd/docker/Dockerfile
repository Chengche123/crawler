FROM mysql:8.0.24
LABEL maintainer="chengche123"
ENV MYSQL_ROOT_PASSWORD=root MYSQL_DATABASE=comic

COPY *.sql /docker-entrypoint-initdb.d/
