run:
	@echo "  >  Running Api DDD..."
	go run main.go

run-watch:
	@echo "  >  Watching Api DDD..."
	air

generate:
	@echo "  >  Building .proto..."
	protoc \
		-I=src/api/proto/ \
		--go_opt=paths=source_relative src/api/proto/*.proto \
		--go_out=src/api/proto/pb

test:
	@echo "  >  Running Tests ..."
	go test ./... -cover

mocks:
	@echo "  >  Generating Mocks..."
	mockery --all