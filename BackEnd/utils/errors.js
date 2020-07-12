import config from './config.js';

export class ErrorMessage {
  constructor(err) {
    this.error = {
      name: err.name,
      msg: err.message,
      stack: config.production
        ? 'Stack not shown in production mode'
        : err.stack,
    };
  }
}

export class RequestError extends Error {
  constructor(message) {
    super(message);
    this.name = 'RequestError';
  }
}

export class DatabaseError extends Error {
  constructor(message) {
    super(message);
    this.name = 'DatabaseError';
  }
}
