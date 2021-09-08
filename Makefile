# This rule creates the container with docker using the postgres image
docker-container-create:
	docker run --name notesApp -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

# This rule runs the postgres container (If it is stopped)
docker-container-start:
	docker container start notesApp

# This rule stops the posgres container
docker-container-stop:
	docker container stop notesApp

# This rule creates the db on the container
docker-postgres-createdb:
	docker container exec -it notesApp createdb --username=postgres --owner=postgres notes-app-db
	# docker container exec -it notesApp bash
	# psql -U postgres
	# \c notes-app-db
	# GRANT SELECT, UPDATE, INSERT, DELETE ON invoiceheaders TO postgres;
# We must grant on all tables

# This rule deletes the db on the container
docker-postgres-dropdb:
	docker exec -it notesApp dropdb notes-app-db

# This rule runs the migrations up
run-migrations-up:
	migrate --path database/migrations --database "postgresql://postgres:postgres@localhost:5432/notes-app-db?sslmode=disable" --verbose up

run-migrations-down:
	migrate --path database/migrations --database "postgresql://postgres:postgres@localhost:5432/notes-app-db?sslmode=disable" --verbose down

# .PHONY tell explicitly to MAKE that those rules are not associated with files
.PHONY: docker-container-create docker-container-start docker-container-stop docker-postgres-createdb docker-postgres-dropdb run-migrations-up run-migrations-down