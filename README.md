# CompanyManagementSystem
This project will expose APIs to create standard company structure

Project Structure:
This Project follows MVC architecture
1. apis : this directory contains controller and api routes.
2. cmd : this directory contains main.go file.
3. config : this directory contains config files and global config model.
4. dal : this directory handles all data access layer operations.
5. initiate : this directory contains logic for dependency injections of different modules.
6. models: this directory contains models for different modules.
7. modules: this directory contains modules which have business logic. this is bascially service layer.
8. pkg : this directory contains logic of external packages integration.
9. scripts: this directory contains different scripts such as for linting or unit tests.
10. shared: this directory contains model for all internal dependencies of the project.
11. sql: this directory contains all sql related data.
Local Setup:
1. Run Postgres DB via docker using below command
docker run --name postgresdb -v local_psql_data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=root_pwd -d postgres:12-bullseye
2. use setup.sql to create DB and relavant Tables.
3. create app.env file with all secrets.
4. just run main.go and application will start on port 8080

Docker setup:
1. create app.env file with all secrets and db.env file to put DB credentials.
2. For docker, set TIER=docker and DB_HOST=db in app.env file
3. docker-compose up --build
4. Application will start on port 80

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