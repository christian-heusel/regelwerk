FROM golang:1.25.1-bookworm
MAINTAINER Christian Heusel <christian@heusel.eu>

WORKDIR /go/src

COPY . .
RUN go mod download

RUN go build -a -o regelwerk
ENTRYPOINT ["./regelwerk"]
