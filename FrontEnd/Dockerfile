### STAGE 1: Build ###
# base image
FROM node:14.3.0-alpine AS build
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
# Create Angular env file
ARG BACKEND_BASE_URL=${BACKEND_BASE_URL}
ARG BACKEND_PORT=${BACKEND_PORT}
RUN ./angular-env-setup.sh
RUN npm run build

### STAGE 2: Run ###
FROM nginx:1.19.0-alpine
COPY --from=build /app/dist/power-budget /usr/share/nginx/html
COPY ./nginx-conf-setup.sh ./nginx-conf-setup.sh
# Copy over application certs
COPY ssl/ /etc/nginx/certs
# Copy over configuration
COPY nginx/ /etc/nginx/conf.d/
RUN ./nginx-conf-setup.sh
# Edit nginx.conf file
ARG FRONTEND_PORT=${FRONTEND_PORT}
RUN sed -i "s/change_me/$FRONTEND_PORT/g" /etc/nginx/conf.d/nginx.conf