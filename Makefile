build:
	@go build -o bin/app

run: build
	@./bin/app

run-node: 
	nodemon --exec go run *.go 

test: 
	@go test -v ./...