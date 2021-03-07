### STAGE 1: Build ###

# base image
FROM node:14.3.0-alpine AS build
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .

# Check lint
RUN npm run lint

# Build Angular app
RUN npm run build

### STAGE 2: Run ###
FROM nginx:1.19.0-alpine
COPY --from=build /app/dist/power-budget /usr/share/nginx/html

# Copy over configuration
COPY nginx/ /etc/nginx/conf.d/
