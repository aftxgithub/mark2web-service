FROM golang:alpine

WORKDIR /app
# goreleaser handles the build and gives us the binary
COPY mark2web-service /go/bin

ARG PORT=9090

RUN mkdir /var/www
ENV GO111MODULE=on M2W_DB_PATH=/var/www M2W_PORT=${PORT}
EXPOSE ${PORT}

ENTRYPOINT ["mark2web-service"]