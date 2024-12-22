test:
	go test -v ./...

tidy: 
	go mod tidy

example-basic:
	go run examples/basic/main.go