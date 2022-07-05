# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
# FROM golang:latest
FROM golang:1.17-alpine

# Add Maintainer Info
LABEL maintainer="M.a_k & Araya"

RUN mkdir ascii-art-web

# Set the Current Working Directory inside the container
WORKDIR /ascii-art-web

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN go build cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/ascii-art-web/main"]