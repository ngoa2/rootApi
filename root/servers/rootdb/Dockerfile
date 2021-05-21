FROM mysql
COPY brands.csv .
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql

EXPOSE 3306