
export default function(app, db) {
  // Verify the server is up and reachable
  app.get('/', function (req, res) {
    res.send({ msg: 'Server is running...' });
  });

  // Get all the available tasks
  app.get('/todo/all', function (req, res) {
    var query = 'SELECT * FROM tasks;';
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
  app.get('/todo/id', function (req, res) {
    var id = req.body.id;
    var query = `SELECT * FROM tasks WHERE id = ${id};`;
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
  app.post('/todo', function (req, res) {
    console.log(req.body);
    var title = req.body.title;
    var completed = req.body.completed;
    var statement = `INSERT INTO tasks(title, completed) VALUES('${title}', '${completed}') RETURNING *;`;

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
  app.put('/todo', function (req, res) {
    var id = req.body.id;
    var title = req.body.title;
    var completed = req.body.completed;
    var statement = `UPDATE tasks SET title = '${title}', completed =  '${completed}' WHERE id = '${id}' RETURNING *;`;

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
  app.delete('/todo/:id', function (req, res) {
    var id = req.params.id;
    var statement = `DELETE FROM tasks WHERE id = ${id};`;

    db.any(statement)
      .then(() => {
        res.send({ msg: 'Delete successful' });
      })
      .catch((error) => {
        console.error('ERROR:', error.message);
        res.status(500).send('Failed to query database');
      });
  });
};
