FROM golang:1.20.5-alpine3.17

RUN apk update && apk add --no-cache make

WORKDIR /app
COPY go.mod ./
COPY Makefile ./
COPY *.go ./

RUN go mod download

RUN make build

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

CMD [ "make", "run" ]
