# Project Name Generator

Can't name that new project you're working on?
Great! Then the project name generator is for you!

This project is based on a similar project I created,
written in Python. I'm rebuilding this project to help
me learn Golang.

## Running the Project

A `docker-compose.yml` file has been provided which can
be run with `docker-compose up` providing Docker and
docker-compose are installed on the host machine. The
front end can be accessed on `localhost:5000`. The api
can be accessed on `localhost:8080/api`. Generated words
are provided on the `api/generate` endpoint.

## API

The API is written in Golang and makes use of one
endpoint at the `/api/generate/` that will give a
random word.

## Client

The client app is written in HTML, CSS and JS and
will display a generated word when the `generate`
button is clicked.
