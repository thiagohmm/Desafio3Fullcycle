# Etapa de build
FROM golang:1.23-alpine AS build

WORKDIR /app

# Copia os arquivos de código e o arquivo .env da pasta correta
COPY . .
COPY ./cmd/ordersystem/.env ./cmd/ordersystem/.env

# Instala o certificado SSL necessário
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Baixa as dependências do projeto
RUN go mod tidy

# Lista os arquivos no diretório de trabalho
RUN ls -la /app

# Correção para o comando de build
RUN echo "Building the application..."
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cleanArch ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

# Etapa final
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/cleanArch .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./cmd/ordersystem/.env .env

ENTRYPOINT ["./cleanArch"]
