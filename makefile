PHONY: build run

all: build run

build:
	@echo "Building docker image..."
	docker build -t test-app:latest .

run:
	@echo "Running docker container..."
	docker run -p 8080:8080 --env-file ./.env test-app:latest 