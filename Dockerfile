# Use an official Golang runtime as a parent image
FROM golang:latest

# Environment variables
ENV PORT="8080" \
    GO_ENV="DEV" 


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
