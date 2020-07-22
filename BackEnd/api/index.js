import express from 'express';

var baseRouter = express.Router();

// Verify the server is up and reachable
baseRouter.get('/', function (req, res) {
  res.send({ msg: 'Server is up...' });
});

export { baseRouter };
