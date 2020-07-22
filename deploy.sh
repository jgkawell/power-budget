#!/bin/bash

# Export the vars in .env into the shell
export $(egrep -v '^#' .prod.env | xargs)

function backend {
    echo "Deploying BackEnd..."

    # Deploy to heroku
    git subtree push --prefix BackEnd heroku master

    echo "BackEnd deployed."
}

function frontend {
    echo "Deploying FrontEnd..."

    FRONTEND_PROD_ENV="FrontEnd/src/environments/environment.prod.ts"

    # Prep prod env file
    sed -i "s/'BACKEND_PRODUCTION'/$BACKEND_PRODUCTION/" $FRONTEND_PROD_ENV
    sed -i "s,BACKEND_URL,$BACKEND_BASE_URL," $FRONTEND_PROD_ENV

    # Build and deploy
    cd FrontEnd
    ng build --prod
    firebase deploy --only hosting:main
    cd ..

    # Reset prod env file
    sed -i "s/$BACKEND_PRODUCTION/'BACKEND_PRODUCTION'/" $FRONTEND_PROD_ENV
    sed -i "s,$BACKEND_BASE_URL,BACKEND_URL," $FRONTEND_PROD_ENV

    echo "FrontEnd deployed."
}

function all {
    echo "Deploying all components..."

    backend
    frontend

    echo "All components deployed."
}

case $1 in
    "--backend")
        backend
        ;;
    "--frontend")
        frontend
        ;;
    "--all")
        all
        ;;
    *)
        echo "Not a valid argument"
        exit
esac