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

build_prod:
	docker build \
		-t namnam206/vca-api:1.1.0 \
		--build-arg="BUILD_MODULE=public" \
		-f ./docker/Dockerfile \
		. && \
		docker push namnam206/vca-api:1.1.0
