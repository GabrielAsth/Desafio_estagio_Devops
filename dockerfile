# Estágio 1: Build da aplicação
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copia absolutamente tudo o que já foi baixado localmente
COPY . .

# Compila direto o binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o http-server-projeto-korp main.go

# Estágio 2: Execução em uma imagem limpa e leve
FROM alpine:3.19

WORKDIR /root/

COPY --from=builder /app/http-server-projeto-korp .

EXPOSE 8080

CMD ["./http-server-projeto-korp"]