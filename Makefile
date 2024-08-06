run:
	go run main.go

run-watch:
	air

test:
	go test ./... -cover

mocks:
	mockery --all