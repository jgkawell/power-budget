#!/bin/sh

if test "$1" = "--local"; then
  # Export the vars in .env into the shell
  export $(egrep -v '^#' ../.env | xargs)
fi

FILE="src/environments/environment.prod.ts"

/bin/cat <<EOM >$FILE
export const environment = {
  production: true,
  serverURL: '$BACKEND_BASE_URL:$BACKEND_PORT',
};
EOM
