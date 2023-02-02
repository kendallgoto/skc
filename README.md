# SKC Lang Server

## Prerequisites
Relies on [antlr.jar](https://www.antlr.org/download.html) to be installed at /antlr.jar

## Usage
Build the application via Docker and connect in your [browser](http://localhost:8080):
```bash
docker-compose up
```

## Running Natively
Go CPython bindings are pinned to Python 3.17 - if you have Python3.17 headers installed, you can build natively (only works on linux AFAIK):
```bash
go get ./...
make && ./bin/skcserver
```
