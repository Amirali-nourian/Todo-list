# builder stage
FROM golang:alpine AS builder

# install git and ca-certificates (needed for go mod download)
RUN apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

# cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /todo ./cmd

# ... (بخش‌های بالا دست نزنید)

# final stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/

COPY --from=builder /todo /todo
COPY web ./web

EXPOSE 8080
CMD ["/todo"]