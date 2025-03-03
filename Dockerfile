FROM golang:latest

# 設置環境變數
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 設置工作目錄
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go mod download
RUN cd ./cmd/worker
RUN go build -o main .
EXPOSE 8000

CMD ["./main"]
