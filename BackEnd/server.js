import express from 'express';
import cors from 'cors';
import config from './config.js';
import bodyParser from 'body-parser';
import { db } from './utils/db.js';
import tasks from './api/tasks.js';
import https from 'https';
import fs from 'fs';

// Set up application
var app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cors());

// Set up routes
tasks(app, db);

// If running in prod, use SSL
if (config.production === 'true') {
  console.log('Running in prod mode (https)');

  var httpsOptions = {
    key: fs.readFileSync('ssl/server.key'),
    cert: fs.readFileSync('ssl/server.crt'),
  };

  https.createServer(httpsOptions, app).listen(config.port);
} else {
  console.log('Running in dev mode (http)');

  app.listen(config.port);
}

// Log the current port
console.log(`Server up on port ${config.port}`);
