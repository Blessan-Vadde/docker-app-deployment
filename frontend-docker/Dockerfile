
# Use the official Node.js image as the base image

FROM node:14.17.0


# Set the working directory inside the container

WORKDIR /app


# Copy the frontend folder into the container

COPY frontend/ .


# Install dependencies

RUN npm install


# Expose the port on which the frontend server runs

EXPOSE 3000


# Run the frontend application

CMD ["npm", "start"]

