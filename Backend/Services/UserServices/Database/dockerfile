FROM postgres:17.2

COPY user.sql /docker-entrypoint-initdb.d/

RUN chmod 644 /docker-entrypoint-initdb.d/user.sql

EXPOSE 5432
