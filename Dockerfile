# ---------- Etapa 1: Compilación ----------
    FROM golang:1.24.2-alpine AS builder

    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
    
    # ---------- Etapa 2: Imagen final minimalista ----------
    FROM scratch
    
    WORKDIR /app
    
    COPY --from=builder /app/main .
    
    EXPOSE 8080
    
    ENTRYPOINT ["./main"]