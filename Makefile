.PHONY: install
install:
	go install .

.PHONY: test
test:
	go test -i .
	go test -v .
