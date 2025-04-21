# Build stage: Utiliza a imagem oficial do Golang para compilar o aplicativo
FROM golang:1.24-alpine AS builder
WORKDIR /app
# Copia todos os arquivos do projeto para o diretório de trabalho no container
COPY . .
# Compila o aplicativo Go, gerando o binário chamado "main"
RUN go build -o main ./cmd/api

# Runtime stage: Utiliza uma imagem Alpine minimalista para executar o aplicativo
FROM alpine:latest
WORKDIR /app
# Copia o binário compilado da etapa anterior para o container final
COPY --from=builder /app/main .
# Expõe a porta 8080 para permitir o acesso ao servidor
EXPOSE 8080
# Define o comando padrão para iniciar o aplicativo
CMD ["./main"]