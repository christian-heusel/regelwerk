FROM golang:1.20.3-bullseye
MAINTAINER Christian Heusel <christian@heusel.eu>

WORKDIR /go/src

COPY . .
RUN go mod download

RUN go build -a -o regelwerk
ENTRYPOINT ["./regelwerk"]
