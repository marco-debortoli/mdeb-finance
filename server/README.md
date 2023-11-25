# MDEB Ledger Backend

This is a go project - so you will need to install the latest version of Go which can be found from the Go website:
- https://go.dev/dl/
- Latest used version: 1.21.4

## MongoDB

I recommend using docker to setup and run a MongoDB instance that the backend will connect to in order to store data. I have attached my `docker-compose.yml` file below that I used to create the instance:

```
version: "3.8"

services:
  mongo:
    image: mongo:latest
    restart: unless-stopped
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: password
    container_name: mongo
    volumes:
      - ./mongodata:/data/db
    ports:
      - 27017:27017
```

## Environment Variables

In this directory, create a `.env` file and add two environment variables:
```
DB_URL=mongodb://root:password@localhost:27017
PORT=8080
```

Feel free to change the port and configure the `DB_URL` based on your deployed mongo instance (the above should work with the `docker-compose.yml` file provided)

## Developing

The easiest way to run the project is to make sure that you are in the same directory as this `README.md`, and then run `go run .`.

Alternatively you can run `go build` and then run the resulting binary.

The server is running properly if you receive the following two messages:
```
<DATE> Connecting to MongoDB database
<DATE> Starting Server on Port <PORT>
```