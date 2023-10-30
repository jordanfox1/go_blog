FROM golang:1.21.1 as builder
WORKDIR /usr/src/app

COPY . .
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./main

FROM scratch as final
COPY --from=builder /usr/src/app /app
ENTRYPOINT ["/app/main"]