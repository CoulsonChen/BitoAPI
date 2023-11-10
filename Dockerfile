# 第一階段：使用 Golang 官方映像進行編譯
FROM golang:latest AS builder
# 設定工作目錄為 /go/src/app
WORKDIR /go/src/app
# 將本地的當前目錄複製到容器的工作目錄
COPY . .
# 下載並安裝應用程式的相依套件
# RUN go get -d -v ./...
COPY go.mod . 
COPY go.sum .
RUN go mod download
# 編譯應用程式
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# 第二階段：使用 Alpine Linux 作為最終基礎映像
FROM alpine:latest
# 前一階段的編譯結果複製到最終基礎映像中
COPY --from=builder /go/src/app/app /app
# 設定工作目錄為 /
WORKDIR /
# 執行應用程式
CMD ["./app"]
# 暴露應用程式所使用的端口
EXPOSE 8080
