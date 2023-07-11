# Use the official Golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app


# Expose the port on which the backend server runs
EXPOSE 8080

# Run the backend application
CMD ["go", "run", "main.go"]


# Install dependencies
RUN go get ./...


# Copy the backend folder into the container
COPY backend/ .






