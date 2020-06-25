FROM golang:alpine as builder

ADD . /src
WORKDIR /src

RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o \ 
    main

FROM alpine

WORKDIR /app

COPY --from=builder /src/main /app

EXPOSE 8080
ENTRYPOINT ./app