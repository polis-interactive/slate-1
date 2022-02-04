
.PHONY: compile

compile:
	protoc ./api/v1/control.proto \
		--proto_path=./api/v1 \
		--go_out=./api/v1/go \
		--go-grpc_out=./api/v1/go \
		--go_opt=paths=source_relative \
        --go-grpc_opt=paths=source_relative