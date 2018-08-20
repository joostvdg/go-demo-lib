FROM golang:1.10 AS build
ADD . /src
WORKDIR /src
RUN go test --cover