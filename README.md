[![Build Status](https://travis-ci.com/jgkawell/power-budget.svg?branch=master)](https://travis-ci.com/jgkawell/power-budget)

# power-budget

A powerful budgeting PWA written in Angular.

- FrontEnd is Angular
- BackEnd is Node.js connected to Postgres DB

# Installation

The installation is fairly straightforward since everything is set up in docker-compose. First, you'll need [Docker](https://docs.docker.com/get-docker/) installed on your machine. Then can run the below commands:

```bash
git clone https://github.com/jgkawell/power-budget.git
cd power-budget
cp example.env .env
docker-compose build
docker-compose up -d
```

# Usage

Once the Docker containers are running, you should be able to navigate to the application in your browser at http://localhost:4200/
