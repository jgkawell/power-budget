import { db } from '../utils/db.js';
import express from 'express';

var creditRouter = express.Router();

// Get all the available credits
creditRouter.get('/all', function (req, res) {
  var query = 'SELECT * FROM credits;';
  db.any(query)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Get a credit entry by its id
creditRouter.get('/id', function (req, res) {
  const props = { id: req.body.id };

  const statement = 'SELECT * FROM credits WHERE id = ${id};';

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

// Insert a new credit entry into the db
creditRouter.post('/', function (req, res) {
  const props = {
    posted_date: req.body.posted_date,
    amount: req.body.amount,
    source: req.body.source,
    purpose: req.body.purpose,
    account: req.body.account,
    budget: req.body.budget,
    notes: req.body.notes || '',
  };

  const statement =
    'INSERT INTO \
    credits(posted_date, amount, source, purpose, account, budget, notes) \
    VALUES(${posted_date}, ${amount}, ${source}, ${purpose}, ${account}, ${budget}, ${notes}) \
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

// Update the title, and status of a credit entry
creditRouter.put('/', function (req, res) {
  const props = req.body;

  // Build statement
  var statement = 'UPDATE credits SET ';
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

// Delete a credit by id
creditRouter.delete('/:id', function (req, res) {
  const props = { id: req.params.id };
  const statement = 'DELETE FROM credits WHERE id = ${id} RETURNING *;';

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

export { creditRouter };
