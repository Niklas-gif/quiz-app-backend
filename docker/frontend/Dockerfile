# Use the official Node.js image as the base image
FROM node:22-alpine

# Set the working directory
WORKDIR /app

# Clone the frontend repository
RUN apk add --no-cache git
RUN git clone -b development https://github.com/Niklas-gif/quiz-app-frontend.git .

# Install dependencies
RUN npm install

# Build the app for production
RUN npm run build

# Expose port 3000
EXPOSE 3000

# Command to run the app
CMD ["npm", "run", "start"]
