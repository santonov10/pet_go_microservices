FROM golang:1.17-buster as build

WORKDIR /app
RUN GOOS=linux CGO_ENABLED=0

COPY . .
RUN go mod download

RUN go build -o main ./cmd/grpc-server/main.go

ENTRYPOINT ["/app/main"]