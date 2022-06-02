FROM golang:alpine

RUN apk update && apk add --no-cache git && apk --no-cache --update add build-base 

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]
