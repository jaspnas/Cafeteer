FROM golang:1.19 as builder
RUN mkdir "/app"
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main ./
EXPOSE 8080
CMD ["./main"]
