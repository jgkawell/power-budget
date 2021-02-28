### STAGE 1: Build ###

# base image
FROM node:14.3.0-alpine AS build
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .

# Create Angular env file
ARG API_BASE_URL=${API_BASE_URL}
ARG API_PORT=${API_PORT}
RUN ./angular-env-setup.sh
RUN npm run build

### STAGE 2: Run ###
FROM nginx:1.19.0-alpine
COPY --from=build /app/dist/power-budget /usr/share/nginx/html

# Copy over configuration
COPY nginx/ /etc/nginx/conf.d/

# Edit nginx.conf file
ARG CLIENT_PORT=${CLIENT_PORT}
RUN sed -i "s/change_me/$CLIENT_PORT/g" /etc/nginx/conf.d/nginx.conf
