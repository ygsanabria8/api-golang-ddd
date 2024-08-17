run:
	@echo "  >  Running Api DDD..."
	go run main.go

run_watch:
	@echo "  >  Watching Api DDD..."
	air

generate_all_pb: generate_api_pb

generate_api_pb:
	@echo "  >  Building Api .proto..."
	protoc \
	-I src/api/proto/ \
	--go_opt=paths=source_relative \
	--go_out=src/api/proto/pb src/api/proto/*.proto \


tests:
	@echo "  >  Running Tests ..."
	go clean -testcache
	go test ./... -cover

mocks:
	@echo "  >  Generating Mocks..."
	mockery --all

commit:
	@echo "  >  Making Commit..."
	yarn commit

build:
	@echo "  >  Building App..."
	go build

install:
	@echo "  >  Installing Dependencies..."
	go get .