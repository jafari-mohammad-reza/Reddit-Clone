FROM postgres:latest

# Set environment variables - not recommended for sensitive data in production
ENV POSTGRES_DB=yourdbname
ENV POSTGRES_USER=youruser
ENV POSTGRES_PASSWORD=yourpassword

# Add custom configurations or extensions here
COPY queries/schema.sql /docker-entrypoint-initdb.d/
