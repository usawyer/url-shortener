run_memory:
	STORAGE_FLAG=memory docker-compose up server redis

run_db:
	STORAGE_FLAG=db docker-compose up server db

lint:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -w .
	gofmt -s -w .
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --out-format colored-line-number -v
