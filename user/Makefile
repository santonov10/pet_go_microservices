#export POSTGRES_DB_SERVER = test

generate:
	protoc --proto_path=api/grpc/ --go_out=api/grpc --go-grpc_out=api/grpc api/grpc/*.proto

migration_up:
	go run ./cmd/migrate/main.go up

run:
	go run .\cmd\grpc-server\main.go