
# Use the official Node.js image as the base image
FROM node:14


# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the frontend folder into the container
COPY frontend/package*.json  /

RUN npm install

COPY frontend/ .

# Expose the port on which the frontend server runs
EXPOSE 3000

# Run the frontend application
CMD ["npm", "start"]






