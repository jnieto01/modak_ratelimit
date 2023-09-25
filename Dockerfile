# Use an official Golang runtime as a parent image
FROM golang:latest

# Environment variables
ENV PORT="8080" \
    GO_ENV="DEV" \
    REDIS_ADDR="modak-redis-container:6379" \
    REDIS_PASSWORD="" \
    REDIS_DB="0"



# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main ./cmd/api/

# Expose a port (e.g., 8080) that your application will run on
EXPOSE 8080

# Define the command to run your application
CMD ["./main"]
