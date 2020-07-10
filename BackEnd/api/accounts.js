import { db } from '../utils/db.js';
import express from 'express';

var accountRouter = express.Router();

// Get all the available accounts
accountRouter.get('/all', function (req, res) {
  var query = 'SELECT * FROM accounts;';
  db.any(query)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Get a account entry by its id
accountRouter.get('/id', function (req, res) {
  const props = { id: req.body.id };

  const statement = 'SELECT * FROM accounts WHERE id = ${id};';

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

// Insert a new account entry into the db
accountRouter.post('/', function (req, res) {
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

  db.any(statement, props)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Update the title, and status of a account entry
accountRouter.put('/', function (req, res) {
  const props = {
    id: req.body.id,
    name: req.body.name,
    type: req.body.type,
    card_number: req.body.card_number,
    account_number: req.body.account_number
  };

  // Build statement
  var statement = 'UPDATE accounts SET ';
  for (const [ key, value ] of Object.entries(props)) {
    if (value && key !== 'id') {
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

// Delete a account by id
accountRouter.delete('/:id', function (req, res) {
  const props = { id: req.params.id };
  const statement = 'DELETE FROM accounts WHERE id = ${id} RETURNING *;';

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

export { accountRouter };
