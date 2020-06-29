var config = {
  production: process.env.BACKEND_PRODUCTION,
  port: process.env.BACKEND_PORT,
};

config.db = {
  host: process.env.DB_HOST,
  port: process.env.DB_PORT,
  database: process.env.DB_DATABASE,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
};

module.exports = config;
