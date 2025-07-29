# Use a imagem oficial do Go como base
FROM golang:1.18-alpine AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo-api .

# Usar uma imagem mínima para a aplicação final
FROM alpine:latest

# Instalar certificados SSL
RUN apk --no-cache add ca-certificates

# Criar diretório de trabalho
WORKDIR /root/

# Copiar o binário da aplicação do estágio anterior
COPY --from=builder /app/todo-api .

# Expor a porta 8080
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./todo-api"]

