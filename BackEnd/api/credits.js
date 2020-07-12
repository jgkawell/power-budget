import { db } from '../utils/db.js';
import { camelKeysArray } from '../utils/convert-case.js';
import express from 'express';
import { BUDGET_NUMS, NO_RECORDS_MESSAGE } from '../utils/constants.js';
import { DatabaseError, RequestError } from '../utils/errors.js';

var creditRouter = express.Router();

// Get all the available credits
creditRouter.get('/all', function (req, res, next) {
  var query = 'SELECT * FROM credits;';
  db.any(query)
    .then((results) => {
      res.send(camelKeysArray(results));
    })
    .catch((error) => {
      next(new DatabaseError(error.message));
    });
});

// Get a credit entry by its id
creditRouter.get('/id', function (req, res, next) {
  const props = { id: req.body.id };
  const statement = 'SELECT * FROM credits WHERE id = ${id};';

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

// Insert a new credit entry into the db
creditRouter.post('/', function (req, res, next) {
  const props = {
    posted_date: req.body.postedDate || new Date(),
    amount: req.body.amount || 0,
    source: req.body.source || '',
    purpose: req.body.purpose || '',
    account: req.body.account || '',
    budget: req.body.budget || 0,
    notes: req.body.notes || '',
  };

  const statement =
    'INSERT INTO \
    credits(posted_date, amount, source, purpose, account, budget, notes) \
    VALUES(${posted_date}, ${amount}, ${source}, ${purpose}, ${account}, ${budget}, ${notes}) \
    RETURNING *;';

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

// Update the values of a credit entry
creditRouter.put('/', function (req, res, next) {
  const props = {
    id: req.body.id,
    posted_date: req.body.postedDate,
    amount: req.body.amount,
    source: req.body.source,
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
  var statement = 'UPDATE credits SET ';
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

// Delete a credit by id
creditRouter.delete('/', function (req, res, next) {
  const props = { id: req.body.id };
  const statement = 'DELETE FROM credits WHERE id = ${id} RETURNING *;';

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

export { creditRouter };
