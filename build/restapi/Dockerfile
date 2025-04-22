# Step 1: Use the official Golang 1.24.2 Alpine image to build the application
FROM golang:1.24.2-alpine3.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Install necessary dependencies
RUN apk update && apk add --no-cache git

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN go build -o /bin/account-service .

# Step 2: Create the final image from a minimal Alpine-based image
FROM alpine:3.21.3

# Install necessary libraries (e.g., for certificates and SSL)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder image to /bin
COPY --from=builder /bin/account-service /bin/

# Expose the port the app will run on
EXPOSE 8080

# Set the default command to run the application
CMD ["/bin/account-service"]
