import { db } from '../utils/db.js';
import { camelKeysArray } from '../utils/convert-case.js';
import express from 'express';
import { NO_RECORDS_MESSAGE } from '../utils/constants.js';
import { DatabaseError, RequestError } from '../utils/errors.js';

var accountRouter = express.Router();

// Get all the available accounts
accountRouter.get('/all', function (req, res, next) {
  var query = 'SELECT * FROM accounts;';
  db.any(query)
    .then((results) => {
      res.send(camelKeysArray(results));
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

// Get a account entry by its id
accountRouter.get('/id', function (req, res, next) {
  const props = { id: req.body.id };
  const statement = 'SELECT * FROM accounts WHERE id = ${id};';

  // Make sure id was given
  if (!props.id) {
    throw new RequestError('id is required');
  }

  db.any(statement, props)
    .then((results) => {
      if (results.length == 0) {
        next(new RequestError(NO_RECORDS_MESSAGE));
      } else {
        res.send(camelKeysArray(results));
      }
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

// Insert a new account entry into the db
accountRouter.post('/', function (req, res, next) {
  const props = {
    name: req.body.name,
    type: req.body.type,
    card_number: req.body.card_number || '',
    account_number: req.body.account_number || ''
  };

  const statement =
    'INSERT INTO \
    accounts(name, type, card_number, account_number, balance, total_in, total_out) \
    VALUES(${name}, ${type}, ${card_number}, ${account_number}, 0, 0, 0) \
    RETURNING *;';

  if (!props.name || !props.type) {
    throw new RequestError('name and type are required');
  }

  db.any(statement, props)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

// Update the title, and status of a account entry
accountRouter.put('/', function (req, res, next) {
  const props = {
    id: req.body.id,
    name: req.body.name,
    type: req.body.type,
    card_number: req.body.card_number,
    account_number: req.body.account_number
  };

  // Make sure id was given
  if (!props.id) {
    throw new RequestError('id is required');
  }

  // Build statement
  var statement = 'UPDATE accounts SET ';
  var hasValues = false;
  for (const [ key, value ] of Object.entries(props)) {
    if (value && key !== 'id') {
      statement += key + ' = ${' + key + '}, ';
      hasValues = true;
    }
  }
  statement = statement.slice(0, -2);
  statement += ' WHERE id = ${id} RETURNING *;';

  if (!hasValues) {
    throw new RequestError('No values to update');
  }

  db.any(statement, props)
    .then((results) => {
      if (results.length == 0) {
        next(new RequestError(NO_RECORDS_MESSAGE));
      } else {
        res.send(camelKeysArray(results));
      }
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

// Delete a account by id
accountRouter.delete('/', function (req, res, next) {
  const props = { id: req.body.id };
  const statement = 'DELETE FROM accounts WHERE id = ${id} RETURNING *;';

  // Make sure id was given
  if (!props.id) {
    throw new RequestError('id is required');
  }

  db.any(statement, props)
    .then((results) => {
      if (results.length == 0) {
        next(new RequestError(NO_RECORDS_MESSAGE));
      } else {
        res.send(camelKeysArray(results));
      }
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

export { accountRouter };
