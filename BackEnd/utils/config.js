export default {
  production: (process.env.BACKEND_PRODUCTION === 'true'),
  deployed: (process.env.BACKEND_DEPLOYED === 'true'),
  port: process.env.PORT || process.env.BACKEND_PORT,
  db: {
    host: process.env.DB_HOST,
    port: process.env.DB_PORT,
    database: process.env.DB_DATABASE,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
  },
};
