# Estágio de construção
FROM golang:1.22 as builder

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN go build -o app-globo-go ./cmd

# Estágio de produção
FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /go/src/app/app-globo-go .

CMD ["./app-globo-go"]

