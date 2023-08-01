# Use the official Go image as the base image
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the entire application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Create a minimal image to run the Go application
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port that the Go application listens on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
