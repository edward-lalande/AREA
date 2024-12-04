FROM postgres:17.2

COPY discord.sql /docker-entrypoint-initdb.d/

RUN chmod 644 /docker-entrypoint-initdb.d/discord.sql

EXPOSE 5434
