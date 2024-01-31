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

api:
	docker build \
		-t namnam206/voice-changer-api:latest \
		--build-arg="BUILD_MODULE=public" \
		-f ./docker/Dockerfile \
		. && \
		docker push namnam206/voice-changer-api:latest

worker:
	docker build \
		-t namnam206/voice-changer-worker:latest \
		--build-arg="BUILD_MODULE=worker" \
		-f ./docker/Dockerfile \
		. && \
		docker push namnam206/voice-changer-worker:latest
