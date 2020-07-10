import { db } from '../utils/db.js';
import express from 'express';

var debitRouter = express.Router();

// Get all the available debits
debitRouter.get('/all', function (req, res) {
  var query = 'SELECT * FROM debits;';
  db.any(query)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Get a debit entry by its id
debitRouter.get('/id', function (req, res) {
  const props = { id: req.body.id };

  const statement = 'SELECT * FROM debits WHERE id = ${id};';

  db.any(statement, props)
    .then((results) => {
      if (results.length == 0) {
        res.send({ msg: 'id did not match any database records' });
      } else {
        res.send(results);
      }
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Insert a new debit entry into the db
debitRouter.post('/', function (req, res) {
  const props = {
    posted_date: req.body.posted_date,
    amount: req.body.amount,
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

  db.any(statement, props)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Update the title, and status of a debit entry
debitRouter.put('/', function (req, res) {
  const props = req.body;

  // Build statement
  var statement = 'UPDATE debits SET ';
  for (const key of Object.keys(props)) {
    if (key != 'id') {
      statement += key + ' = ${' + key + '}, ';
    }
  }
  statement = statement.slice(0, -2);
  statement += ' WHERE id = ${id} RETURNING *;';

  db.any(statement, props)
    .then((results) => {
      if (results.length == 0) {
        res.send({ msg: 'id did not match any database records' });
      } else {
        res.send(results);
      }
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Delete a debit by id
debitRouter.delete('/:id', function (req, res) {
  const props = { id: req.params.id };
  const statement = 'DELETE FROM debits WHERE id = ${id} RETURNING *;';

  db.any(statement, props)
    .then((results) => {
      if (results.length == 0) {
        res.send({ msg: 'id did not match any database records' });
      } else {
        res.send(results);
      }
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

export { debitRouter };
