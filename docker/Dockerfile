FROM golang:1.19-bullseye
MAINTAINER Christian Heusel <christian@heusel.eu>

WORKDIR /go/src

COPY . .
RUN go mod download

RUN go build -a -o regelwerk
ENTRYPOINT ["./regelwerk"]
