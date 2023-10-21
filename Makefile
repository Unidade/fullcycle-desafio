generate-proto:
	protoc --go_out=application/grpc/pb \
    --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
    --go-grpc_out=application/grpc/pb \
    --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto