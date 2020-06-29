FROM node:12.18.0-alpine
COPY package.json ./
RUN npm install
COPY . .
CMD node server.js