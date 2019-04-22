# Build Container
FROM golang:latest as builder
WORKDIR /go/src/github.com/tozastation/clasick-core
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app main.go

# Runtime Container
FROM alpine
ENV DB_VENDOR="mysql"
ENV USER="clasick"
ENV	PASS="clasick"
#ENV PROTOCOL="tcp(127.0.0.1:3306)"
ENV PROTOCOL="tcp(cloud-proxy)"
ENV	DBNAME="clasick"
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/tozastation/clasick-core/app /app
EXPOSE 3001
ENTRYPOINT ["/app"]