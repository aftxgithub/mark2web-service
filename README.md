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

## Usage
Start the service with sensible defaults:
```shell
$ ./mark2web-service
TRAC[0000] creating a new server for addr :8080         
.
.
.
INFO[0000] starting server on :8080 
```

Request a webpage for a test markdown file:
```shell
$ curl localhost:8080 -F "file=@test.md"
localhost:8080/f005ed3e144fc4de6adbd90292e77d557e88ce73
```
Visit the link in your browser to see the rendered webpage.

![image](https://user-images.githubusercontent.com/42256651/123470336-2c781f00-d5ec-11eb-85ac-8cff88c88613.png)


### Config
Start the service with optional environment configurations:
    
- M2W_PORT: Specify the port the server runs on
- M2W_LOG_LEVEL: Choose preferred log level. See [Logrus log levels](https://pkg.go.dev/github.com/sirupsen/logrus#readme-level-logging)

#### Data store
The service offers two storage options: ```Filesystem``` or ```Firebase```. Configure using one of these env vars:
- M2W_DB_PATH: Path to store the filesystem database
- GOOGLE_APPLICATION_CREDENTIALS: Path to firebase service account json

If GOOGLE_APPLICATION_CREDENTIALS is set in the environment, it takes precedence over the filesystem.

Start the service on a custom port:
```shell
$ M2W_PORT=9000 ./mark2web-service
TRAC[0000] creating a new server for addr :9000         
.
.
.
INFO[0000] starting server on :9000
```

Change filesystem database directory:
```shell
$ M2W_DB_PATH=/home/var/www ./mark2web-service
```

Set TRACE log level:
```shell
$ M2W_LOG_LEVEL=trace ./mark2web-service
```

## Thanks
- [@Meghatronics](https://github.com/Meghatronics): Responsible for the visual aesthetics.
- [@EmmanuelOluwafemi](https://github.com/EmmanuelOluwafemi): Initial work on the stylesheet.