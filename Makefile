up:
	docker-compose up

down:
	docker-compose down

tidy:
	cd src/public && go mod tidy
	cd src/worker && go mod tidy
	cd src/core && go mod tidy
	cd src/adapter && go mod tidy

swagger-public:
	cd src/public && swag init --parseDependency --parseDepth=3

test:
	cd src/public && go test ./...

build-api:
	docker build \
		-t namnam206/voice-changer-api:latest \
		--build-arg="BUILD_MODULE=public" \
		-f ./docker/Dockerfile \
		.

push-api:
	docker push namnam206/voice-changer-api:latest

build-worker:
	docker build \
		-t namnam206/voice-changer-worker:latest \
		--build-arg="BUILD_MODULE=worker" \
		-f ./docker/Dockerfile \
		.

push-worker:
	docker push namnam206/voice-changer-worker:latest
