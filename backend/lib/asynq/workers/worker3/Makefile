.PHONY: rebuild
rebuild:
	docker build -t ngoctd/asynq-worker3:latest . && \
	docker push ngoctd/asynq-worker3

.PHONY: redeploy
redeploy:
	kubectl rollout restart deployment depl-worker3-service
