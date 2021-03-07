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
