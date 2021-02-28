#!/bin/sh

FILE="src/environments/environment.prod.ts"

/bin/cat <<EOM >$FILE
export const environment = {
  production: true,
  serverURL: '$API_BASE_URL:$API_PORT',
};
EOM
