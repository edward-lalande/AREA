FROM postgres:17.2

COPY time.sql /docker-entrypoint-initdb.d/

RUN chmod 644 /docker-entrypoint-initdb.d/time.sql

EXPOSE 5433
