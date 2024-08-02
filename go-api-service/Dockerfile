# Use official golang image
FROM golang:1.22-alpine

# Create and set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Copy environment file
COPY .env ./

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 3001

# Start the application
CMD ["./main"]