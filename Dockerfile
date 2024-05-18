# Dockerfile
FROM golang:1.22.2

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install air for live reloading (optional but recommended for development)
RUN go install github.com/cosmtrek/air@latest

# Command to run the executable with air for live reloading
CMD ["air"]
