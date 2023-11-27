.PHONY: rebuild
rebuild:
	docker build -t ngoctd/asynq-worker2:latest . && \
	docker push ngoctd/asynq-worker2

.PHONY: redeploy
redeploy:
	kubectl rollout restart deployment depl-worker2-service
