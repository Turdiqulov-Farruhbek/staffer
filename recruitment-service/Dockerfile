# Base image sifatida Go'ni tanlaymiz
FROM golang:1.21 AS builder

# Ishchi katalogni yaratamiz
WORKDIR /app

# Modullari va dependency-larni qo'shamiz
COPY go.mod go.sum ./
RUN go mod download

# Barcha loyihani nusxalaymiz
COPY . .

# Loyiha kodini build qilamiz
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Minimalistic run-time image yaratamiz
FROM alpine:latest

# Ishchi katalogni yaratamiz
WORKDIR /root/      

# Build natijasida olingan binary'ni nusxalaymiz
COPY --from=builder /app/main .
COPY .env .env

# Tarmoq portini ochamiz
EXPOSE 9090

# Ilovani ishga tushiramiz
CMD ["./main"]
