# ビルドステージ
FROM golang:1.22-alpine AS builder

WORKDIR /app

# 依存関係をキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてバイナリをビルド
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# 実行ステージ
FROM gcr.io/distroless/static:nonroot

WORKDIR /

# ビルド済みバイナリをコピー
COPY --from=builder /app/main /

# ポート番号の公開
EXPOSE 8080

# 非rootユーザーで実行
USER nonroot:nonroot

# 実行コマンド
CMD ["/main"]
