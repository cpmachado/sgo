SRC=$(shell find . -type f -name "*.go")

reverse: $(SRC)
	go build ./cmd/reverse

clean:
	rm -r reverse testdata

test:
	go test ./...

fuzztest:
	go test -fuzz=Fuzz -fuzztime=30s 

.PHONY: test fuzztest
