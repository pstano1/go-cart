protoc:
	@cd proto && protoc --go_out=../protopb --go_opt=paths=source_relative \
	--go-grpc_out=../protopb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protopb --grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt=generate_unbound_methods=true \
	./**/*.proto
	