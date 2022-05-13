[![CodeFactor](https://www.codefactor.io/repository/github/beschlz/memeclub-api/badge)](https://www.codefactor.io/repository/github/beschlz/memeclub-api)
# memeclub-api

Memeclub is a mere after-work project to create a website where memes can be exchanged privately.

## Running memeclub-api

This project is dockerized. You can simply use docker compose to get up an running.

### Docker compose
You can use docker compose to run this project. Be advised that docker compose does not rebuild the images every time you run it.

First create volumes for this application.

Run

`docker volume create minio_volume` and
`docker volume create memeclub_postgres`

Finally run the stack with

`docker compose -p memeclub -f ./docker/docker-compose/docker-compose.yml up`

To rebuild all the images you can use

` docker compose -f ./docker/docker-compose/docker-compose.yml build
`

In Intellj simply run the ***Start memeclub-api*** run goal
