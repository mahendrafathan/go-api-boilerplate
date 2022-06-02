FROM golang:alpine

RUN apk update && apk add --no-cache git && apk --no-cache --update add build-base 

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 9090

# ENTRYPOINT ["/app/binary"]
CMD ["/app/binary"]
