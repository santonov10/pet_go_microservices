FROM golang:1.17-buster as build

WORKDIR /app
RUN GOOS=linux CGO_ENABLED=0

COPY . .
RUN go mod download

RUN go build -o main ./cmd/grpc-server/main.go
RUN go build -o migrate ./cmd/migrate/main.go

#FROM alpine:3.15
#RUN apk add --no-cache bash
#
#COPY --from=build /app/migrations/. /app/migrations
#COPY --from=build /app/config/. /app/config
#COPY --from=build /app/bin/. /app/bin
#COPY --from=build /app/main /app/
#COPY --from=build /app/migrate /app/

ENTRYPOINT ["/app/bin/start.sh"]