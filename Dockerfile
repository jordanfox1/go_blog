FROM alpine:3.15 as root-certs
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:1.21.1 as builder
WORKDIR /usr/src/app
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


COPY . .
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./main

FROM scratch as final
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=1001:1001 --from=builder /usr/src/app /app
USER app
ENTRYPOINT ["/main"]