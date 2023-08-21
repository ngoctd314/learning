.PHONY: TestAll
TestAll:
	go test -v ./...

.PHONY: TestShouldUpdateStats
TestShouldUpdateStats:
	go test -timeout 30s -run ^TestShouldUpdateStats g-sqlmock -v

.PHONY: TestShouldRollbackStatUpdatesOnFailure
TestShouldRollbackStatUpdatesOnFailure:
	go test -timeout 30s -run ^TestShouldRollbackStatUpdatesOnFailure g-sqlmock -v

