# ============================================================
# Build image cho vanthu-backend (Go + Gin + pgx).
# Build:  docker build -t vanthu-backend .
# Run:    docker run -p 10000:10000 -e DATABASE_URL=postgres://... vanthu-backend
# KHÔNG bake .env vào image — DATABASE_URL truyền qua biến môi
# trường lúc chạy (xem .dockerignore).
# ============================================================

# ---- Giai đoạn build ----
FROM golang:1.25-alpine AS build

WORKDIR /src

# Cache layer dependencies: chỉ tải lại khi go.mod/go.sum đổi.
COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

# Binary tĩnh (CGO off) để chạy được trên alpine trần.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" \
    -o /out/server ./cmd/server

# ---- Giai đoạn chạy ----
FROM alpine:3.20

# ca-certificates: bắt buộc để nối TLS tới Neon.
# tzdata: cho timestamp trả về đúng múi giờ khi đặt TZ.
RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -u 10001 app

ENV GIN_MODE=release \
    TZ=Asia/Ho_Chi_Minh

COPY --from=build /out/server /usr/local/bin/server

USER app

# App đọc PORT từ env, mặc định 10000 (internal/config/config.go).
EXPOSE 10000

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s \
    CMD wget -qO- "http://127.0.0.1:${PORT:-10000}/healthz" || exit 1

ENTRYPOINT ["/usr/local/bin/server"]
