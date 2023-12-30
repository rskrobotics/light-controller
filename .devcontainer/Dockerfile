# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
# COPY go.mod go.sum ./

# Download dependencies (if using go.mod)
# RUN go mod download

# Copy the source code into the container
COPY . .

# Command to run the application
CMD ["go", "run", "./main.go"]
