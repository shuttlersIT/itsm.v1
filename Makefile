build:
	@go build -o itsm-mvp

run: build
	@./itsm-mvp

test:
	@go test -v ./...
