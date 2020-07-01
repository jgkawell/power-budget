import { db } from '../utils/db.js';
import express from 'express';

var todoRouter = express.Router();

// Get all the available todos
todoRouter.get('/all', function (req, res) {
  var query = 'SELECT * FROM todos;';
  db.any(query)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Get a todo item by its id
todoRouter.get('/id', function (req, res) {
  var id = req.body.id;
  var query = `SELECT * FROM todos WHERE id = ${id};`;
  db.any(query)
    .then((results) => {
      res.send(results);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Insert a new todo item into the db
todoRouter.post('/', function (req, res) {
  console.log(req.body);
  var title = req.body.title;
  var completed = req.body.completed;
  var statement = `INSERT INTO todos(title, completed) VALUES('${title}', '${completed}') RETURNING *;`;

  db.any(statement)
    .then((results) => {
      res.send(results[0]);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Update the title, and status of a todo item
todoRouter.put('/', function (req, res) {
  var id = req.body.id;
  var title = req.body.title;
  var completed = req.body.completed;
  var statement = `UPDATE todos SET title = '${title}', completed =  '${completed}' WHERE id = '${id}' RETURNING *;`;

  db.any(statement)
    .then((results) => {
      res.send(results[0]);
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

// Delete a todo by id
todoRouter.delete('/:id', function (req, res) {
  var id = req.params.id;
  var statement = `DELETE FROM todos WHERE id = ${id};`;

  db.any(statement)
    .then(() => {
      res.send({ msg: 'Delete successful' });
    })
    .catch((error) => {
      console.error('ERROR:', error.message);
      res.status(500).send('Failed to query database');
    });
});

export { todoRouter };
