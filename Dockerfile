# Use a slim GoLang runtime as the base image
FROM golang:1.21.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local source code to the container's working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose a port that the application will run on
EXPOSE 3030

# Command to run the executable when the container starts
CMD ["./main"]
