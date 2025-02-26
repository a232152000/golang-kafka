FROM golang:latest

# 設置必要的環境變數
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o ./cmd/worker/main ./cmd/worker
RUN chmod +x ./cmd/worker/main

EXPOSE 8000

CMD ["./cmd/worker/main"]
