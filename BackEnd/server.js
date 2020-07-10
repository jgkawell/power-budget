import bodyParser from 'body-parser';
import config from './config.js';
import cors from 'cors';
import express from 'express';
import fs from 'fs';
import https from 'https';
import { baseRouter } from './api/index.js';
import { accountRouter } from './api/accounts.js';
import { creditRouter } from './api/credits.js';
import { debitRouter } from './api/debits.js';

// Set up application
var app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cors());

// Initialize routing
app.use('/', baseRouter);
app.use('/account', accountRouter)
app.use('/credit', creditRouter);
app.use('/debit', debitRouter);

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
