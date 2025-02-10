FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o bucket-router .
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bucket-router .
EXPOSE 5000
CMD ["./bucket-router"]
