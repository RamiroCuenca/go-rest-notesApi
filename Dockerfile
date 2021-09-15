# BUILD STAGE
# The base image we are going to use
FROM golang:1.17.1-alpine3.14 AS builder

# Declare the current directory inside the image
WORKDIR /app

# Copy all necessary files inside the image
# The first dot means copy all from current directory and
# the second means for where to paste it (Inside the image)
# A dot means from the root of the project, in this case: /app
COPY . .

# Next, wi want to build our app to a single binary executable file
# The -o means output
RUN go build -o build/ ./...

# RUN STAGE
FROM alpine:3.14

WORKDIR /app

# Copy the executable binary file into this run stage image
# We use the same COPY command but we add "from" flag to 
# tell docker where to copy the file from
# The second dot represents the workdir
COPY --from=builder /app/build .

# It is a good practice to add an EXPOSE instruction in order to 
# inform docker that the container listens on the specified network port
EXPOSE 8000

# Define the default command that will be executed when the container starts
# It's an array of CMD arguments
CMD "./cmd"
# ./app/build

# docker build -t notes-app:latest . 
