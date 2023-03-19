FROM golang AS builder

WORKDIR /app
COPY auth_proxy.go /app
RUN go build auth_proxy.go

FROM rclone/rclone
COPY --from=builder /app/auth_proxy /usr/bin/auth_proxy
EXPOSE 8080
ENV AUTH_DATA_FILE=/data/auth_data.json
CMD ["serve", "webdav", "--auth-proxy", "/usr/bin/auth_proxy", "--addr", ":8080", "--dir-cache-time", "30s"]
