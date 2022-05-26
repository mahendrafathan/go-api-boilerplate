FROM golang:alpine as builder

WORKDIR /
COPY . .
RUN go mod tidy
RUN go build -o /app .

FROM scratch
COPY --from=builder /app /app
EXPOSE 4000

ENTRYPOINT ["/app"]
