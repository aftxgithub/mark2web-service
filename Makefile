.PHONY: test
test: 
	go test ./...

.PHONY: clean
clean:
	go clean