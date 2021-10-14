# BUILD STAGE
# The base image we are going to use
FROM postgres:latest

# We are going to copy the Makefile into the image so
# that we can execute the database creation commands
COPY ./../Makefile .

# Expose 5432 port
EXPOSE 5432

# Define the default command that will be executed when the container starts
# It's an array of CMD arguments
CMD ["make docker-postgres-createdb", "make run-migrations-up"]

# docker build -t notes-app:latest . 
