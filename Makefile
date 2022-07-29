generate:
	@clear
	@echo "\n ----- Build Protobuf ----- \n\c"
	@echo $(PATH) | grep -i $(HOME)/go/bin --count > /dev/null && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto && echo "OK\n\c" || echo "Error\n\c"