# This rule compiles the project
golang-build:
	go build -o ./build/ ./...

# This rule creates the container with docker using the postgres image
docker-container-create:
	docker run --name backend-postgres-db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

# This rule runs the postgres container (If it is stopped)
docker-container-start:
	docker container start backend-postgres-db

# This rule stops the posgres container
docker-container-stop:
	docker container stop backend-postgres-db

# This rule creates the db on the container
docker-postgres-createdb:
	docker container exec -it backend-postgres-db createdb --username=postgres --owner=postgres go-notes-database
# docker container exec -it backend-postgres-db bash
# psql -U postgres
# \c go-notes-database
# GRANT SELECT, UPDATE, INSERT, DELETE ON invoiceheaders TO postgres;
# We must grant on all tables

# This rule deletes the db on the container
docker-postgres-dropdb:
	docker exec -it backend-postgres-db dropdb go-notes-database

# This rule runs the migrations up
run-migrations-up:
	migrate --path database/migrations --database "postgresql://postgres:postgres@localhost:5432/go-notes-database?sslmode=disable" --verbose up

run-migrations-down:
	migrate --path database/migrations --database "postgresql://postgres:postgres@localhost:5432/go-notes-database?sslmode=disable" --verbose down

# .PHONY tell explicitly to MAKE that those rules are not associated with files
.PHONY: docker-container-create docker-container-start docker-container-stop docker-postgres-createdb docker-postgres-dropdb run-migrations-up run-migrations-down