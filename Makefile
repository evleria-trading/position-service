compose:
	docker-compose build && docker-compose up -d
compose-down:
	docker-compose down
lint:
	golangci-lint run ./internal/... && golangci-lint run ./main.go
import:
	goimports -local "github.com/evleria/cats-app" -w .
protoc:
	 rm -rf ./protocol/pb && protoc --proto_path=protocol/proto protocol/proto/*.proto --go_out=./protocol --go-grpc_out=./protocol
grpcui:
	grpcui -plaintext localhost:$(PORT)

.PHONY: compose-build, compose-up, compose-down, lint, import, protoc, grpcui