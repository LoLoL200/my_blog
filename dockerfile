# # ---------- Stage 1: Builder ----------
# FROM golang:1.25-alpine AS builder

# # Installing dependencies
# RUN apk add --no-cache git

# # Working directory
# WORKDIR /app

# # Copy go.mod and go.sum separately (for cache)
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the entire project
# COPY . .

# # Build the binary
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog

# # ---------- Stage 2: Runtime ----------
# FROM alpine:3.20

# # Install CA certificates (if there are HTTP requests)
# RUN apk add --no-cache ca-certificates

# WORKDIR /app

# # Copy the binary from builder
# COPY --from=builder /app/blog .

# # Folder for data (articles)
# RUN mkdir -p /app/data

# # Environment variables (defoult)
# ENV PORT=8080
# ENV DATA_DIR=/app/data

# # Open port
# EXPOSE 8080

# # Start
# CMD ["./blog"]
# ---------- Stage 1: Builder ----------
# ---------- Stage 1: Builder ----------
FROM golang:1.25-alpine AS builder

# Installing dependencies
RUN apk add --no-cache git

# Working directory
WORKDIR /app

# Copy go.mod and go.sum separately (for cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog

# ---------- Stage 2: Runtime ----------
FROM alpine:3.20

# Install CA certificates (if there are HTTP requests)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/blog .

# 👇 ЭТО КРИТИЧНО ДЛЯ ТВОЕГО ПРОЕКТА
COPY --from=builder /app/articles ./articles
COPY --from=builder /app/models ./models
COPY --from=builder /app/templates ./templates

# Folder for data (articles)
RUN mkdir -p /app/data

# Environment variables (defoult)
ENV PORT=8080
ENV DATA_DIR=/app/data

# Open port
EXPOSE 8080

# Start
CMD ["./blog"]