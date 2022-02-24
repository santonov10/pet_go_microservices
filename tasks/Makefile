generate:
	protoc --proto_path=api/grpc/ --go_out=api/grpc --go-grpc_out=api/grpc api/grpc/*.proto

run:
	go run .\cmd\grpc-server\main.go