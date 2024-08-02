# Stage 1: Build the React application
FROM node:14 AS build

# Create and set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Serve the built application using Nginx
FROM nginx:alpine

# Copy the build output to Nginx's public folder
COPY --from=build /app/dist /usr/share/nginx/html

# Expose the port
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
