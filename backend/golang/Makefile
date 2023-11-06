.PHONY: TestAcmeData
TestAcmeData:
	@echo "Testing AcmeData"
	go test -v -run ${name} go-learn/acme/internal/modules/data

.PHONY: benchmark
benchmark:
	go test -benchmem -run=^$ -bench ^${name}\$ go-learn -v

.PHONY: run
run:
	go run main.go
