BINARY = gocli

build:
	go build -o $(BINARY) .

clean:
	rm -f $(BINARY)

.PHONY: build clean
