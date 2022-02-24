FROM golang:1.17-buster as build

WORKDIR /app
RUN GOOS=linux CGO_ENABLED=0

COPY . .
RUN go mod download

RUN go build -o main ./cmd/http-server/main.go


#FROM alpine:3.15
#RUN apk add --no-cache bash
#
#COPY --from=build /app/main /app/

ENTRYPOINT ["/app/main"]