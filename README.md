# mark2web-service

[![Go](https://github.com/thealamu/mark2web-service/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/thealamu/mark2web-service/actions/workflows/go.yml)

A Markdown to Webpage service. Mark2Web receives markdown and returns a link to it rendered as a static webpage.

## Installation
To get started, you can download a binary for your platform from [releases](https://github.com/thealamu/mark2web-service/releases) or pull the latest Docker image.

### Run on Docker
Pull the latest image from Dockerhub:
```shell
$ docker pull vague369/mark2web-service:latest
```
Start an interactive container, exposing port 9090:
```shell
$ docker run --it -p9090:9090 vague369/mark2web-service
```
Or expose the service on a custom port using the M2W_PORT environment variable:
```shell
$ docker run --it -p8000:8000 -e M2W_PORT=8000 vague369/mark2web-service
```

### Run a release
Obtain a binary from [releases](https://github.com/thealamu/mark2web-service/releases) or build from source.
#### Build from source
To build from source, you must have installed, [Go](https://golang.org) >= 1.16
```shell
$ git clone https://github.com/thealamu/mark2web-service
$ cd mark2web-service
$ go build ./cmd/mark2web-service
```
