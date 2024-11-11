FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY docker .
RUN go build -o main ./cmd/api/main.go
CMD ["/app/main"]
