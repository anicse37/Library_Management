# Use the official Golang image as a base
FROM golang:1.23.10

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main .

# Expose the port your app listens on
EXPOSE 8080

# Start the Go application
CMD ["./main"]
