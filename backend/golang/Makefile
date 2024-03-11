.PHONY: TestAcmeData
TestAcmeData:
	@echo "Testing AcmeData"
	go test -v -run ${name} go-learn/acme/internal/modules/data

.PHONY: benchmark
benchmark:
	go test -bench "^${name}" go-learn -benchmem -v

.PHONY: run
run:
	go run main.go
