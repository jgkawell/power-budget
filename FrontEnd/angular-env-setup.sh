#!/bin/sh

FILE="src/environments/environment.prod.ts"

/bin/cat <<EOM >$FILE
export const environment = {
  production: true,
  serverURL: '$BACKEND_BASE_URL:$BACKEND_PORT/todo',
};
EOM
