build:
	docker build -t ngoctd/learn-kafka:latest .

redeploy:
	kubectl delete -f service.yaml && \
	kubectl apply -f service.yaml

.PHONY: benchmark
benchmark:
	go test -bench "^${name}" lkafka/${pkg} -benchmem -v
