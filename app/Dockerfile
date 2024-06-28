# Estágio de compilação
FROM golang:1.22.1 AS builder

WORKDIR /app

# Copiar os arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte
COPY . .

# Compilar o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Estágio final
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o executável compilado do estágio de compilação
COPY --from=builder /app/main .

# Copiar os templates e arquivos estáticos
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Expor a porta que o aplicativo usa
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./main"]