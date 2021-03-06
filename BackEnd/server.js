import bodyParser from 'body-parser';
import config from './utils/config.js';
import cors from 'cors';
import express from 'express';
import { baseRouter } from './api/index.js';
import { accountRouter } from './api/accounts.js';
import { creditRouter } from './api/credits.js';
import { debitRouter } from './api/debits.js';
import { ErrorMessage } from './utils/errors.js';

// Set up application
var app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cors());

// Initialize routing
app.use('/api', baseRouter);
app.use('/api/account', accountRouter);
app.use('/api/credit', creditRouter);
app.use('/api/debit', debitRouter);

// Error handling
app.use(function (err, req, res, next) { // eslint-disable-line no-unused-vars
  if (err.name === 'RequestError') {
    res.status(400).json(new ErrorMessage(err));
  }
  if (err.name === 'DatabaseError') {
    res.status(500).json(new ErrorMessage(err));
  }
});

console.log('Running in http mode');
app.listen(config.port);

// Log the current port
console.log(`Server up on port ${config.port}`);
