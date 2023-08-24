.PHONY: TestAll
TestAll:
	go test -v ./...

.PHONY: TestShouldUpdateStats
TestShouldUpdateStats:
	go test -timeout 30s -run ^TestShouldUpdateStats g-sqlmock -v

.PHONY: TestShouldRollbackStatUpdatesOnFailure
TestShouldRollbackStatUpdatesOnFailure:
	go test -timeout 30s -run ^TestShouldRollbackStatUpdatesOnFailure g-sqlmock -v

.PHONY: TestShouldGetPosts
TestShouldGetPosts:
	go test -timeout 30s -run ^TestShouldGetPosts g-sqlmock -v

.PHONY: TestShouldResponseWithErrorOnFailure
TestShouldResponseWithErrorOnFailure:
	go test -timeout 30s -run ^TestShouldResponseWithErrorOnFailure g-sqlmock -v
