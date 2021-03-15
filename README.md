![API Local](https://github.com/jgkawell/power-budget/actions/workflows/api-local.yml/badge.svg?branch=main)
![API Docker](https://github.com/jgkawell/power-budget/actions/workflows/api-docker.yml/badge.svg?branch=main)
![Database Docker](https://github.com/jgkawell/power-budget/actions/workflows/database-docker.yml/badge.svg?branch=main)
![Client Local](https://github.com/jgkawell/power-budget/actions/workflows/client-local.yml/badge.svg?branch=main)
![Client Docker](https://github.com/jgkawell/power-budget/actions/workflows/client-docker.yml/badge.svg?branch=main)


# power-budget

A powerful budgeting PWA served by a locally hosted API and database:

- Client is Angular
- API is Golang
- Database is Postgres DB

**This is still very early days for this project. I welcome feedback (in the form of issues) or PRs if you want to contribute.**

This application will one day replace the (quite extensive) Excel spreadsheet I use to keep track of my personal finances. It should eventually have many (or more) of the below features:

- Debit/credit transaction tracking
- Budgeting categories
- Charts/graphs of spending and saving habits
- Import/export of data to CSV
- Financial account management
- Others?

This project is meant to be hosted locally to avoid the issues of security/privacy. This would mean connecting to it on your local network *only* which is the safest way to go about it. It will also be possible to host it locally using a reverse-proxy, dynamic DNS, and router port-forwarding to make it accessible everywhere. Instructions on how to do all this will be made later down the road.

# Installation

The installation is fairly straightforward since everything is set up in docker-compose. First, you'll need [Docker](https://docs.docker.com/get-docker/) installed on your machine. Then you can run the below commands:

```bash
git clone https://github.com/jgkawell/power-budget.git
cd power-budget
cp example.env .env
docker-compose build
docker-compose up -d
```

# Usage

Once the Docker containers are running, you should be able to navigate to the application in your browser at http://localhost:4200/

# Development

## API

The API for Power Budget is written in Go and resides under the `/api` directory of the repository.

A list of needed tools to develop the API:

- go `v1.16` (I suggest using [gvm](https://github.com/moovweb/gvm) as your go version manager)
- [Docker](https://docs.docker.com/get-docker/) (and `docker-compose`)

First copy the `example.env` file to the root (for docker-compose to read from) and to the `/api` directory (for go to read from):

```bash
cp ./example.env .env
cp ./example.env ./api/.env
```

Then start up the Database:

```bash
docker-compose up -d database
```

Then install the needed go modules and run the application:

```bash
cd ./api
go mod download
make run
```

You should now be able to interact with the API at `http://localhost:8080`

## Client

The Client for Power Budget is written in Angluar (typescript) and resides under the `/client` directory of the repository.

A list of needed tools to develop the Client:

- Node (I recommend using [nvm](https://github.com/nvm-sh/nvm) as your node version manager)
- [Angular CLI](https://cli.angular.io/) `v1.9.1`
- [Docker](https://docs.docker.com/get-docker/) (and `docker-compose`)

First, go ahead and start up the API and Database in Docker:

```bash
docker-compose up -d database api
```

Then install the needed modules and run the application:

```bash
npm install
npm run start
```
