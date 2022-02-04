
.PHONY: compile

compile:
	protoc ./api/v1/control.proto \
		--proto_path=./api/v1 \
		--go_out=./api/v1/go \
		--go-grpc_out=./api/v1/go \
		--go_opt=paths=source_relative \
        --go-grpc_opt=paths=source_relative

.PHONY: cert

cert:
	cfssl gencert \
		-initca config/ca-csr.json | cfssljson -bare ca

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=config/ca-config.json \
		-profile=server \
		config/server-csr.json | cfssljson -bare server

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=config/ca-config.json \
		-profile=client \
		config/client-csr.json | cfssljson -bare client

	mv *.pem *.csr ./certs