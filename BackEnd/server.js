var express = require('express');
var app = express();

var cors = require('cors');
var config = require('./config');

var bodyParser = require('body-parser');
app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cors());

require('./api/db')(app);

// If running in prod, use SSL
if (config.production === 'true') {
  console.log('Running in prod mode (https)');

  var https = require('https');
  var fs = require('fs');

  var httpsOptions = {
    key: fs.readFileSync('ssl/server.key'),
    cert: fs.readFileSync('ssl/server.crt'),
  };

  https.createServer(httpsOptions, app).listen(config.port);
} else {
  console.log('Running in dev mode (http)');

  app.listen(config.port);
}

// Log the current port
console.log(`Server up on port ${config.port}`);
