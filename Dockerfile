FROM golang:1.17 as base

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

FROM golang:1.17 as builder

WORKDIR /app

COPY . .
RUN make build

FROM scratch

COPY --from=builder /app/bin/go-login-example /app/bin/go-login-example

CMD ["/app/bin/go-login-example"]
