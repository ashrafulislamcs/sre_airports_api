# Step 1: Build the Go app
FROM golang:1.16 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download the Go modules (dependencies)
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Step 2: Prepare a minimal runtime environment
FROM alpine:latest

# Install necessary CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app will run on
EXPOSE 8080

# Run the Go app
CMD ["./main"]
