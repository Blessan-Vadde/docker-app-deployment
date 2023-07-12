# Use the official Golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the backend folder into the container
COPY backend/ .

# Install dependencies
RUN go get ./...

# Setup Environment Variables
ENV DB_HOST="dev-database.cx2v0znhuvnq.eu-central-1.rds.amazonaws.com"
ENV DB_USER="postgres"
ENV DB_PASSWORD="adminpassword"
ENV DB_NAME="nspglobal"
ENV DB_PORT="5432"
ENV ALLOWED_ORIGINS = "true"

# Expose the port on which the backend server runs
EXPOSE 8080

# Run the backend application
CMD ["go", "run", "main.go"]










