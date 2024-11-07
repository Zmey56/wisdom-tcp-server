PROJECT_NAME := word-of-wisdom
SERVER_IMAGE := $(PROJECT_NAME)-server
CLIENT_IMAGE := $(PROJECT_NAME)-client
TEST_FLAGS := -v

# Build Docker Images
.PHONY: build-server build-client
build-server:
	docker build -t word-of-wisdom-server -f cmd/server/Dockerfile .

build-client:
	docker build -t word-of-wisdom-client -f cmd/client/Dockerfile .

# Run Containers
.PHONY: run-server run-client
run-server:
	docker run --rm -p 8080:8080 $(SERVER_IMAGE)

run-client:
	docker run --rm --network="host" $(CLIENT_IMAGE)

# Run Tests
.PHONY: test
test:
	go test -v ./... -tags=integration

# Run Tests with Testcontainers-Go
.PHONY: test-testcontainers
test-testcontainers:
	go test $(TEST_FLAGS) ./internal/server -run TestHandleConnection

# Clean up Docker Containers and Images
.PHONY: clean
clean:
	docker rmi -f $(SERVER_IMAGE) $(CLIENT_IMAGE)
	docker system prune -f

# Full Cycle: Build, Test, and Run
.PHONY: all
all: build-server build-client test run-server

