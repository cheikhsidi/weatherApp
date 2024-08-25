# Variables
APP_NAME = go_weather_app
DOCKER_IMAGE = $(APP_NAME):latest
DOCKER_CONTAINER = $(APP_NAME)_container

# Default target to build the application
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	@echo "Building the application..."
	go build -o main ./cmd/server

# Run the application locally
.PHONY: run
run: build
	@echo "Running the application..."
	./main

# Run the tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# Build the Docker image
.PHONY: docker-build
docker-build:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
.PHONY: docker-run
docker-run:
	@echo "Running the Docker container..."
	docker run -d --name $(DOCKER_CONTAINER) -p 3001:3001 $(DOCKER_IMAGE)

# Stop the Docker container
.PHONY: docker-stop
docker-stop:
	@echo "Stopping the Docker container..."
	docker stop $(DOCKER_CONTAINER)

# Clean up the build and Docker artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f main
	docker rm -f $(DOCKER_CONTAINER) || true
	docker rmi -f $(DOCKER_IMAGE) || true
