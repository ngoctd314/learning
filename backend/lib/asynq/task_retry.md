# Task Retry

By default, asynq will retry a task up to 25 times. Every time a task is retried is uses an exponential backoff strategy to caculate the retry delay. If a task exhausts all of its retry count (default: 25), the task will moved to the archive for debugging and inspection purposes and won't be automatically retried (You can still manually run task using CLI or WebUI).

The following properties of task-retry can be customized:

- Max retry count per task
- Time duration to wait (i.e. delay) before a failed task can be retried again
- Whether to consume the retry count for the task
- Whether to skip retry and send the task directly to the archive
