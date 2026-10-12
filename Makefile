.PHONY: build test bench clean docker

BINARY=evm-sentinel

build:
	go build -v -o $(BINARY) cmd/evmsentinel/main.go

test:
	go test -v ./...

bench:
	go test -bench=. -benchmem ./pkg/scanner/

clean:
	rm -f $(BINARY)
	rm -f *.out

docker:
	docker build -t $(BINARY):latest .
