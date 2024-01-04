FROM postgres:latest

ENV POSTGRES_DB=reddit-clone
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin


COPY queries/schema.sql /docker-entrypoint-initdb.d/
