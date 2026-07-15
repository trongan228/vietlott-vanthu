# ============================================================
# Build frontend (Vite) rồi phát bằng nginx, đồng thời proxy
# /api sang backend cùng mạng docker (same-origin, khỏi CORS).
# Build qua docker-compose.yml ở gốc repo.
# ============================================================

# ---- Giai đoạn build ----
FROM node:22-alpine AS build

WORKDIR /src

COPY package.json package-lock.json ./
RUN npm ci

COPY index.html vite.config.js ./
COPY src ./src

# VITE_API_URL rỗng → frontend gọi /api tương đối, nginx lo proxy.
# (.env dev bị loại khỏi build context qua .dockerignore)
ENV VITE_API_URL=""
RUN npm run build

# ---- Giai đoạn chạy ----
FROM nginx:1.27-alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /src/dist /usr/share/nginx/html

EXPOSE 80
