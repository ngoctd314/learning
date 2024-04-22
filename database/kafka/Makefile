build:
	docker build -t ngoctd/learn-kafka:latest .

redeploy:
	kubectl delete -f service.yaml && \
	kubectl apply -f service.yaml
