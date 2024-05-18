# Dockerfile
FROM golang:1.22.2 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN GOARCH=arm64 GOOS=linux go build -o main cmd/server/main.go

# # Start a new stage from scratch
# FROM alpine:latest

# # Add ca-certificates for SSL
# RUN apk --no-cache add ca-certificates

# # Set the Current Working Directory inside the container
# WORKDIR /root/

# # Copy the Pre-built binary file from the previous stage
# COPY --from=builder /app/main .
# COPY config/config.yaml /root/config/config.yaml

# # Command to run the executable
CMD ["./main"]
