# CompanyBuilder
This project will expose APIs to create standard company structure

DB setup on local:
docker run --name postgresdb -v local_psql_data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=root_pwd -d postgres:12-bullseye

use setup.sql to create Table in company DB and relavant Tables.

create app.env file with all secrets and db.env file to put DB credentials.

just run main.go and application will start on port 80

Docker setup:
1. create app.env file with all secrets and db.env file to put DB credentials.
2. For docker, set TIER=docker and DB_HOST=db in app.env file
2. docker-compose up --build

APIs:
1. Register : This API will register the User.
curl --location --request POST 'localhost:80/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "raman",
    "password": "raman@123",
    "isActive": true
}'

2. Login : This API will authenticate the user and returns token on successfull response
curl --location --request POST 'localhost:80/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "raman",
    "password": "raman@123"
}'

For all companies related API JWT token needs to be passed as Bearer Token which comes from login API

1. GetAllCompanies : This API will return all companies created in the System
curl --location --request GET 'localhost:80/company' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTAwOTYzMTksInVzZXJfaWQiOjF9.56m6hja-Mq9ciNcauvzKIjqF8L8rQTE1c_EhAvOk9h0'

2. GetCompany : This API will return Company information by company id.
curl --location --request GET 'localhost:80/company/499622eb-8712-49d9-9bd3-cbfb0b038525' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTAwOTYzMTksInVzZXJfaWQiOjF9.56m6hja-Mq9ciNcauvzKIjqF8L8rQTE1c_EhAvOk9h0'

3. CreateCompany: This API will create Company in the system
curl --location --request POST 'localhost:80/company/' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTAwOTYzMTksInVzZXJfaWQiOjF9.56m6hja-Mq9ciNcauvzKIjqF8L8rQTE1c_EhAvOk9h0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "apple.com",
    "description": "apple company",
    "employees": 16337,
    "registered": true,
    "type": "IT"
}'

4. UpdateCompany : This API will update Company information in the system
curl --location --request PATCH 'localhost:80/company/499622eb-8712-49d9-9bd3-cbfb0b038525' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTAwOTYzMTksInVzZXJfaWQiOjF9.56m6hja-Mq9ciNcauvzKIjqF8L8rQTE1c_EhAvOk9h0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "description": "google company",
    "registered": false,
    "employees": 894630
}'

5. Delete Company : This API will delete Company data from the system by company id
curl --location --request DELETE 'localhost:80/company/499622eb-8712-49d9-9bd3-cbfb0b038525' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTAwOTYzMTksInVzZXJfaWQiOjF9.56m6hja-Mq9ciNcauvzKIjqF8L8rQTE1c_EhAvOk9h0' \
--data-raw ''