# CompanyBuilder
This project will expose APIs to create standard company structure

DB setup on local:
docker run --name postgresdb -v local_psql_data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=root_pwd -d postgres:12-bullseye

use setup.sql to create Table in company DB.

just run main.go and application will start on port 8080

Docker setup:
1. create .env file
2. docker-compose up --build