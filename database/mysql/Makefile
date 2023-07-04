DSN=mysql://admin:secret@tcp(192.168.49.2:30200)/db
.PHONY: create
create:
	migrate create -ext sql -dir migration -seq ${name}

.PHONY: up, down, force
up:
	migrate -path migration -database "${DSN}" -verbose up
down:
	migrate -path migration -database "${DSN}" -verbose down