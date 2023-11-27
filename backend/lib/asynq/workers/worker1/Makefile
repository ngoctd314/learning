.PHONY: rebuild
rebuild:
	docker build -t ngoctd/asynq-worker1:latest . && \
	docker push ngoctd/asynq-worker1

.PHONY: redeploy
redeploy:
	kubectl rollout restart deployment depl-worker-service
