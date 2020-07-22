#!/bin/sh

if test "$1" = "--local"; then
  # Export the vars in .env into the shell
  export $(egrep -v '^#' ../.env | xargs)
fi

FRONTEND_PROD_ENV="src/environments/environment.prod.ts"

# Prep prod env file
sed -i "s/'BACKEND_PRODUCTION'/$BACKEND_PRODUCTION/" $FRONTEND_PROD_ENV
sed -i "s,BACKEND_URL,$BACKEND_BASE_URL:$BACKEND_PORT," $FRONTEND_PROD_ENV
