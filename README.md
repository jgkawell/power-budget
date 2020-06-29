[![Build Status](https://travis-ci.com/jgkawell/power-budget.svg?branch=master)](https://travis-ci.com/jgkawell/power-budget)

# power-budget

A powerful budgeting PWA written in Angular.

- FrontEnd is Angular
- BackEnd is Node.js connected to Postgres DB

# Installation (dev)

The installation is fairly straightforward since everything is set up in docker-compose (the images are already in [Docker Hub](https://hub.docker.com/u/jgkawell) as well so you shouldn't even have to build them). First, you'll need [Docker](https://docs.docker.com/docker-for-windows/install/) installed on your machine. Then you'll need to run the compose file:

```bash
git clone https://github.com/jgkawell/power-budget.git
cd power-budget
cp example.env .env
docker-compose pull
docker-compose up -d
```

If you run into issues with the prebuilt images you can build them locally and then bring up the containers:

```bash
docker-compose build
docker-compose up -d
```

# Usage (dev)

Once the Docker containers are running, you should be able to navigate to the application in your browser at http://localhost:4200/

# Installation (prod)

If you'd rather run things more securely with HTTPS, you'll have to do a little more set up.

First, you'll need to create your `.crt` and `.key` files. You can do this easily with the script by Ruben Vermeulen in [this repo](https://github.com/RubenVermeulen/generate-trusted-ssl-certificate). Make sure you have it set up to generate with an `alt_name` with your localhost and IP. For example:

```
[alt_names]
DNS.1 = *.localhost
DNS.2 = localhost
IP.1 = 192.168.0.***
```

Once you've generated these files, move copies of them into `BackEnd/ssl/` and `FrontEnd/ssl/` with the names `server.crt` and `server.key`. The result should be like this:

```
.
+-- BackEnd
|   +-- ssl
|       +-- server.crt
|       +-- server.key
+-- FrontEnd
|   +-- ssl
|       +-- server.crt
|       +-- server.key
```

You also need to make sure and install the certificate to your machine that will be accessing the application. On Windows the steps are as follows:

- Double click on the certificate (server.crt)
- Click on the button “Install Certificate …”
- Select whether you want to store it on user level or on machine level
- Click “Next”
- Select “Place all certificates in the following store”
- Click “Browse”
- Select “Trusted Root Certification Authorities”
- Click “Ok”
- Click “Next”
- Click “Finish”
- If you get a prompt, click “Yes”

Once that is all done, you'll need to setup your `.env` file (copy it from the example file) to point to your IP and set the build flag `BACKEND_PRODUCTION=true`. The resulting `.env` file will look like this where the IP address is the one you used to generate the certificate above (remember to make it `https` not `http`):

```
BACKEND_BASE_URL=https://192.168.0.***
BACKEND_PRODUCTION=true
...
```

Now you can build the images locally and bring up the containers:

```bash
docker-compose build
docker-compose up -d
```

# Usage (prod)

Once the Docker containers are running, you should be able to navigate to the application in your browser at https://localhost:4200/
