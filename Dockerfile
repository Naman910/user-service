FROM golang:1.16 AS builder

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o user-service .

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/user-service .

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./user-service"]
