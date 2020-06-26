# Image that your app will build
FROM golang:alpine as builder

ADD . /src
WORKDIR /src

RUN apk update && \
    apk --no-cache add gcc git
RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o \ 
    main

# Image that you app will run
FROM alpine

WORKDIR /app
COPY --from=builder /src/main /app

EXPOSE 8080
ENTRYPOINT ./main