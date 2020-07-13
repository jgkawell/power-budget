import config from './config.js';
import pgPromise from 'pg-promise';

// Log db connection errors
const initOptions = {
  error(error, e) {
    if (e.cn) {
      console.log('CN:', e.cn);
      console.log('EVENT:', error.message || error);
    }
  },
};

// Create Database Connection
const pgp = pgPromise(initOptions);
const db = pgp(config.db);

// Test connection
console.log('Setting up DB...');
db.connect()
  .then((obj) => {
    console.log('Connected to database');
    obj.done(); // success, release connection;
  })
  .catch((error) => {
    console.error('ERROR:', error.message);
  });

export { db };
