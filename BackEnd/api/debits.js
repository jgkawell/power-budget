import { db } from '../utils/db.js';
import { camelKeysArray } from '../utils/convert-case.js';
import express from 'express';
import { BUDGET_NUMS, NO_RECORDS_MESSAGE } from '../utils/constants.js';
import { DatabaseError, RequestError } from '../utils/errors.js';

var debitRouter = express.Router();

// Get all the available debits
debitRouter.get('/all', function (req, res, next) {
  var query = 'SELECT * FROM debits;';
  db.any(query)
    .then((results) => {
      res.send(camelKeysArray(results));
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

// Get a debit entry by its id
debitRouter.get('/id', function (req, res, next) {
  const props = { id: req.body.id };
  const statement = 'SELECT * FROM debits WHERE id = ${id};';

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

// Insert a new debit entry into the db
debitRouter.post('/', function (req, res, next) {
  const props = {
    posted_date: req.body.postedDate || new Date(),
    amount: req.body.amount || 0,
    vendor: req.body.vendor,
    purpose: req.body.purpose,
    account: req.body.account,
    budget: req.body.budget,
    notes: req.body.notes || '',
  };

  const statement =
    'INSERT INTO \
    debits(posted_date, amount, vendor, purpose, account, budget, notes) \
    VALUES(${posted_date}, ${amount}, ${vendor}, ${purpose}, ${account}, ${budget}, ${notes}) \
    RETURNING *;';

  if (!props.vendor || !props.purpose || !props.account || !props.budget) {
    throw new RequestError('vendor, purpose, account, and budget are required');
  }

  // Make sure budget number is valid
  if (!BUDGET_NUMS.includes(parseInt(props.budget))) {
    throw new RequestError(`Invalid budget number: ${props.budget}`);
  }

  db.any(statement, props)
    .then((results) => {
      res.send(camelKeysArray(results));
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

// Update the values of a debit entry
debitRouter.put('/', function (req, res, next) {
  const props = {
    id: req.body.id,
    posted_date: req.body.postedDate,
    amount: req.body.amount,
    vendor: req.body.vendor,
    purpose: req.body.purpose,
    account: req.body.account,
    budget: req.body.budget,
    notes: req.body.notes,
  };

  // Make sure id was given
  if (!props.id) {
    throw new RequestError('id is required');
  }

  // Make sure budget number is valid
  if (props.budget && !BUDGET_NUMS.includes(parseInt(props.budget))) {
    throw new RequestError(`Invalid budget number: ${props.budget}`);
  }

  // Build statement
  var statement = 'UPDATE debits SET ';
  var hasValues = false;
  for (const [key, value] of Object.entries(props)) {
    if (value && key != 'id') {
      statement += key + ' = ${' + key + '}, ';
      hasValues = true;
    }
  }

  if (!hasValues) {
    throw new RequestError('No values to update');
  }

  statement = statement.slice(0, -2);
  statement += ' WHERE id = ${id} RETURNING *;';

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

// Delete a debit by id
debitRouter.delete('/', function (req, res, next) {
  const props = { id: req.body.id };
  const statement = 'DELETE FROM debits WHERE id = ${id} RETURNING *;';

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

export { debitRouter };
